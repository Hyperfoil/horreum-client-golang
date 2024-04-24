package horreum

import (
	"net/http"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type HorreumCredentials struct {
	Username *string
	Password *string
}

func NewDefaultHorreumCredentials() HorreumCredentials {
	return HorreumCredentials{
		Username: nil,
		Password: nil,
	}
}

type ClientConfiguration struct {
	HttpClient            *http.Client
	ParentTransport       http.RoundTripper
	UseDefaultMiddlewares bool
	Options               []abstractions.RequestOption
}

func NewDefaultClientConfiguration() ClientConfiguration {
	return ClientConfiguration{
		HttpClient:            nil,
		ParentTransport:       nil,
		UseDefaultMiddlewares: true,
		Options:               []abstractions.RequestOption{},
	}
}
