package client

const (
	SERVICE_UPGRADE_STRATEGY_TYPE = "serviceUpgradeStrategy"
)

type ServiceUpgradeStrategy struct {
	Resource

	BatchSize int64 `json:"batchSize,omitempty" yaml:"batch_size,omitempty"`

	IntervalMillis int64 `json:"intervalMillis,omitempty" yaml:"interval_millis,omitempty"`
}

type ServiceUpgradeStrategyCollection struct {
	Collection
	Data   []ServiceUpgradeStrategy `json:"data,omitempty"`
	client *ServiceUpgradeStrategyClient
}

type ServiceUpgradeStrategyClient struct {
	kuladoClient *KuladoClient
}

type ServiceUpgradeStrategyOperations interface {
	List(opts *ListOpts) (*ServiceUpgradeStrategyCollection, error)
	Create(opts *ServiceUpgradeStrategy) (*ServiceUpgradeStrategy, error)
	Update(existing *ServiceUpgradeStrategy, updates interface{}) (*ServiceUpgradeStrategy, error)
	ById(id string) (*ServiceUpgradeStrategy, error)
	Delete(container *ServiceUpgradeStrategy) error
}

func newServiceUpgradeStrategyClient(kuladoClient *KuladoClient) *ServiceUpgradeStrategyClient {
	return &ServiceUpgradeStrategyClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ServiceUpgradeStrategyClient) Create(container *ServiceUpgradeStrategy) (*ServiceUpgradeStrategy, error) {
	resp := &ServiceUpgradeStrategy{}
	err := c.kuladoClient.doCreate(SERVICE_UPGRADE_STRATEGY_TYPE, container, resp)
	return resp, err
}

func (c *ServiceUpgradeStrategyClient) Update(existing *ServiceUpgradeStrategy, updates interface{}) (*ServiceUpgradeStrategy, error) {
	resp := &ServiceUpgradeStrategy{}
	err := c.kuladoClient.doUpdate(SERVICE_UPGRADE_STRATEGY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceUpgradeStrategyClient) List(opts *ListOpts) (*ServiceUpgradeStrategyCollection, error) {
	resp := &ServiceUpgradeStrategyCollection{}
	err := c.kuladoClient.doList(SERVICE_UPGRADE_STRATEGY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceUpgradeStrategyCollection) Next() (*ServiceUpgradeStrategyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceUpgradeStrategyCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceUpgradeStrategyClient) ById(id string) (*ServiceUpgradeStrategy, error) {
	resp := &ServiceUpgradeStrategy{}
	err := c.kuladoClient.doById(SERVICE_UPGRADE_STRATEGY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ServiceUpgradeStrategyClient) Delete(container *ServiceUpgradeStrategy) error {
	return c.kuladoClient.doResourceDelete(SERVICE_UPGRADE_STRATEGY_TYPE, &container.Resource)
}
