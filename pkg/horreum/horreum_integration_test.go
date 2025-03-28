//go:build integration
// +build integration

package horreum

import (
	"context"
	"net/url"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/authentication"

	"github.com/hyperfoil/horreum-client-golang/pkg/raw_client/api"
	"github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

const authorizationHeader = "Authorization"

var (
	ctx = context.Background()
)

func setupAnonymousClient(t *testing.T) (*assert.Assertions, *HorreumClient) {
	client, err := NewHorreumClient("http://localhost:8080", nil, nil)
	assertion := assert.New(t)
	assertion.Nil(err)

	checkHorreumConnection(t, client)

	return assertion, client
}

func setupAuthenticatedClient(t *testing.T) (*assert.Assertions, *HorreumClient) {
	client, err := NewHorreumClient("http://localhost:8080", &HorreumCredentials{
		Username: &username,
		Password: &password,
	}, nil)
	assertion := assert.New(t)
	assertion.Nil(err)

	checkHorreumConnection(t, client)

	return assertion, client
}

func checkHorreumConnection(t *testing.T, client *HorreumClient) {
	_, err := client.RawClient.Api().Config().Version().Get(ctx, nil)
	assert.Nil(t, err, "unable to fetch Horreum version, is Horreum running in the background?")
}

func TestMissingTokenWithAnonymous(t *testing.T) {
	a, client := setupAnonymousClient(t)

	req := &abstractions.RequestInformation{
		Method:      abstractions.GET,
		UrlTemplate: "/api",
	}
	err := client.AuthProvider.AuthenticateRequest(ctx, req, nil)

	a.Nil(err)
	a.Nil(req.Headers)
}

func TestExistingTokenWithAuthentication(t *testing.T) {
	a, client := setupAuthenticatedClient(t)

	req := &abstractions.RequestInformation{
		Method: abstractions.GET,
	}
	req.SetUri(url.URL{Scheme: "https", Host: "localhost:8080", Path: "api/"})
	err := client.AuthProvider.AuthenticateRequest(ctx, req, nil)

	a.Nil(err)
	a.NotNil(req.Headers)
	a.Greater(len(req.Headers.Get(authorizationHeader)), 0)
}

func TestGetAnonymousServerVersion(t *testing.T) {
	a, client := setupAnonymousClient(t)

	version, err := client.RawClient.Api().Config().Version().Get(context.Background(), nil)
	a.Nil(err)
	a.NotNil(version.GetVersion())
	a.NotEmpty(*version.GetVersion())
	a.NotNil(version.GetStartTimestamp())
	a.Greater(*version.GetStartTimestamp(), int64(0))
}

func TestGetAuthenticatedServerVersion(t *testing.T) {
	a, client := setupAuthenticatedClient(t)

	version, err := client.RawClient.Api().Config().Version().Get(context.Background(), nil)
	a.Nil(err)
	a.NotNil(version.GetVersion())
	a.NotEmpty(*version.GetVersion())
	a.NotNil(version.GetStartTimestamp())
	a.Greater(*version.GetStartTimestamp(), int64(0))
}

func TestEmptyTestsList(t *testing.T) {
	a, client := setupAuthenticatedClient(t)

	res, err := client.RawClient.Api().Test().Get(ctx, &api.TestRequestBuilderGetRequestConfiguration{
		QueryParameters: &api.TestRequestBuilderGetQueryParameters{
			Limit: of(int32(1)),
			Page:  of(int32(0)),
		},
	})
	a.Nil(err)
	a.NotNil(res.GetCount())
	a.EqualValues(0, *res.GetCount())
}

func TestAddAndDeleteTest(t *testing.T) {
	a, client := setupAuthenticatedClient(t)

	body := models.NewTest()
	body.SetName(of("TestName"))
	body.SetDescription(of("Simple Test"))
	body.SetOwner(of("dev-team"))
	body.SetAccess(of(models.PROTECTED_TEST_ACCESS))

	created, err := client.RawClient.Api().Test().Post(ctx, body, nil)
	a.Nil(err)
	a.NotNil(created.GetId())

	list, err := client.RawClient.Api().Test().Get(ctx, nil)
	a.Nil(err)
	a.EqualValues(1, *list.GetCount())

	err = client.RawClient.Api().Test().ByIdInteger(*created.GetId()).Delete(ctx, nil)
	a.Nil(err)

	list, err = client.RawClient.Api().Test().Get(ctx, nil)
	a.Nil(err)
	a.EqualValues(0, *list.GetCount())
}

func TestAPIKeyAuthSetup(t *testing.T) {
	client, _ := NewHorreumClient("http://localhost:8080", &HorreumCredentials{
		Username: &username,
		Password: &password,
	}, &ClientConfiguration{
		AuthMethod: BEARER,
	})

	keyRequest := api.NewUserApikeyPostRequestBody()
	keyRequest.SetName(of("Go client test key"))
	keyRequest.SetTypeEscaped(of(models.USER_KEYTYPE))

	key, err := client.RawClient.Api().User().Apikey().Post(context.Background(), keyRequest, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, key)

	apiClient, err := NewHorreumClient("http://localhost:8080", &HorreumCredentials{
		ApiKey: key,
	}, &ClientConfiguration{AuthMethod: API_KEY})

	assert.IsType(t, &authentication.ApiKeyAuthenticationProvider{}, apiClient.AuthProvider)

	// make sure the API Key header is added correctly to a dummy request
	req := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, "https://localhost:8080", make(map[string]string))
	apiClient.AuthProvider.AuthenticateRequest(context.Background(), req, nil)
	assert.True(t, req.Headers.ContainsKey("X-Horreum-API-Key"))
	assert.Equal(t, *key, req.Headers.Get("X-Horreum-API-Key")[0])

	// kiota go client requires https which prevents using the API Key for integration test (with a dev Horreum instance)
	_, apiErr := apiClient.RawClient.Api().User().Apikey().Get(context.Background(), nil)
	assert.NotNil(t, apiErr)
}

// of returns a pointer to the provided literal/const input
func of[E any](e E) *E {
	return &e
}
