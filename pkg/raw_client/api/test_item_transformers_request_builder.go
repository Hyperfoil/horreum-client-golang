package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TestItemTransformersRequestBuilder builds and executes requests for operations under \api\test\{id}\transformers
type TestItemTransformersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemTransformersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemTransformersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTestItemTransformersRequestBuilderInternal instantiates a new TestItemTransformersRequestBuilder and sets the default values.
func NewTestItemTransformersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemTransformersRequestBuilder) {
    m := &TestItemTransformersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/transformers", pathParameters),
    }
    return m
}
// NewTestItemTransformersRequestBuilder instantiates a new TestItemTransformersRequestBuilder and sets the default values.
func NewTestItemTransformersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemTransformersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemTransformersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update transformers for Test
func (m *TestItemTransformersRequestBuilder) Post(ctx context.Context, body []int32, requestConfiguration *TestItemTransformersRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation update transformers for Test
// returns a *RequestInformation when successful
func (m *TestItemTransformersRequestBuilder) ToPostRequestInformation(ctx context.Context, body []int32, requestConfiguration *TestItemTransformersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    cast := make([]interface{}, len(body))
    for i, v := range body {
        cast[i] = v
    }
    requestInfo.SetContentFromScalarCollection(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", cast)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestItemTransformersRequestBuilder when successful
func (m *TestItemTransformersRequestBuilder) WithUrl(rawUrl string)(*TestItemTransformersRequestBuilder) {
    return NewTestItemTransformersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
