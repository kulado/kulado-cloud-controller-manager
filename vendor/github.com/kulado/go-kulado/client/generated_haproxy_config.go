package client

const (
	HAPROXY_CONFIG_TYPE = "haproxyConfig"
)

type HaproxyConfig struct {
	Resource

	Defaults string `json:"defaults,omitempty" yaml:"defaults,omitempty"`

	Global string `json:"global,omitempty" yaml:"global,omitempty"`
}

type HaproxyConfigCollection struct {
	Collection
	Data   []HaproxyConfig `json:"data,omitempty"`
	client *HaproxyConfigClient
}

type HaproxyConfigClient struct {
	kuladoClient *KuladoClient
}

type HaproxyConfigOperations interface {
	List(opts *ListOpts) (*HaproxyConfigCollection, error)
	Create(opts *HaproxyConfig) (*HaproxyConfig, error)
	Update(existing *HaproxyConfig, updates interface{}) (*HaproxyConfig, error)
	ById(id string) (*HaproxyConfig, error)
	Delete(container *HaproxyConfig) error
}

func newHaproxyConfigClient(kuladoClient *KuladoClient) *HaproxyConfigClient {
	return &HaproxyConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *HaproxyConfigClient) Create(container *HaproxyConfig) (*HaproxyConfig, error) {
	resp := &HaproxyConfig{}
	err := c.kuladoClient.doCreate(HAPROXY_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *HaproxyConfigClient) Update(existing *HaproxyConfig, updates interface{}) (*HaproxyConfig, error) {
	resp := &HaproxyConfig{}
	err := c.kuladoClient.doUpdate(HAPROXY_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HaproxyConfigClient) List(opts *ListOpts) (*HaproxyConfigCollection, error) {
	resp := &HaproxyConfigCollection{}
	err := c.kuladoClient.doList(HAPROXY_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HaproxyConfigCollection) Next() (*HaproxyConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HaproxyConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HaproxyConfigClient) ById(id string) (*HaproxyConfig, error) {
	resp := &HaproxyConfig{}
	err := c.kuladoClient.doById(HAPROXY_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *HaproxyConfigClient) Delete(container *HaproxyConfig) error {
	return c.kuladoClient.doResourceDelete(HAPROXY_CONFIG_TYPE, &container.Resource)
}
