package client

const (
	SERVICE_RESTART_TYPE = "serviceRestart"
)

type ServiceRestart struct {
	Resource

	RollingRestartStrategy RollingRestartStrategy `json:"rollingRestartStrategy,omitempty" yaml:"rolling_restart_strategy,omitempty"`
}

type ServiceRestartCollection struct {
	Collection
	Data   []ServiceRestart `json:"data,omitempty"`
	client *ServiceRestartClient
}

type ServiceRestartClient struct {
	kuladoClient *KuladoClient
}

type ServiceRestartOperations interface {
	List(opts *ListOpts) (*ServiceRestartCollection, error)
	Create(opts *ServiceRestart) (*ServiceRestart, error)
	Update(existing *ServiceRestart, updates interface{}) (*ServiceRestart, error)
	ById(id string) (*ServiceRestart, error)
	Delete(container *ServiceRestart) error
}

func newServiceRestartClient(kuladoClient *KuladoClient) *ServiceRestartClient {
	return &ServiceRestartClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ServiceRestartClient) Create(container *ServiceRestart) (*ServiceRestart, error) {
	resp := &ServiceRestart{}
	err := c.kuladoClient.doCreate(SERVICE_RESTART_TYPE, container, resp)
	return resp, err
}

func (c *ServiceRestartClient) Update(existing *ServiceRestart, updates interface{}) (*ServiceRestart, error) {
	resp := &ServiceRestart{}
	err := c.kuladoClient.doUpdate(SERVICE_RESTART_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceRestartClient) List(opts *ListOpts) (*ServiceRestartCollection, error) {
	resp := &ServiceRestartCollection{}
	err := c.kuladoClient.doList(SERVICE_RESTART_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceRestartCollection) Next() (*ServiceRestartCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceRestartCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceRestartClient) ById(id string) (*ServiceRestart, error) {
	resp := &ServiceRestart{}
	err := c.kuladoClient.doById(SERVICE_RESTART_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ServiceRestartClient) Delete(container *ServiceRestart) error {
	return c.kuladoClient.doResourceDelete(SERVICE_RESTART_TYPE, &container.Resource)
}
