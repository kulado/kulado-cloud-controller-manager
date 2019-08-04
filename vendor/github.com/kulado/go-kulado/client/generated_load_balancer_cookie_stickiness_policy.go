package client

const (
	LOAD_BALANCER_COOKIE_STICKINESS_POLICY_TYPE = "loadBalancerCookieStickinessPolicy"
)

type LoadBalancerCookieStickinessPolicy struct {
	Resource

	Cookie string `json:"cookie,omitempty" yaml:"cookie,omitempty"`

	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	Indirect bool `json:"indirect,omitempty" yaml:"indirect,omitempty"`

	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Nocache bool `json:"nocache,omitempty" yaml:"nocache,omitempty"`

	Postonly bool `json:"postonly,omitempty" yaml:"postonly,omitempty"`
}

type LoadBalancerCookieStickinessPolicyCollection struct {
	Collection
	Data   []LoadBalancerCookieStickinessPolicy `json:"data,omitempty"`
	client *LoadBalancerCookieStickinessPolicyClient
}

type LoadBalancerCookieStickinessPolicyClient struct {
	kuladoClient *KuladoClient
}

type LoadBalancerCookieStickinessPolicyOperations interface {
	List(opts *ListOpts) (*LoadBalancerCookieStickinessPolicyCollection, error)
	Create(opts *LoadBalancerCookieStickinessPolicy) (*LoadBalancerCookieStickinessPolicy, error)
	Update(existing *LoadBalancerCookieStickinessPolicy, updates interface{}) (*LoadBalancerCookieStickinessPolicy, error)
	ById(id string) (*LoadBalancerCookieStickinessPolicy, error)
	Delete(container *LoadBalancerCookieStickinessPolicy) error
}

func newLoadBalancerCookieStickinessPolicyClient(kuladoClient *KuladoClient) *LoadBalancerCookieStickinessPolicyClient {
	return &LoadBalancerCookieStickinessPolicyClient{
		kuladoClient: kuladoClient,
	}
}

func (c *LoadBalancerCookieStickinessPolicyClient) Create(container *LoadBalancerCookieStickinessPolicy) (*LoadBalancerCookieStickinessPolicy, error) {
	resp := &LoadBalancerCookieStickinessPolicy{}
	err := c.kuladoClient.doCreate(LOAD_BALANCER_COOKIE_STICKINESS_POLICY_TYPE, container, resp)
	return resp, err
}

func (c *LoadBalancerCookieStickinessPolicyClient) Update(existing *LoadBalancerCookieStickinessPolicy, updates interface{}) (*LoadBalancerCookieStickinessPolicy, error) {
	resp := &LoadBalancerCookieStickinessPolicy{}
	err := c.kuladoClient.doUpdate(LOAD_BALANCER_COOKIE_STICKINESS_POLICY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LoadBalancerCookieStickinessPolicyClient) List(opts *ListOpts) (*LoadBalancerCookieStickinessPolicyCollection, error) {
	resp := &LoadBalancerCookieStickinessPolicyCollection{}
	err := c.kuladoClient.doList(LOAD_BALANCER_COOKIE_STICKINESS_POLICY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LoadBalancerCookieStickinessPolicyCollection) Next() (*LoadBalancerCookieStickinessPolicyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LoadBalancerCookieStickinessPolicyCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LoadBalancerCookieStickinessPolicyClient) ById(id string) (*LoadBalancerCookieStickinessPolicy, error) {
	resp := &LoadBalancerCookieStickinessPolicy{}
	err := c.kuladoClient.doById(LOAD_BALANCER_COOKIE_STICKINESS_POLICY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *LoadBalancerCookieStickinessPolicyClient) Delete(container *LoadBalancerCookieStickinessPolicy) error {
	return c.kuladoClient.doResourceDelete(LOAD_BALANCER_COOKIE_STICKINESS_POLICY_TYPE, &container.Resource)
}
