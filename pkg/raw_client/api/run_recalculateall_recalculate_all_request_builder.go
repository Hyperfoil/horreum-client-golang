package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunRecalculateallRecalculateAllRequestBuilder builds and executes requests for operations under \api\run\recalculateAll
type RunRecalculateallRecalculateAllRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunRecalculateallRecalculateAllRequestBuilderPostQueryParameters recalculate Datasets for Runs between two dates
type RunRecalculateallRecalculateAllRequestBuilderPostQueryParameters struct {
    // start timestamp
    From *string `uriparametername:"from"`
    // end timestamp
    To *string `uriparametername:"to"`
}
// RunRecalculateallRecalculateAllRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunRecalculateallRecalculateAllRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunRecalculateallRecalculateAllRequestBuilderPostQueryParameters
}
// NewRunRecalculateallRecalculateAllRequestBuilderInternal instantiates a new RunRecalculateallRecalculateAllRequestBuilder and sets the default values.
func NewRunRecalculateallRecalculateAllRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunRecalculateallRecalculateAllRequestBuilder) {
    m := &RunRecalculateallRecalculateAllRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/recalculateAll{?from*,to*}", pathParameters),
    }
    return m
}
// NewRunRecalculateallRecalculateAllRequestBuilder instantiates a new RunRecalculateallRecalculateAllRequestBuilder and sets the default values.
func NewRunRecalculateallRecalculateAllRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunRecalculateallRecalculateAllRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunRecalculateallRecalculateAllRequestBuilderInternal(urlParams, requestAdapter)
}
// Post recalculate Datasets for Runs between two dates
func (m *RunRecalculateallRecalculateAllRequestBuilder) Post(ctx context.Context, requestConfiguration *RunRecalculateallRecalculateAllRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation recalculate Datasets for Runs between two dates
// returns a *RequestInformation when successful
func (m *RunRecalculateallRecalculateAllRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RunRecalculateallRecalculateAllRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunRecalculateallRecalculateAllRequestBuilder when successful
func (m *RunRecalculateallRecalculateAllRequestBuilder) WithUrl(rawUrl string)(*RunRecalculateallRecalculateAllRequestBuilder) {
    return NewRunRecalculateallRecalculateAllRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
