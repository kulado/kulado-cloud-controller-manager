package client

const (
	LB_TARGET_CONFIG_TYPE = "lbTargetConfig"
)

type LbTargetConfig struct {
	Resource

	PortRules []TargetPortRule `json:"portRules,omitempty" yaml:"port_rules,omitempty"`
}

type LbTargetConfigCollection struct {
	Collection
	Data   []LbTargetConfig `json:"data,omitempty"`
	client *LbTargetConfigClient
}

type LbTargetConfigClient struct {
	kuladoClient *KuladoClient
}

type LbTargetConfigOperations interface {
	List(opts *ListOpts) (*LbTargetConfigCollection, error)
	Create(opts *LbTargetConfig) (*LbTargetConfig, error)
	Update(existing *LbTargetConfig, updates interface{}) (*LbTargetConfig, error)
	ById(id string) (*LbTargetConfig, error)
	Delete(container *LbTargetConfig) error
}

func newLbTargetConfigClient(kuladoClient *KuladoClient) *LbTargetConfigClient {
	return &LbTargetConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *LbTargetConfigClient) Create(container *LbTargetConfig) (*LbTargetConfig, error) {
	resp := &LbTargetConfig{}
	err := c.kuladoClient.doCreate(LB_TARGET_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *LbTargetConfigClient) Update(existing *LbTargetConfig, updates interface{}) (*LbTargetConfig, error) {
	resp := &LbTargetConfig{}
	err := c.kuladoClient.doUpdate(LB_TARGET_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LbTargetConfigClient) List(opts *ListOpts) (*LbTargetConfigCollection, error) {
	resp := &LbTargetConfigCollection{}
	err := c.kuladoClient.doList(LB_TARGET_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LbTargetConfigCollection) Next() (*LbTargetConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LbTargetConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LbTargetConfigClient) ById(id string) (*LbTargetConfig, error) {
	resp := &LbTargetConfig{}
	err := c.kuladoClient.doById(LB_TARGET_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *LbTargetConfigClient) Delete(container *LbTargetConfig) error {
	return c.kuladoClient.doResourceDelete(LB_TARGET_CONFIG_TYPE, &container.Resource)
}
