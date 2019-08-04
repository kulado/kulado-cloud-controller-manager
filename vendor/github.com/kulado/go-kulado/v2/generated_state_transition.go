package client

const (
	STATE_TRANSITION_TYPE = "stateTransition"
)

type StateTransition struct {
	Resource
}

type StateTransitionCollection struct {
	Collection
	Data   []StateTransition `json:"data,omitempty"`
	client *StateTransitionClient
}

type StateTransitionClient struct {
	kuladoClient *KuladoClient
}

type StateTransitionOperations interface {
	List(opts *ListOpts) (*StateTransitionCollection, error)
	Create(opts *StateTransition) (*StateTransition, error)
	Update(existing *StateTransition, updates interface{}) (*StateTransition, error)
	ById(id string) (*StateTransition, error)
	Delete(container *StateTransition) error
}

func newStateTransitionClient(kuladoClient *KuladoClient) *StateTransitionClient {
	return &StateTransitionClient{
		kuladoClient: kuladoClient,
	}
}

func (c *StateTransitionClient) Create(container *StateTransition) (*StateTransition, error) {
	resp := &StateTransition{}
	err := c.kuladoClient.doCreate(STATE_TRANSITION_TYPE, container, resp)
	return resp, err
}

func (c *StateTransitionClient) Update(existing *StateTransition, updates interface{}) (*StateTransition, error) {
	resp := &StateTransition{}
	err := c.kuladoClient.doUpdate(STATE_TRANSITION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *StateTransitionClient) List(opts *ListOpts) (*StateTransitionCollection, error) {
	resp := &StateTransitionCollection{}
	err := c.kuladoClient.doList(STATE_TRANSITION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *StateTransitionCollection) Next() (*StateTransitionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &StateTransitionCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *StateTransitionClient) ById(id string) (*StateTransition, error) {
	resp := &StateTransition{}
	err := c.kuladoClient.doById(STATE_TRANSITION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *StateTransitionClient) Delete(container *StateTransition) error {
	return c.kuladoClient.doResourceDelete(STATE_TRANSITION_TYPE, &container.Resource)
}
