/*
Copyright 2020 The Kubernetes Authors.

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

package provider

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/trusts"
	"net/http"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	osclient "github.com/gophercloud/utils/client"
	"github.com/gophercloud/utils/openstack/clientconfig"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	infrav1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha5"
)

type  NewAuthInfo struct {
	clientconfig.AuthInfo
	TrustID string `yaml:"trust_id,omitempty" json:"trust_id,omitempty"`
}

// NewCloud represents an entry in a clouds.yaml/public-clouds.yaml/secure.yaml file.
type NewCloud struct {
	Cloud      string    `yaml:"cloud,omitempty" json:"cloud,omitempty"`
	Profile    string    `yaml:"profile,omitempty" json:"profile,omitempty"`
	AuthInfo   *NewAuthInfo `yaml:"auth,omitempty" json:"auth,omitempty"`
	AuthType   clientconfig.AuthType  `yaml:"auth_type,omitempty" json:"auth_type,omitempty"`
	RegionName string    `yaml:"region_name,omitempty" json:"region_name,omitempty"`
	Regions    []clientconfig.Region  `yaml:"regions,omitempty" json:"regions,omitempty"`

	// EndpointType and Interface both specify whether to use the public, internal,
	// or admin interface of a service. They should be considered synonymous, but
	// EndpointType will take precedence when both are specified.
	EndpointType string `yaml:"endpoint_type,omitempty" json:"endpoint_type,omitempty"`
	Interface    string `yaml:"interface,omitempty" json:"interface,omitempty"`

	// API Version overrides.
	IdentityAPIVersion string `yaml:"identity_api_version,omitempty" json:"identity_api_version,omitempty"`
	VolumeAPIVersion   string `yaml:"volume_api_version,omitempty" json:"volume_api_version,omitempty"`

	// Verify whether or not SSL API requests should be verified.
	Verify *bool `yaml:"verify,omitempty" json:"verify,omitempty"`

	// CACertFile a path to a CA Cert bundle that can be used as part of
	// verifying SSL API requests.
	CACertFile string `yaml:"cacert,omitempty" json:"cacert,omitempty"`

	// ClientCertFile a path to a client certificate to use as part of the SSL
	// transaction.
	ClientCertFile string `yaml:"cert,omitempty" json:"cert,omitempty"`

	// ClientKeyFile a path to a client key to use as part of the SSL
	// transaction.
	ClientKeyFile string `yaml:"key,omitempty" json:"key,omitempty"`
}

type NewClouds struct {
	Clouds map[string]NewCloud `yaml:"clouds" json:"clouds"`
}

const (
	cloudsSecretKey = "clouds.yaml"
	caSecretKey     = "cacert"
)


func NewClientFromMachine(ctx context.Context, ctrlClient client.Client, openStackMachine *infrav1.OpenStackMachine) (*gophercloud.ProviderClient, *clientconfig.ClientOpts, string, error) {
	var cloud NewCloud
	var caCert []byte

	if openStackMachine.Spec.IdentityRef != nil {
		var err error
		cloud, caCert, err = getCloudFromSecret(ctx, ctrlClient, openStackMachine.Namespace, openStackMachine.Spec.IdentityRef.Name, openStackMachine.Spec.CloudName)
		if err != nil {
			return nil, nil, "", err
		}
	}
	return NewClient(cloud, caCert)
}

func NewClientFromCluster(ctx context.Context, ctrlClient client.Client, openStackCluster *infrav1.OpenStackCluster) (*gophercloud.ProviderClient, *clientconfig.ClientOpts, string, error) {
	var cloud NewCloud
	var caCert []byte

	if openStackCluster.Spec.IdentityRef != nil {
		var err error
		cloud, caCert, err = getCloudFromSecret(ctx, ctrlClient, openStackCluster.Namespace, openStackCluster.Spec.IdentityRef.Name, openStackCluster.Spec.CloudName)
		if err != nil {
			return nil, nil, "", err
		}
	}
	return NewClient(cloud, caCert)
}

func NewClient(cloud NewCloud, caCert []byte) (*gophercloud.ProviderClient, *clientconfig.ClientOpts, string, error) {
	clientOpts := new(clientconfig.ClientOpts)
	if cloud.AuthInfo != nil {
		clientOpts.AuthInfo = &cloud.AuthInfo.AuthInfo
		clientOpts.AuthType = cloud.AuthType
		clientOpts.RegionName = cloud.RegionName
	}

	opts, err := clientconfig.AuthOptions(clientOpts)
	if err != nil {
		return nil, nil, "", fmt.Errorf("auth option failed for cloud %v: %v", cloud.Cloud, err)
	}
	opts.AllowReauth = true


	provider, err := openstack.NewClient(opts.IdentityEndpoint)
	if err != nil {
		return nil, nil, "", fmt.Errorf("create providerClient err: %v", err)
	}
	config := &tls.Config{
		RootCAs:    x509.NewCertPool(),
		MinVersion: tls.VersionTLS12,
	}
	if cloud.Verify != nil {
		config.InsecureSkipVerify = !*cloud.Verify
	}
	if caCert != nil {
		config.RootCAs.AppendCertsFromPEM(caCert)
	}

	provider.HTTPClient.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment, TLSClientConfig: config}
	provider.HTTPClient.Transport = &osclient.RoundTripper{
		Rt:     provider.HTTPClient.Transport,
		Logger: &defaultLogger{},
	}
	if cloud.AuthInfo.TrustID!="" {
		tokenauth:=tokens.AuthOptions{}
		tokenauth.IdentityEndpoint=opts.IdentityEndpoint
		tokenauth.UserID=opts.UserID
		tokenauth.Username=opts.Username
		tokenauth.Password=opts.Password
		tokenauth.DomainID=opts.DomainID
		tokenauth.DomainName=opts.DomainName
		tokenauth.ApplicationCredentialID=opts.ApplicationCredentialID
		tokenauth.ApplicationCredentialName=opts.ApplicationCredentialName
		tokenauth.ApplicationCredentialSecret=opts.ApplicationCredentialSecret
		tokenauth.AllowReauth=opts.AllowReauth
		if opts.Scope!=nil {
			tokenauth.Scope.ProjectID=opts.Scope.ProjectID
			tokenauth.Scope.ProjectName=opts.Scope.ProjectName
			tokenauth.Scope.DomainName=opts.Scope.DomainName
			tokenauth.Scope.DomainID=opts.Scope.DomainID
		}
		authOptsExt := trusts.AuthOptsExt{
			TrustID:            cloud.AuthInfo.TrustID,
			AuthOptionsBuilder: &tokenauth,
		}
		err = openstack.AuthenticateV3(provider, authOptsExt, gophercloud.EndpointOpts{})
		if err != nil {
			return nil, nil, "", fmt.Errorf("providerClient authentication err: %v", err)
		}
		projectID, err := getProjectIDFromAuthResult(provider.GetAuthResult())
		if err != nil {
			return nil, nil, "", err
		}
		return provider,clientOpts,projectID,nil
	}
	err = openstack.Authenticate(provider, *opts)
	if err != nil {
		return nil, nil, "", fmt.Errorf("providerClient authentication err: %v", err)
	}
	projectID, err := getProjectIDFromAuthResult(provider.GetAuthResult())
	if err != nil {
		return nil, nil, "", err
	}

	return provider, clientOpts, projectID, nil
}

type defaultLogger struct{}

// Printf is a default Printf method.
func (defaultLogger) Printf(format string, args ...interface{}) {
	klog.V(6).Infof(format, args...)
}

// getCloudFromSecret extract a Cloud from the given namespace:secretName.
func getCloudFromSecret(ctx context.Context, ctrlClient client.Client, secretNamespace string, secretName string, cloudName string) (NewCloud, []byte, error) {
	emptyCloud := NewCloud{}

	if secretName == "" {
		return emptyCloud, nil, nil
	}

	if cloudName == "" {
		return emptyCloud, nil, fmt.Errorf("secret name set to %v but no cloud was specified. Please set cloud_name in your machine spec", secretName)
	}

	secret := &corev1.Secret{}
	err := ctrlClient.Get(ctx, types.NamespacedName{
		Namespace: secretNamespace,
		Name:      secretName,
	}, secret)
	if err != nil {
		return emptyCloud, nil, err
	}

	content, ok := secret.Data[cloudsSecretKey]
	if !ok {
		return emptyCloud, nil, fmt.Errorf("OpenStack credentials secret %v did not contain key %v",
			secretName, cloudsSecretKey)
	}
	var clouds NewClouds
	if err = yaml.Unmarshal(content, &clouds); err != nil {
		return emptyCloud, nil, fmt.Errorf("failed to unmarshal clouds credentials stored in secret %v: %v", secretName, err)
	}

	// get caCert
	caCert, ok := secret.Data[caSecretKey]
	if !ok {
		return clouds.Clouds[cloudName], nil, nil
	}

	return clouds.Clouds[cloudName], caCert, nil
}

// getProjectIDFromAuthResult handles different auth mechanisms to retrieve the
// current project id. Usually we use the Identity v3 Token mechanism that
// returns the project id in the response to the initial auth request.
func getProjectIDFromAuthResult(authResult gophercloud.AuthResult) (string, error) {
	switch authResult := authResult.(type) {
	case tokens.CreateResult:
		project, err := authResult.ExtractProject()
		if err != nil {
			return "", fmt.Errorf("unable to extract project from CreateResult: %v", err)
		}

		return project.ID, nil

	default:
		return "", fmt.Errorf("unable to get the project id from auth response with type %T", authResult)
	}
}
