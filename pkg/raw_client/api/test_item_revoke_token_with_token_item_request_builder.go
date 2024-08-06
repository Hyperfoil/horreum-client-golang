package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TestItemRevokeTokenWithTokenItemRequestBuilder builds and executes requests for operations under \api\test\{id}\revokeToken\{tokenId}
type TestItemRevokeTokenWithTokenItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemRevokeTokenWithTokenItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemRevokeTokenWithTokenItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTestItemRevokeTokenWithTokenItemRequestBuilderInternal instantiates a new TestItemRevokeTokenWithTokenItemRequestBuilder and sets the default values.
func NewTestItemRevokeTokenWithTokenItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemRevokeTokenWithTokenItemRequestBuilder) {
    m := &TestItemRevokeTokenWithTokenItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/revokeToken/{tokenId}", pathParameters),
    }
    return m
}
// NewTestItemRevokeTokenWithTokenItemRequestBuilder instantiates a new TestItemRevokeTokenWithTokenItemRequestBuilder and sets the default values.
func NewTestItemRevokeTokenWithTokenItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemRevokeTokenWithTokenItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemRevokeTokenWithTokenItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete revoke a Token defined for a Test
func (m *TestItemRevokeTokenWithTokenItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TestItemRevokeTokenWithTokenItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TestItemRevokeTokenWithTokenItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TestItemRevokeTokenWithTokenItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestItemRevokeTokenWithTokenItemRequestBuilder when successful
func (m *TestItemRevokeTokenWithTokenItemRequestBuilder) WithUrl(rawUrl string)(*TestItemRevokeTokenWithTokenItemRequestBuilder) {
    return NewTestItemRevokeTokenWithTokenItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
