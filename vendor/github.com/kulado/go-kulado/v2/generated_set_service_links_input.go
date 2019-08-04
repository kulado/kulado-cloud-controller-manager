package client

const (
	SET_SERVICE_LINKS_INPUT_TYPE = "setServiceLinksInput"
)

type SetServiceLinksInput struct {
	Resource

	ServiceLinks []ServiceLink `json:"serviceLinks,omitempty" yaml:"service_links,omitempty"`
}

type SetServiceLinksInputCollection struct {
	Collection
	Data   []SetServiceLinksInput `json:"data,omitempty"`
	client *SetServiceLinksInputClient
}

type SetServiceLinksInputClient struct {
	kuladoClient *KuladoClient
}

type SetServiceLinksInputOperations interface {
	List(opts *ListOpts) (*SetServiceLinksInputCollection, error)
	Create(opts *SetServiceLinksInput) (*SetServiceLinksInput, error)
	Update(existing *SetServiceLinksInput, updates interface{}) (*SetServiceLinksInput, error)
	ById(id string) (*SetServiceLinksInput, error)
	Delete(container *SetServiceLinksInput) error
}

func newSetServiceLinksInputClient(kuladoClient *KuladoClient) *SetServiceLinksInputClient {
	return &SetServiceLinksInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *SetServiceLinksInputClient) Create(container *SetServiceLinksInput) (*SetServiceLinksInput, error) {
	resp := &SetServiceLinksInput{}
	err := c.kuladoClient.doCreate(SET_SERVICE_LINKS_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *SetServiceLinksInputClient) Update(existing *SetServiceLinksInput, updates interface{}) (*SetServiceLinksInput, error) {
	resp := &SetServiceLinksInput{}
	err := c.kuladoClient.doUpdate(SET_SERVICE_LINKS_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SetServiceLinksInputClient) List(opts *ListOpts) (*SetServiceLinksInputCollection, error) {
	resp := &SetServiceLinksInputCollection{}
	err := c.kuladoClient.doList(SET_SERVICE_LINKS_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SetServiceLinksInputCollection) Next() (*SetServiceLinksInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SetServiceLinksInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SetServiceLinksInputClient) ById(id string) (*SetServiceLinksInput, error) {
	resp := &SetServiceLinksInput{}
	err := c.kuladoClient.doById(SET_SERVICE_LINKS_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SetServiceLinksInputClient) Delete(container *SetServiceLinksInput) error {
	return c.kuladoClient.doResourceDelete(SET_SERVICE_LINKS_INPUT_TYPE, &container.Resource)
}
