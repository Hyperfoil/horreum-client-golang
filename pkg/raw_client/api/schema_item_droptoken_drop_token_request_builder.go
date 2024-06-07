package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaItemDroptokenDropTokenRequestBuilder builds and executes requests for operations under \api\schema\{-id}\dropToken
type SchemaItemDroptokenDropTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaItemDroptokenDropTokenRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaItemDroptokenDropTokenRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaItemDroptokenDropTokenRequestBuilderInternal instantiates a new SchemaItemDroptokenDropTokenRequestBuilder and sets the default values.
func NewSchemaItemDroptokenDropTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemDroptokenDropTokenRequestBuilder) {
    m := &SchemaItemDroptokenDropTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/{%2Did}/dropToken", pathParameters),
    }
    return m
}
// NewSchemaItemDroptokenDropTokenRequestBuilder instantiates a new SchemaItemDroptokenDropTokenRequestBuilder and sets the default values.
func NewSchemaItemDroptokenDropTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemDroptokenDropTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaItemDroptokenDropTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove access token for schema
func (m *SchemaItemDroptokenDropTokenRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchemaItemDroptokenDropTokenRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *SchemaItemDroptokenDropTokenRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SchemaItemDroptokenDropTokenRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaItemDroptokenDropTokenRequestBuilder when successful
func (m *SchemaItemDroptokenDropTokenRequestBuilder) WithUrl(rawUrl string)(*SchemaItemDroptokenDropTokenRequestBuilder) {
    return NewSchemaItemDroptokenDropTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
