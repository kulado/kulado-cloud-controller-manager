package client

const (
	TASK_INSTANCE_TYPE = "taskInstance"
)

type TaskInstance struct {
	Resource

	EndTime string `json:"endTime,omitempty" yaml:"end_time,omitempty"`

	Exception string `json:"exception,omitempty" yaml:"exception,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	ServerId string `json:"serverId,omitempty" yaml:"server_id,omitempty"`

	StartTime string `json:"startTime,omitempty" yaml:"start_time,omitempty"`

	TaskId string `json:"taskId,omitempty" yaml:"task_id,omitempty"`
}

type TaskInstanceCollection struct {
	Collection
	Data   []TaskInstance `json:"data,omitempty"`
	client *TaskInstanceClient
}

type TaskInstanceClient struct {
	kuladoClient *KuladoClient
}

type TaskInstanceOperations interface {
	List(opts *ListOpts) (*TaskInstanceCollection, error)
	Create(opts *TaskInstance) (*TaskInstance, error)
	Update(existing *TaskInstance, updates interface{}) (*TaskInstance, error)
	ById(id string) (*TaskInstance, error)
	Delete(container *TaskInstance) error
}

func newTaskInstanceClient(kuladoClient *KuladoClient) *TaskInstanceClient {
	return &TaskInstanceClient{
		kuladoClient: kuladoClient,
	}
}

func (c *TaskInstanceClient) Create(container *TaskInstance) (*TaskInstance, error) {
	resp := &TaskInstance{}
	err := c.kuladoClient.doCreate(TASK_INSTANCE_TYPE, container, resp)
	return resp, err
}

func (c *TaskInstanceClient) Update(existing *TaskInstance, updates interface{}) (*TaskInstance, error) {
	resp := &TaskInstance{}
	err := c.kuladoClient.doUpdate(TASK_INSTANCE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TaskInstanceClient) List(opts *ListOpts) (*TaskInstanceCollection, error) {
	resp := &TaskInstanceCollection{}
	err := c.kuladoClient.doList(TASK_INSTANCE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *TaskInstanceCollection) Next() (*TaskInstanceCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TaskInstanceCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TaskInstanceClient) ById(id string) (*TaskInstance, error) {
	resp := &TaskInstance{}
	err := c.kuladoClient.doById(TASK_INSTANCE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *TaskInstanceClient) Delete(container *TaskInstance) error {
	return c.kuladoClient.doResourceDelete(TASK_INSTANCE_TYPE, &container.Resource)
}
