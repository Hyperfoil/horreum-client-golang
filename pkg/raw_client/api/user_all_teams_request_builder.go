package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserAllTeamsRequestBuilder builds and executes requests for operations under \api\user\allTeams
type UserAllTeamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserAllTeamsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserAllTeamsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserAllTeamsRequestBuilderInternal instantiates a new UserAllTeamsRequestBuilder and sets the default values.
func NewUserAllTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserAllTeamsRequestBuilder) {
    m := &UserAllTeamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/allTeams", pathParameters),
    }
    return m
}
// NewUserAllTeamsRequestBuilder instantiates a new UserAllTeamsRequestBuilder and sets the default values.
func NewUserAllTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserAllTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserAllTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get list of all teams.
// returns a []string when successful
func (m *UserAllTeamsRequestBuilder) Get(ctx context.Context, requestConfiguration *UserAllTeamsRequestBuilderGetRequestConfiguration)([]string, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", nil)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*string))
        }
    }
    return val, nil
}
// ToGetRequestInformation get list of all teams.
// returns a *RequestInformation when successful
func (m *UserAllTeamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserAllTeamsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserAllTeamsRequestBuilder when successful
func (m *UserAllTeamsRequestBuilder) WithUrl(rawUrl string)(*UserAllTeamsRequestBuilder) {
    return NewUserAllTeamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
