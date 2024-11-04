package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemMetadataRequestBuilder builds and executes requests for operations under \api\run\{id}\metadata
type RunItemMetadataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemMetadataRequestBuilderGetQueryParameters get Run  meta data by Run ID
type RunItemMetadataRequestBuilderGetQueryParameters struct {
    // Filter by Schmea URI
    SchemaUri *string `uriparametername:"schemaUri"`
}
// RunItemMetadataRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemMetadataRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunItemMetadataRequestBuilderGetQueryParameters
}
// NewRunItemMetadataRequestBuilderInternal instantiates a new RunItemMetadataRequestBuilder and sets the default values.
func NewRunItemMetadataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemMetadataRequestBuilder) {
    m := &RunItemMetadataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/metadata{?schemaUri*}", pathParameters),
    }
    return m
}
// NewRunItemMetadataRequestBuilder instantiates a new RunItemMetadataRequestBuilder and sets the default values.
func NewRunItemMetadataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemMetadataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemMetadataRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get Run  meta data by Run ID
// Deprecated: This method is obsolete. Use GetAsMetadataGetResponse instead.
// returns a RunItemMetadataResponseable when successful
func (m *RunItemMetadataRequestBuilder) Get(ctx context.Context, requestConfiguration *RunItemMetadataRequestBuilderGetRequestConfiguration)(RunItemMetadataResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRunItemMetadataResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RunItemMetadataResponseable), nil
}
// GetAsMetadataGetResponse get Run  meta data by Run ID
// returns a RunItemMetadataGetResponseable when successful
func (m *RunItemMetadataRequestBuilder) GetAsMetadataGetResponse(ctx context.Context, requestConfiguration *RunItemMetadataRequestBuilderGetRequestConfiguration)(RunItemMetadataGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRunItemMetadataGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RunItemMetadataGetResponseable), nil
}
// ToGetRequestInformation get Run  meta data by Run ID
// returns a *RequestInformation when successful
func (m *RunItemMetadataRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunItemMetadataRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemMetadataRequestBuilder when successful
func (m *RunItemMetadataRequestBuilder) WithUrl(rawUrl string)(*RunItemMetadataRequestBuilder) {
    return NewRunItemMetadataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
