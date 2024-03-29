package client

const (
	RESOURCE_DEFINITION_TYPE = "resourceDefinition"
)

type ResourceDefinition struct {
	Resource

	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type ResourceDefinitionCollection struct {
	Collection
	Data   []ResourceDefinition `json:"data,omitempty"`
	client *ResourceDefinitionClient
}

type ResourceDefinitionClient struct {
	kuladoClient *KuladoClient
}

type ResourceDefinitionOperations interface {
	List(opts *ListOpts) (*ResourceDefinitionCollection, error)
	Create(opts *ResourceDefinition) (*ResourceDefinition, error)
	Update(existing *ResourceDefinition, updates interface{}) (*ResourceDefinition, error)
	ById(id string) (*ResourceDefinition, error)
	Delete(container *ResourceDefinition) error
}

func newResourceDefinitionClient(kuladoClient *KuladoClient) *ResourceDefinitionClient {
	return &ResourceDefinitionClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ResourceDefinitionClient) Create(container *ResourceDefinition) (*ResourceDefinition, error) {
	resp := &ResourceDefinition{}
	err := c.kuladoClient.doCreate(RESOURCE_DEFINITION_TYPE, container, resp)
	return resp, err
}

func (c *ResourceDefinitionClient) Update(existing *ResourceDefinition, updates interface{}) (*ResourceDefinition, error) {
	resp := &ResourceDefinition{}
	err := c.kuladoClient.doUpdate(RESOURCE_DEFINITION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ResourceDefinitionClient) List(opts *ListOpts) (*ResourceDefinitionCollection, error) {
	resp := &ResourceDefinitionCollection{}
	err := c.kuladoClient.doList(RESOURCE_DEFINITION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ResourceDefinitionCollection) Next() (*ResourceDefinitionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ResourceDefinitionCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ResourceDefinitionClient) ById(id string) (*ResourceDefinition, error) {
	resp := &ResourceDefinition{}
	err := c.kuladoClient.doById(RESOURCE_DEFINITION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ResourceDefinitionClient) Delete(container *ResourceDefinition) error {
	return c.kuladoClient.doResourceDelete(RESOURCE_DEFINITION_TYPE, &container.Resource)
}
