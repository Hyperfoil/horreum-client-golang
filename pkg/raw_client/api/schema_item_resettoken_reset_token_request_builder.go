package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaItemResettokenResetTokenRequestBuilder builds and executes requests for operations under \api\schema\{-id}\resetToken
type SchemaItemResettokenResetTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaItemResettokenResetTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaItemResettokenResetTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaItemResettokenResetTokenRequestBuilderInternal instantiates a new SchemaItemResettokenResetTokenRequestBuilder and sets the default values.
func NewSchemaItemResettokenResetTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemResettokenResetTokenRequestBuilder) {
    m := &SchemaItemResettokenResetTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/{%2Did}/resetToken", pathParameters),
    }
    return m
}
// NewSchemaItemResettokenResetTokenRequestBuilder instantiates a new SchemaItemResettokenResetTokenRequestBuilder and sets the default values.
func NewSchemaItemResettokenResetTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemResettokenResetTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaItemResettokenResetTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post regenerate access token for schema
// returns a *string when successful
func (m *SchemaItemResettokenResetTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *SchemaItemResettokenResetTokenRequestBuilderPostRequestConfiguration)(*string, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*string), nil
}
// ToPostRequestInformation regenerate access token for schema
// returns a *RequestInformation when successful
func (m *SchemaItemResettokenResetTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SchemaItemResettokenResetTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaItemResettokenResetTokenRequestBuilder when successful
func (m *SchemaItemResettokenResetTokenRequestBuilder) WithUrl(rawUrl string)(*SchemaItemResettokenResetTokenRequestBuilder) {
    return NewSchemaItemResettokenResetTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
