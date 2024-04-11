package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunItemLabelValuesRequestBuilder builds and executes requests for operations under \api\run\{id}\labelValues
type RunItemLabelValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemLabelValuesRequestBuilderGetQueryParameters get all the label values for the run
type RunItemLabelValuesRequestBuilderGetQueryParameters struct {
    // either Ascending or Descending
    Direction *string `uriparametername:"direction"`
    // label name(s) to exclude from the result as scalar or comma separated
    Exclude []string `uriparametername:"exclude"`
    // either a required json sub-document or path expression
    Filter *string `uriparametername:"filter"`
    // label name(s) to include in the result as scalar or comma separated
    Include []string `uriparametername:"include"`
    // the maximum number of results to include
    Limit *int32 `uriparametername:"limit"`
    // which page to skip to when using a limit
    Page *int32 `uriparametername:"page"`
    // label name for sorting
    Sort *string `uriparametername:"sort"`
}
// RunItemLabelValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemLabelValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunItemLabelValuesRequestBuilderGetQueryParameters
}
// NewRunItemLabelValuesRequestBuilderInternal instantiates a new RunItemLabelValuesRequestBuilder and sets the default values.
func NewRunItemLabelValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemLabelValuesRequestBuilder) {
    m := &RunItemLabelValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/labelValues{?direction*,exclude*,filter*,include*,limit*,page*,sort*}", pathParameters),
    }
    return m
}
// NewRunItemLabelValuesRequestBuilder instantiates a new RunItemLabelValuesRequestBuilder and sets the default values.
func NewRunItemLabelValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemLabelValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemLabelValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all the label values for the run
// returns a []ExportedLabelValuesable when successful
func (m *RunItemLabelValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *RunItemLabelValuesRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ExportedLabelValuesable, error) {
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
// ToGetRequestInformation get all the label values for the run
// returns a *RequestInformation when successful
func (m *RunItemLabelValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunItemLabelValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunItemLabelValuesRequestBuilder when successful
func (m *RunItemLabelValuesRequestBuilder) WithUrl(rawUrl string)(*RunItemLabelValuesRequestBuilder) {
    return NewRunItemLabelValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
