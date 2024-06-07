package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TestItemRevoketokenWithTokenItemRequestBuilder builds and executes requests for operations under \api\test\{id}\revokeToken\{tokenId}
type TestItemRevoketokenWithTokenItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemRevoketokenWithTokenItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemRevoketokenWithTokenItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTestItemRevoketokenWithTokenItemRequestBuilderInternal instantiates a new TestItemRevoketokenWithTokenItemRequestBuilder and sets the default values.
func NewTestItemRevoketokenWithTokenItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemRevoketokenWithTokenItemRequestBuilder) {
    m := &TestItemRevoketokenWithTokenItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/revokeToken/{tokenId}", pathParameters),
    }
    return m
}
// NewTestItemRevoketokenWithTokenItemRequestBuilder instantiates a new TestItemRevoketokenWithTokenItemRequestBuilder and sets the default values.
func NewTestItemRevoketokenWithTokenItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemRevoketokenWithTokenItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemRevoketokenWithTokenItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete revoke a Token defined for a Test
func (m *TestItemRevoketokenWithTokenItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TestItemRevoketokenWithTokenItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation revoke a Token defined for a Test
// returns a *RequestInformation when successful
func (m *TestItemRevoketokenWithTokenItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TestItemRevoketokenWithTokenItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestItemRevoketokenWithTokenItemRequestBuilder when successful
func (m *TestItemRevoketokenWithTokenItemRequestBuilder) WithUrl(rawUrl string)(*TestItemRevoketokenWithTokenItemRequestBuilder) {
    return NewTestItemRevoketokenWithTokenItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
