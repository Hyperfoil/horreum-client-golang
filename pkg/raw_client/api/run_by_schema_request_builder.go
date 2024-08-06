package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunBySchemaRequestBuilder builds and executes requests for operations under \api\run\bySchema
type RunBySchemaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunBySchemaRequestBuilderGetQueryParameters retrieve a paginated list of Runs with available count for a given Schema URI
type RunBySchemaRequestBuilderGetQueryParameters struct {
    // Sort direction
    // Deprecated: This property is deprecated, use DirectionAsSortDirection instead
    Direction *string `uriparametername:"direction"`
    // Sort direction
    DirectionAsSortDirection *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SortDirection `uriparametername:"direction"`
    // limit the number of results
    Limit *int32 `uriparametername:"limit"`
    // filter by page number of a paginated list of Tests
    Page *int32 `uriparametername:"page"`
    // Field name to sort results
    Sort *string `uriparametername:"sort"`
    // Schema URI
    Uri *string `uriparametername:"uri"`
}
// RunBySchemaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunBySchemaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunBySchemaRequestBuilderGetQueryParameters
}
// NewRunBySchemaRequestBuilderInternal instantiates a new RunBySchemaRequestBuilder and sets the default values.
func NewRunBySchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunBySchemaRequestBuilder) {
    m := &RunBySchemaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/bySchema?uri={uri}{&direction*,limit*,page*,sort*}", pathParameters),
    }
    return m
}
// NewRunBySchemaRequestBuilder instantiates a new RunBySchemaRequestBuilder and sets the default values.
func NewRunBySchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunBySchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunBySchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a paginated list of Runs with available count for a given Schema URI
// returns a RunsSummaryable when successful
func (m *RunBySchemaRequestBuilder) Get(ctx context.Context, requestConfiguration *RunBySchemaRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunsSummaryable, error) {
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
// ToGetRequestInformation retrieve a paginated list of Runs with available count for a given Schema URI
// returns a *RequestInformation when successful
func (m *RunBySchemaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunBySchemaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunBySchemaRequestBuilder when successful
func (m *RunBySchemaRequestBuilder) WithUrl(rawUrl string)(*RunBySchemaRequestBuilder) {
    return NewRunBySchemaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
