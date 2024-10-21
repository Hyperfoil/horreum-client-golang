package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserApikeyItemRenameRequestBuilder builds and executes requests for operations under \api\user\apikey\{id}\rename
type UserApikeyItemRenameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserApikeyItemRenameRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserApikeyItemRenameRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserApikeyItemRenameRequestBuilderInternal instantiates a new UserApikeyItemRenameRequestBuilder and sets the default values.
func NewUserApikeyItemRenameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserApikeyItemRenameRequestBuilder) {
    m := &UserApikeyItemRenameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/apikey/{id}/rename", pathParameters),
    }
    return m
}
// NewUserApikeyItemRenameRequestBuilder instantiates a new UserApikeyItemRenameRequestBuilder and sets the default values.
func NewUserApikeyItemRenameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserApikeyItemRenameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserApikeyItemRenameRequestBuilderInternal(urlParams, requestAdapter)
}
// Put rename API key.
func (m *UserApikeyItemRenameRequestBuilder) Put(ctx context.Context, body *string, requestConfiguration *UserApikeyItemRenameRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPutRequestInformation rename API key.
// returns a *RequestInformation when successful
func (m *UserApikeyItemRenameRequestBuilder) ToPutRequestInformation(ctx context.Context, body *string, requestConfiguration *UserApikeyItemRenameRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.SetContentFromScalar(ctx, m.BaseRequestBuilder.RequestAdapter, "text/plain", body)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserApikeyItemRenameRequestBuilder when successful
func (m *UserApikeyItemRenameRequestBuilder) WithUrl(rawUrl string)(*UserApikeyItemRenameRequestBuilder) {
    return NewUserApikeyItemRenameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
