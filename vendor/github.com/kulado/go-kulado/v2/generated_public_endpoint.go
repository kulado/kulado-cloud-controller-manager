package client

const (
	PUBLIC_ENDPOINT_TYPE = "publicEndpoint"
)

type PublicEndpoint struct {
	Resource

	HostId string `json:"hostId,omitempty" yaml:"host_id,omitempty"`

	InstanceId string `json:"instanceId,omitempty" yaml:"instance_id,omitempty"`

	IpAddress string `json:"ipAddress,omitempty" yaml:"ip_address,omitempty"`

	Port int64 `json:"port,omitempty" yaml:"port,omitempty"`

	ServiceId string `json:"serviceId,omitempty" yaml:"service_id,omitempty"`
}

type PublicEndpointCollection struct {
	Collection
	Data   []PublicEndpoint `json:"data,omitempty"`
	client *PublicEndpointClient
}

type PublicEndpointClient struct {
	kuladoClient *KuladoClient
}

type PublicEndpointOperations interface {
	List(opts *ListOpts) (*PublicEndpointCollection, error)
	Create(opts *PublicEndpoint) (*PublicEndpoint, error)
	Update(existing *PublicEndpoint, updates interface{}) (*PublicEndpoint, error)
	ById(id string) (*PublicEndpoint, error)
	Delete(container *PublicEndpoint) error
}

func newPublicEndpointClient(kuladoClient *KuladoClient) *PublicEndpointClient {
	return &PublicEndpointClient{
		kuladoClient: kuladoClient,
	}
}

func (c *PublicEndpointClient) Create(container *PublicEndpoint) (*PublicEndpoint, error) {
	resp := &PublicEndpoint{}
	err := c.kuladoClient.doCreate(PUBLIC_ENDPOINT_TYPE, container, resp)
	return resp, err
}

func (c *PublicEndpointClient) Update(existing *PublicEndpoint, updates interface{}) (*PublicEndpoint, error) {
	resp := &PublicEndpoint{}
	err := c.kuladoClient.doUpdate(PUBLIC_ENDPOINT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PublicEndpointClient) List(opts *ListOpts) (*PublicEndpointCollection, error) {
	resp := &PublicEndpointCollection{}
	err := c.kuladoClient.doList(PUBLIC_ENDPOINT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PublicEndpointCollection) Next() (*PublicEndpointCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PublicEndpointCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PublicEndpointClient) ById(id string) (*PublicEndpoint, error) {
	resp := &PublicEndpoint{}
	err := c.kuladoClient.doById(PUBLIC_ENDPOINT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *PublicEndpointClient) Delete(container *PublicEndpoint) error {
	return c.kuladoClient.doResourceDelete(PUBLIC_ENDPOINT_TYPE, &container.Resource)
}
