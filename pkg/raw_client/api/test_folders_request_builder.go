package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TestFoldersRequestBuilder builds and executes requests for operations under \api\test\folders
type TestFoldersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestFoldersRequestBuilderGetQueryParameters retrieve a list of all folders
type TestFoldersRequestBuilderGetQueryParameters struct {
    // "__my", "__all" or a comma delimited  list of roles
    Roles *string `uriparametername:"roles"`
}
// TestFoldersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestFoldersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TestFoldersRequestBuilderGetQueryParameters
}
// NewTestFoldersRequestBuilderInternal instantiates a new TestFoldersRequestBuilder and sets the default values.
func NewTestFoldersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestFoldersRequestBuilder) {
    m := &TestFoldersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/folders{?roles*}", pathParameters),
    }
    return m
}
// NewTestFoldersRequestBuilder instantiates a new TestFoldersRequestBuilder and sets the default values.
func NewTestFoldersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestFoldersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestFoldersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a list of all folders
// returns a []string when successful
func (m *TestFoldersRequestBuilder) Get(ctx context.Context, requestConfiguration *TestFoldersRequestBuilderGetRequestConfiguration)([]string, error) {
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
// ToGetRequestInformation retrieve a list of all folders
// returns a *RequestInformation when successful
func (m *TestFoldersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestFoldersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestFoldersRequestBuilder when successful
func (m *TestFoldersRequestBuilder) WithUrl(rawUrl string)(*TestFoldersRequestBuilder) {
    return NewTestFoldersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
