package client

const (
	SET_LOAD_BALANCER_SERVICE_LINKS_INPUT_TYPE = "setLoadBalancerServiceLinksInput"
)

type SetLoadBalancerServiceLinksInput struct {
	Resource

	ServiceLinks []interface{} `json:"serviceLinks,omitempty" yaml:"service_links,omitempty"`
}

type SetLoadBalancerServiceLinksInputCollection struct {
	Collection
	Data   []SetLoadBalancerServiceLinksInput `json:"data,omitempty"`
	client *SetLoadBalancerServiceLinksInputClient
}

type SetLoadBalancerServiceLinksInputClient struct {
	kuladoClient *KuladoClient
}

type SetLoadBalancerServiceLinksInputOperations interface {
	List(opts *ListOpts) (*SetLoadBalancerServiceLinksInputCollection, error)
	Create(opts *SetLoadBalancerServiceLinksInput) (*SetLoadBalancerServiceLinksInput, error)
	Update(existing *SetLoadBalancerServiceLinksInput, updates interface{}) (*SetLoadBalancerServiceLinksInput, error)
	ById(id string) (*SetLoadBalancerServiceLinksInput, error)
	Delete(container *SetLoadBalancerServiceLinksInput) error
}

func newSetLoadBalancerServiceLinksInputClient(kuladoClient *KuladoClient) *SetLoadBalancerServiceLinksInputClient {
	return &SetLoadBalancerServiceLinksInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *SetLoadBalancerServiceLinksInputClient) Create(container *SetLoadBalancerServiceLinksInput) (*SetLoadBalancerServiceLinksInput, error) {
	resp := &SetLoadBalancerServiceLinksInput{}
	err := c.kuladoClient.doCreate(SET_LOAD_BALANCER_SERVICE_LINKS_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *SetLoadBalancerServiceLinksInputClient) Update(existing *SetLoadBalancerServiceLinksInput, updates interface{}) (*SetLoadBalancerServiceLinksInput, error) {
	resp := &SetLoadBalancerServiceLinksInput{}
	err := c.kuladoClient.doUpdate(SET_LOAD_BALANCER_SERVICE_LINKS_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SetLoadBalancerServiceLinksInputClient) List(opts *ListOpts) (*SetLoadBalancerServiceLinksInputCollection, error) {
	resp := &SetLoadBalancerServiceLinksInputCollection{}
	err := c.kuladoClient.doList(SET_LOAD_BALANCER_SERVICE_LINKS_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SetLoadBalancerServiceLinksInputCollection) Next() (*SetLoadBalancerServiceLinksInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SetLoadBalancerServiceLinksInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SetLoadBalancerServiceLinksInputClient) ById(id string) (*SetLoadBalancerServiceLinksInput, error) {
	resp := &SetLoadBalancerServiceLinksInput{}
	err := c.kuladoClient.doById(SET_LOAD_BALANCER_SERVICE_LINKS_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SetLoadBalancerServiceLinksInputClient) Delete(container *SetLoadBalancerServiceLinksInput) error {
	return c.kuladoClient.doResourceDelete(SET_LOAD_BALANCER_SERVICE_LINKS_INPUT_TYPE, &container.Resource)
}
