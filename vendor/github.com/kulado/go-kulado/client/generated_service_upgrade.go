package client

const (
	SERVICE_UPGRADE_TYPE = "serviceUpgrade"
)

type ServiceUpgrade struct {
	Resource

	InServiceStrategy *InServiceUpgradeStrategy `json:"inServiceStrategy,omitempty" yaml:"in_service_strategy,omitempty"`

	ToServiceStrategy *ToServiceUpgradeStrategy `json:"toServiceStrategy,omitempty" yaml:"to_service_strategy,omitempty"`
}

type ServiceUpgradeCollection struct {
	Collection
	Data   []ServiceUpgrade `json:"data,omitempty"`
	client *ServiceUpgradeClient
}

type ServiceUpgradeClient struct {
	kuladoClient *KuladoClient
}

type ServiceUpgradeOperations interface {
	List(opts *ListOpts) (*ServiceUpgradeCollection, error)
	Create(opts *ServiceUpgrade) (*ServiceUpgrade, error)
	Update(existing *ServiceUpgrade, updates interface{}) (*ServiceUpgrade, error)
	ById(id string) (*ServiceUpgrade, error)
	Delete(container *ServiceUpgrade) error
}

func newServiceUpgradeClient(kuladoClient *KuladoClient) *ServiceUpgradeClient {
	return &ServiceUpgradeClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ServiceUpgradeClient) Create(container *ServiceUpgrade) (*ServiceUpgrade, error) {
	resp := &ServiceUpgrade{}
	err := c.kuladoClient.doCreate(SERVICE_UPGRADE_TYPE, container, resp)
	return resp, err
}

func (c *ServiceUpgradeClient) Update(existing *ServiceUpgrade, updates interface{}) (*ServiceUpgrade, error) {
	resp := &ServiceUpgrade{}
	err := c.kuladoClient.doUpdate(SERVICE_UPGRADE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceUpgradeClient) List(opts *ListOpts) (*ServiceUpgradeCollection, error) {
	resp := &ServiceUpgradeCollection{}
	err := c.kuladoClient.doList(SERVICE_UPGRADE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceUpgradeCollection) Next() (*ServiceUpgradeCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceUpgradeCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceUpgradeClient) ById(id string) (*ServiceUpgrade, error) {
	resp := &ServiceUpgrade{}
	err := c.kuladoClient.doById(SERVICE_UPGRADE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ServiceUpgradeClient) Delete(container *ServiceUpgrade) error {
	return c.kuladoClient.doResourceDelete(SERVICE_UPGRADE_TYPE, &container.Resource)
}
