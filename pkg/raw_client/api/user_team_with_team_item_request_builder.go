package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserTeamWithTeamItemRequestBuilder builds and executes requests for operations under \api\user\team\{team}
type UserTeamWithTeamItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserTeamWithTeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserTeamWithTeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserTeamWithTeamItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserTeamWithTeamItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserTeamWithTeamItemRequestBuilderInternal instantiates a new UserTeamWithTeamItemRequestBuilder and sets the default values.
func NewUserTeamWithTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserTeamWithTeamItemRequestBuilder) {
    m := &UserTeamWithTeamItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/team/{team}", pathParameters),
    }
    return m
}
// NewUserTeamWithTeamItemRequestBuilder instantiates a new UserTeamWithTeamItemRequestBuilder and sets the default values.
func NewUserTeamWithTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserTeamWithTeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserTeamWithTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove existing team.
func (m *UserTeamWithTeamItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserTeamWithTeamItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Members the members property
// returns a *UserTeamItemMembersRequestBuilder when successful
func (m *UserTeamWithTeamItemRequestBuilder) Members()(*UserTeamItemMembersRequestBuilder) {
    return NewUserTeamItemMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new team.
func (m *UserTeamWithTeamItemRequestBuilder) Post(ctx context.Context, requestConfiguration *UserTeamWithTeamItemRequestBuilderPostRequestConfiguration)(error) {
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
// ToDeleteRequestInformation remove existing team.
// returns a *RequestInformation when successful
func (m *UserTeamWithTeamItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserTeamWithTeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create new team.
// returns a *RequestInformation when successful
func (m *UserTeamWithTeamItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *UserTeamWithTeamItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserTeamWithTeamItemRequestBuilder when successful
func (m *UserTeamWithTeamItemRequestBuilder) WithUrl(rawUrl string)(*UserTeamWithTeamItemRequestBuilder) {
    return NewUserTeamWithTeamItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
