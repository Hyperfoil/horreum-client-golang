package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserRolesRequestBuilder builds and executes requests for operations under \api\user\roles
type UserRolesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserRolesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserRolesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserRolesRequestBuilderInternal instantiates a new UserRolesRequestBuilder and sets the default values.
func NewUserRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRolesRequestBuilder) {
    m := &UserRolesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/roles", pathParameters),
    }
    return m
}
// NewUserRolesRequestBuilder instantiates a new UserRolesRequestBuilder and sets the default values.
func NewUserRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get roles for the authenticated user.
// returns a []string when successful
func (m *UserRolesRequestBuilder) Get(ctx context.Context, requestConfiguration *UserRolesRequestBuilderGetRequestConfiguration)([]string, error) {
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
// ToGetRequestInformation get roles for the authenticated user.
// returns a *RequestInformation when successful
func (m *UserRolesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserRolesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserRolesRequestBuilder when successful
func (m *UserRolesRequestBuilder) WithUrl(rawUrl string)(*UserRolesRequestBuilder) {
    return NewUserRolesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
