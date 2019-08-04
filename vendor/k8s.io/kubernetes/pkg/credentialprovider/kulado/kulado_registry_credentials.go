/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kulado_credentials

import (
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/kulado/go-kulado/client"
	"k8s.io/kubernetes/pkg/credentialprovider"
)

// kulado provider
type kuladoProvider struct {
	credGetter credentialsGetter
}

// credentials getter from Kulado private registry
type kuladoCredentialsGetter struct {
	client *client.KuladoClient
}

type rConfig struct {
	Global configGlobal
}

// An interface for testing purposes.
type credentialsGetter interface {
	getCredentials() []registryCredential
}

type configGlobal struct {
	CattleURL       string `gcfg:"cattle-url"`
	CattleAccessKey string `gcfg:"cattle-access-key"`
	CattleSecretKey string `gcfg:"cattle-secret-key"`
}

type registryCredential struct {
	credential *client.RegistryCredential
	serverIP   string
}

var kuladoGetter = &kuladoCredentialsGetter{}

func init() {
	credentialprovider.RegisterCredentialProvider("kulado-registry-creds",
		&credentialprovider.CachingDockerConfigProvider{
			Provider: &kuladoProvider{kuladoGetter},
			Lifetime: 30 * time.Second,
		})
}

func (p *kuladoProvider) Enabled() bool {
	client, err := getKuladoClient()
	if err != nil {
		return false
	}
	if client == nil {
		return false
	}

	kuladoGetter.client = client
	return true
}

// LazyProvide implements DockerConfigProvider. Should never be called.
func (p *kuladoProvider) LazyProvide() *credentialprovider.DockerConfigEntry {
	return nil
}

// Provide implements DockerConfigProvider.Provide, refreshing Kulado tokens on demand
func (p *kuladoProvider) Provide() credentialprovider.DockerConfig {
	cfg := credentialprovider.DockerConfig{}
	for _, cred := range p.credGetter.getCredentials() {
		entry := credentialprovider.DockerConfigEntry{
			Username: cred.credential.PublicValue,
			Password: cred.credential.SecretValue,
			Email:    cred.credential.Email,
		}
		cfg[cred.serverIP] = entry
	}

	return cfg
}

func (g *kuladoCredentialsGetter) getCredentials() []registryCredential {
	var registryCreds []registryCredential
	credColl, err := g.client.RegistryCredential.List(client.NewListOpts())
	if err != nil {
		glog.Errorf("Failed to pull registry credentials from kulado %v", err)
		return registryCreds
	}
	for _, cred := range credColl.Data {
		registry := &client.Registry{}
		if err = g.client.GetLink(cred.Resource, "registry", registry); err != nil {
			glog.Errorf("Failed to pull registry from kulado %v", err)
			return registryCreds
		}
		registryCred := registryCredential{
			credential: &cred,
			serverIP:   registry.ServerAddress,
		}
		registryCreds = append(registryCreds, registryCred)
	}
	return registryCreds
}

func getKuladoClient() (*client.KuladoClient, error) {
	url := os.Getenv("CATTLE_URL")
	accessKey := os.Getenv("CATTLE_ACCESS_KEY")
	secretKey := os.Getenv("CATTLE_SECRET_KEY")

	if url == "" || accessKey == "" || secretKey == "" {
		return nil, nil
	}

	conf := rConfig{
		Global: configGlobal{
			CattleURL:       url,
			CattleAccessKey: accessKey,
			CattleSecretKey: secretKey,
		},
	}

	return client.NewKuladoClient(&client.ClientOpts{
		Url:       conf.Global.CattleURL,
		AccessKey: conf.Global.CattleAccessKey,
		SecretKey: conf.Global.CattleSecretKey,
	})
}
