package client

const (
	SET_PROJECT_MEMBERS_INPUT_TYPE = "setProjectMembersInput"
)

type SetProjectMembersInput struct {
	Resource

	Members []interface{} `json:"members,omitempty" yaml:"members,omitempty"`
}

type SetProjectMembersInputCollection struct {
	Collection
	Data   []SetProjectMembersInput `json:"data,omitempty"`
	client *SetProjectMembersInputClient
}

type SetProjectMembersInputClient struct {
	kuladoClient *KuladoClient
}

type SetProjectMembersInputOperations interface {
	List(opts *ListOpts) (*SetProjectMembersInputCollection, error)
	Create(opts *SetProjectMembersInput) (*SetProjectMembersInput, error)
	Update(existing *SetProjectMembersInput, updates interface{}) (*SetProjectMembersInput, error)
	ById(id string) (*SetProjectMembersInput, error)
	Delete(container *SetProjectMembersInput) error
}

func newSetProjectMembersInputClient(kuladoClient *KuladoClient) *SetProjectMembersInputClient {
	return &SetProjectMembersInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *SetProjectMembersInputClient) Create(container *SetProjectMembersInput) (*SetProjectMembersInput, error) {
	resp := &SetProjectMembersInput{}
	err := c.kuladoClient.doCreate(SET_PROJECT_MEMBERS_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *SetProjectMembersInputClient) Update(existing *SetProjectMembersInput, updates interface{}) (*SetProjectMembersInput, error) {
	resp := &SetProjectMembersInput{}
	err := c.kuladoClient.doUpdate(SET_PROJECT_MEMBERS_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SetProjectMembersInputClient) List(opts *ListOpts) (*SetProjectMembersInputCollection, error) {
	resp := &SetProjectMembersInputCollection{}
	err := c.kuladoClient.doList(SET_PROJECT_MEMBERS_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SetProjectMembersInputCollection) Next() (*SetProjectMembersInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SetProjectMembersInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SetProjectMembersInputClient) ById(id string) (*SetProjectMembersInput, error) {
	resp := &SetProjectMembersInput{}
	err := c.kuladoClient.doById(SET_PROJECT_MEMBERS_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SetProjectMembersInputClient) Delete(container *SetProjectMembersInput) error {
	return c.kuladoClient.doResourceDelete(SET_PROJECT_MEMBERS_INPUT_TYPE, &container.Resource)
}
