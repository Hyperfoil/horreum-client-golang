package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunRecalculateAllRequestBuilder builds and executes requests for operations under \api\run\recalculateAll
type RunRecalculateAllRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunRecalculateAllRequestBuilderPostQueryParameters recalculate Datasets for Runs between two dates
type RunRecalculateAllRequestBuilderPostQueryParameters struct {
    // start timestamp
    From *string `uriparametername:"from"`
    // end timestamp
    To *string `uriparametername:"to"`
}
// RunRecalculateAllRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunRecalculateAllRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunRecalculateAllRequestBuilderPostQueryParameters
}
// NewRunRecalculateAllRequestBuilderInternal instantiates a new RunRecalculateAllRequestBuilder and sets the default values.
func NewRunRecalculateAllRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunRecalculateAllRequestBuilder) {
    m := &RunRecalculateAllRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/recalculateAll{?from*,to*}", pathParameters),
    }
    return m
}
// NewRunRecalculateAllRequestBuilder instantiates a new RunRecalculateAllRequestBuilder and sets the default values.
func NewRunRecalculateAllRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunRecalculateAllRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunRecalculateAllRequestBuilderInternal(urlParams, requestAdapter)
}
// Post recalculate Datasets for Runs between two dates
func (m *RunRecalculateAllRequestBuilder) Post(ctx context.Context, requestConfiguration *RunRecalculateAllRequestBuilderPostRequestConfiguration)(error) {
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
func (m *RunRecalculateAllRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RunRecalculateAllRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunRecalculateAllRequestBuilder when successful
func (m *RunRecalculateAllRequestBuilder) WithUrl(rawUrl string)(*RunRecalculateAllRequestBuilder) {
    return NewRunRecalculateAllRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
