package client

const (
	SET_LABELS_INPUT_TYPE = "setLabelsInput"
)

type SetLabelsInput struct {
	Resource

	Labels interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`
}

type SetLabelsInputCollection struct {
	Collection
	Data   []SetLabelsInput `json:"data,omitempty"`
	client *SetLabelsInputClient
}

type SetLabelsInputClient struct {
	kuladoClient *KuladoClient
}

type SetLabelsInputOperations interface {
	List(opts *ListOpts) (*SetLabelsInputCollection, error)
	Create(opts *SetLabelsInput) (*SetLabelsInput, error)
	Update(existing *SetLabelsInput, updates interface{}) (*SetLabelsInput, error)
	ById(id string) (*SetLabelsInput, error)
	Delete(container *SetLabelsInput) error
}

func newSetLabelsInputClient(kuladoClient *KuladoClient) *SetLabelsInputClient {
	return &SetLabelsInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *SetLabelsInputClient) Create(container *SetLabelsInput) (*SetLabelsInput, error) {
	resp := &SetLabelsInput{}
	err := c.kuladoClient.doCreate(SET_LABELS_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *SetLabelsInputClient) Update(existing *SetLabelsInput, updates interface{}) (*SetLabelsInput, error) {
	resp := &SetLabelsInput{}
	err := c.kuladoClient.doUpdate(SET_LABELS_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SetLabelsInputClient) List(opts *ListOpts) (*SetLabelsInputCollection, error) {
	resp := &SetLabelsInputCollection{}
	err := c.kuladoClient.doList(SET_LABELS_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SetLabelsInputCollection) Next() (*SetLabelsInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SetLabelsInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SetLabelsInputClient) ById(id string) (*SetLabelsInput, error) {
	resp := &SetLabelsInput{}
	err := c.kuladoClient.doById(SET_LABELS_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SetLabelsInputClient) Delete(container *SetLabelsInput) error {
	return c.kuladoClient.doResourceDelete(SET_LABELS_INPUT_TYPE, &container.Resource)
}
