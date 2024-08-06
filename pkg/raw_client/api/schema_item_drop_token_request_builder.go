package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaItemDropTokenRequestBuilder builds and executes requests for operations under \api\schema\{-id}\dropToken
type SchemaItemDropTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaItemDropTokenRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaItemDropTokenRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaItemDropTokenRequestBuilderInternal instantiates a new SchemaItemDropTokenRequestBuilder and sets the default values.
func NewSchemaItemDropTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemDropTokenRequestBuilder) {
    m := &SchemaItemDropTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/{%2Did}/dropToken", pathParameters),
    }
    return m
}
// NewSchemaItemDropTokenRequestBuilder instantiates a new SchemaItemDropTokenRequestBuilder and sets the default values.
func NewSchemaItemDropTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemDropTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaItemDropTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove access token for schema
func (m *SchemaItemDropTokenRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchemaItemDropTokenRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation remove access token for schema
// returns a *RequestInformation when successful
func (m *SchemaItemDropTokenRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SchemaItemDropTokenRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaItemDropTokenRequestBuilder when successful
func (m *SchemaItemDropTokenRequestBuilder) WithUrl(rawUrl string)(*SchemaItemDropTokenRequestBuilder) {
    return NewSchemaItemDropTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
