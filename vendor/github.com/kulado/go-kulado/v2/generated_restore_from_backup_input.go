package client

const (
	RESTORE_FROM_BACKUP_INPUT_TYPE = "restoreFromBackupInput"
)

type RestoreFromBackupInput struct {
	Resource

	BackupId string `json:"backupId,omitempty" yaml:"backup_id,omitempty"`
}

type RestoreFromBackupInputCollection struct {
	Collection
	Data   []RestoreFromBackupInput `json:"data,omitempty"`
	client *RestoreFromBackupInputClient
}

type RestoreFromBackupInputClient struct {
	kuladoClient *KuladoClient
}

type RestoreFromBackupInputOperations interface {
	List(opts *ListOpts) (*RestoreFromBackupInputCollection, error)
	Create(opts *RestoreFromBackupInput) (*RestoreFromBackupInput, error)
	Update(existing *RestoreFromBackupInput, updates interface{}) (*RestoreFromBackupInput, error)
	ById(id string) (*RestoreFromBackupInput, error)
	Delete(container *RestoreFromBackupInput) error
}

func newRestoreFromBackupInputClient(kuladoClient *KuladoClient) *RestoreFromBackupInputClient {
	return &RestoreFromBackupInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *RestoreFromBackupInputClient) Create(container *RestoreFromBackupInput) (*RestoreFromBackupInput, error) {
	resp := &RestoreFromBackupInput{}
	err := c.kuladoClient.doCreate(RESTORE_FROM_BACKUP_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *RestoreFromBackupInputClient) Update(existing *RestoreFromBackupInput, updates interface{}) (*RestoreFromBackupInput, error) {
	resp := &RestoreFromBackupInput{}
	err := c.kuladoClient.doUpdate(RESTORE_FROM_BACKUP_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RestoreFromBackupInputClient) List(opts *ListOpts) (*RestoreFromBackupInputCollection, error) {
	resp := &RestoreFromBackupInputCollection{}
	err := c.kuladoClient.doList(RESTORE_FROM_BACKUP_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *RestoreFromBackupInputCollection) Next() (*RestoreFromBackupInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &RestoreFromBackupInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *RestoreFromBackupInputClient) ById(id string) (*RestoreFromBackupInput, error) {
	resp := &RestoreFromBackupInput{}
	err := c.kuladoClient.doById(RESTORE_FROM_BACKUP_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *RestoreFromBackupInputClient) Delete(container *RestoreFromBackupInput) error {
	return c.kuladoClient.doResourceDelete(RESTORE_FROM_BACKUP_INPUT_TYPE, &container.Resource)
}
