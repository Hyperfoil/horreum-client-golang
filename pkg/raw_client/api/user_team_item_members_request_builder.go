package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserTeamItemMembersRequestBuilder builds and executes requests for operations under \api\user\team\{team}\members
type UserTeamItemMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserTeamItemMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserTeamItemMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserTeamItemMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserTeamItemMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserTeamItemMembersRequestBuilderInternal instantiates a new UserTeamItemMembersRequestBuilder and sets the default values.
func NewUserTeamItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserTeamItemMembersRequestBuilder) {
    m := &UserTeamItemMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/team/{team}/members", pathParameters),
    }
    return m
}
// NewUserTeamItemMembersRequestBuilder instantiates a new UserTeamItemMembersRequestBuilder and sets the default values.
func NewUserTeamItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserTeamItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserTeamItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the membership of a given team.
// Deprecated: This method is obsolete. Use GetAsMembersGetResponse instead.
// returns a UserTeamItemMembersResponseable when successful
func (m *UserTeamItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *UserTeamItemMembersRequestBuilderGetRequestConfiguration)(UserTeamItemMembersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserTeamItemMembersResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserTeamItemMembersResponseable), nil
}
// GetAsMembersGetResponse get the membership of a given team.
// returns a UserTeamItemMembersGetResponseable when successful
func (m *UserTeamItemMembersRequestBuilder) GetAsMembersGetResponse(ctx context.Context, requestConfiguration *UserTeamItemMembersRequestBuilderGetRequestConfiguration)(UserTeamItemMembersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserTeamItemMembersGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserTeamItemMembersGetResponseable), nil
}
// Post set the membership of a given team.
func (m *UserTeamItemMembersRequestBuilder) Post(ctx context.Context, body UserTeamItemMembersPostRequestBodyable, requestConfiguration *UserTeamItemMembersRequestBuilderPostRequestConfiguration)(error) {
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
// ToGetRequestInformation get the membership of a given team.
// returns a *RequestInformation when successful
func (m *UserTeamItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserTeamItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation set the membership of a given team.
// returns a *RequestInformation when successful
func (m *UserTeamItemMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body UserTeamItemMembersPostRequestBodyable, requestConfiguration *UserTeamItemMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserTeamItemMembersRequestBuilder when successful
func (m *UserTeamItemMembersRequestBuilder) WithUrl(rawUrl string)(*UserTeamItemMembersRequestBuilder) {
    return NewUserTeamItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}