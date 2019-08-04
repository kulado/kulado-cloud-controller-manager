package catalog

const (
	TEMPLATE_VERSION_TYPE = "templateVersion"
)

type TemplateVersion struct {
	Resource

	Actions map[string]interface{} `json:"actions,omitempty" yaml:"actions,omitempty"`

	Bindings map[string]interface{} `json:"bindings,omitempty" yaml:"bindings,omitempty"`

	CatalogId string `json:"catalogId,omitempty" yaml:"catalog_id,omitempty"`

	Category string `json:"category,omitempty" yaml:"category,omitempty"`

	DefaultVersion string `json:"defaultVersion,omitempty" yaml:"default_version,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Files map[string]interface{} `json:"files,omitempty" yaml:"files,omitempty"`

	IconLink string `json:"iconLink,omitempty" yaml:"icon_link,omitempty"`

	IsSystem string `json:"isSystem,omitempty" yaml:"is_system,omitempty"`

	Labels map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`

	License string `json:"license,omitempty" yaml:"license,omitempty"`

	Links map[string]interface{} `json:"links,omitempty" yaml:"links,omitempty"`

	Maintainer string `json:"maintainer,omitempty" yaml:"maintainer,omitempty"`

	MaximumKuladoVersion string `json:"maximumKuladoVersion,omitempty" yaml:"maximum_kulado_version,omitempty"`

	MinimumKuladoVersion string `json:"minimumKuladoVersion,omitempty" yaml:"minimum_kulado_version,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	ProjectURL string `json:"projectURL,omitempty" yaml:"project_url,omitempty"`

	Questions []Question `json:"questions,omitempty" yaml:"questions,omitempty"`

	ReadmeLink string `json:"readmeLink,omitempty" yaml:"readme_link,omitempty"`

	TemplateBase string `json:"templateBase,omitempty" yaml:"template_base,omitempty"`

	TemplateVersionKuladoVersion map[string]interface{} `json:"templateVersionKuladoVersion,omitempty" yaml:"template_version_kulado_version,omitempty"`

	TemplateVersionKuladoVersionGte map[string]interface{} `json:"templateVersionKuladoVersionGte,omitempty" yaml:"template_version_kulado_version_gte,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	UpgradeFrom string `json:"upgradeFrom,omitempty" yaml:"upgrade_from,omitempty"`

	UpgradeVersionLinks map[string]interface{} `json:"upgradeVersionLinks,omitempty" yaml:"upgrade_version_links,omitempty"`

	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	VersionLinks map[string]interface{} `json:"versionLinks,omitempty" yaml:"version_links,omitempty"`
}

type TemplateVersionCollection struct {
	Collection
	Data   []TemplateVersion `json:"data,omitempty"`
	client *TemplateVersionClient
}

type TemplateVersionClient struct {
	kuladoClient *KuladoClient
}

type TemplateVersionOperations interface {
	List(opts *ListOpts) (*TemplateVersionCollection, error)
	Create(opts *TemplateVersion) (*TemplateVersion, error)
	Update(existing *TemplateVersion, updates interface{}) (*TemplateVersion, error)
	ById(id string) (*TemplateVersion, error)
	Delete(container *TemplateVersion) error
}

func newTemplateVersionClient(kuladoClient *KuladoClient) *TemplateVersionClient {
	return &TemplateVersionClient{
		kuladoClient: kuladoClient,
	}
}

func (c *TemplateVersionClient) Create(container *TemplateVersion) (*TemplateVersion, error) {
	resp := &TemplateVersion{}
	err := c.kuladoClient.doCreate(TEMPLATE_VERSION_TYPE, container, resp)
	return resp, err
}

func (c *TemplateVersionClient) Update(existing *TemplateVersion, updates interface{}) (*TemplateVersion, error) {
	resp := &TemplateVersion{}
	err := c.kuladoClient.doUpdate(TEMPLATE_VERSION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TemplateVersionClient) List(opts *ListOpts) (*TemplateVersionCollection, error) {
	resp := &TemplateVersionCollection{}
	err := c.kuladoClient.doList(TEMPLATE_VERSION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *TemplateVersionCollection) Next() (*TemplateVersionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TemplateVersionCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TemplateVersionClient) ById(id string) (*TemplateVersion, error) {
	resp := &TemplateVersion{}
	err := c.kuladoClient.doById(TEMPLATE_VERSION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *TemplateVersionClient) Delete(container *TemplateVersion) error {
	return c.kuladoClient.doResourceDelete(TEMPLATE_VERSION_TYPE, &container.Resource)
}
