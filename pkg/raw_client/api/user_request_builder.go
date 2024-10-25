package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserRequestBuilder builds and executes requests for operations under \api\user
type UserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Administrators the administrators property
// returns a *UserAdministratorsRequestBuilder when successful
func (m *UserRequestBuilder) Administrators()(*UserAdministratorsRequestBuilder) {
    return NewUserAdministratorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AllTeams the allTeams property
// returns a *UserAllTeamsRequestBuilder when successful
func (m *UserRequestBuilder) AllTeams()(*UserAllTeamsRequestBuilder) {
    return NewUserAllTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apikey the apikey property
// returns a *UserApikeyRequestBuilder when successful
func (m *UserRequestBuilder) Apikey()(*UserApikeyRequestBuilder) {
    return NewUserApikeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByUsername gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.user.item collection
// returns a *UserWithUsernameItemRequestBuilder when successful
func (m *UserRequestBuilder) ByUsername(username string)(*UserWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewUserWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user", pathParameters),
    }
    return m
}
// NewUserRequestBuilder instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateUser the createUser property
// returns a *UserCreateUserRequestBuilder when successful
func (m *UserRequestBuilder) CreateUser()(*UserCreateUserRequestBuilder) {
    return NewUserCreateUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DefaultTeam the defaultTeam property
// returns a *UserDefaultTeamRequestBuilder when successful
func (m *UserRequestBuilder) DefaultTeam()(*UserDefaultTeamRequestBuilder) {
    return NewUserDefaultTeamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Info the info property
// returns a *UserInfoRequestBuilder when successful
func (m *UserRequestBuilder) Info()(*UserInfoRequestBuilder) {
    return NewUserInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Roles the roles property
// returns a *UserRolesRequestBuilder when successful
func (m *UserRequestBuilder) Roles()(*UserRolesRequestBuilder) {
    return NewUserRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
// returns a *UserSearchRequestBuilder when successful
func (m *UserRequestBuilder) Search()(*UserSearchRequestBuilder) {
    return NewUserSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Team the team property
// returns a *UserTeamRequestBuilder when successful
func (m *UserRequestBuilder) Team()(*UserTeamRequestBuilder) {
    return NewUserTeamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
// returns a *UserTeamsRequestBuilder when successful
func (m *UserRequestBuilder) Teams()(*UserTeamsRequestBuilder) {
    return NewUserTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}