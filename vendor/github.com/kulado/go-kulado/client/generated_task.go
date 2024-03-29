package client

const (
	TASK_TYPE = "task"
)

type Task struct {
	Resource

	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type TaskCollection struct {
	Collection
	Data   []Task `json:"data,omitempty"`
	client *TaskClient
}

type TaskClient struct {
	kuladoClient *KuladoClient
}

type TaskOperations interface {
	List(opts *ListOpts) (*TaskCollection, error)
	Create(opts *Task) (*Task, error)
	Update(existing *Task, updates interface{}) (*Task, error)
	ById(id string) (*Task, error)
	Delete(container *Task) error

	ActionExecute(*Task) (*Task, error)
}

func newTaskClient(kuladoClient *KuladoClient) *TaskClient {
	return &TaskClient{
		kuladoClient: kuladoClient,
	}
}

func (c *TaskClient) Create(container *Task) (*Task, error) {
	resp := &Task{}
	err := c.kuladoClient.doCreate(TASK_TYPE, container, resp)
	return resp, err
}

func (c *TaskClient) Update(existing *Task, updates interface{}) (*Task, error) {
	resp := &Task{}
	err := c.kuladoClient.doUpdate(TASK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TaskClient) List(opts *ListOpts) (*TaskCollection, error) {
	resp := &TaskCollection{}
	err := c.kuladoClient.doList(TASK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *TaskCollection) Next() (*TaskCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TaskCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TaskClient) ById(id string) (*Task, error) {
	resp := &Task{}
	err := c.kuladoClient.doById(TASK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *TaskClient) Delete(container *Task) error {
	return c.kuladoClient.doResourceDelete(TASK_TYPE, &container.Resource)
}

func (c *TaskClient) ActionExecute(resource *Task) (*Task, error) {

	resp := &Task{}

	err := c.kuladoClient.doAction(TASK_TYPE, "execute", &resource.Resource, nil, resp)

	return resp, err
}
