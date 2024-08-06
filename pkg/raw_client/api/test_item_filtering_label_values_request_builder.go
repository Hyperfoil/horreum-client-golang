package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TestItemFilteringLabelValuesRequestBuilder builds and executes requests for operations under \api\test\{id}\filteringLabelValues
type TestItemFilteringLabelValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemFilteringLabelValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemFilteringLabelValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTestItemFilteringLabelValuesRequestBuilderInternal instantiates a new TestItemFilteringLabelValuesRequestBuilder and sets the default values.
func NewTestItemFilteringLabelValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemFilteringLabelValuesRequestBuilder) {
    m := &TestItemFilteringLabelValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/filteringLabelValues", pathParameters),
    }
    return m
}
// NewTestItemFilteringLabelValuesRequestBuilder instantiates a new TestItemFilteringLabelValuesRequestBuilder and sets the default values.
func NewTestItemFilteringLabelValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemFilteringLabelValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemFilteringLabelValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all unique Label Values for a Test
// returns a UntypedNodeable when successful
func (m *TestItemFilteringLabelValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *TestItemFilteringLabelValuesRequestBuilderGetRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable), nil
}
// ToGetRequestInformation list all unique Label Values for a Test
// returns a *RequestInformation when successful
func (m *TestItemFilteringLabelValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestItemFilteringLabelValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestItemFilteringLabelValuesRequestBuilder when successful
func (m *TestItemFilteringLabelValuesRequestBuilder) WithUrl(rawUrl string)(*TestItemFilteringLabelValuesRequestBuilder) {
    return NewTestItemFilteringLabelValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
