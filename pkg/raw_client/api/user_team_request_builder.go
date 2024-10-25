package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserTeamRequestBuilder builds and executes requests for operations under \api\user\team
type UserTeamRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTeam gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.user.team.item collection
// returns a *UserTeamWithTeamItemRequestBuilder when successful
func (m *UserTeamRequestBuilder) ByTeam(team string)(*UserTeamWithTeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if team != "" {
        urlTplParams["team"] = team
    }
    return NewUserTeamWithTeamItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserTeamRequestBuilderInternal instantiates a new UserTeamRequestBuilder and sets the default values.
func NewUserTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserTeamRequestBuilder) {
    m := &UserTeamRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/team", pathParameters),
    }
    return m
}
// NewUserTeamRequestBuilder instantiates a new UserTeamRequestBuilder and sets the default values.
func NewUserTeamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserTeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserTeamRequestBuilderInternal(urlParams, requestAdapter)
}