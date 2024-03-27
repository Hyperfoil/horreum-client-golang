package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TestItemMoveRequestBuilder builds and executes requests for operations under \api\test\{id}\move
type TestItemMoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemMoveRequestBuilderPostQueryParameters update the folder for a Test. Tests can be moved to different folders
type TestItemMoveRequestBuilderPostQueryParameters struct {
    // New folder to store the tests
    Folder *string `uriparametername:"folder"`
}
// TestItemMoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemMoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TestItemMoveRequestBuilderPostQueryParameters
}
// NewTestItemMoveRequestBuilderInternal instantiates a new TestItemMoveRequestBuilder and sets the default values.
func NewTestItemMoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemMoveRequestBuilder) {
    m := &TestItemMoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/move{?folder*}", pathParameters),
    }
    return m
}
// NewTestItemMoveRequestBuilder instantiates a new TestItemMoveRequestBuilder and sets the default values.
func NewTestItemMoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemMoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemMoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the folder for a Test. Tests can be moved to different folders
func (m *TestItemMoveRequestBuilder) Post(ctx context.Context, requestConfiguration *TestItemMoveRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation update the folder for a Test. Tests can be moved to different folders
// returns a *RequestInformation when successful
func (m *TestItemMoveRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TestItemMoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TestItemMoveRequestBuilder when successful
func (m *TestItemMoveRequestBuilder) WithUrl(rawUrl string)(*TestItemMoveRequestBuilder) {
    return NewTestItemMoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
