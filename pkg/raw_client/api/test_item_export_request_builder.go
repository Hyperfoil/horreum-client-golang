package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b "github.com/hyperfoil/horreum/pkg/raw_client/models"
)

// TestItemExportRequestBuilder builds and executes requests for operations under \api\test\{id}\export
type TestItemExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemExportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemExportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTestItemExportRequestBuilderInternal instantiates a new TestItemExportRequestBuilder and sets the default values.
func NewTestItemExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemExportRequestBuilder) {
    m := &TestItemExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/export", pathParameters),
    }
    return m
}
// NewTestItemExportRequestBuilder instantiates a new TestItemExportRequestBuilder and sets the default values.
func NewTestItemExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemExportRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a TestExportable when successful
func (m *TestItemExportRequestBuilder) Get(ctx context.Context, requestConfiguration *TestItemExportRequestBuilderGetRequestConfiguration)(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TestExportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.CreateTestExportFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TestExportable), nil
}
// returns a *RequestInformation when successful
func (m *TestItemExportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestItemExportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestItemExportRequestBuilder when successful
func (m *TestItemExportRequestBuilder) WithUrl(rawUrl string)(*TestItemExportRequestBuilder) {
    return NewTestItemExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
