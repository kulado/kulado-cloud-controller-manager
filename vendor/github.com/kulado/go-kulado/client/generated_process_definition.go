package client

const (
	PROCESS_DEFINITION_TYPE = "processDefinition"
)

type ProcessDefinition struct {
	Resource

	ExtensionBased bool `json:"extensionBased,omitempty" yaml:"extension_based,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	PostProcessListeners interface{} `json:"postProcessListeners,omitempty" yaml:"post_process_listeners,omitempty"`

	PreProcessListeners interface{} `json:"preProcessListeners,omitempty" yaml:"pre_process_listeners,omitempty"`

	ProcessHandlers interface{} `json:"processHandlers,omitempty" yaml:"process_handlers,omitempty"`

	ResourceType string `json:"resourceType,omitempty" yaml:"resource_type,omitempty"`

	StateTransitions []interface{} `json:"stateTransitions,omitempty" yaml:"state_transitions,omitempty"`
}

type ProcessDefinitionCollection struct {
	Collection
	Data   []ProcessDefinition `json:"data,omitempty"`
	client *ProcessDefinitionClient
}

type ProcessDefinitionClient struct {
	kuladoClient *KuladoClient
}

type ProcessDefinitionOperations interface {
	List(opts *ListOpts) (*ProcessDefinitionCollection, error)
	Create(opts *ProcessDefinition) (*ProcessDefinition, error)
	Update(existing *ProcessDefinition, updates interface{}) (*ProcessDefinition, error)
	ById(id string) (*ProcessDefinition, error)
	Delete(container *ProcessDefinition) error
}

func newProcessDefinitionClient(kuladoClient *KuladoClient) *ProcessDefinitionClient {
	return &ProcessDefinitionClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ProcessDefinitionClient) Create(container *ProcessDefinition) (*ProcessDefinition, error) {
	resp := &ProcessDefinition{}
	err := c.kuladoClient.doCreate(PROCESS_DEFINITION_TYPE, container, resp)
	return resp, err
}

func (c *ProcessDefinitionClient) Update(existing *ProcessDefinition, updates interface{}) (*ProcessDefinition, error) {
	resp := &ProcessDefinition{}
	err := c.kuladoClient.doUpdate(PROCESS_DEFINITION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProcessDefinitionClient) List(opts *ListOpts) (*ProcessDefinitionCollection, error) {
	resp := &ProcessDefinitionCollection{}
	err := c.kuladoClient.doList(PROCESS_DEFINITION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProcessDefinitionCollection) Next() (*ProcessDefinitionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProcessDefinitionCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProcessDefinitionClient) ById(id string) (*ProcessDefinition, error) {
	resp := &ProcessDefinition{}
	err := c.kuladoClient.doById(PROCESS_DEFINITION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ProcessDefinitionClient) Delete(container *ProcessDefinition) error {
	return c.kuladoClient.doResourceDelete(PROCESS_DEFINITION_TYPE, &container.Resource)
}
