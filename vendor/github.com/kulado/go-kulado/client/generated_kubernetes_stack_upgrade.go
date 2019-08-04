package client

const (
	KUBERNETES_STACK_UPGRADE_TYPE = "kubernetesStackUpgrade"
)

type KubernetesStackUpgrade struct {
	Resource

	Environment map[string]interface{} `json:"environment,omitempty" yaml:"environment,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	Templates map[string]interface{} `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type KubernetesStackUpgradeCollection struct {
	Collection
	Data   []KubernetesStackUpgrade `json:"data,omitempty"`
	client *KubernetesStackUpgradeClient
}

type KubernetesStackUpgradeClient struct {
	kuladoClient *KuladoClient
}

type KubernetesStackUpgradeOperations interface {
	List(opts *ListOpts) (*KubernetesStackUpgradeCollection, error)
	Create(opts *KubernetesStackUpgrade) (*KubernetesStackUpgrade, error)
	Update(existing *KubernetesStackUpgrade, updates interface{}) (*KubernetesStackUpgrade, error)
	ById(id string) (*KubernetesStackUpgrade, error)
	Delete(container *KubernetesStackUpgrade) error
}

func newKubernetesStackUpgradeClient(kuladoClient *KuladoClient) *KubernetesStackUpgradeClient {
	return &KubernetesStackUpgradeClient{
		kuladoClient: kuladoClient,
	}
}

func (c *KubernetesStackUpgradeClient) Create(container *KubernetesStackUpgrade) (*KubernetesStackUpgrade, error) {
	resp := &KubernetesStackUpgrade{}
	err := c.kuladoClient.doCreate(KUBERNETES_STACK_UPGRADE_TYPE, container, resp)
	return resp, err
}

func (c *KubernetesStackUpgradeClient) Update(existing *KubernetesStackUpgrade, updates interface{}) (*KubernetesStackUpgrade, error) {
	resp := &KubernetesStackUpgrade{}
	err := c.kuladoClient.doUpdate(KUBERNETES_STACK_UPGRADE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *KubernetesStackUpgradeClient) List(opts *ListOpts) (*KubernetesStackUpgradeCollection, error) {
	resp := &KubernetesStackUpgradeCollection{}
	err := c.kuladoClient.doList(KUBERNETES_STACK_UPGRADE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *KubernetesStackUpgradeCollection) Next() (*KubernetesStackUpgradeCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &KubernetesStackUpgradeCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *KubernetesStackUpgradeClient) ById(id string) (*KubernetesStackUpgrade, error) {
	resp := &KubernetesStackUpgrade{}
	err := c.kuladoClient.doById(KUBERNETES_STACK_UPGRADE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *KubernetesStackUpgradeClient) Delete(container *KubernetesStackUpgrade) error {
	return c.kuladoClient.doResourceDelete(KUBERNETES_STACK_UPGRADE_TYPE, &container.Resource)
}
