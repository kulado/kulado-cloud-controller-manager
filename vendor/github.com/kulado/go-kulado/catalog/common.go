package catalog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

const (
	SELF       = "self"
	COLLECTION = "collection"
)

var (
	debug  = false
	dialer = &websocket.Dialer{}
)

type ClientOpts struct {
	Url       string
	AccessKey string
	SecretKey string
	Timeout   time.Duration
}

type ApiError struct {
	StatusCode int
	Url        string
	Msg        string
	Status     string
	Body       string
}

func (e *ApiError) Error() string {
	return e.Msg
}

func IsNotFound(err error) bool {
	apiError, ok := err.(*ApiError)
	if !ok {
		return false
	}

	return apiError.StatusCode == http.StatusNotFound
}

func newApiError(resp *http.Response, url string) *ApiError {
	contents, err := ioutil.ReadAll(resp.Body)
	var body string
	if err != nil {
		body = "Unreadable body."
	} else {
		body = string(contents)
	}

	data := map[string]interface{}{}
	if json.Unmarshal(contents, &data) == nil {
		delete(data, "id")
		delete(data, "links")
		delete(data, "actions")
		delete(data, "type")
		delete(data, "status")
		buf := &bytes.Buffer{}
		for k, v := range data {
			if v == nil {
				continue
			}
			if buf.Len() > 0 {
				buf.WriteString(", ")
			}
			fmt.Fprintf(buf, "%s=%v", k, v)
		}
		body = buf.String()
	}
	formattedMsg := fmt.Sprintf("Bad response statusCode [%d]. Status [%s]. Body: [%s] from [%s]",
		resp.StatusCode, resp.Status, body, url)
	return &ApiError{
		Url:        url,
		Msg:        formattedMsg,
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       body,
	}
}

func contains(array []string, item string) bool {
	for _, check := range array {
		if check == item {
			return true
		}
	}

	return false
}

func appendFilters(urlString string, filters map[string]interface{}) (string, error) {
	if len(filters) == 0 {
		return urlString, nil
	}

	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range filters {
		if l, ok := v.([]string); ok {
			for _, v := range l {
				q.Add(k, v)
			}
		} else {
			q.Add(k, fmt.Sprintf("%v", v))
		}
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

func setupKuladoBaseClient(kuladoClient *KuladoBaseClientImpl, opts *ClientOpts) error {
	u, err := url.Parse(opts.Url)
	if err != nil {
		return err
	}

	if u.Path == "" || u.Path == "/" {
		u.Path = "v2-beta"
	} else if u.Path == "/v1" || strings.HasPrefix(u.Path, "/v1/") {
		u.Path = strings.Replace(u.Path, "/v1", "/v2-beta", 1)
	}
	opts.Url = u.String()

	if opts.Timeout == 0 {
		opts.Timeout = time.Second * 10
	}
	client := &http.Client{Timeout: opts.Timeout}
	req, err := http.NewRequest("GET", opts.Url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(opts.AccessKey, opts.SecretKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return newApiError(resp, opts.Url)
	}

	schemasUrls := resp.Header.Get("X-API-Schemas")
	if len(schemasUrls) == 0 {
		return errors.New("Failed to find schema at [" + opts.Url + "]")
	}

	if schemasUrls != opts.Url {
		req, err = http.NewRequest("GET", schemasUrls, nil)
		req.SetBasicAuth(opts.AccessKey, opts.SecretKey)
		if err != nil {
			return err
		}

		resp, err = client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return newApiError(resp, opts.Url)
		}
	}

	var schemas Schemas
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &schemas)
	if err != nil {
		return err
	}

	kuladoClient.Opts = opts
	kuladoClient.Schemas = &schemas

	for _, schema := range schemas.Data {
		kuladoClient.Types[schema.Id] = schema
	}

	return nil
}

func NewListOpts() *ListOpts {
	return &ListOpts{
		Filters: map[string]interface{}{},
	}
}

func (kuladoClient *KuladoBaseClientImpl) setupRequest(req *http.Request) {
	req.SetBasicAuth(kuladoClient.Opts.AccessKey, kuladoClient.Opts.SecretKey)
}

func (kuladoClient *KuladoBaseClientImpl) newHttpClient() *http.Client {
	if kuladoClient.Opts.Timeout == 0 {
		kuladoClient.Opts.Timeout = time.Second * 10
	}
	return &http.Client{Timeout: kuladoClient.Opts.Timeout}
}

func (kuladoClient *KuladoBaseClientImpl) doDelete(url string) error {
	client := kuladoClient.newHttpClient()
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	kuladoClient.setupRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(ioutil.Discard, resp.Body)

	if resp.StatusCode >= 300 {
		return newApiError(resp, url)
	}

	return nil
}

func (kuladoClient *KuladoBaseClientImpl) Websocket(url string, headers map[string][]string) (*websocket.Conn, *http.Response, error) {
	return dialer.Dial(url, http.Header(headers))
}

func (kuladoClient *KuladoBaseClientImpl) doGet(url string, opts *ListOpts, respObject interface{}) error {
	if opts == nil {
		opts = NewListOpts()
	}
	url, err := appendFilters(url, opts.Filters)
	if err != nil {
		return err
	}

	if debug {
		fmt.Println("GET " + url)
	}

	client := kuladoClient.newHttpClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	kuladoClient.setupRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return newApiError(resp, url)
	}

	byteContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if debug {
		fmt.Println("Response <= " + string(byteContent))
	}

	if err := json.Unmarshal(byteContent, respObject); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to parse: %s", byteContent))
	}

	return nil
}

func (kuladoClient *KuladoBaseClientImpl) List(schemaType string, opts *ListOpts, respObject interface{}) error {
	return kuladoClient.doList(schemaType, opts, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) doList(schemaType string, opts *ListOpts, respObject interface{}) error {
	schema, ok := kuladoClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.CollectionMethods, "GET") {
		return errors.New("Resource type [" + schemaType + "] is not listable")
	}

	collectionUrl, ok := schema.Links[COLLECTION]
	if !ok {
		return errors.New("Failed to find collection URL for [" + schemaType + "]")
	}

	return kuladoClient.doGet(collectionUrl, opts, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) doNext(nextUrl string, respObject interface{}) error {
	return kuladoClient.doGet(nextUrl, nil, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) Post(url string, createObj interface{}, respObject interface{}) error {
	return kuladoClient.doModify("POST", url, createObj, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) GetLink(resource Resource, link string, respObject interface{}) error {
	url := resource.Links[link]
	if url == "" {
		return fmt.Errorf("Failed to find link: %s", link)
	}

	return kuladoClient.doGet(url, &ListOpts{}, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) doModify(method string, url string, createObj interface{}, respObject interface{}) error {
	bodyContent, err := json.Marshal(createObj)
	if err != nil {
		return err
	}

	if debug {
		fmt.Println(method + " " + url)
		fmt.Println("Request => " + string(bodyContent))
	}

	client := kuladoClient.newHttpClient()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyContent))
	if err != nil {
		return err
	}

	kuladoClient.setupRequest(req)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return newApiError(resp, url)
	}

	byteContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(byteContent) > 0 {
		if debug {
			fmt.Println("Response <= " + string(byteContent))
		}
		return json.Unmarshal(byteContent, respObject)
	}

	return nil
}

func (kuladoClient *KuladoBaseClientImpl) Create(schemaType string, createObj interface{}, respObject interface{}) error {
	return kuladoClient.doCreate(schemaType, createObj, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) doCreate(schemaType string, createObj interface{}, respObject interface{}) error {
	if createObj == nil {
		createObj = map[string]string{}
	}
	if respObject == nil {
		respObject = &map[string]interface{}{}
	}
	schema, ok := kuladoClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.CollectionMethods, "POST") {
		return errors.New("Resource type [" + schemaType + "] is not creatable")
	}

	var collectionUrl string
	collectionUrl, ok = schema.Links[COLLECTION]
	if !ok {
		// return errors.New("Failed to find collection URL for [" + schemaType + "]")
		// This is a hack to address https://github.com/kulado/cattle/issues/254
		re := regexp.MustCompile("schemas.*")
		collectionUrl = re.ReplaceAllString(schema.Links[SELF], schema.PluralName)
	}

	return kuladoClient.doModify("POST", collectionUrl, createObj, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) Update(schemaType string, existing *Resource, updates interface{}, respObject interface{}) error {
	return kuladoClient.doUpdate(schemaType, existing, updates, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) doUpdate(schemaType string, existing *Resource, updates interface{}, respObject interface{}) error {
	if existing == nil {
		return errors.New("Existing object is nil")
	}

	selfUrl, ok := existing.Links[SELF]
	if !ok {
		return errors.New(fmt.Sprintf("Failed to find self URL of [%v]", existing))
	}

	if updates == nil {
		updates = map[string]string{}
	}

	if respObject == nil {
		respObject = &map[string]interface{}{}
	}

	schema, ok := kuladoClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.ResourceMethods, "PUT") {
		return errors.New("Resource type [" + schemaType + "] is not updatable")
	}

	return kuladoClient.doModify("PUT", selfUrl, updates, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) ById(schemaType string, id string, respObject interface{}) error {
	return kuladoClient.doById(schemaType, id, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) doById(schemaType string, id string, respObject interface{}) error {
	schema, ok := kuladoClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.ResourceMethods, "GET") {
		return errors.New("Resource type [" + schemaType + "] can not be looked up by ID")
	}

	collectionUrl, ok := schema.Links[COLLECTION]
	if !ok {
		return errors.New("Failed to find collection URL for [" + schemaType + "]")
	}

	err := kuladoClient.doGet(collectionUrl+"/"+id, nil, respObject)
	//TODO check for 404 and return nil, nil
	return err
}

func (kuladoClient *KuladoBaseClientImpl) Delete(existing *Resource) error {
	if existing == nil {
		return nil
	}
	return kuladoClient.doResourceDelete(existing.Type, existing)
}

func (kuladoClient *KuladoBaseClientImpl) doResourceDelete(schemaType string, existing *Resource) error {
	schema, ok := kuladoClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.ResourceMethods, "DELETE") {
		return errors.New("Resource type [" + schemaType + "] can not be deleted")
	}

	selfUrl, ok := existing.Links[SELF]
	if !ok {
		return errors.New(fmt.Sprintf("Failed to find self URL of [%v]", existing))
	}

	return kuladoClient.doDelete(selfUrl)
}

func (kuladoClient *KuladoBaseClientImpl) Reload(existing *Resource, output interface{}) error {
	selfUrl, ok := existing.Links[SELF]
	if !ok {
		return errors.New(fmt.Sprintf("Failed to find self URL of [%v]", existing))
	}

	return kuladoClient.doGet(selfUrl, NewListOpts(), output)
}

func (kuladoClient *KuladoBaseClientImpl) Action(schemaType string, action string,
	existing *Resource, inputObject, respObject interface{}) error {
	return kuladoClient.doAction(schemaType, action, existing, inputObject, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) doAction(schemaType string, action string,
	existing *Resource, inputObject, respObject interface{}) error {

	if existing == nil {
		return errors.New("Existing object is nil")
	}

	actionUrl, ok := existing.Actions[action]
	if !ok {
		return errors.New(fmt.Sprintf("Action [%v] not available on [%v]", action, existing))
	}

	_, ok = kuladoClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	var input io.Reader

	if inputObject != nil {
		bodyContent, err := json.Marshal(inputObject)
		if err != nil {
			return err
		}
		if debug {
			fmt.Println("Request => " + string(bodyContent))
		}
		input = bytes.NewBuffer(bodyContent)
	}

	client := kuladoClient.newHttpClient()
	req, err := http.NewRequest("POST", actionUrl, input)
	if err != nil {
		return err
	}

	kuladoClient.setupRequest(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", "0")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return newApiError(resp, actionUrl)
	}

	byteContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if debug {
		fmt.Println("Response <= " + string(byteContent))
	}

	return json.Unmarshal(byteContent, respObject)
}

func (kuladoClient *KuladoBaseClientImpl) GetOpts() *ClientOpts {
	return kuladoClient.Opts
}

func (kuladoClient *KuladoBaseClientImpl) GetSchemas() *Schemas {
	return kuladoClient.Schemas
}

func (kuladoClient *KuladoBaseClientImpl) GetTypes() map[string]Schema {
	return kuladoClient.Types
}

func init() {
	debug = os.Getenv("RANCHER_CLIENT_DEBUG") == "true"
	if debug {
		fmt.Println("Kulado client debug on")
	}
}
