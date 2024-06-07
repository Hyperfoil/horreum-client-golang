package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemDroptokenDropTokenRequestBuilder builds and executes requests for operations under \api\run\{id}\dropToken
type RunItemDroptokenDropTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemDroptokenDropTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemDroptokenDropTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRunItemDroptokenDropTokenRequestBuilderInternal instantiates a new RunItemDroptokenDropTokenRequestBuilder and sets the default values.
func NewRunItemDroptokenDropTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemDroptokenDropTokenRequestBuilder) {
    m := &RunItemDroptokenDropTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/dropToken", pathParameters),
    }
    return m
}
// NewRunItemDroptokenDropTokenRequestBuilder instantiates a new RunItemDroptokenDropTokenRequestBuilder and sets the default values.
func NewRunItemDroptokenDropTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemDroptokenDropTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemDroptokenDropTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove access token for Run
// returns a *string when successful
func (m *RunItemDroptokenDropTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *RunItemDroptokenDropTokenRequestBuilderPostRequestConfiguration)(*string, error) {
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
func (m *RunItemDroptokenDropTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RunItemDroptokenDropTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemDroptokenDropTokenRequestBuilder when successful
func (m *RunItemDroptokenDropTokenRequestBuilder) WithUrl(rawUrl string)(*RunItemDroptokenDropTokenRequestBuilder) {
    return NewRunItemDroptokenDropTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
