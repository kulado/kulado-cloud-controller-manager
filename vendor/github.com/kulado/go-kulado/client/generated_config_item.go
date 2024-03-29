package client

const (
	CONFIG_ITEM_TYPE = "configItem"
)

type ConfigItem struct {
	Resource

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	SourceVersion string `json:"sourceVersion,omitempty" yaml:"source_version,omitempty"`
}

type ConfigItemCollection struct {
	Collection
	Data   []ConfigItem `json:"data,omitempty"`
	client *ConfigItemClient
}

type ConfigItemClient struct {
	kuladoClient *KuladoClient
}

type ConfigItemOperations interface {
	List(opts *ListOpts) (*ConfigItemCollection, error)
	Create(opts *ConfigItem) (*ConfigItem, error)
	Update(existing *ConfigItem, updates interface{}) (*ConfigItem, error)
	ById(id string) (*ConfigItem, error)
	Delete(container *ConfigItem) error
}

func newConfigItemClient(kuladoClient *KuladoClient) *ConfigItemClient {
	return &ConfigItemClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ConfigItemClient) Create(container *ConfigItem) (*ConfigItem, error) {
	resp := &ConfigItem{}
	err := c.kuladoClient.doCreate(CONFIG_ITEM_TYPE, container, resp)
	return resp, err
}

func (c *ConfigItemClient) Update(existing *ConfigItem, updates interface{}) (*ConfigItem, error) {
	resp := &ConfigItem{}
	err := c.kuladoClient.doUpdate(CONFIG_ITEM_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ConfigItemClient) List(opts *ListOpts) (*ConfigItemCollection, error) {
	resp := &ConfigItemCollection{}
	err := c.kuladoClient.doList(CONFIG_ITEM_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ConfigItemCollection) Next() (*ConfigItemCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ConfigItemCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ConfigItemClient) ById(id string) (*ConfigItem, error) {
	resp := &ConfigItem{}
	err := c.kuladoClient.doById(CONFIG_ITEM_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ConfigItemClient) Delete(container *ConfigItem) error {
	return c.kuladoClient.doResourceDelete(CONFIG_ITEM_TYPE, &container.Resource)
}
