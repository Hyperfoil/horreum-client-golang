package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserCreateUserRequestBuilder builds and executes requests for operations under \api\user\createUser
type UserCreateUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserCreateUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserCreateUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserCreateUserRequestBuilderInternal instantiates a new UserCreateUserRequestBuilder and sets the default values.
func NewUserCreateUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserCreateUserRequestBuilder) {
    m := &UserCreateUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/createUser", pathParameters),
    }
    return m
}
// NewUserCreateUserRequestBuilder instantiates a new UserCreateUserRequestBuilder and sets the default values.
func NewUserCreateUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserCreateUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserCreateUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create new user.
func (m *UserCreateUserRequestBuilder) Post(ctx context.Context, body UserCreateUserPostRequestBodyable, requestConfiguration *UserCreateUserRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation create new user.
// returns a *RequestInformation when successful
func (m *UserCreateUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body UserCreateUserPostRequestBodyable, requestConfiguration *UserCreateUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserCreateUserRequestBuilder when successful
func (m *UserCreateUserRequestBuilder) WithUrl(rawUrl string)(*UserCreateUserRequestBuilder) {
    return NewUserCreateUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}