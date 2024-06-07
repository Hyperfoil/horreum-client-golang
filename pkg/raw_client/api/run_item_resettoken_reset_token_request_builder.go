package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemResettokenResetTokenRequestBuilder builds and executes requests for operations under \api\run\{id}\resetToken
type RunItemResettokenResetTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemResettokenResetTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemResettokenResetTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRunItemResettokenResetTokenRequestBuilderInternal instantiates a new RunItemResettokenResetTokenRequestBuilder and sets the default values.
func NewRunItemResettokenResetTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemResettokenResetTokenRequestBuilder) {
    m := &RunItemResettokenResetTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/resetToken", pathParameters),
    }
    return m
}
// NewRunItemResettokenResetTokenRequestBuilder instantiates a new RunItemResettokenResetTokenRequestBuilder and sets the default values.
func NewRunItemResettokenResetTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemResettokenResetTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemResettokenResetTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post regenerate access token for Run
// returns a *string when successful
func (m *RunItemResettokenResetTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *RunItemResettokenResetTokenRequestBuilderPostRequestConfiguration)(*string, error) {
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
// ToPostRequestInformation regenerate access token for Run
// returns a *RequestInformation when successful
func (m *RunItemResettokenResetTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RunItemResettokenResetTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemResettokenResetTokenRequestBuilder when successful
func (m *RunItemResettokenResetTokenRequestBuilder) WithUrl(rawUrl string)(*RunItemResettokenResetTokenRequestBuilder) {
    return NewRunItemResettokenResetTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
