package horreum

import (
	"net/http"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type HorreumCredentials struct {
	Username *string
	Password *string
	ApiKey   *string
}

func NewDefaultHorreumCredentials() HorreumCredentials {
	return HorreumCredentials{
		Username: nil,
		Password: nil,
		ApiKey:   nil,
	}
}

type AuthMethod int64

const (
	// BEARER authentication method where a token is provided from the OIDC server
	BEARER = iota
	// BASIC encodes username and password in the HTTP request
	BASIC
	// API_KEY authenticate with Horreum tokens
	API_KEY
)

type ClientConfiguration struct {
	HttpClient            *http.Client
	ParentTransport       http.RoundTripper
	UseDefaultMiddlewares bool
	AuthMethod            AuthMethod
	Options               []abstractions.RequestOption
}

func NewDefaultClientConfiguration() ClientConfiguration {
	return ClientConfiguration{
		HttpClient:            nil,
		ParentTransport:       nil,
		UseDefaultMiddlewares: true,
		AuthMethod:            BEARER,
		Options:               []abstractions.RequestOption{},
	}
}
