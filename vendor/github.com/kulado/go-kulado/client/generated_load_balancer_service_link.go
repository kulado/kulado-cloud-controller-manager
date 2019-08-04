package client

const (
	LOAD_BALANCER_SERVICE_LINK_TYPE = "loadBalancerServiceLink"
)

type LoadBalancerServiceLink struct {
	Resource

	Ports []string `json:"ports,omitempty" yaml:"ports,omitempty"`

	ServiceId string `json:"serviceId,omitempty" yaml:"service_id,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type LoadBalancerServiceLinkCollection struct {
	Collection
	Data   []LoadBalancerServiceLink `json:"data,omitempty"`
	client *LoadBalancerServiceLinkClient
}

type LoadBalancerServiceLinkClient struct {
	kuladoClient *KuladoClient
}

type LoadBalancerServiceLinkOperations interface {
	List(opts *ListOpts) (*LoadBalancerServiceLinkCollection, error)
	Create(opts *LoadBalancerServiceLink) (*LoadBalancerServiceLink, error)
	Update(existing *LoadBalancerServiceLink, updates interface{}) (*LoadBalancerServiceLink, error)
	ById(id string) (*LoadBalancerServiceLink, error)
	Delete(container *LoadBalancerServiceLink) error
}

func newLoadBalancerServiceLinkClient(kuladoClient *KuladoClient) *LoadBalancerServiceLinkClient {
	return &LoadBalancerServiceLinkClient{
		kuladoClient: kuladoClient,
	}
}

func (c *LoadBalancerServiceLinkClient) Create(container *LoadBalancerServiceLink) (*LoadBalancerServiceLink, error) {
	resp := &LoadBalancerServiceLink{}
	err := c.kuladoClient.doCreate(LOAD_BALANCER_SERVICE_LINK_TYPE, container, resp)
	return resp, err
}

func (c *LoadBalancerServiceLinkClient) Update(existing *LoadBalancerServiceLink, updates interface{}) (*LoadBalancerServiceLink, error) {
	resp := &LoadBalancerServiceLink{}
	err := c.kuladoClient.doUpdate(LOAD_BALANCER_SERVICE_LINK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LoadBalancerServiceLinkClient) List(opts *ListOpts) (*LoadBalancerServiceLinkCollection, error) {
	resp := &LoadBalancerServiceLinkCollection{}
	err := c.kuladoClient.doList(LOAD_BALANCER_SERVICE_LINK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LoadBalancerServiceLinkCollection) Next() (*LoadBalancerServiceLinkCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LoadBalancerServiceLinkCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LoadBalancerServiceLinkClient) ById(id string) (*LoadBalancerServiceLink, error) {
	resp := &LoadBalancerServiceLink{}
	err := c.kuladoClient.doById(LOAD_BALANCER_SERVICE_LINK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *LoadBalancerServiceLinkClient) Delete(container *LoadBalancerServiceLink) error {
	return c.kuladoClient.doResourceDelete(LOAD_BALANCER_SERVICE_LINK_TYPE, &container.Resource)
}
