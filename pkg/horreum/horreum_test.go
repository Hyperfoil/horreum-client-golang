package horreum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	username = "user"
	password = "secret"
)

func setup(t *testing.T) (*assert.Assertions, *HorreumClient) {
	client, err := NewHorreumClient("http://localhost:8080", nil, nil)
	assertion := assert.New(t)
	assertion.Nil(err)

	return assertion, client
}

func TestMissingMissingPasswordWithUsername(t *testing.T) {
	_, err := NewHorreumClient("http://localhost:8080", nil, &password)
	assert.NotNil(t, err)
	assert.Equal(t, "providing password without username, have you missed something?", err.Error())
}

func TestAuthProviderSetupFailure(t *testing.T) {
	_, err := NewHorreumClient("http://localhost:9999", &username, &password)
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
