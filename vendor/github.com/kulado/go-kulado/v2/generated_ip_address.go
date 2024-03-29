package client

const (
	IP_ADDRESS_TYPE = "ipAddress"
)

type IpAddress struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	NetworkId string `json:"networkId,omitempty" yaml:"network_id,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type IpAddressCollection struct {
	Collection
	Data   []IpAddress `json:"data,omitempty"`
	client *IpAddressClient
}

type IpAddressClient struct {
	kuladoClient *KuladoClient
}

type IpAddressOperations interface {
	List(opts *ListOpts) (*IpAddressCollection, error)
	Create(opts *IpAddress) (*IpAddress, error)
	Update(existing *IpAddress, updates interface{}) (*IpAddress, error)
	ById(id string) (*IpAddress, error)
	Delete(container *IpAddress) error

	ActionActivate(*IpAddress) (*IpAddress, error)

	ActionCreate(*IpAddress) (*IpAddress, error)

	ActionDeactivate(*IpAddress) (*IpAddress, error)

	ActionDisassociate(*IpAddress) (*IpAddress, error)

	ActionPurge(*IpAddress) (*IpAddress, error)

	ActionRemove(*IpAddress) (*IpAddress, error)

	ActionRestore(*IpAddress) (*IpAddress, error)

	ActionUpdate(*IpAddress) (*IpAddress, error)
}

func newIpAddressClient(kuladoClient *KuladoClient) *IpAddressClient {
	return &IpAddressClient{
		kuladoClient: kuladoClient,
	}
}

func (c *IpAddressClient) Create(container *IpAddress) (*IpAddress, error) {
	resp := &IpAddress{}
	err := c.kuladoClient.doCreate(IP_ADDRESS_TYPE, container, resp)
	return resp, err
}

func (c *IpAddressClient) Update(existing *IpAddress, updates interface{}) (*IpAddress, error) {
	resp := &IpAddress{}
	err := c.kuladoClient.doUpdate(IP_ADDRESS_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *IpAddressClient) List(opts *ListOpts) (*IpAddressCollection, error) {
	resp := &IpAddressCollection{}
	err := c.kuladoClient.doList(IP_ADDRESS_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *IpAddressCollection) Next() (*IpAddressCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &IpAddressCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *IpAddressClient) ById(id string) (*IpAddress, error) {
	resp := &IpAddress{}
	err := c.kuladoClient.doById(IP_ADDRESS_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *IpAddressClient) Delete(container *IpAddress) error {
	return c.kuladoClient.doResourceDelete(IP_ADDRESS_TYPE, &container.Resource)
}

func (c *IpAddressClient) ActionActivate(resource *IpAddress) (*IpAddress, error) {

	resp := &IpAddress{}

	err := c.kuladoClient.doAction(IP_ADDRESS_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *IpAddressClient) ActionCreate(resource *IpAddress) (*IpAddress, error) {

	resp := &IpAddress{}

	err := c.kuladoClient.doAction(IP_ADDRESS_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *IpAddressClient) ActionDeactivate(resource *IpAddress) (*IpAddress, error) {

	resp := &IpAddress{}

	err := c.kuladoClient.doAction(IP_ADDRESS_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *IpAddressClient) ActionDisassociate(resource *IpAddress) (*IpAddress, error) {

	resp := &IpAddress{}

	err := c.kuladoClient.doAction(IP_ADDRESS_TYPE, "disassociate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *IpAddressClient) ActionPurge(resource *IpAddress) (*IpAddress, error) {

	resp := &IpAddress{}

	err := c.kuladoClient.doAction(IP_ADDRESS_TYPE, "purge", &resource.Resource, nil, resp)

	return resp, err
}

func (c *IpAddressClient) ActionRemove(resource *IpAddress) (*IpAddress, error) {

	resp := &IpAddress{}

	err := c.kuladoClient.doAction(IP_ADDRESS_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *IpAddressClient) ActionRestore(resource *IpAddress) (*IpAddress, error) {

	resp := &IpAddress{}

	err := c.kuladoClient.doAction(IP_ADDRESS_TYPE, "restore", &resource.Resource, nil, resp)

	return resp, err
}

func (c *IpAddressClient) ActionUpdate(resource *IpAddress) (*IpAddress, error) {

	resp := &IpAddress{}

	err := c.kuladoClient.doAction(IP_ADDRESS_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
