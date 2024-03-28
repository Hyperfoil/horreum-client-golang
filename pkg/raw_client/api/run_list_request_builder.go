package api

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunListRequestBuilder builds and executes requests for operations under \api\run\list
type RunListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunListRequestBuilderGetQueryParameters retrieve a paginated list of Runs with available count
type RunListRequestBuilderGetQueryParameters struct {
    // Sort direction
    // Deprecated: This property is deprecated, use DirectionAsSortDirection instead
    Direction *string `uriparametername:"direction"`
    // Sort direction
    DirectionAsSortDirection *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SortDirection `uriparametername:"direction"`
    // limit the number of results
    Limit *int32 `uriparametername:"limit"`
    // match all Runs?
    MatchAll *bool `uriparametername:"matchAll"`
    // filter by page number of a paginated list of Tests
    Page *int32 `uriparametername:"page"`
    // query string to filter runs
    Query *string `uriparametername:"query"`
    // __my, __all or a comma delimited  list of roles
    Roles *string `uriparametername:"roles"`
    // Field name to sort results
    Sort *string `uriparametername:"sort"`
    // show trashed runs
    Trashed *bool `uriparametername:"trashed"`
}
// RunListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunListRequestBuilderGetQueryParameters
}
// ByTestId gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.run.list.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *RunListWithTestItemRequestBuilder when successful
func (m *RunListRequestBuilder) ByTestId(testId string)(*RunListWithTestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if testId != "" {
        urlTplParams["testId"] = testId
    }
    return NewRunListWithTestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByTestIdInteger gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.run.list.item collection
// returns a *RunListWithTestItemRequestBuilder when successful
func (m *RunListRequestBuilder) ByTestIdInteger(testId int32)(*RunListWithTestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["testId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(testId), 10)
    return NewRunListWithTestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRunListRequestBuilderInternal instantiates a new RunListRequestBuilder and sets the default values.
func NewRunListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunListRequestBuilder) {
    m := &RunListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/list{?direction*,limit*,matchAll*,page*,query*,roles*,sort*,trashed*}", pathParameters),
    }
    return m
}
// NewRunListRequestBuilder instantiates a new RunListRequestBuilder and sets the default values.
func NewRunListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunListRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a paginated list of Runs with available count
// returns a RunsSummaryable when successful
func (m *RunListRequestBuilder) Get(ctx context.Context, requestConfiguration *RunListRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunsSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateRunsSummaryFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunsSummaryable), nil
}
// ToGetRequestInformation retrieve a paginated list of Runs with available count
// returns a *RequestInformation when successful
func (m *RunListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunListRequestBuilder when successful
func (m *RunListRequestBuilder) WithUrl(rawUrl string)(*RunListRequestBuilder) {
    return NewRunListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
