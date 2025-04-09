package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ConfigDatastoreDatastoreItemRequestBuilder builds and executes requests for operations under \api\config\datastore\{id}
type ConfigDatastoreDatastoreItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigDatastoreDatastoreItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigDatastoreDatastoreItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigDatastoreDatastoreItemRequestBuilderInternal instantiates a new ConfigDatastoreDatastoreItemRequestBuilder and sets the default values.
func NewConfigDatastoreDatastoreItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigDatastoreDatastoreItemRequestBuilder) {
    m := &ConfigDatastoreDatastoreItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/config/datastore/{id}", pathParameters),
    }
    return m
}
// NewConfigDatastoreDatastoreItemRequestBuilder instantiates a new ConfigDatastoreDatastoreItemRequestBuilder and sets the default values.
func NewConfigDatastoreDatastoreItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigDatastoreDatastoreItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigDatastoreDatastoreItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete test a Datastore
func (m *ConfigDatastoreDatastoreItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConfigDatastoreDatastoreItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Test the test property
// returns a *ConfigDatastoreItemTestRequestBuilder when successful
func (m *ConfigDatastoreDatastoreItemRequestBuilder) Test()(*ConfigDatastoreItemTestRequestBuilder) {
    return NewConfigDatastoreItemTestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation test a Datastore
// returns a *RequestInformation when successful
func (m *ConfigDatastoreDatastoreItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConfigDatastoreDatastoreItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConfigDatastoreDatastoreItemRequestBuilder when successful
func (m *ConfigDatastoreDatastoreItemRequestBuilder) WithUrl(rawUrl string)(*ConfigDatastoreDatastoreItemRequestBuilder) {
    return NewConfigDatastoreDatastoreItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
