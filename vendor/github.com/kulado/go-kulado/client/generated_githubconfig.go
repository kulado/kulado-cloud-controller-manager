package client

const (
	GITHUBCONFIG_TYPE = "githubconfig"
)

type Githubconfig struct {
	Resource

	AccessMode string `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`

	AllowedIdentities []interface{} `json:"allowedIdentities,omitempty" yaml:"allowed_identities,omitempty"`

	ClientId string `json:"clientId,omitempty" yaml:"client_id,omitempty"`

	ClientSecret string `json:"clientSecret,omitempty" yaml:"client_secret,omitempty"`

	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

type GithubconfigCollection struct {
	Collection
	Data   []Githubconfig `json:"data,omitempty"`
	client *GithubconfigClient
}

type GithubconfigClient struct {
	kuladoClient *KuladoClient
}

type GithubconfigOperations interface {
	List(opts *ListOpts) (*GithubconfigCollection, error)
	Create(opts *Githubconfig) (*Githubconfig, error)
	Update(existing *Githubconfig, updates interface{}) (*Githubconfig, error)
	ById(id string) (*Githubconfig, error)
	Delete(container *Githubconfig) error
}

func newGithubconfigClient(kuladoClient *KuladoClient) *GithubconfigClient {
	return &GithubconfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *GithubconfigClient) Create(container *Githubconfig) (*Githubconfig, error) {
	resp := &Githubconfig{}
	err := c.kuladoClient.doCreate(GITHUBCONFIG_TYPE, container, resp)
	return resp, err
}

func (c *GithubconfigClient) Update(existing *Githubconfig, updates interface{}) (*Githubconfig, error) {
	resp := &Githubconfig{}
	err := c.kuladoClient.doUpdate(GITHUBCONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GithubconfigClient) List(opts *ListOpts) (*GithubconfigCollection, error) {
	resp := &GithubconfigCollection{}
	err := c.kuladoClient.doList(GITHUBCONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *GithubconfigCollection) Next() (*GithubconfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &GithubconfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *GithubconfigClient) ById(id string) (*Githubconfig, error) {
	resp := &Githubconfig{}
	err := c.kuladoClient.doById(GITHUBCONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *GithubconfigClient) Delete(container *Githubconfig) error {
	return c.kuladoClient.doResourceDelete(GITHUBCONFIG_TYPE, &container.Resource)
}
