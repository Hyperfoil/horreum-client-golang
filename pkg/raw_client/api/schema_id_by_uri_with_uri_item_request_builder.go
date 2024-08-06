package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaIdByUriWithUriItemRequestBuilder builds and executes requests for operations under \api\schema\idByUri\{uri}
type SchemaIdByUriWithUriItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaIdByUriWithUriItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaIdByUriWithUriItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaIdByUriWithUriItemRequestBuilderInternal instantiates a new SchemaIdByUriWithUriItemRequestBuilder and sets the default values.
func NewSchemaIdByUriWithUriItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdByUriWithUriItemRequestBuilder) {
    m := &SchemaIdByUriWithUriItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/idByUri/{uri}", pathParameters),
    }
    return m
}
// NewSchemaIdByUriWithUriItemRequestBuilder instantiates a new SchemaIdByUriWithUriItemRequestBuilder and sets the default values.
func NewSchemaIdByUriWithUriItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdByUriWithUriItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaIdByUriWithUriItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve Schema ID by uri
// returns a *int32 when successful
func (m *SchemaIdByUriWithUriItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaIdByUriWithUriItemRequestBuilderGetRequestConfiguration)(*int32, error) {
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
func (m *SchemaIdByUriWithUriItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaIdByUriWithUriItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaIdByUriWithUriItemRequestBuilder when successful
func (m *SchemaIdByUriWithUriItemRequestBuilder) WithUrl(rawUrl string)(*SchemaIdByUriWithUriItemRequestBuilder) {
    return NewSchemaIdByUriWithUriItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
