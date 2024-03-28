package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemRecalculateRequestBuilder builds and executes requests for operations under \api\run\{id}\recalculate
type RunItemRecalculateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemRecalculateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemRecalculateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRunItemRecalculateRequestBuilderInternal instantiates a new RunItemRecalculateRequestBuilder and sets the default values.
func NewRunItemRecalculateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemRecalculateRequestBuilder) {
    m := &RunItemRecalculateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/recalculate", pathParameters),
    }
    return m
}
// NewRunItemRecalculateRequestBuilder instantiates a new RunItemRecalculateRequestBuilder and sets the default values.
func NewRunItemRecalculateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemRecalculateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemRecalculateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post recalculate Datasets for Run
// returns a []int32 when successful
func (m *RunItemRecalculateRequestBuilder) Post(ctx context.Context, requestConfiguration *RunItemRecalculateRequestBuilderPostRequestConfiguration)([]int32, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    val := make([]int32, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*int32))
        }
    }
    return val, nil
}
// ToPostRequestInformation recalculate Datasets for Run
// returns a *RequestInformation when successful
func (m *RunItemRecalculateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RunItemRecalculateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemRecalculateRequestBuilder when successful
func (m *RunItemRecalculateRequestBuilder) WithUrl(rawUrl string)(*RunItemRecalculateRequestBuilder) {
    return NewRunItemRecalculateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
