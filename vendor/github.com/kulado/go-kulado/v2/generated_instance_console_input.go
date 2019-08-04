package client

const (
	INSTANCE_CONSOLE_INPUT_TYPE = "instanceConsoleInput"
)

type InstanceConsoleInput struct {
	Resource
}

type InstanceConsoleInputCollection struct {
	Collection
	Data   []InstanceConsoleInput `json:"data,omitempty"`
	client *InstanceConsoleInputClient
}

type InstanceConsoleInputClient struct {
	kuladoClient *KuladoClient
}

type InstanceConsoleInputOperations interface {
	List(opts *ListOpts) (*InstanceConsoleInputCollection, error)
	Create(opts *InstanceConsoleInput) (*InstanceConsoleInput, error)
	Update(existing *InstanceConsoleInput, updates interface{}) (*InstanceConsoleInput, error)
	ById(id string) (*InstanceConsoleInput, error)
	Delete(container *InstanceConsoleInput) error
}

func newInstanceConsoleInputClient(kuladoClient *KuladoClient) *InstanceConsoleInputClient {
	return &InstanceConsoleInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *InstanceConsoleInputClient) Create(container *InstanceConsoleInput) (*InstanceConsoleInput, error) {
	resp := &InstanceConsoleInput{}
	err := c.kuladoClient.doCreate(INSTANCE_CONSOLE_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *InstanceConsoleInputClient) Update(existing *InstanceConsoleInput, updates interface{}) (*InstanceConsoleInput, error) {
	resp := &InstanceConsoleInput{}
	err := c.kuladoClient.doUpdate(INSTANCE_CONSOLE_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InstanceConsoleInputClient) List(opts *ListOpts) (*InstanceConsoleInputCollection, error) {
	resp := &InstanceConsoleInputCollection{}
	err := c.kuladoClient.doList(INSTANCE_CONSOLE_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *InstanceConsoleInputCollection) Next() (*InstanceConsoleInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &InstanceConsoleInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *InstanceConsoleInputClient) ById(id string) (*InstanceConsoleInput, error) {
	resp := &InstanceConsoleInput{}
	err := c.kuladoClient.doById(INSTANCE_CONSOLE_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *InstanceConsoleInputClient) Delete(container *InstanceConsoleInput) error {
	return c.kuladoClient.doResourceDelete(INSTANCE_CONSOLE_INPUT_TYPE, &container.Resource)
}
