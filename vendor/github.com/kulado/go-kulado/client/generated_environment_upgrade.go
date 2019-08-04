package client

const (
	ENVIRONMENT_UPGRADE_TYPE = "environmentUpgrade"
)

type EnvironmentUpgrade struct {
	Resource

	DockerCompose string `json:"dockerCompose,omitempty" yaml:"docker_compose,omitempty"`

	Environment map[string]interface{} `json:"environment,omitempty" yaml:"environment,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	KuladoCompose string `json:"kuladoCompose,omitempty" yaml:"kulado_compose,omitempty"`
}

type EnvironmentUpgradeCollection struct {
	Collection
	Data   []EnvironmentUpgrade `json:"data,omitempty"`
	client *EnvironmentUpgradeClient
}

type EnvironmentUpgradeClient struct {
	kuladoClient *KuladoClient
}

type EnvironmentUpgradeOperations interface {
	List(opts *ListOpts) (*EnvironmentUpgradeCollection, error)
	Create(opts *EnvironmentUpgrade) (*EnvironmentUpgrade, error)
	Update(existing *EnvironmentUpgrade, updates interface{}) (*EnvironmentUpgrade, error)
	ById(id string) (*EnvironmentUpgrade, error)
	Delete(container *EnvironmentUpgrade) error
}

func newEnvironmentUpgradeClient(kuladoClient *KuladoClient) *EnvironmentUpgradeClient {
	return &EnvironmentUpgradeClient{
		kuladoClient: kuladoClient,
	}
}

func (c *EnvironmentUpgradeClient) Create(container *EnvironmentUpgrade) (*EnvironmentUpgrade, error) {
	resp := &EnvironmentUpgrade{}
	err := c.kuladoClient.doCreate(ENVIRONMENT_UPGRADE_TYPE, container, resp)
	return resp, err
}

func (c *EnvironmentUpgradeClient) Update(existing *EnvironmentUpgrade, updates interface{}) (*EnvironmentUpgrade, error) {
	resp := &EnvironmentUpgrade{}
	err := c.kuladoClient.doUpdate(ENVIRONMENT_UPGRADE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *EnvironmentUpgradeClient) List(opts *ListOpts) (*EnvironmentUpgradeCollection, error) {
	resp := &EnvironmentUpgradeCollection{}
	err := c.kuladoClient.doList(ENVIRONMENT_UPGRADE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *EnvironmentUpgradeCollection) Next() (*EnvironmentUpgradeCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &EnvironmentUpgradeCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *EnvironmentUpgradeClient) ById(id string) (*EnvironmentUpgrade, error) {
	resp := &EnvironmentUpgrade{}
	err := c.kuladoClient.doById(ENVIRONMENT_UPGRADE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *EnvironmentUpgradeClient) Delete(container *EnvironmentUpgrade) error {
	return c.kuladoClient.doResourceDelete(ENVIRONMENT_UPGRADE_TYPE, &container.Resource)
}
