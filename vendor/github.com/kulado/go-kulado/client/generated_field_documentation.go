package client

const (
	FIELD_DOCUMENTATION_TYPE = "fieldDocumentation"
)

type FieldDocumentation struct {
	Resource

	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type FieldDocumentationCollection struct {
	Collection
	Data   []FieldDocumentation `json:"data,omitempty"`
	client *FieldDocumentationClient
}

type FieldDocumentationClient struct {
	kuladoClient *KuladoClient
}

type FieldDocumentationOperations interface {
	List(opts *ListOpts) (*FieldDocumentationCollection, error)
	Create(opts *FieldDocumentation) (*FieldDocumentation, error)
	Update(existing *FieldDocumentation, updates interface{}) (*FieldDocumentation, error)
	ById(id string) (*FieldDocumentation, error)
	Delete(container *FieldDocumentation) error
}

func newFieldDocumentationClient(kuladoClient *KuladoClient) *FieldDocumentationClient {
	return &FieldDocumentationClient{
		kuladoClient: kuladoClient,
	}
}

func (c *FieldDocumentationClient) Create(container *FieldDocumentation) (*FieldDocumentation, error) {
	resp := &FieldDocumentation{}
	err := c.kuladoClient.doCreate(FIELD_DOCUMENTATION_TYPE, container, resp)
	return resp, err
}

func (c *FieldDocumentationClient) Update(existing *FieldDocumentation, updates interface{}) (*FieldDocumentation, error) {
	resp := &FieldDocumentation{}
	err := c.kuladoClient.doUpdate(FIELD_DOCUMENTATION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *FieldDocumentationClient) List(opts *ListOpts) (*FieldDocumentationCollection, error) {
	resp := &FieldDocumentationCollection{}
	err := c.kuladoClient.doList(FIELD_DOCUMENTATION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *FieldDocumentationCollection) Next() (*FieldDocumentationCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &FieldDocumentationCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *FieldDocumentationClient) ById(id string) (*FieldDocumentation, error) {
	resp := &FieldDocumentation{}
	err := c.kuladoClient.doById(FIELD_DOCUMENTATION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *FieldDocumentationClient) Delete(container *FieldDocumentation) error {
	return c.kuladoClient.doResourceDelete(FIELD_DOCUMENTATION_TYPE, &container.Resource)
}
