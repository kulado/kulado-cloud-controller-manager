package client

const (
	SUBSCRIBE_TYPE = "subscribe"
)

type Subscribe struct {
	Resource

	AgentId string `json:"agentId,omitempty" yaml:"agent_id,omitempty"`

	EventNames []string `json:"eventNames,omitempty" yaml:"event_names,omitempty"`
}

type SubscribeCollection struct {
	Collection
	Data   []Subscribe `json:"data,omitempty"`
	client *SubscribeClient
}

type SubscribeClient struct {
	kuladoClient *KuladoClient
}

type SubscribeOperations interface {
	List(opts *ListOpts) (*SubscribeCollection, error)
	Create(opts *Subscribe) (*Subscribe, error)
	Update(existing *Subscribe, updates interface{}) (*Subscribe, error)
	ById(id string) (*Subscribe, error)
	Delete(container *Subscribe) error
}

func newSubscribeClient(kuladoClient *KuladoClient) *SubscribeClient {
	return &SubscribeClient{
		kuladoClient: kuladoClient,
	}
}

func (c *SubscribeClient) Create(container *Subscribe) (*Subscribe, error) {
	resp := &Subscribe{}
	err := c.kuladoClient.doCreate(SUBSCRIBE_TYPE, container, resp)
	return resp, err
}

func (c *SubscribeClient) Update(existing *Subscribe, updates interface{}) (*Subscribe, error) {
	resp := &Subscribe{}
	err := c.kuladoClient.doUpdate(SUBSCRIBE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SubscribeClient) List(opts *ListOpts) (*SubscribeCollection, error) {
	resp := &SubscribeCollection{}
	err := c.kuladoClient.doList(SUBSCRIBE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SubscribeCollection) Next() (*SubscribeCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SubscribeCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SubscribeClient) ById(id string) (*Subscribe, error) {
	resp := &Subscribe{}
	err := c.kuladoClient.doById(SUBSCRIBE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SubscribeClient) Delete(container *Subscribe) error {
	return c.kuladoClient.doResourceDelete(SUBSCRIBE_TYPE, &container.Resource)
}
