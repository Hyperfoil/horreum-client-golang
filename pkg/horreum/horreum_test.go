package horreum

import (
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
	assert.Equal(t, "providing password without username, have you missed something?", err.Error())
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

func TestGetClientVersion(t *testing.T) {
	a, client := setup(t)

	version := client.Version()
	a.NotEmpty(version)
}
