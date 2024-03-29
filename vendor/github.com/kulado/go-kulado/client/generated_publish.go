package client

const (
	PUBLISH_TYPE = "publish"
)

type Publish struct {
	Resource

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	PreviousIds []string `json:"previousIds,omitempty" yaml:"previous_ids,omitempty"`

	Publisher string `json:"publisher,omitempty" yaml:"publisher,omitempty"`

	ResourceId string `json:"resourceId,omitempty" yaml:"resource_id,omitempty"`

	ResourceType string `json:"resourceType,omitempty" yaml:"resource_type,omitempty"`

	Time int64 `json:"time,omitempty" yaml:"time,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningInternalMessage string `json:"transitioningInternalMessage,omitempty" yaml:"transitioning_internal_message,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`
}

type PublishCollection struct {
	Collection
	Data   []Publish `json:"data,omitempty"`
	client *PublishClient
}

type PublishClient struct {
	kuladoClient *KuladoClient
}

type PublishOperations interface {
	List(opts *ListOpts) (*PublishCollection, error)
	Create(opts *Publish) (*Publish, error)
	Update(existing *Publish, updates interface{}) (*Publish, error)
	ById(id string) (*Publish, error)
	Delete(container *Publish) error
}

func newPublishClient(kuladoClient *KuladoClient) *PublishClient {
	return &PublishClient{
		kuladoClient: kuladoClient,
	}
}

func (c *PublishClient) Create(container *Publish) (*Publish, error) {
	resp := &Publish{}
	err := c.kuladoClient.doCreate(PUBLISH_TYPE, container, resp)
	return resp, err
}

func (c *PublishClient) Update(existing *Publish, updates interface{}) (*Publish, error) {
	resp := &Publish{}
	err := c.kuladoClient.doUpdate(PUBLISH_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PublishClient) List(opts *ListOpts) (*PublishCollection, error) {
	resp := &PublishCollection{}
	err := c.kuladoClient.doList(PUBLISH_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PublishCollection) Next() (*PublishCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PublishCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PublishClient) ById(id string) (*Publish, error) {
	resp := &Publish{}
	err := c.kuladoClient.doById(PUBLISH_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *PublishClient) Delete(container *Publish) error {
	return c.kuladoClient.doResourceDelete(PUBLISH_TYPE, &container.Resource)
}
