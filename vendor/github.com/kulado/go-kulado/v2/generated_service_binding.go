package client

const (
	SERVICE_BINDING_TYPE = "serviceBinding"
)

type ServiceBinding struct {
	Resource

	Labels map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`

	Ports []string `json:"ports,omitempty" yaml:"ports,omitempty"`
}

type ServiceBindingCollection struct {
	Collection
	Data   []ServiceBinding `json:"data,omitempty"`
	client *ServiceBindingClient
}

type ServiceBindingClient struct {
	kuladoClient *KuladoClient
}

type ServiceBindingOperations interface {
	List(opts *ListOpts) (*ServiceBindingCollection, error)
	Create(opts *ServiceBinding) (*ServiceBinding, error)
	Update(existing *ServiceBinding, updates interface{}) (*ServiceBinding, error)
	ById(id string) (*ServiceBinding, error)
	Delete(container *ServiceBinding) error
}

func newServiceBindingClient(kuladoClient *KuladoClient) *ServiceBindingClient {
	return &ServiceBindingClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ServiceBindingClient) Create(container *ServiceBinding) (*ServiceBinding, error) {
	resp := &ServiceBinding{}
	err := c.kuladoClient.doCreate(SERVICE_BINDING_TYPE, container, resp)
	return resp, err
}

func (c *ServiceBindingClient) Update(existing *ServiceBinding, updates interface{}) (*ServiceBinding, error) {
	resp := &ServiceBinding{}
	err := c.kuladoClient.doUpdate(SERVICE_BINDING_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceBindingClient) List(opts *ListOpts) (*ServiceBindingCollection, error) {
	resp := &ServiceBindingCollection{}
	err := c.kuladoClient.doList(SERVICE_BINDING_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceBindingCollection) Next() (*ServiceBindingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceBindingCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceBindingClient) ById(id string) (*ServiceBinding, error) {
	resp := &ServiceBinding{}
	err := c.kuladoClient.doById(SERVICE_BINDING_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ServiceBindingClient) Delete(container *ServiceBinding) error {
	return c.kuladoClient.doResourceDelete(SERVICE_BINDING_TYPE, &container.Resource)
}
