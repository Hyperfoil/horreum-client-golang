package horreum

import (
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	username = "horreum.bootstrap"
	password = "secret"
)

func setup(t *testing.T) (*assert.Assertions, *HorreumClient) {
	client, err := NewHorreumClient("http://localhost:8080", nil, nil)
	assertion := assert.New(t)
	assertion.Nil(err)

	return assertion, client
}

func TestMissingMissingPasswordWithUsername(t *testing.T) {
	_, err := NewHorreumClient("http://localhost:8080", &HorreumCredentials{
		Username: nil,
		Password: &password,
	}, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "provided password without username", err.Error())
}

func TestAuthProviderSetupFailure(t *testing.T) {
	_, err := NewHorreumClient("http://localhost:9999", &HorreumCredentials{
		Username: &username,
		Password: &password,
	}, nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error setting up keycloak provider")
	assert.Contains(t, err.Error(), "error retrieving keycloak configuration")
	assert.Contains(t, err.Error(), "connection refused")
}

func TestBasicAuthSetup(t *testing.T) {
	client, _ := NewHorreumClient("http://localhost:9999", &HorreumCredentials{
		Username: &username,
		Password: &password,
	}, &ClientConfiguration{
		AuthMethod: BASIC,
	})
	assert.IsType(t, &authentication.ApiKeyAuthenticationProvider{}, client.AuthProvider)
}

func TestGetClientVersion(t *testing.T) {
	a, client := setup(t)

	version := client.Version()
	a.NotEmpty(version)
}
