package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunItemSummaryRequestBuilder builds and executes requests for operations under \api\run\{id}\summary
type RunItemSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemSummaryRequestBuilderGetQueryParameters get Run Summary information by Run ID
type RunItemSummaryRequestBuilderGetQueryParameters struct {
    // Run API token
    Token *string `uriparametername:"token"`
}
// RunItemSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunItemSummaryRequestBuilderGetQueryParameters
}
// NewRunItemSummaryRequestBuilderInternal instantiates a new RunItemSummaryRequestBuilder and sets the default values.
func NewRunItemSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemSummaryRequestBuilder) {
    m := &RunItemSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/summary{?token*}", pathParameters),
    }
    return m
}
// NewRunItemSummaryRequestBuilder instantiates a new RunItemSummaryRequestBuilder and sets the default values.
func NewRunItemSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get Run Summary information by Run ID
// returns a RunSummaryable when successful
func (m *RunItemSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *RunItemSummaryRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateRunSummaryFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunSummaryable), nil
}
// ToGetRequestInformation get Run Summary information by Run ID
// returns a *RequestInformation when successful
func (m *RunItemSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunItemSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunItemSummaryRequestBuilder when successful
func (m *RunItemSummaryRequestBuilder) WithUrl(rawUrl string)(*RunItemSummaryRequestBuilder) {
    return NewRunItemSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
