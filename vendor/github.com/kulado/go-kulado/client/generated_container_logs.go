package client

const (
	CONTAINER_LOGS_TYPE = "containerLogs"
)

type ContainerLogs struct {
	Resource

	Follow bool `json:"follow,omitempty" yaml:"follow,omitempty"`

	Lines int64 `json:"lines,omitempty" yaml:"lines,omitempty"`
}

type ContainerLogsCollection struct {
	Collection
	Data   []ContainerLogs `json:"data,omitempty"`
	client *ContainerLogsClient
}

type ContainerLogsClient struct {
	kuladoClient *KuladoClient
}

type ContainerLogsOperations interface {
	List(opts *ListOpts) (*ContainerLogsCollection, error)
	Create(opts *ContainerLogs) (*ContainerLogs, error)
	Update(existing *ContainerLogs, updates interface{}) (*ContainerLogs, error)
	ById(id string) (*ContainerLogs, error)
	Delete(container *ContainerLogs) error
}

func newContainerLogsClient(kuladoClient *KuladoClient) *ContainerLogsClient {
	return &ContainerLogsClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ContainerLogsClient) Create(container *ContainerLogs) (*ContainerLogs, error) {
	resp := &ContainerLogs{}
	err := c.kuladoClient.doCreate(CONTAINER_LOGS_TYPE, container, resp)
	return resp, err
}

func (c *ContainerLogsClient) Update(existing *ContainerLogs, updates interface{}) (*ContainerLogs, error) {
	resp := &ContainerLogs{}
	err := c.kuladoClient.doUpdate(CONTAINER_LOGS_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ContainerLogsClient) List(opts *ListOpts) (*ContainerLogsCollection, error) {
	resp := &ContainerLogsCollection{}
	err := c.kuladoClient.doList(CONTAINER_LOGS_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ContainerLogsCollection) Next() (*ContainerLogsCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ContainerLogsCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ContainerLogsClient) ById(id string) (*ContainerLogs, error) {
	resp := &ContainerLogs{}
	err := c.kuladoClient.doById(CONTAINER_LOGS_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ContainerLogsClient) Delete(container *ContainerLogs) error {
	return c.kuladoClient.doResourceDelete(CONTAINER_LOGS_TYPE, &container.Resource)
}
