package client

const (
	REVERT_TO_SNAPSHOT_INPUT_TYPE = "revertToSnapshotInput"
)

type RevertToSnapshotInput struct {
	Resource

	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshot_id,omitempty"`
}

type RevertToSnapshotInputCollection struct {
	Collection
	Data   []RevertToSnapshotInput `json:"data,omitempty"`
	client *RevertToSnapshotInputClient
}

type RevertToSnapshotInputClient struct {
	kuladoClient *KuladoClient
}

type RevertToSnapshotInputOperations interface {
	List(opts *ListOpts) (*RevertToSnapshotInputCollection, error)
	Create(opts *RevertToSnapshotInput) (*RevertToSnapshotInput, error)
	Update(existing *RevertToSnapshotInput, updates interface{}) (*RevertToSnapshotInput, error)
	ById(id string) (*RevertToSnapshotInput, error)
	Delete(container *RevertToSnapshotInput) error
}

func newRevertToSnapshotInputClient(kuladoClient *KuladoClient) *RevertToSnapshotInputClient {
	return &RevertToSnapshotInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *RevertToSnapshotInputClient) Create(container *RevertToSnapshotInput) (*RevertToSnapshotInput, error) {
	resp := &RevertToSnapshotInput{}
	err := c.kuladoClient.doCreate(REVERT_TO_SNAPSHOT_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *RevertToSnapshotInputClient) Update(existing *RevertToSnapshotInput, updates interface{}) (*RevertToSnapshotInput, error) {
	resp := &RevertToSnapshotInput{}
	err := c.kuladoClient.doUpdate(REVERT_TO_SNAPSHOT_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RevertToSnapshotInputClient) List(opts *ListOpts) (*RevertToSnapshotInputCollection, error) {
	resp := &RevertToSnapshotInputCollection{}
	err := c.kuladoClient.doList(REVERT_TO_SNAPSHOT_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *RevertToSnapshotInputCollection) Next() (*RevertToSnapshotInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &RevertToSnapshotInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *RevertToSnapshotInputClient) ById(id string) (*RevertToSnapshotInput, error) {
	resp := &RevertToSnapshotInput{}
	err := c.kuladoClient.doById(REVERT_TO_SNAPSHOT_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *RevertToSnapshotInputClient) Delete(container *RevertToSnapshotInput) error {
	return c.kuladoClient.doResourceDelete(REVERT_TO_SNAPSHOT_INPUT_TYPE, &container.Resource)
}
