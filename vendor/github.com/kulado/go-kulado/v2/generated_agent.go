package client

const (
	AGENT_TYPE = "agent"
)

type Agent struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	ManagedConfig bool `json:"managedConfig,omitempty" yaml:"managed_config,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type AgentCollection struct {
	Collection
	Data   []Agent `json:"data,omitempty"`
	client *AgentClient
}

type AgentClient struct {
	kuladoClient *KuladoClient
}

type AgentOperations interface {
	List(opts *ListOpts) (*AgentCollection, error)
	Create(opts *Agent) (*Agent, error)
	Update(existing *Agent, updates interface{}) (*Agent, error)
	ById(id string) (*Agent, error)
	Delete(container *Agent) error

	ActionActivate(*Agent) (*Agent, error)

	ActionCreate(*Agent) (*Agent, error)

	ActionDeactivate(*Agent) (*Agent, error)

	ActionDisconnect(*Agent) (*Agent, error)

	ActionFinishreconnect(*Agent) (*Agent, error)

	ActionPurge(*Agent) (*Agent, error)

	ActionReconnect(*Agent) (*Agent, error)

	ActionRemove(*Agent) (*Agent, error)

	ActionRestore(*Agent) (*Agent, error)

	ActionUpdate(*Agent) (*Agent, error)
}

func newAgentClient(kuladoClient *KuladoClient) *AgentClient {
	return &AgentClient{
		kuladoClient: kuladoClient,
	}
}

func (c *AgentClient) Create(container *Agent) (*Agent, error) {
	resp := &Agent{}
	err := c.kuladoClient.doCreate(AGENT_TYPE, container, resp)
	return resp, err
}

func (c *AgentClient) Update(existing *Agent, updates interface{}) (*Agent, error) {
	resp := &Agent{}
	err := c.kuladoClient.doUpdate(AGENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AgentClient) List(opts *ListOpts) (*AgentCollection, error) {
	resp := &AgentCollection{}
	err := c.kuladoClient.doList(AGENT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *AgentCollection) Next() (*AgentCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AgentCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AgentClient) ById(id string) (*Agent, error) {
	resp := &Agent{}
	err := c.kuladoClient.doById(AGENT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *AgentClient) Delete(container *Agent) error {
	return c.kuladoClient.doResourceDelete(AGENT_TYPE, &container.Resource)
}

func (c *AgentClient) ActionActivate(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *AgentClient) ActionCreate(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *AgentClient) ActionDeactivate(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *AgentClient) ActionDisconnect(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "disconnect", &resource.Resource, nil, resp)

	return resp, err
}

func (c *AgentClient) ActionFinishreconnect(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "finishreconnect", &resource.Resource, nil, resp)

	return resp, err
}

func (c *AgentClient) ActionPurge(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "purge", &resource.Resource, nil, resp)

	return resp, err
}

func (c *AgentClient) ActionReconnect(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "reconnect", &resource.Resource, nil, resp)

	return resp, err
}

func (c *AgentClient) ActionRemove(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *AgentClient) ActionRestore(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "restore", &resource.Resource, nil, resp)

	return resp, err
}

func (c *AgentClient) ActionUpdate(resource *Agent) (*Agent, error) {

	resp := &Agent{}

	err := c.kuladoClient.doAction(AGENT_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
