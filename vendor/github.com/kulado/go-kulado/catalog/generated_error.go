package catalog

const (
	ERROR_TYPE = "error"
)

type Error struct {
	Resource

	Actions map[string]interface{} `json:"actions,omitempty" yaml:"actions,omitempty"`

	Links map[string]interface{} `json:"links,omitempty" yaml:"links,omitempty"`

	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type ErrorCollection struct {
	Collection
	Data   []Error `json:"data,omitempty"`
	client *ErrorClient
}

type ErrorClient struct {
	kuladoClient *KuladoClient
}

type ErrorOperations interface {
	List(opts *ListOpts) (*ErrorCollection, error)
	Create(opts *Error) (*Error, error)
	Update(existing *Error, updates interface{}) (*Error, error)
	ById(id string) (*Error, error)
	Delete(container *Error) error
}

func newErrorClient(kuladoClient *KuladoClient) *ErrorClient {
	return &ErrorClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ErrorClient) Create(container *Error) (*Error, error) {
	resp := &Error{}
	err := c.kuladoClient.doCreate(ERROR_TYPE, container, resp)
	return resp, err
}

func (c *ErrorClient) Update(existing *Error, updates interface{}) (*Error, error) {
	resp := &Error{}
	err := c.kuladoClient.doUpdate(ERROR_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ErrorClient) List(opts *ListOpts) (*ErrorCollection, error) {
	resp := &ErrorCollection{}
	err := c.kuladoClient.doList(ERROR_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ErrorCollection) Next() (*ErrorCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ErrorCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ErrorClient) ById(id string) (*Error, error) {
	resp := &Error{}
	err := c.kuladoClient.doById(ERROR_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ErrorClient) Delete(container *Error) error {
	return c.kuladoClient.doResourceDelete(ERROR_TYPE, &container.Resource)
}
