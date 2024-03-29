package client

const (
	SCALE_POLICY_TYPE = "scalePolicy"
)

type ScalePolicy struct {
	Resource

	Increment int64 `json:"increment,omitempty" yaml:"increment,omitempty"`

	Max int64 `json:"max,omitempty" yaml:"max,omitempty"`

	Min int64 `json:"min,omitempty" yaml:"min,omitempty"`
}

type ScalePolicyCollection struct {
	Collection
	Data   []ScalePolicy `json:"data,omitempty"`
	client *ScalePolicyClient
}

type ScalePolicyClient struct {
	kuladoClient *KuladoClient
}

type ScalePolicyOperations interface {
	List(opts *ListOpts) (*ScalePolicyCollection, error)
	Create(opts *ScalePolicy) (*ScalePolicy, error)
	Update(existing *ScalePolicy, updates interface{}) (*ScalePolicy, error)
	ById(id string) (*ScalePolicy, error)
	Delete(container *ScalePolicy) error
}

func newScalePolicyClient(kuladoClient *KuladoClient) *ScalePolicyClient {
	return &ScalePolicyClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ScalePolicyClient) Create(container *ScalePolicy) (*ScalePolicy, error) {
	resp := &ScalePolicy{}
	err := c.kuladoClient.doCreate(SCALE_POLICY_TYPE, container, resp)
	return resp, err
}

func (c *ScalePolicyClient) Update(existing *ScalePolicy, updates interface{}) (*ScalePolicy, error) {
	resp := &ScalePolicy{}
	err := c.kuladoClient.doUpdate(SCALE_POLICY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ScalePolicyClient) List(opts *ListOpts) (*ScalePolicyCollection, error) {
	resp := &ScalePolicyCollection{}
	err := c.kuladoClient.doList(SCALE_POLICY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ScalePolicyCollection) Next() (*ScalePolicyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ScalePolicyCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ScalePolicyClient) ById(id string) (*ScalePolicy, error) {
	resp := &ScalePolicy{}
	err := c.kuladoClient.doById(SCALE_POLICY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ScalePolicyClient) Delete(container *ScalePolicy) error {
	return c.kuladoClient.doResourceDelete(SCALE_POLICY_TYPE, &container.Resource)
}
