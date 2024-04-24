package horreum

import (
	"context"
	"fmt"

	nethttp "net/http"

	"github.com/hyperfoil/horreum-client-golang/pkg/raw_client"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"

	_ "embed"
)

//go:embed version.txt
var version string

type HorreumClient struct {
	baseUrl      string
	credentials  *HorreumCredentials
	clientConfig *ClientConfiguration

	RawClient   *raw_client.HorreumRawClient
	AuthProvder authentication.AuthenticationProvider
}

func setupAuthProvider(baseUrl string, username string, password string) (authentication.AccessTokenProvider, error) {
	anonymousProvider := &authentication.AnonymousAuthenticationProvider{}
	anonymousAdapter, err := http.NewNetHttpRequestAdapter(anonymousProvider)
	if err != nil {
		return nil, fmt.Errorf("error creating anonymous adapter: %w", err)
	}

	anonymousAdapter.SetBaseUrl(baseUrl)
	client := raw_client.NewHorreumRawClient(anonymousAdapter)

	config, err := client.Api().Config().Keycloak().Get(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error retrieving keycloak configuration: %w", err)
	}

	return NewKeycloakAccessProvider(config, username, password)
}

func NewHorreumClient(baseUrl string, credentials *HorreumCredentials, clientConfig *ClientConfiguration) (*HorreumClient, error) {
	client := &HorreumClient{
		baseUrl:      baseUrl,
		credentials:  credentials,
		clientConfig: clientConfig,

		// By default, set auth provider to anonymous one
		AuthProvder: &authentication.AnonymousAuthenticationProvider{},
	}

	if credentials != nil {
		if credentials.Username != nil && credentials.Password != nil {
			provider, err := setupAuthProvider(baseUrl, *credentials.Username, *credentials.Password)
			if err != nil {
				return nil, fmt.Errorf("error setting up keycloak provider: %w", err)
			}
			client.AuthProvder = authentication.NewBaseBearerTokenAuthenticationProvider(provider)
		} else if credentials.Password != nil {
			return nil, fmt.Errorf("providing password without username, have you missed something?")
		}
	}

	// Disable gzip compression
	// To avoid "HTTP 400 Bad Request Illegal character ((CTRL-CHAR, code 31)): only regular white space (\r, \n, \t) is allowed between tokens"
	compressOption := http.NewCompressionOptions(false)
	middlewareOptions := []abstractions.RequestOption{
		&compressOption,
	}
	// Add additional middleware options if provided
	if clientConfig != nil && len(clientConfig.Options) > 0 {
		middlewareOptions = append(middlewareOptions, clientConfig.Options...)
	}

	// Create default middlewares with provided options
	middlewares, err := http.GetDefaultMiddlewaresWithOptions(middlewareOptions...)
	if err != nil {
		return nil, fmt.Errorf("error creating default middlewares with custom options: %w", err)
	}

	var httpClient *nethttp.Client
	if clientConfig != nil && clientConfig.HttpClient != nil && clientConfig.UseDefaultMiddlewares {
		httpClient = clientConfig.HttpClient
		// Set middlewares
		httpClient.Transport = http.NewCustomTransportWithParentTransport(clientConfig.ParentTransport, middlewares...)
	} else {
		httpClient = http.GetDefaultClient(middlewares...)
	}

	adapter, err := http.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(client.AuthProvder, nil, nil, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error creating client adapter: %w", err)
	}
	adapter.SetBaseUrl(baseUrl)
	client.RawClient = raw_client.NewHorreumRawClient(adapter)

	return client, nil
}

func (h *HorreumClient) Version() string {
	return version
}
