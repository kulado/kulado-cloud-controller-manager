package catalog

type KuladoClient struct {
	KuladoBaseClient

	ApiVersion      ApiVersionOperations
	Question        QuestionOperations
	Template        TemplateOperations
	TemplateVersion TemplateVersionOperations
	Catalog         CatalogOperations
	Error           ErrorOperations
}

func constructClient(kuladoBaseClient *KuladoBaseClientImpl) *KuladoClient {
	client := &KuladoClient{
		KuladoBaseClient: kuladoBaseClient,
	}

	client.ApiVersion = newApiVersionClient(client)
	client.Question = newQuestionClient(client)
	client.Template = newTemplateClient(client)
	client.TemplateVersion = newTemplateVersionClient(client)
	client.Catalog = newCatalogClient(client)
	client.Error = newErrorClient(client)

	return client
}

func NewKuladoClient(opts *ClientOpts) (*KuladoClient, error) {
	kuladoBaseClient := &KuladoBaseClientImpl{
		Types: map[string]Schema{},
	}
	client := constructClient(kuladoBaseClient)

	err := setupKuladoBaseClient(kuladoBaseClient, opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}
