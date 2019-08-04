package client

const (
	SERVICES_PORT_RANGE_TYPE = "servicesPortRange"
)

type ServicesPortRange struct {
	Resource

	EndPort int64 `json:"endPort,omitempty" yaml:"end_port,omitempty"`

	StartPort int64 `json:"startPort,omitempty" yaml:"start_port,omitempty"`
}

type ServicesPortRangeCollection struct {
	Collection
	Data   []ServicesPortRange `json:"data,omitempty"`
	client *ServicesPortRangeClient
}

type ServicesPortRangeClient struct {
	kuladoClient *KuladoClient
}

type ServicesPortRangeOperations interface {
	List(opts *ListOpts) (*ServicesPortRangeCollection, error)
	Create(opts *ServicesPortRange) (*ServicesPortRange, error)
	Update(existing *ServicesPortRange, updates interface{}) (*ServicesPortRange, error)
	ById(id string) (*ServicesPortRange, error)
	Delete(container *ServicesPortRange) error
}

func newServicesPortRangeClient(kuladoClient *KuladoClient) *ServicesPortRangeClient {
	return &ServicesPortRangeClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ServicesPortRangeClient) Create(container *ServicesPortRange) (*ServicesPortRange, error) {
	resp := &ServicesPortRange{}
	err := c.kuladoClient.doCreate(SERVICES_PORT_RANGE_TYPE, container, resp)
	return resp, err
}

func (c *ServicesPortRangeClient) Update(existing *ServicesPortRange, updates interface{}) (*ServicesPortRange, error) {
	resp := &ServicesPortRange{}
	err := c.kuladoClient.doUpdate(SERVICES_PORT_RANGE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServicesPortRangeClient) List(opts *ListOpts) (*ServicesPortRangeCollection, error) {
	resp := &ServicesPortRangeCollection{}
	err := c.kuladoClient.doList(SERVICES_PORT_RANGE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServicesPortRangeCollection) Next() (*ServicesPortRangeCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServicesPortRangeCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServicesPortRangeClient) ById(id string) (*ServicesPortRange, error) {
	resp := &ServicesPortRange{}
	err := c.kuladoClient.doById(SERVICES_PORT_RANGE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ServicesPortRangeClient) Delete(container *ServicesPortRange) error {
	return c.kuladoClient.doResourceDelete(SERVICES_PORT_RANGE_TYPE, &container.Resource)
}
