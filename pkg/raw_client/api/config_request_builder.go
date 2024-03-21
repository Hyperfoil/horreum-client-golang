package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ConfigRequestBuilder builds and executes requests for operations under \api\config
type ConfigRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewConfigRequestBuilderInternal instantiates a new ConfigRequestBuilder and sets the default values.
func NewConfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigRequestBuilder) {
    m := &ConfigRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/config", pathParameters),
    }
    return m
}
// NewConfigRequestBuilder instantiates a new ConfigRequestBuilder and sets the default values.
func NewConfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Datastore the datastore property
// returns a *ConfigDatastoreRequestBuilder when successful
func (m *ConfigRequestBuilder) Datastore()(*ConfigDatastoreRequestBuilder) {
    return NewConfigDatastoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Keycloak the keycloak property
// returns a *ConfigKeycloakRequestBuilder when successful
func (m *ConfigRequestBuilder) Keycloak()(*ConfigKeycloakRequestBuilder) {
    return NewConfigKeycloakRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Version the version property
// returns a *ConfigVersionRequestBuilder when successful
func (m *ConfigRequestBuilder) Version()(*ConfigVersionRequestBuilder) {
    return NewConfigVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
