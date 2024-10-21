package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserApikeyItemRevokeRequestBuilder builds and executes requests for operations under \api\user\apikey\{id}\revoke
type UserApikeyItemRevokeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserApikeyItemRevokeRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserApikeyItemRevokeRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserApikeyItemRevokeRequestBuilderInternal instantiates a new UserApikeyItemRevokeRequestBuilder and sets the default values.
func NewUserApikeyItemRevokeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserApikeyItemRevokeRequestBuilder) {
    m := &UserApikeyItemRevokeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/apikey/{id}/revoke", pathParameters),
    }
    return m
}
// NewUserApikeyItemRevokeRequestBuilder instantiates a new UserApikeyItemRevokeRequestBuilder and sets the default values.
func NewUserApikeyItemRevokeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserApikeyItemRevokeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserApikeyItemRevokeRequestBuilderInternal(urlParams, requestAdapter)
}
// Put revoke API key.
func (m *UserApikeyItemRevokeRequestBuilder) Put(ctx context.Context, requestConfiguration *UserApikeyItemRevokeRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPutRequestInformation revoke API key.
// returns a *RequestInformation when successful
func (m *UserApikeyItemRevokeRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *UserApikeyItemRevokeRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserApikeyItemRevokeRequestBuilder when successful
func (m *UserApikeyItemRevokeRequestBuilder) WithUrl(rawUrl string)(*UserApikeyItemRevokeRequestBuilder) {
    return NewUserApikeyItemRevokeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
