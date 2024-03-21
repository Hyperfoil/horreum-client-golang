package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b "github.com/hyperfoil/horreum/pkg/raw_client/models"
)

// TestSummaryRequestBuilder builds and executes requests for operations under \api\test\summary
type TestSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestSummaryRequestBuilderGetQueryParameters retrieve a summary of Tests in a folder
type TestSummaryRequestBuilderGetQueryParameters struct {
    // Sort direction
    Direction *string `uriparametername:"direction"`
    // name of the Folder containing the Tests
    Folder *string `uriparametername:"folder"`
    // limit the result count
    Limit *int32 `uriparametername:"limit"`
    // filter by page number of a paginated list of 
    Page *int32 `uriparametername:"page"`
    // "__my", "__all" or a comma delimited  list of roles
    Roles *string `uriparametername:"roles"`
}
// TestSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TestSummaryRequestBuilderGetQueryParameters
}
// NewTestSummaryRequestBuilderInternal instantiates a new TestSummaryRequestBuilder and sets the default values.
func NewTestSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestSummaryRequestBuilder) {
    m := &TestSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/summary{?direction*,folder*,limit*,page*,roles*}", pathParameters),
    }
    return m
}
// NewTestSummaryRequestBuilder instantiates a new TestSummaryRequestBuilder and sets the default values.
func NewTestSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a summary of Tests in a folder
// returns a TestListingable when successful
func (m *TestSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *TestSummaryRequestBuilderGetRequestConfiguration)(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TestListingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.CreateTestListingFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i474b066f9576d008d7de8ccd52cbe2ceff5e0826fad92c1bbc3202f77dfa272b.TestListingable), nil
}
// ToGetRequestInformation retrieve a summary of Tests in a folder
// returns a *RequestInformation when successful
func (m *TestSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TestSummaryRequestBuilder when successful
func (m *TestSummaryRequestBuilder) WithUrl(rawUrl string)(*TestSummaryRequestBuilder) {
    return NewTestSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
