package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunCountRequestBuilder builds and executes requests for operations under \api\run\count
type RunCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunCountRequestBuilderGetQueryParameters run count summary for given Test ID
type RunCountRequestBuilderGetQueryParameters struct {
    // Test ID
    TestId *int32 `uriparametername:"testId"`
}
// RunCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunCountRequestBuilderGetQueryParameters
}
// NewRunCountRequestBuilderInternal instantiates a new RunCountRequestBuilder and sets the default values.
func NewRunCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunCountRequestBuilder) {
    m := &RunCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/count?testId={testId}", pathParameters),
    }
    return m
}
// NewRunCountRequestBuilder instantiates a new RunCountRequestBuilder and sets the default values.
func NewRunCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get run count summary for given Test ID
// returns a RunCountable when successful
func (m *RunCountRequestBuilder) Get(ctx context.Context, requestConfiguration *RunCountRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunCountable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateRunCountFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunCountable), nil
}
// ToGetRequestInformation run count summary for given Test ID
// returns a *RequestInformation when successful
func (m *RunCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunCountRequestBuilder when successful
func (m *RunCountRequestBuilder) WithUrl(rawUrl string)(*RunCountRequestBuilder) {
    return NewRunCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
