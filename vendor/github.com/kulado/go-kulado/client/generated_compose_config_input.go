package client

const (
	COMPOSE_CONFIG_INPUT_TYPE = "composeConfigInput"
)

type ComposeConfigInput struct {
	Resource

	ServiceIds []string `json:"serviceIds,omitempty" yaml:"service_ids,omitempty"`
}

type ComposeConfigInputCollection struct {
	Collection
	Data   []ComposeConfigInput `json:"data,omitempty"`
	client *ComposeConfigInputClient
}

type ComposeConfigInputClient struct {
	kuladoClient *KuladoClient
}

type ComposeConfigInputOperations interface {
	List(opts *ListOpts) (*ComposeConfigInputCollection, error)
	Create(opts *ComposeConfigInput) (*ComposeConfigInput, error)
	Update(existing *ComposeConfigInput, updates interface{}) (*ComposeConfigInput, error)
	ById(id string) (*ComposeConfigInput, error)
	Delete(container *ComposeConfigInput) error
}

func newComposeConfigInputClient(kuladoClient *KuladoClient) *ComposeConfigInputClient {
	return &ComposeConfigInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ComposeConfigInputClient) Create(container *ComposeConfigInput) (*ComposeConfigInput, error) {
	resp := &ComposeConfigInput{}
	err := c.kuladoClient.doCreate(COMPOSE_CONFIG_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *ComposeConfigInputClient) Update(existing *ComposeConfigInput, updates interface{}) (*ComposeConfigInput, error) {
	resp := &ComposeConfigInput{}
	err := c.kuladoClient.doUpdate(COMPOSE_CONFIG_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ComposeConfigInputClient) List(opts *ListOpts) (*ComposeConfigInputCollection, error) {
	resp := &ComposeConfigInputCollection{}
	err := c.kuladoClient.doList(COMPOSE_CONFIG_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ComposeConfigInputCollection) Next() (*ComposeConfigInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ComposeConfigInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ComposeConfigInputClient) ById(id string) (*ComposeConfigInput, error) {
	resp := &ComposeConfigInput{}
	err := c.kuladoClient.doById(COMPOSE_CONFIG_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ComposeConfigInputClient) Delete(container *ComposeConfigInput) error {
	return c.kuladoClient.doResourceDelete(COMPOSE_CONFIG_INPUT_TYPE, &container.Resource)
}
