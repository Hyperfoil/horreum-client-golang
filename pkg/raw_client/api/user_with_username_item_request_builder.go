package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserWithUsernameItemRequestBuilder builds and executes requests for operations under \api\user\{username}
type UserWithUsernameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserWithUsernameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserWithUsernameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserWithUsernameItemRequestBuilderInternal instantiates a new UserWithUsernameItemRequestBuilder and sets the default values.
func NewUserWithUsernameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserWithUsernameItemRequestBuilder) {
    m := &UserWithUsernameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/{username}", pathParameters),
    }
    return m
}
// NewUserWithUsernameItemRequestBuilder instantiates a new UserWithUsernameItemRequestBuilder and sets the default values.
func NewUserWithUsernameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserWithUsernameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserWithUsernameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove existing user.
func (m *UserWithUsernameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserWithUsernameItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation remove existing user.
// returns a *RequestInformation when successful
func (m *UserWithUsernameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserWithUsernameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserWithUsernameItemRequestBuilder when successful
func (m *UserWithUsernameItemRequestBuilder) WithUrl(rawUrl string)(*UserWithUsernameItemRequestBuilder) {
    return NewUserWithUsernameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
