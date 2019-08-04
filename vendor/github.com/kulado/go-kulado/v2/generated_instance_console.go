package client

const (
	INSTANCE_CONSOLE_TYPE = "instanceConsole"
)

type InstanceConsole struct {
	Resource

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}

type InstanceConsoleCollection struct {
	Collection
	Data   []InstanceConsole `json:"data,omitempty"`
	client *InstanceConsoleClient
}

type InstanceConsoleClient struct {
	kuladoClient *KuladoClient
}

type InstanceConsoleOperations interface {
	List(opts *ListOpts) (*InstanceConsoleCollection, error)
	Create(opts *InstanceConsole) (*InstanceConsole, error)
	Update(existing *InstanceConsole, updates interface{}) (*InstanceConsole, error)
	ById(id string) (*InstanceConsole, error)
	Delete(container *InstanceConsole) error
}

func newInstanceConsoleClient(kuladoClient *KuladoClient) *InstanceConsoleClient {
	return &InstanceConsoleClient{
		kuladoClient: kuladoClient,
	}
}

func (c *InstanceConsoleClient) Create(container *InstanceConsole) (*InstanceConsole, error) {
	resp := &InstanceConsole{}
	err := c.kuladoClient.doCreate(INSTANCE_CONSOLE_TYPE, container, resp)
	return resp, err
}

func (c *InstanceConsoleClient) Update(existing *InstanceConsole, updates interface{}) (*InstanceConsole, error) {
	resp := &InstanceConsole{}
	err := c.kuladoClient.doUpdate(INSTANCE_CONSOLE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InstanceConsoleClient) List(opts *ListOpts) (*InstanceConsoleCollection, error) {
	resp := &InstanceConsoleCollection{}
	err := c.kuladoClient.doList(INSTANCE_CONSOLE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *InstanceConsoleCollection) Next() (*InstanceConsoleCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &InstanceConsoleCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *InstanceConsoleClient) ById(id string) (*InstanceConsole, error) {
	resp := &InstanceConsole{}
	err := c.kuladoClient.doById(INSTANCE_CONSOLE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *InstanceConsoleClient) Delete(container *InstanceConsole) error {
	return c.kuladoClient.doResourceDelete(INSTANCE_CONSOLE_TYPE, &container.Resource)
}
