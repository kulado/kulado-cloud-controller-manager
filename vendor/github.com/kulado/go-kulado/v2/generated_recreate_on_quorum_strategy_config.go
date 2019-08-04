package client

const (
	RECREATE_ON_QUORUM_STRATEGY_CONFIG_TYPE = "recreateOnQuorumStrategyConfig"
)

type RecreateOnQuorumStrategyConfig struct {
	Resource

	Quorum int64 `json:"quorum,omitempty" yaml:"quorum,omitempty"`
}

type RecreateOnQuorumStrategyConfigCollection struct {
	Collection
	Data   []RecreateOnQuorumStrategyConfig `json:"data,omitempty"`
	client *RecreateOnQuorumStrategyConfigClient
}

type RecreateOnQuorumStrategyConfigClient struct {
	kuladoClient *KuladoClient
}

type RecreateOnQuorumStrategyConfigOperations interface {
	List(opts *ListOpts) (*RecreateOnQuorumStrategyConfigCollection, error)
	Create(opts *RecreateOnQuorumStrategyConfig) (*RecreateOnQuorumStrategyConfig, error)
	Update(existing *RecreateOnQuorumStrategyConfig, updates interface{}) (*RecreateOnQuorumStrategyConfig, error)
	ById(id string) (*RecreateOnQuorumStrategyConfig, error)
	Delete(container *RecreateOnQuorumStrategyConfig) error
}

func newRecreateOnQuorumStrategyConfigClient(kuladoClient *KuladoClient) *RecreateOnQuorumStrategyConfigClient {
	return &RecreateOnQuorumStrategyConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *RecreateOnQuorumStrategyConfigClient) Create(container *RecreateOnQuorumStrategyConfig) (*RecreateOnQuorumStrategyConfig, error) {
	resp := &RecreateOnQuorumStrategyConfig{}
	err := c.kuladoClient.doCreate(RECREATE_ON_QUORUM_STRATEGY_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *RecreateOnQuorumStrategyConfigClient) Update(existing *RecreateOnQuorumStrategyConfig, updates interface{}) (*RecreateOnQuorumStrategyConfig, error) {
	resp := &RecreateOnQuorumStrategyConfig{}
	err := c.kuladoClient.doUpdate(RECREATE_ON_QUORUM_STRATEGY_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RecreateOnQuorumStrategyConfigClient) List(opts *ListOpts) (*RecreateOnQuorumStrategyConfigCollection, error) {
	resp := &RecreateOnQuorumStrategyConfigCollection{}
	err := c.kuladoClient.doList(RECREATE_ON_QUORUM_STRATEGY_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *RecreateOnQuorumStrategyConfigCollection) Next() (*RecreateOnQuorumStrategyConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &RecreateOnQuorumStrategyConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *RecreateOnQuorumStrategyConfigClient) ById(id string) (*RecreateOnQuorumStrategyConfig, error) {
	resp := &RecreateOnQuorumStrategyConfig{}
	err := c.kuladoClient.doById(RECREATE_ON_QUORUM_STRATEGY_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *RecreateOnQuorumStrategyConfigClient) Delete(container *RecreateOnQuorumStrategyConfig) error {
	return c.kuladoClient.doResourceDelete(RECREATE_ON_QUORUM_STRATEGY_CONFIG_TYPE, &container.Resource)
}
