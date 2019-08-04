package client

const (
	HOST_API_PROXY_TOKEN_TYPE = "hostApiProxyToken"
)

type HostApiProxyToken struct {
	Resource

	ReportedUuid string `json:"reportedUuid,omitempty" yaml:"reported_uuid,omitempty"`

	Token string `json:"token,omitempty" yaml:"token,omitempty"`

	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}

type HostApiProxyTokenCollection struct {
	Collection
	Data   []HostApiProxyToken `json:"data,omitempty"`
	client *HostApiProxyTokenClient
}

type HostApiProxyTokenClient struct {
	kuladoClient *KuladoClient
}

type HostApiProxyTokenOperations interface {
	List(opts *ListOpts) (*HostApiProxyTokenCollection, error)
	Create(opts *HostApiProxyToken) (*HostApiProxyToken, error)
	Update(existing *HostApiProxyToken, updates interface{}) (*HostApiProxyToken, error)
	ById(id string) (*HostApiProxyToken, error)
	Delete(container *HostApiProxyToken) error
}

func newHostApiProxyTokenClient(kuladoClient *KuladoClient) *HostApiProxyTokenClient {
	return &HostApiProxyTokenClient{
		kuladoClient: kuladoClient,
	}
}

func (c *HostApiProxyTokenClient) Create(container *HostApiProxyToken) (*HostApiProxyToken, error) {
	resp := &HostApiProxyToken{}
	err := c.kuladoClient.doCreate(HOST_API_PROXY_TOKEN_TYPE, container, resp)
	return resp, err
}

func (c *HostApiProxyTokenClient) Update(existing *HostApiProxyToken, updates interface{}) (*HostApiProxyToken, error) {
	resp := &HostApiProxyToken{}
	err := c.kuladoClient.doUpdate(HOST_API_PROXY_TOKEN_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HostApiProxyTokenClient) List(opts *ListOpts) (*HostApiProxyTokenCollection, error) {
	resp := &HostApiProxyTokenCollection{}
	err := c.kuladoClient.doList(HOST_API_PROXY_TOKEN_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HostApiProxyTokenCollection) Next() (*HostApiProxyTokenCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HostApiProxyTokenCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HostApiProxyTokenClient) ById(id string) (*HostApiProxyToken, error) {
	resp := &HostApiProxyToken{}
	err := c.kuladoClient.doById(HOST_API_PROXY_TOKEN_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *HostApiProxyTokenClient) Delete(container *HostApiProxyToken) error {
	return c.kuladoClient.doResourceDelete(HOST_API_PROXY_TOKEN_TYPE, &container.Resource)
}
