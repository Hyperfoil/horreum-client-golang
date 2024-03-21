package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TestItemNotificationsRequestBuilder builds and executes requests for operations under \api\test\{id}\notifications
type TestItemNotificationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemNotificationsRequestBuilderPostQueryParameters update notifications for a Test. It is possible to disable notifications for a Test, so that no notifications are sent to subscribers
type TestItemNotificationsRequestBuilderPostQueryParameters struct {
    // Whether notifications are enabled
    Enabled *bool `uriparametername:"enabled"`
}
// TestItemNotificationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemNotificationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TestItemNotificationsRequestBuilderPostQueryParameters
}
// NewTestItemNotificationsRequestBuilderInternal instantiates a new TestItemNotificationsRequestBuilder and sets the default values.
func NewTestItemNotificationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemNotificationsRequestBuilder) {
    m := &TestItemNotificationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/notifications?enabled={enabled}", pathParameters),
    }
    return m
}
// NewTestItemNotificationsRequestBuilder instantiates a new TestItemNotificationsRequestBuilder and sets the default values.
func NewTestItemNotificationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemNotificationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemNotificationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update notifications for a Test. It is possible to disable notifications for a Test, so that no notifications are sent to subscribers
func (m *TestItemNotificationsRequestBuilder) Post(ctx context.Context, requestConfiguration *TestItemNotificationsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation update notifications for a Test. It is possible to disable notifications for a Test, so that no notifications are sent to subscribers
// returns a *RequestInformation when successful
func (m *TestItemNotificationsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TestItemNotificationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TestItemNotificationsRequestBuilder when successful
func (m *TestItemNotificationsRequestBuilder) WithUrl(rawUrl string)(*TestItemNotificationsRequestBuilder) {
    return NewTestItemNotificationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
