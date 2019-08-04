package client

const (
	DATABASECHANGELOG_TYPE = "databasechangelog"
)

type Databasechangelog struct {
	Resource

	Author string `json:"author,omitempty" yaml:"author,omitempty"`

	Comments string `json:"comments,omitempty" yaml:"comments,omitempty"`

	Dateexecuted string `json:"dateexecuted,omitempty" yaml:"dateexecuted,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Exectype string `json:"exectype,omitempty" yaml:"exectype,omitempty"`

	Filename string `json:"filename,omitempty" yaml:"filename,omitempty"`

	Liquibase string `json:"liquibase,omitempty" yaml:"liquibase,omitempty"`

	Md5sum string `json:"md5sum,omitempty" yaml:"md5sum,omitempty"`

	Orderexecuted int64 `json:"orderexecuted,omitempty" yaml:"orderexecuted,omitempty"`

	Tag string `json:"tag,omitempty" yaml:"tag,omitempty"`
}

type DatabasechangelogCollection struct {
	Collection
	Data   []Databasechangelog `json:"data,omitempty"`
	client *DatabasechangelogClient
}

type DatabasechangelogClient struct {
	kuladoClient *KuladoClient
}

type DatabasechangelogOperations interface {
	List(opts *ListOpts) (*DatabasechangelogCollection, error)
	Create(opts *Databasechangelog) (*Databasechangelog, error)
	Update(existing *Databasechangelog, updates interface{}) (*Databasechangelog, error)
	ById(id string) (*Databasechangelog, error)
	Delete(container *Databasechangelog) error
}

func newDatabasechangelogClient(kuladoClient *KuladoClient) *DatabasechangelogClient {
	return &DatabasechangelogClient{
		kuladoClient: kuladoClient,
	}
}

func (c *DatabasechangelogClient) Create(container *Databasechangelog) (*Databasechangelog, error) {
	resp := &Databasechangelog{}
	err := c.kuladoClient.doCreate(DATABASECHANGELOG_TYPE, container, resp)
	return resp, err
}

func (c *DatabasechangelogClient) Update(existing *Databasechangelog, updates interface{}) (*Databasechangelog, error) {
	resp := &Databasechangelog{}
	err := c.kuladoClient.doUpdate(DATABASECHANGELOG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DatabasechangelogClient) List(opts *ListOpts) (*DatabasechangelogCollection, error) {
	resp := &DatabasechangelogCollection{}
	err := c.kuladoClient.doList(DATABASECHANGELOG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *DatabasechangelogCollection) Next() (*DatabasechangelogCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &DatabasechangelogCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *DatabasechangelogClient) ById(id string) (*Databasechangelog, error) {
	resp := &Databasechangelog{}
	err := c.kuladoClient.doById(DATABASECHANGELOG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *DatabasechangelogClient) Delete(container *Databasechangelog) error {
	return c.kuladoClient.doResourceDelete(DATABASECHANGELOG_TYPE, &container.Resource)
}
