package client

const (
	TYPE_DOCUMENTATION_TYPE = "typeDocumentation"
)

type TypeDocumentation struct {
	Resource

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	ResourceFields map[string]interface{} `json:"resourceFields,omitempty" yaml:"resource_fields,omitempty"`
}

type TypeDocumentationCollection struct {
	Collection
	Data   []TypeDocumentation `json:"data,omitempty"`
	client *TypeDocumentationClient
}

type TypeDocumentationClient struct {
	kuladoClient *KuladoClient
}

type TypeDocumentationOperations interface {
	List(opts *ListOpts) (*TypeDocumentationCollection, error)
	Create(opts *TypeDocumentation) (*TypeDocumentation, error)
	Update(existing *TypeDocumentation, updates interface{}) (*TypeDocumentation, error)
	ById(id string) (*TypeDocumentation, error)
	Delete(container *TypeDocumentation) error
}

func newTypeDocumentationClient(kuladoClient *KuladoClient) *TypeDocumentationClient {
	return &TypeDocumentationClient{
		kuladoClient: kuladoClient,
	}
}

func (c *TypeDocumentationClient) Create(container *TypeDocumentation) (*TypeDocumentation, error) {
	resp := &TypeDocumentation{}
	err := c.kuladoClient.doCreate(TYPE_DOCUMENTATION_TYPE, container, resp)
	return resp, err
}

func (c *TypeDocumentationClient) Update(existing *TypeDocumentation, updates interface{}) (*TypeDocumentation, error) {
	resp := &TypeDocumentation{}
	err := c.kuladoClient.doUpdate(TYPE_DOCUMENTATION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TypeDocumentationClient) List(opts *ListOpts) (*TypeDocumentationCollection, error) {
	resp := &TypeDocumentationCollection{}
	err := c.kuladoClient.doList(TYPE_DOCUMENTATION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *TypeDocumentationCollection) Next() (*TypeDocumentationCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TypeDocumentationCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TypeDocumentationClient) ById(id string) (*TypeDocumentation, error) {
	resp := &TypeDocumentation{}
	err := c.kuladoClient.doById(TYPE_DOCUMENTATION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *TypeDocumentationClient) Delete(container *TypeDocumentation) error {
	return c.kuladoClient.doResourceDelete(TYPE_DOCUMENTATION_TYPE, &container.Resource)
}
