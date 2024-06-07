package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaIdbyuriWithUriItemRequestBuilder builds and executes requests for operations under \api\schema\idByUri\{uri}
type SchemaIdbyuriWithUriItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaIdbyuriWithUriItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaIdbyuriWithUriItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaIdbyuriWithUriItemRequestBuilderInternal instantiates a new SchemaIdbyuriWithUriItemRequestBuilder and sets the default values.
func NewSchemaIdbyuriWithUriItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdbyuriWithUriItemRequestBuilder) {
    m := &SchemaIdbyuriWithUriItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/idByUri/{uri}", pathParameters),
    }
    return m
}
// NewSchemaIdbyuriWithUriItemRequestBuilder instantiates a new SchemaIdbyuriWithUriItemRequestBuilder and sets the default values.
func NewSchemaIdbyuriWithUriItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdbyuriWithUriItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaIdbyuriWithUriItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve Schema ID by uri
// returns a *int32 when successful
func (m *SchemaIdbyuriWithUriItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaIdbyuriWithUriItemRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation retrieve Schema ID by uri
// returns a *RequestInformation when successful
func (m *SchemaIdbyuriWithUriItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaIdbyuriWithUriItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaIdbyuriWithUriItemRequestBuilder when successful
func (m *SchemaIdbyuriWithUriItemRequestBuilder) WithUrl(rawUrl string)(*SchemaIdbyuriWithUriItemRequestBuilder) {
    return NewSchemaIdbyuriWithUriItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
