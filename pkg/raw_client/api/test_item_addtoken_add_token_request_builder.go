package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// TestItemAddtokenAddTokenRequestBuilder builds and executes requests for operations under \api\test\{id}\addToken
type TestItemAddtokenAddTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemAddtokenAddTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemAddtokenAddTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTestItemAddtokenAddTokenRequestBuilderInternal instantiates a new TestItemAddtokenAddTokenRequestBuilder and sets the default values.
func NewTestItemAddtokenAddTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemAddtokenAddTokenRequestBuilder) {
    m := &TestItemAddtokenAddTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/addToken", pathParameters),
    }
    return m
}
// NewTestItemAddtokenAddTokenRequestBuilder instantiates a new TestItemAddtokenAddTokenRequestBuilder and sets the default values.
func NewTestItemAddtokenAddTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemAddtokenAddTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemAddtokenAddTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add a Test API Token for access to provide access to a test data for integrated tooling, e.g. reporting services
// returns a *int32 when successful
func (m *TestItemAddtokenAddTokenRequestBuilder) Post(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.TestTokenable, requestConfiguration *TestItemAddtokenAddTokenRequestBuilderPostRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToPostRequestInformation add a Test API Token for access to provide access to a test data for integrated tooling, e.g. reporting services
// returns a *RequestInformation when successful
func (m *TestItemAddtokenAddTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.TestTokenable, requestConfiguration *TestItemAddtokenAddTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestItemAddtokenAddTokenRequestBuilder when successful
func (m *TestItemAddtokenAddTokenRequestBuilder) WithUrl(rawUrl string)(*TestItemAddtokenAddTokenRequestBuilder) {
    return NewTestItemAddtokenAddTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
