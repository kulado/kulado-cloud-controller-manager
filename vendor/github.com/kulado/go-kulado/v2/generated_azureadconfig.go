package client

const (
	AZUREADCONFIG_TYPE = "azureadconfig"
)

type Azureadconfig struct {
	Resource

	AccessMode string `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`

	AdminAccountPassword string `json:"adminAccountPassword,omitempty" yaml:"admin_account_password,omitempty"`

	AdminAccountUsername string `json:"adminAccountUsername,omitempty" yaml:"admin_account_username,omitempty"`

	ClientId string `json:"clientId,omitempty" yaml:"client_id,omitempty"`

	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	TenantId string `json:"tenantId,omitempty" yaml:"tenant_id,omitempty"`
}

type AzureadconfigCollection struct {
	Collection
	Data   []Azureadconfig `json:"data,omitempty"`
	client *AzureadconfigClient
}

type AzureadconfigClient struct {
	kuladoClient *KuladoClient
}

type AzureadconfigOperations interface {
	List(opts *ListOpts) (*AzureadconfigCollection, error)
	Create(opts *Azureadconfig) (*Azureadconfig, error)
	Update(existing *Azureadconfig, updates interface{}) (*Azureadconfig, error)
	ById(id string) (*Azureadconfig, error)
	Delete(container *Azureadconfig) error
}

func newAzureadconfigClient(kuladoClient *KuladoClient) *AzureadconfigClient {
	return &AzureadconfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *AzureadconfigClient) Create(container *Azureadconfig) (*Azureadconfig, error) {
	resp := &Azureadconfig{}
	err := c.kuladoClient.doCreate(AZUREADCONFIG_TYPE, container, resp)
	return resp, err
}

func (c *AzureadconfigClient) Update(existing *Azureadconfig, updates interface{}) (*Azureadconfig, error) {
	resp := &Azureadconfig{}
	err := c.kuladoClient.doUpdate(AZUREADCONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AzureadconfigClient) List(opts *ListOpts) (*AzureadconfigCollection, error) {
	resp := &AzureadconfigCollection{}
	err := c.kuladoClient.doList(AZUREADCONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *AzureadconfigCollection) Next() (*AzureadconfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AzureadconfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AzureadconfigClient) ById(id string) (*Azureadconfig, error) {
	resp := &Azureadconfig{}
	err := c.kuladoClient.doById(AZUREADCONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *AzureadconfigClient) Delete(container *Azureadconfig) error {
	return c.kuladoClient.doResourceDelete(AZUREADCONFIG_TYPE, &container.Resource)
}
