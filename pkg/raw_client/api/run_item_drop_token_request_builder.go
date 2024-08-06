package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemDropTokenRequestBuilder builds and executes requests for operations under \api\run\{id}\dropToken
type RunItemDropTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemDropTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemDropTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRunItemDropTokenRequestBuilderInternal instantiates a new RunItemDropTokenRequestBuilder and sets the default values.
func NewRunItemDropTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemDropTokenRequestBuilder) {
    m := &RunItemDropTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/dropToken", pathParameters),
    }
    return m
}
// NewRunItemDropTokenRequestBuilder instantiates a new RunItemDropTokenRequestBuilder and sets the default values.
func NewRunItemDropTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemDropTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemDropTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove access token for Run
// returns a *string when successful
func (m *RunItemDropTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *RunItemDropTokenRequestBuilderPostRequestConfiguration)(*string, error) {
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
// ToPostRequestInformation remove access token for Run
// returns a *RequestInformation when successful
func (m *RunItemDropTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RunItemDropTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemDropTokenRequestBuilder when successful
func (m *RunItemDropTokenRequestBuilder) WithUrl(rawUrl string)(*RunItemDropTokenRequestBuilder) {
    return NewRunItemDropTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
