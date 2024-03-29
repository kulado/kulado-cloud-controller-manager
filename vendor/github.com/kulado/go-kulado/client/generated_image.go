package client

const (
	IMAGE_TYPE = "image"
)

type Image struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ImageCollection struct {
	Collection
	Data   []Image `json:"data,omitempty"`
	client *ImageClient
}

type ImageClient struct {
	kuladoClient *KuladoClient
}

type ImageOperations interface {
	List(opts *ListOpts) (*ImageCollection, error)
	Create(opts *Image) (*Image, error)
	Update(existing *Image, updates interface{}) (*Image, error)
	ById(id string) (*Image, error)
	Delete(container *Image) error

	ActionActivate(*Image) (*Image, error)

	ActionCreate(*Image) (*Image, error)

	ActionDeactivate(*Image) (*Image, error)

	ActionPurge(*Image) (*Image, error)

	ActionRemove(*Image) (*Image, error)

	ActionRestore(*Image) (*Image, error)

	ActionUpdate(*Image) (*Image, error)
}

func newImageClient(kuladoClient *KuladoClient) *ImageClient {
	return &ImageClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ImageClient) Create(container *Image) (*Image, error) {
	resp := &Image{}
	err := c.kuladoClient.doCreate(IMAGE_TYPE, container, resp)
	return resp, err
}

func (c *ImageClient) Update(existing *Image, updates interface{}) (*Image, error) {
	resp := &Image{}
	err := c.kuladoClient.doUpdate(IMAGE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ImageClient) List(opts *ListOpts) (*ImageCollection, error) {
	resp := &ImageCollection{}
	err := c.kuladoClient.doList(IMAGE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ImageCollection) Next() (*ImageCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ImageCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ImageClient) ById(id string) (*Image, error) {
	resp := &Image{}
	err := c.kuladoClient.doById(IMAGE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ImageClient) Delete(container *Image) error {
	return c.kuladoClient.doResourceDelete(IMAGE_TYPE, &container.Resource)
}

func (c *ImageClient) ActionActivate(resource *Image) (*Image, error) {

	resp := &Image{}

	err := c.kuladoClient.doAction(IMAGE_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ImageClient) ActionCreate(resource *Image) (*Image, error) {

	resp := &Image{}

	err := c.kuladoClient.doAction(IMAGE_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ImageClient) ActionDeactivate(resource *Image) (*Image, error) {

	resp := &Image{}

	err := c.kuladoClient.doAction(IMAGE_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ImageClient) ActionPurge(resource *Image) (*Image, error) {

	resp := &Image{}

	err := c.kuladoClient.doAction(IMAGE_TYPE, "purge", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ImageClient) ActionRemove(resource *Image) (*Image, error) {

	resp := &Image{}

	err := c.kuladoClient.doAction(IMAGE_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ImageClient) ActionRestore(resource *Image) (*Image, error) {

	resp := &Image{}

	err := c.kuladoClient.doAction(IMAGE_TYPE, "restore", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ImageClient) ActionUpdate(resource *Image) (*Image, error) {

	resp := &Image{}

	err := c.kuladoClient.doAction(IMAGE_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
