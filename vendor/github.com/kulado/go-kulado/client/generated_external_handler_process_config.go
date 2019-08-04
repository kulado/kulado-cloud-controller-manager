package client

const (
	EXTERNAL_HANDLER_PROCESS_CONFIG_TYPE = "externalHandlerProcessConfig"
)

type ExternalHandlerProcessConfig struct {
	Resource

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	OnError string `json:"onError,omitempty" yaml:"on_error,omitempty"`
}

type ExternalHandlerProcessConfigCollection struct {
	Collection
	Data   []ExternalHandlerProcessConfig `json:"data,omitempty"`
	client *ExternalHandlerProcessConfigClient
}

type ExternalHandlerProcessConfigClient struct {
	kuladoClient *KuladoClient
}

type ExternalHandlerProcessConfigOperations interface {
	List(opts *ListOpts) (*ExternalHandlerProcessConfigCollection, error)
	Create(opts *ExternalHandlerProcessConfig) (*ExternalHandlerProcessConfig, error)
	Update(existing *ExternalHandlerProcessConfig, updates interface{}) (*ExternalHandlerProcessConfig, error)
	ById(id string) (*ExternalHandlerProcessConfig, error)
	Delete(container *ExternalHandlerProcessConfig) error
}

func newExternalHandlerProcessConfigClient(kuladoClient *KuladoClient) *ExternalHandlerProcessConfigClient {
	return &ExternalHandlerProcessConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExternalHandlerProcessConfigClient) Create(container *ExternalHandlerProcessConfig) (*ExternalHandlerProcessConfig, error) {
	resp := &ExternalHandlerProcessConfig{}
	err := c.kuladoClient.doCreate(EXTERNAL_HANDLER_PROCESS_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *ExternalHandlerProcessConfigClient) Update(existing *ExternalHandlerProcessConfig, updates interface{}) (*ExternalHandlerProcessConfig, error) {
	resp := &ExternalHandlerProcessConfig{}
	err := c.kuladoClient.doUpdate(EXTERNAL_HANDLER_PROCESS_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalHandlerProcessConfigClient) List(opts *ListOpts) (*ExternalHandlerProcessConfigCollection, error) {
	resp := &ExternalHandlerProcessConfigCollection{}
	err := c.kuladoClient.doList(EXTERNAL_HANDLER_PROCESS_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExternalHandlerProcessConfigCollection) Next() (*ExternalHandlerProcessConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalHandlerProcessConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalHandlerProcessConfigClient) ById(id string) (*ExternalHandlerProcessConfig, error) {
	resp := &ExternalHandlerProcessConfig{}
	err := c.kuladoClient.doById(EXTERNAL_HANDLER_PROCESS_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExternalHandlerProcessConfigClient) Delete(container *ExternalHandlerProcessConfig) error {
	return c.kuladoClient.doResourceDelete(EXTERNAL_HANDLER_PROCESS_CONFIG_TYPE, &container.Resource)
}
