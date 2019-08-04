package client

const (
	STACK_UPGRADE_TYPE = "stackUpgrade"
)

type StackUpgrade struct {
	Resource

	DockerCompose string `json:"dockerCompose,omitempty" yaml:"docker_compose,omitempty"`

	Environment map[string]interface{} `json:"environment,omitempty" yaml:"environment,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	KuladoCompose string `json:"kuladoCompose,omitempty" yaml:"kulado_compose,omitempty"`
}

type StackUpgradeCollection struct {
	Collection
	Data   []StackUpgrade `json:"data,omitempty"`
	client *StackUpgradeClient
}

type StackUpgradeClient struct {
	kuladoClient *KuladoClient
}

type StackUpgradeOperations interface {
	List(opts *ListOpts) (*StackUpgradeCollection, error)
	Create(opts *StackUpgrade) (*StackUpgrade, error)
	Update(existing *StackUpgrade, updates interface{}) (*StackUpgrade, error)
	ById(id string) (*StackUpgrade, error)
	Delete(container *StackUpgrade) error
}

func newStackUpgradeClient(kuladoClient *KuladoClient) *StackUpgradeClient {
	return &StackUpgradeClient{
		kuladoClient: kuladoClient,
	}
}

func (c *StackUpgradeClient) Create(container *StackUpgrade) (*StackUpgrade, error) {
	resp := &StackUpgrade{}
	err := c.kuladoClient.doCreate(STACK_UPGRADE_TYPE, container, resp)
	return resp, err
}

func (c *StackUpgradeClient) Update(existing *StackUpgrade, updates interface{}) (*StackUpgrade, error) {
	resp := &StackUpgrade{}
	err := c.kuladoClient.doUpdate(STACK_UPGRADE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *StackUpgradeClient) List(opts *ListOpts) (*StackUpgradeCollection, error) {
	resp := &StackUpgradeCollection{}
	err := c.kuladoClient.doList(STACK_UPGRADE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *StackUpgradeCollection) Next() (*StackUpgradeCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &StackUpgradeCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *StackUpgradeClient) ById(id string) (*StackUpgrade, error) {
	resp := &StackUpgrade{}
	err := c.kuladoClient.doById(STACK_UPGRADE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *StackUpgradeClient) Delete(container *StackUpgrade) error {
	return c.kuladoClient.doResourceDelete(STACK_UPGRADE_TYPE, &container.Resource)
}
