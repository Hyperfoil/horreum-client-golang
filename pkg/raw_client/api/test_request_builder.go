package api

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// TestRequestBuilder builds and executes requests for operations under \api\test
type TestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestRequestBuilderGetQueryParameters retrieve a paginated list of Tests with available count
type TestRequestBuilderGetQueryParameters struct {
    // Sort direction
    // Deprecated: This property is deprecated, use DirectionAsSortDirection instead
    Direction *string `uriparametername:"direction"`
    // Sort direction
    DirectionAsSortDirection *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SortDirection `uriparametername:"direction"`
    // limit the number of results
    Limit *int32 `uriparametername:"limit"`
    // filter by page number of a paginated list of Tests
    Page *int32 `uriparametername:"page"`
    // __my, __all or a comma delimited  list of roles
    Roles *string `uriparametername:"roles"`
    // Field name to sort results
    Sort *string `uriparametername:"sort"`
}
// TestRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TestRequestBuilderGetQueryParameters
}
// TestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.test.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *TestTestItemRequestBuilder when successful
func (m *TestRequestBuilder) ById(id string)(*TestTestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewTestTestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdInteger gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.test.item collection
// returns a *TestTestItemRequestBuilder when successful
func (m *TestRequestBuilder) ByIdInteger(id int32)(*TestTestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(id), 10)
    return NewTestTestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByName the byName property
// returns a *TestByNameRequestBuilder when successful
func (m *TestRequestBuilder) ByName()(*TestByNameRequestBuilder) {
    return NewTestByNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTestRequestBuilderInternal instantiates a new TestRequestBuilder and sets the default values.
func NewTestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestRequestBuilder) {
    m := &TestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test{?direction*,limit*,page*,roles*,sort*}", pathParameters),
    }
    return m
}
// NewTestRequestBuilder instantiates a new TestRequestBuilder and sets the default values.
func NewTestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestRequestBuilderInternal(urlParams, requestAdapter)
}
// Folders the folders property
// returns a *TestFoldersRequestBuilder when successful
func (m *TestRequestBuilder) Folders()(*TestFoldersRequestBuilder) {
    return NewTestFoldersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a paginated list of Tests with available count
// returns a TestQueryResultable when successful
func (m *TestRequestBuilder) Get(ctx context.Context, requestConfiguration *TestRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.TestQueryResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateTestQueryResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.TestQueryResultable), nil
}
// ImportEscaped the import property
// returns a *TestImportRequestBuilder when successful
func (m *TestRequestBuilder) ImportEscaped()(*TestImportRequestBuilder) {
    return NewTestImportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create a new test
// returns a Testable when successful
func (m *TestRequestBuilder) Post(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Testable, requestConfiguration *TestRequestBuilderPostRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Testable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateTestFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Testable), nil
}
// Summary the summary property
// returns a *TestSummaryRequestBuilder when successful
func (m *TestRequestBuilder) Summary()(*TestSummaryRequestBuilder) {
    return NewTestSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieve a paginated list of Tests with available count
// returns a *RequestInformation when successful
func (m *TestRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new test
// returns a *RequestInformation when successful
func (m *TestRequestBuilder) ToPostRequestInformation(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Testable, requestConfiguration *TestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestRequestBuilder when successful
func (m *TestRequestBuilder) WithUrl(rawUrl string)(*TestRequestBuilder) {
    return NewTestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
