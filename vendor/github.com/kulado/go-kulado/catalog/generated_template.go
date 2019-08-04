package catalog

const (
	TEMPLATE_TYPE = "template"
)

type Template struct {
	Resource

	Actions map[string]interface{} `json:"actions,omitempty" yaml:"actions,omitempty"`

	Bindings map[string]interface{} `json:"bindings,omitempty" yaml:"bindings,omitempty"`

	CatalogId string `json:"catalogId,omitempty" yaml:"catalog_id,omitempty"`

	Category string `json:"category,omitempty" yaml:"category,omitempty"`

	DefaultVersion string `json:"defaultVersion,omitempty" yaml:"default_version,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Files map[string]interface{} `json:"files,omitempty" yaml:"files,omitempty"`

	IsSystem string `json:"isSystem,omitempty" yaml:"is_system,omitempty"`

	Labels map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`

	License string `json:"license,omitempty" yaml:"license,omitempty"`

	Links map[string]interface{} `json:"links,omitempty" yaml:"links,omitempty"`

	Maintainer string `json:"maintainer,omitempty" yaml:"maintainer,omitempty"`

	MaximumKuladoVersion string `json:"maximumKuladoVersion,omitempty" yaml:"maximum_kulado_version,omitempty"`

	MinimumKuladoVersion string `json:"minimumKuladoVersion,omitempty" yaml:"minimum_kulado_version,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	TemplateBase string `json:"templateBase,omitempty" yaml:"template_base,omitempty"`

	TemplateVersionKuladoVersion map[string]interface{} `json:"templateVersionKuladoVersion,omitempty" yaml:"template_version_kulado_version,omitempty"`

	TemplateVersionKuladoVersionGte map[string]interface{} `json:"templateVersionKuladoVersionGte,omitempty" yaml:"template_version_kulado_version_gte,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	UpgradeFrom string `json:"upgradeFrom,omitempty" yaml:"upgrade_from,omitempty"`

	UpgradeVersionLinks map[string]interface{} `json:"upgradeVersionLinks,omitempty" yaml:"upgrade_version_links,omitempty"`

	VersionLinks map[string]interface{} `json:"versionLinks,omitempty" yaml:"version_links,omitempty"`
}

type TemplateCollection struct {
	Collection
	Data   []Template `json:"data,omitempty"`
	client *TemplateClient
}

type TemplateClient struct {
	kuladoClient *KuladoClient
}

type TemplateOperations interface {
	List(opts *ListOpts) (*TemplateCollection, error)
	Create(opts *Template) (*Template, error)
	Update(existing *Template, updates interface{}) (*Template, error)
	ById(id string) (*Template, error)
	Delete(container *Template) error
}

func newTemplateClient(kuladoClient *KuladoClient) *TemplateClient {
	return &TemplateClient{
		kuladoClient: kuladoClient,
	}
}

func (c *TemplateClient) Create(container *Template) (*Template, error) {
	resp := &Template{}
	err := c.kuladoClient.doCreate(TEMPLATE_TYPE, container, resp)
	return resp, err
}

func (c *TemplateClient) Update(existing *Template, updates interface{}) (*Template, error) {
	resp := &Template{}
	err := c.kuladoClient.doUpdate(TEMPLATE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TemplateClient) List(opts *ListOpts) (*TemplateCollection, error) {
	resp := &TemplateCollection{}
	err := c.kuladoClient.doList(TEMPLATE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *TemplateCollection) Next() (*TemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TemplateCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TemplateClient) ById(id string) (*Template, error) {
	resp := &Template{}
	err := c.kuladoClient.doById(TEMPLATE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *TemplateClient) Delete(container *Template) error {
	return c.kuladoClient.doResourceDelete(TEMPLATE_TYPE, &container.Resource)
}
