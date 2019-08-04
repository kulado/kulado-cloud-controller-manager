package client

const (
	LOCAL_AUTH_CONFIG_TYPE = "localAuthConfig"
)

type LocalAuthConfig struct {
	Resource

	AccessMode string `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`

	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

type LocalAuthConfigCollection struct {
	Collection
	Data   []LocalAuthConfig `json:"data,omitempty"`
	client *LocalAuthConfigClient
}

type LocalAuthConfigClient struct {
	kuladoClient *KuladoClient
}

type LocalAuthConfigOperations interface {
	List(opts *ListOpts) (*LocalAuthConfigCollection, error)
	Create(opts *LocalAuthConfig) (*LocalAuthConfig, error)
	Update(existing *LocalAuthConfig, updates interface{}) (*LocalAuthConfig, error)
	ById(id string) (*LocalAuthConfig, error)
	Delete(container *LocalAuthConfig) error
}

func newLocalAuthConfigClient(kuladoClient *KuladoClient) *LocalAuthConfigClient {
	return &LocalAuthConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *LocalAuthConfigClient) Create(container *LocalAuthConfig) (*LocalAuthConfig, error) {
	resp := &LocalAuthConfig{}
	err := c.kuladoClient.doCreate(LOCAL_AUTH_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *LocalAuthConfigClient) Update(existing *LocalAuthConfig, updates interface{}) (*LocalAuthConfig, error) {
	resp := &LocalAuthConfig{}
	err := c.kuladoClient.doUpdate(LOCAL_AUTH_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LocalAuthConfigClient) List(opts *ListOpts) (*LocalAuthConfigCollection, error) {
	resp := &LocalAuthConfigCollection{}
	err := c.kuladoClient.doList(LOCAL_AUTH_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LocalAuthConfigCollection) Next() (*LocalAuthConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LocalAuthConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LocalAuthConfigClient) ById(id string) (*LocalAuthConfig, error) {
	resp := &LocalAuthConfig{}
	err := c.kuladoClient.doById(LOCAL_AUTH_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *LocalAuthConfigClient) Delete(container *LocalAuthConfig) error {
	return c.kuladoClient.doResourceDelete(LOCAL_AUTH_CONFIG_TYPE, &container.Resource)
}
