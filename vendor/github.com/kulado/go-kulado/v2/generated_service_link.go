package client

const (
	SERVICE_LINK_TYPE = "serviceLink"
)

type ServiceLink struct {
	Resource

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	ServiceId string `json:"serviceId,omitempty" yaml:"service_id,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ServiceLinkCollection struct {
	Collection
	Data   []ServiceLink `json:"data,omitempty"`
	client *ServiceLinkClient
}

type ServiceLinkClient struct {
	kuladoClient *KuladoClient
}

type ServiceLinkOperations interface {
	List(opts *ListOpts) (*ServiceLinkCollection, error)
	Create(opts *ServiceLink) (*ServiceLink, error)
	Update(existing *ServiceLink, updates interface{}) (*ServiceLink, error)
	ById(id string) (*ServiceLink, error)
	Delete(container *ServiceLink) error
}

func newServiceLinkClient(kuladoClient *KuladoClient) *ServiceLinkClient {
	return &ServiceLinkClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ServiceLinkClient) Create(container *ServiceLink) (*ServiceLink, error) {
	resp := &ServiceLink{}
	err := c.kuladoClient.doCreate(SERVICE_LINK_TYPE, container, resp)
	return resp, err
}

func (c *ServiceLinkClient) Update(existing *ServiceLink, updates interface{}) (*ServiceLink, error) {
	resp := &ServiceLink{}
	err := c.kuladoClient.doUpdate(SERVICE_LINK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceLinkClient) List(opts *ListOpts) (*ServiceLinkCollection, error) {
	resp := &ServiceLinkCollection{}
	err := c.kuladoClient.doList(SERVICE_LINK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceLinkCollection) Next() (*ServiceLinkCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceLinkCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceLinkClient) ById(id string) (*ServiceLink, error) {
	resp := &ServiceLink{}
	err := c.kuladoClient.doById(SERVICE_LINK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ServiceLinkClient) Delete(container *ServiceLink) error {
	return c.kuladoClient.doResourceDelete(SERVICE_LINK_TYPE, &container.Resource)
}
