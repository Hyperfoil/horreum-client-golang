package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemDescriptionRequestBuilder builds and executes requests for operations under \api\run\{id}\description
type RunItemDescriptionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemDescriptionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemDescriptionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRunItemDescriptionRequestBuilderInternal instantiates a new RunItemDescriptionRequestBuilder and sets the default values.
func NewRunItemDescriptionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemDescriptionRequestBuilder) {
    m := &RunItemDescriptionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/description", pathParameters),
    }
    return m
}
// NewRunItemDescriptionRequestBuilder instantiates a new RunItemDescriptionRequestBuilder and sets the default values.
func NewRunItemDescriptionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemDescriptionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemDescriptionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update Run description
func (m *RunItemDescriptionRequestBuilder) Post(ctx context.Context, body *string, requestConfiguration *RunItemDescriptionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation update Run description
// returns a *RequestInformation when successful
func (m *RunItemDescriptionRequestBuilder) ToPostRequestInformation(ctx context.Context, body *string, requestConfiguration *RunItemDescriptionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.SetContentFromScalar(ctx, m.BaseRequestBuilder.RequestAdapter, "text/plain", body)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemDescriptionRequestBuilder when successful
func (m *RunItemDescriptionRequestBuilder) WithUrl(rawUrl string)(*RunItemDescriptionRequestBuilder) {
    return NewRunItemDescriptionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
