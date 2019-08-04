package client

const (
	VOLUME_SNAPSHOT_INPUT_TYPE = "volumeSnapshotInput"
)

type VolumeSnapshotInput struct {
	Resource

	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type VolumeSnapshotInputCollection struct {
	Collection
	Data   []VolumeSnapshotInput `json:"data,omitempty"`
	client *VolumeSnapshotInputClient
}

type VolumeSnapshotInputClient struct {
	kuladoClient *KuladoClient
}

type VolumeSnapshotInputOperations interface {
	List(opts *ListOpts) (*VolumeSnapshotInputCollection, error)
	Create(opts *VolumeSnapshotInput) (*VolumeSnapshotInput, error)
	Update(existing *VolumeSnapshotInput, updates interface{}) (*VolumeSnapshotInput, error)
	ById(id string) (*VolumeSnapshotInput, error)
	Delete(container *VolumeSnapshotInput) error
}

func newVolumeSnapshotInputClient(kuladoClient *KuladoClient) *VolumeSnapshotInputClient {
	return &VolumeSnapshotInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *VolumeSnapshotInputClient) Create(container *VolumeSnapshotInput) (*VolumeSnapshotInput, error) {
	resp := &VolumeSnapshotInput{}
	err := c.kuladoClient.doCreate(VOLUME_SNAPSHOT_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *VolumeSnapshotInputClient) Update(existing *VolumeSnapshotInput, updates interface{}) (*VolumeSnapshotInput, error) {
	resp := &VolumeSnapshotInput{}
	err := c.kuladoClient.doUpdate(VOLUME_SNAPSHOT_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VolumeSnapshotInputClient) List(opts *ListOpts) (*VolumeSnapshotInputCollection, error) {
	resp := &VolumeSnapshotInputCollection{}
	err := c.kuladoClient.doList(VOLUME_SNAPSHOT_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *VolumeSnapshotInputCollection) Next() (*VolumeSnapshotInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &VolumeSnapshotInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *VolumeSnapshotInputClient) ById(id string) (*VolumeSnapshotInput, error) {
	resp := &VolumeSnapshotInput{}
	err := c.kuladoClient.doById(VOLUME_SNAPSHOT_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *VolumeSnapshotInputClient) Delete(container *VolumeSnapshotInput) error {
	return c.kuladoClient.doResourceDelete(VOLUME_SNAPSHOT_INPUT_TYPE, &container.Resource)
}
