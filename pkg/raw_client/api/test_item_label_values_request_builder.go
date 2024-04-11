package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// TestItemLabelValuesRequestBuilder builds and executes requests for operations under \api\test\{id}\labelValues
type TestItemLabelValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemLabelValuesRequestBuilderGetQueryParameters list all Label Values for a Test
type TestItemLabelValuesRequestBuilderGetQueryParameters struct {
    // ISO-like date time string or epoch millis
    After *string `uriparametername:"after"`
    // ISO-like date time string or epoch millis
    Before *string `uriparametername:"before"`
    // either Ascending or Descending
    Direction *string `uriparametername:"direction"`
    // label name(s) to exclude from the result as scalar or comma separated
    Exclude []string `uriparametername:"exclude"`
    // either a required json sub-document or path expression
    Filter *string `uriparametername:"filter"`
    // Retrieve values for Filtering Labels
    Filtering *bool `uriparametername:"filtering"`
    // label name(s) to include in the result as scalar or comma separated
    Include []string `uriparametername:"include"`
    // the maximum number of results to include
    Limit *int32 `uriparametername:"limit"`
    // Retrieve values for Metric Labels
    Metrics *bool `uriparametername:"metrics"`
    // which page to skip to when using a limit
    Page *int32 `uriparametername:"page"`
    // json path to sortable value or start or stop for sorting by time
    Sort *string `uriparametername:"sort"`
}
// TestItemLabelValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemLabelValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TestItemLabelValuesRequestBuilderGetQueryParameters
}
// NewTestItemLabelValuesRequestBuilderInternal instantiates a new TestItemLabelValuesRequestBuilder and sets the default values.
func NewTestItemLabelValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemLabelValuesRequestBuilder) {
    m := &TestItemLabelValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/labelValues{?after*,before*,direction*,exclude*,filter*,filtering*,include*,limit*,metrics*,page*,sort*}", pathParameters),
    }
    return m
}
// NewTestItemLabelValuesRequestBuilder instantiates a new TestItemLabelValuesRequestBuilder and sets the default values.
func NewTestItemLabelValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemLabelValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemLabelValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all Label Values for a Test
// returns a []ExportedLabelValuesable when successful
func (m *TestItemLabelValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *TestItemLabelValuesRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ExportedLabelValuesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateExportedLabelValuesFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ExportedLabelValuesable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ExportedLabelValuesable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list all Label Values for a Test
// returns a *RequestInformation when successful
func (m *TestItemLabelValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestItemLabelValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TestItemLabelValuesRequestBuilder when successful
func (m *TestItemLabelValuesRequestBuilder) WithUrl(rawUrl string)(*TestItemLabelValuesRequestBuilder) {
    return NewTestItemLabelValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
