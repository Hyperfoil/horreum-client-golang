package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemDataRequestBuilder builds and executes requests for operations under \api\run\{id}\data
type RunItemDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemDataRequestBuilderGetQueryParameters get Run data by Run ID
type RunItemDataRequestBuilderGetQueryParameters struct {
    // FIlter by Schmea URI
    SchemaUri *string `uriparametername:"schemaUri"`
}
// RunItemDataRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemDataRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunItemDataRequestBuilderGetQueryParameters
}
// NewRunItemDataRequestBuilderInternal instantiates a new RunItemDataRequestBuilder and sets the default values.
func NewRunItemDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemDataRequestBuilder) {
    m := &RunItemDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/data{?schemaUri*}", pathParameters),
    }
    return m
}
// NewRunItemDataRequestBuilder instantiates a new RunItemDataRequestBuilder and sets the default values.
func NewRunItemDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get Run data by Run ID
// Deprecated: This method is obsolete. Use GetAsDataGetResponse instead.
// returns a RunItemDataResponseable when successful
func (m *RunItemDataRequestBuilder) Get(ctx context.Context, requestConfiguration *RunItemDataRequestBuilderGetRequestConfiguration)(RunItemDataResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRunItemDataResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RunItemDataResponseable), nil
}
// GetAsDataGetResponse get Run data by Run ID
// returns a RunItemDataGetResponseable when successful
func (m *RunItemDataRequestBuilder) GetAsDataGetResponse(ctx context.Context, requestConfiguration *RunItemDataRequestBuilderGetRequestConfiguration)(RunItemDataGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRunItemDataGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RunItemDataGetResponseable), nil
}
// ToGetRequestInformation get Run data by Run ID
// returns a *RequestInformation when successful
func (m *RunItemDataRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunItemDataRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunItemDataRequestBuilder when successful
func (m *RunItemDataRequestBuilder) WithUrl(rawUrl string)(*RunItemDataRequestBuilder) {
    return NewRunItemDataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
