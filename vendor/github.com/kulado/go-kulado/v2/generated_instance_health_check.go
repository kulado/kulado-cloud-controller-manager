package client

const (
	INSTANCE_HEALTH_CHECK_TYPE = "instanceHealthCheck"
)

type InstanceHealthCheck struct {
	Resource

	HealthyThreshold int64 `json:"healthyThreshold,omitempty" yaml:"healthy_threshold,omitempty"`

	InitializingTimeout int64 `json:"initializingTimeout,omitempty" yaml:"initializing_timeout,omitempty"`

	Interval int64 `json:"interval,omitempty" yaml:"interval,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Port int64 `json:"port,omitempty" yaml:"port,omitempty"`

	RecreateOnQuorumStrategyConfig *RecreateOnQuorumStrategyConfig `json:"recreateOnQuorumStrategyConfig,omitempty" yaml:"recreate_on_quorum_strategy_config,omitempty"`

	ReinitializingTimeout int64 `json:"reinitializingTimeout,omitempty" yaml:"reinitializing_timeout,omitempty"`

	RequestLine string `json:"requestLine,omitempty" yaml:"request_line,omitempty"`

	ResponseTimeout int64 `json:"responseTimeout,omitempty" yaml:"response_timeout,omitempty"`

	Strategy string `json:"strategy,omitempty" yaml:"strategy,omitempty"`

	UnhealthyThreshold int64 `json:"unhealthyThreshold,omitempty" yaml:"unhealthy_threshold,omitempty"`
}

type InstanceHealthCheckCollection struct {
	Collection
	Data   []InstanceHealthCheck `json:"data,omitempty"`
	client *InstanceHealthCheckClient
}

type InstanceHealthCheckClient struct {
	kuladoClient *KuladoClient
}

type InstanceHealthCheckOperations interface {
	List(opts *ListOpts) (*InstanceHealthCheckCollection, error)
	Create(opts *InstanceHealthCheck) (*InstanceHealthCheck, error)
	Update(existing *InstanceHealthCheck, updates interface{}) (*InstanceHealthCheck, error)
	ById(id string) (*InstanceHealthCheck, error)
	Delete(container *InstanceHealthCheck) error
}

func newInstanceHealthCheckClient(kuladoClient *KuladoClient) *InstanceHealthCheckClient {
	return &InstanceHealthCheckClient{
		kuladoClient: kuladoClient,
	}
}

func (c *InstanceHealthCheckClient) Create(container *InstanceHealthCheck) (*InstanceHealthCheck, error) {
	resp := &InstanceHealthCheck{}
	err := c.kuladoClient.doCreate(INSTANCE_HEALTH_CHECK_TYPE, container, resp)
	return resp, err
}

func (c *InstanceHealthCheckClient) Update(existing *InstanceHealthCheck, updates interface{}) (*InstanceHealthCheck, error) {
	resp := &InstanceHealthCheck{}
	err := c.kuladoClient.doUpdate(INSTANCE_HEALTH_CHECK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InstanceHealthCheckClient) List(opts *ListOpts) (*InstanceHealthCheckCollection, error) {
	resp := &InstanceHealthCheckCollection{}
	err := c.kuladoClient.doList(INSTANCE_HEALTH_CHECK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *InstanceHealthCheckCollection) Next() (*InstanceHealthCheckCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &InstanceHealthCheckCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *InstanceHealthCheckClient) ById(id string) (*InstanceHealthCheck, error) {
	resp := &InstanceHealthCheck{}
	err := c.kuladoClient.doById(INSTANCE_HEALTH_CHECK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *InstanceHealthCheckClient) Delete(container *InstanceHealthCheck) error {
	return c.kuladoClient.doResourceDelete(INSTANCE_HEALTH_CHECK_TYPE, &container.Resource)
}
