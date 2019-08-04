package client

const (
	LB_CONFIG_TYPE = "lbConfig"
)

type LbConfig struct {
	Resource

	CertificateIds []string `json:"certificateIds,omitempty" yaml:"certificate_ids,omitempty"`

	Config string `json:"config,omitempty" yaml:"config,omitempty"`

	DefaultCertificateId string `json:"defaultCertificateId,omitempty" yaml:"default_certificate_id,omitempty"`

	PortRules []PortRule `json:"portRules,omitempty" yaml:"port_rules,omitempty"`

	StickinessPolicy *LoadBalancerCookieStickinessPolicy `json:"stickinessPolicy,omitempty" yaml:"stickiness_policy,omitempty"`
}

type LbConfigCollection struct {
	Collection
	Data   []LbConfig `json:"data,omitempty"`
	client *LbConfigClient
}

type LbConfigClient struct {
	kuladoClient *KuladoClient
}

type LbConfigOperations interface {
	List(opts *ListOpts) (*LbConfigCollection, error)
	Create(opts *LbConfig) (*LbConfig, error)
	Update(existing *LbConfig, updates interface{}) (*LbConfig, error)
	ById(id string) (*LbConfig, error)
	Delete(container *LbConfig) error
}

func newLbConfigClient(kuladoClient *KuladoClient) *LbConfigClient {
	return &LbConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *LbConfigClient) Create(container *LbConfig) (*LbConfig, error) {
	resp := &LbConfig{}
	err := c.kuladoClient.doCreate(LB_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *LbConfigClient) Update(existing *LbConfig, updates interface{}) (*LbConfig, error) {
	resp := &LbConfig{}
	err := c.kuladoClient.doUpdate(LB_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LbConfigClient) List(opts *ListOpts) (*LbConfigCollection, error) {
	resp := &LbConfigCollection{}
	err := c.kuladoClient.doList(LB_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LbConfigCollection) Next() (*LbConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LbConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LbConfigClient) ById(id string) (*LbConfig, error) {
	resp := &LbConfig{}
	err := c.kuladoClient.doById(LB_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *LbConfigClient) Delete(container *LbConfig) error {
	return c.kuladoClient.doResourceDelete(LB_CONFIG_TYPE, &container.Resource)
}
