package client

const (
	ACTIVE_SETTING_TYPE = "activeSetting"
)

type ActiveSetting struct {
	Resource

	ActiveValue interface{} `json:"activeValue,omitempty" yaml:"active_value,omitempty"`

	InDb bool `json:"inDb,omitempty" yaml:"in_db,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

type ActiveSettingCollection struct {
	Collection
	Data   []ActiveSetting `json:"data,omitempty"`
	client *ActiveSettingClient
}

type ActiveSettingClient struct {
	kuladoClient *KuladoClient
}

type ActiveSettingOperations interface {
	List(opts *ListOpts) (*ActiveSettingCollection, error)
	Create(opts *ActiveSetting) (*ActiveSetting, error)
	Update(existing *ActiveSetting, updates interface{}) (*ActiveSetting, error)
	ById(id string) (*ActiveSetting, error)
	Delete(container *ActiveSetting) error
}

func newActiveSettingClient(kuladoClient *KuladoClient) *ActiveSettingClient {
	return &ActiveSettingClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ActiveSettingClient) Create(container *ActiveSetting) (*ActiveSetting, error) {
	resp := &ActiveSetting{}
	err := c.kuladoClient.doCreate(ACTIVE_SETTING_TYPE, container, resp)
	return resp, err
}

func (c *ActiveSettingClient) Update(existing *ActiveSetting, updates interface{}) (*ActiveSetting, error) {
	resp := &ActiveSetting{}
	err := c.kuladoClient.doUpdate(ACTIVE_SETTING_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ActiveSettingClient) List(opts *ListOpts) (*ActiveSettingCollection, error) {
	resp := &ActiveSettingCollection{}
	err := c.kuladoClient.doList(ACTIVE_SETTING_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ActiveSettingCollection) Next() (*ActiveSettingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ActiveSettingCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ActiveSettingClient) ById(id string) (*ActiveSetting, error) {
	resp := &ActiveSetting{}
	err := c.kuladoClient.doById(ACTIVE_SETTING_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ActiveSettingClient) Delete(container *ActiveSetting) error {
	return c.kuladoClient.doResourceDelete(ACTIVE_SETTING_TYPE, &container.Resource)
}
