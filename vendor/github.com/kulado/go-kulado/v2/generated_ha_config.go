package client

const (
	HA_CONFIG_TYPE = "haConfig"
)

type HaConfig struct {
	Resource

	ClusterSize int64 `json:"clusterSize,omitempty" yaml:"cluster_size,omitempty"`

	DbHost string `json:"dbHost,omitempty" yaml:"db_host,omitempty"`

	DbSize int64 `json:"dbSize,omitempty" yaml:"db_size,omitempty"`

	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type HaConfigCollection struct {
	Collection
	Data   []HaConfig `json:"data,omitempty"`
	client *HaConfigClient
}

type HaConfigClient struct {
	kuladoClient *KuladoClient
}

type HaConfigOperations interface {
	List(opts *ListOpts) (*HaConfigCollection, error)
	Create(opts *HaConfig) (*HaConfig, error)
	Update(existing *HaConfig, updates interface{}) (*HaConfig, error)
	ById(id string) (*HaConfig, error)
	Delete(container *HaConfig) error
}

func newHaConfigClient(kuladoClient *KuladoClient) *HaConfigClient {
	return &HaConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *HaConfigClient) Create(container *HaConfig) (*HaConfig, error) {
	resp := &HaConfig{}
	err := c.kuladoClient.doCreate(HA_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *HaConfigClient) Update(existing *HaConfig, updates interface{}) (*HaConfig, error) {
	resp := &HaConfig{}
	err := c.kuladoClient.doUpdate(HA_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HaConfigClient) List(opts *ListOpts) (*HaConfigCollection, error) {
	resp := &HaConfigCollection{}
	err := c.kuladoClient.doList(HA_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HaConfigCollection) Next() (*HaConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HaConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HaConfigClient) ById(id string) (*HaConfig, error) {
	resp := &HaConfig{}
	err := c.kuladoClient.doById(HA_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *HaConfigClient) Delete(container *HaConfig) error {
	return c.kuladoClient.doResourceDelete(HA_CONFIG_TYPE, &container.Resource)
}
