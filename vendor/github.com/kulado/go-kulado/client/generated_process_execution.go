package client

const (
	PROCESS_EXECUTION_TYPE = "processExecution"
)

type ProcessExecution struct {
	Resource

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Log map[string]interface{} `json:"log,omitempty" yaml:"log,omitempty"`

	ProcessInstanceId string `json:"processInstanceId,omitempty" yaml:"process_instance_id,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ProcessExecutionCollection struct {
	Collection
	Data   []ProcessExecution `json:"data,omitempty"`
	client *ProcessExecutionClient
}

type ProcessExecutionClient struct {
	kuladoClient *KuladoClient
}

type ProcessExecutionOperations interface {
	List(opts *ListOpts) (*ProcessExecutionCollection, error)
	Create(opts *ProcessExecution) (*ProcessExecution, error)
	Update(existing *ProcessExecution, updates interface{}) (*ProcessExecution, error)
	ById(id string) (*ProcessExecution, error)
	Delete(container *ProcessExecution) error
}

func newProcessExecutionClient(kuladoClient *KuladoClient) *ProcessExecutionClient {
	return &ProcessExecutionClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ProcessExecutionClient) Create(container *ProcessExecution) (*ProcessExecution, error) {
	resp := &ProcessExecution{}
	err := c.kuladoClient.doCreate(PROCESS_EXECUTION_TYPE, container, resp)
	return resp, err
}

func (c *ProcessExecutionClient) Update(existing *ProcessExecution, updates interface{}) (*ProcessExecution, error) {
	resp := &ProcessExecution{}
	err := c.kuladoClient.doUpdate(PROCESS_EXECUTION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProcessExecutionClient) List(opts *ListOpts) (*ProcessExecutionCollection, error) {
	resp := &ProcessExecutionCollection{}
	err := c.kuladoClient.doList(PROCESS_EXECUTION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProcessExecutionCollection) Next() (*ProcessExecutionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProcessExecutionCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProcessExecutionClient) ById(id string) (*ProcessExecution, error) {
	resp := &ProcessExecution{}
	err := c.kuladoClient.doById(PROCESS_EXECUTION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ProcessExecutionClient) Delete(container *ProcessExecution) error {
	return c.kuladoClient.doResourceDelete(PROCESS_EXECUTION_TYPE, &container.Resource)
}
