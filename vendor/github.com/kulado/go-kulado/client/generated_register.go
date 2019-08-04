package client

const (
	REGISTER_TYPE = "register"
)

type Register struct {
	Resource

	AccessKey string `json:"accessKey,omitempty" yaml:"access_key,omitempty"`

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	SecretKey string `json:"secretKey,omitempty" yaml:"secret_key,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type RegisterCollection struct {
	Collection
	Data   []Register `json:"data,omitempty"`
	client *RegisterClient
}

type RegisterClient struct {
	kuladoClient *KuladoClient
}

type RegisterOperations interface {
	List(opts *ListOpts) (*RegisterCollection, error)
	Create(opts *Register) (*Register, error)
	Update(existing *Register, updates interface{}) (*Register, error)
	ById(id string) (*Register, error)
	Delete(container *Register) error

	ActionStop(*Register, *InstanceStop) (*Instance, error)
}

func newRegisterClient(kuladoClient *KuladoClient) *RegisterClient {
	return &RegisterClient{
		kuladoClient: kuladoClient,
	}
}

func (c *RegisterClient) Create(container *Register) (*Register, error) {
	resp := &Register{}
	err := c.kuladoClient.doCreate(REGISTER_TYPE, container, resp)
	return resp, err
}

func (c *RegisterClient) Update(existing *Register, updates interface{}) (*Register, error) {
	resp := &Register{}
	err := c.kuladoClient.doUpdate(REGISTER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RegisterClient) List(opts *ListOpts) (*RegisterCollection, error) {
	resp := &RegisterCollection{}
	err := c.kuladoClient.doList(REGISTER_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *RegisterCollection) Next() (*RegisterCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &RegisterCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *RegisterClient) ById(id string) (*Register, error) {
	resp := &Register{}
	err := c.kuladoClient.doById(REGISTER_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *RegisterClient) Delete(container *Register) error {
	return c.kuladoClient.doResourceDelete(REGISTER_TYPE, &container.Resource)
}

func (c *RegisterClient) ActionStop(resource *Register, input *InstanceStop) (*Instance, error) {

	resp := &Instance{}

	err := c.kuladoClient.doAction(REGISTER_TYPE, "stop", &resource.Resource, input, resp)

	return resp, err
}
