package client

const (
	SERVICE_CONSUME_MAP_TYPE = "serviceConsumeMap"
)

type ServiceConsumeMap struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	ConsumedServiceId string `json:"consumedServiceId,omitempty" yaml:"consumed_service_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Ports []string `json:"ports,omitempty" yaml:"ports,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	ServiceId string `json:"serviceId,omitempty" yaml:"service_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ServiceConsumeMapCollection struct {
	Collection
	Data   []ServiceConsumeMap `json:"data,omitempty"`
	client *ServiceConsumeMapClient
}

type ServiceConsumeMapClient struct {
	kuladoClient *KuladoClient
}

type ServiceConsumeMapOperations interface {
	List(opts *ListOpts) (*ServiceConsumeMapCollection, error)
	Create(opts *ServiceConsumeMap) (*ServiceConsumeMap, error)
	Update(existing *ServiceConsumeMap, updates interface{}) (*ServiceConsumeMap, error)
	ById(id string) (*ServiceConsumeMap, error)
	Delete(container *ServiceConsumeMap) error

	ActionCreate(*ServiceConsumeMap) (*ServiceConsumeMap, error)

	ActionRemove(*ServiceConsumeMap) (*ServiceConsumeMap, error)

	ActionUpdate(*ServiceConsumeMap) (*ServiceConsumeMap, error)
}

func newServiceConsumeMapClient(kuladoClient *KuladoClient) *ServiceConsumeMapClient {
	return &ServiceConsumeMapClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ServiceConsumeMapClient) Create(container *ServiceConsumeMap) (*ServiceConsumeMap, error) {
	resp := &ServiceConsumeMap{}
	err := c.kuladoClient.doCreate(SERVICE_CONSUME_MAP_TYPE, container, resp)
	return resp, err
}

func (c *ServiceConsumeMapClient) Update(existing *ServiceConsumeMap, updates interface{}) (*ServiceConsumeMap, error) {
	resp := &ServiceConsumeMap{}
	err := c.kuladoClient.doUpdate(SERVICE_CONSUME_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceConsumeMapClient) List(opts *ListOpts) (*ServiceConsumeMapCollection, error) {
	resp := &ServiceConsumeMapCollection{}
	err := c.kuladoClient.doList(SERVICE_CONSUME_MAP_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceConsumeMapCollection) Next() (*ServiceConsumeMapCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceConsumeMapCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceConsumeMapClient) ById(id string) (*ServiceConsumeMap, error) {
	resp := &ServiceConsumeMap{}
	err := c.kuladoClient.doById(SERVICE_CONSUME_MAP_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ServiceConsumeMapClient) Delete(container *ServiceConsumeMap) error {
	return c.kuladoClient.doResourceDelete(SERVICE_CONSUME_MAP_TYPE, &container.Resource)
}

func (c *ServiceConsumeMapClient) ActionCreate(resource *ServiceConsumeMap) (*ServiceConsumeMap, error) {

	resp := &ServiceConsumeMap{}

	err := c.kuladoClient.doAction(SERVICE_CONSUME_MAP_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ServiceConsumeMapClient) ActionRemove(resource *ServiceConsumeMap) (*ServiceConsumeMap, error) {

	resp := &ServiceConsumeMap{}

	err := c.kuladoClient.doAction(SERVICE_CONSUME_MAP_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ServiceConsumeMapClient) ActionUpdate(resource *ServiceConsumeMap) (*ServiceConsumeMap, error) {

	resp := &ServiceConsumeMap{}

	err := c.kuladoClient.doAction(SERVICE_CONSUME_MAP_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
