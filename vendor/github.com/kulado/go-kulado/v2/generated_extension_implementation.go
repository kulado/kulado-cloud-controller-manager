package client

const (
	EXTENSION_IMPLEMENTATION_TYPE = "extensionImplementation"
)

type ExtensionImplementation struct {
	Resource

	ClassName string `json:"className,omitempty" yaml:"class_name,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Properties map[string]interface{} `json:"properties,omitempty" yaml:"properties,omitempty"`
}

type ExtensionImplementationCollection struct {
	Collection
	Data   []ExtensionImplementation `json:"data,omitempty"`
	client *ExtensionImplementationClient
}

type ExtensionImplementationClient struct {
	kuladoClient *KuladoClient
}

type ExtensionImplementationOperations interface {
	List(opts *ListOpts) (*ExtensionImplementationCollection, error)
	Create(opts *ExtensionImplementation) (*ExtensionImplementation, error)
	Update(existing *ExtensionImplementation, updates interface{}) (*ExtensionImplementation, error)
	ById(id string) (*ExtensionImplementation, error)
	Delete(container *ExtensionImplementation) error
}

func newExtensionImplementationClient(kuladoClient *KuladoClient) *ExtensionImplementationClient {
	return &ExtensionImplementationClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExtensionImplementationClient) Create(container *ExtensionImplementation) (*ExtensionImplementation, error) {
	resp := &ExtensionImplementation{}
	err := c.kuladoClient.doCreate(EXTENSION_IMPLEMENTATION_TYPE, container, resp)
	return resp, err
}

func (c *ExtensionImplementationClient) Update(existing *ExtensionImplementation, updates interface{}) (*ExtensionImplementation, error) {
	resp := &ExtensionImplementation{}
	err := c.kuladoClient.doUpdate(EXTENSION_IMPLEMENTATION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExtensionImplementationClient) List(opts *ListOpts) (*ExtensionImplementationCollection, error) {
	resp := &ExtensionImplementationCollection{}
	err := c.kuladoClient.doList(EXTENSION_IMPLEMENTATION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExtensionImplementationCollection) Next() (*ExtensionImplementationCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExtensionImplementationCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExtensionImplementationClient) ById(id string) (*ExtensionImplementation, error) {
	resp := &ExtensionImplementation{}
	err := c.kuladoClient.doById(EXTENSION_IMPLEMENTATION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExtensionImplementationClient) Delete(container *ExtensionImplementation) error {
	return c.kuladoClient.doResourceDelete(EXTENSION_IMPLEMENTATION_TYPE, &container.Resource)
}
