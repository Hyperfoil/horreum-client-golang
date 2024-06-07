package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunByschemaBySchemaRequestBuilder builds and executes requests for operations under \api\run\bySchema
type RunByschemaBySchemaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunByschemaBySchemaRequestBuilderGetQueryParameters retrieve a paginated list of Runs with available count for a given Schema URI
type RunByschemaBySchemaRequestBuilderGetQueryParameters struct {
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
// RunByschemaBySchemaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunByschemaBySchemaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunByschemaBySchemaRequestBuilderGetQueryParameters
}
// NewRunByschemaBySchemaRequestBuilderInternal instantiates a new RunByschemaBySchemaRequestBuilder and sets the default values.
func NewRunByschemaBySchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunByschemaBySchemaRequestBuilder) {
    m := &RunByschemaBySchemaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/bySchema?uri={uri}{&direction*,limit*,page*,sort*}", pathParameters),
    }
    return m
}
// NewRunByschemaBySchemaRequestBuilder instantiates a new RunByschemaBySchemaRequestBuilder and sets the default values.
func NewRunByschemaBySchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunByschemaBySchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunByschemaBySchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a paginated list of Runs with available count for a given Schema URI
// returns a RunsSummaryable when successful
func (m *RunByschemaBySchemaRequestBuilder) Get(ctx context.Context, requestConfiguration *RunByschemaBySchemaRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunsSummaryable, error) {
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
func (m *RunByschemaBySchemaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunByschemaBySchemaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunByschemaBySchemaRequestBuilder when successful
func (m *RunByschemaBySchemaRequestBuilder) WithUrl(rawUrl string)(*RunByschemaBySchemaRequestBuilder) {
    return NewRunByschemaBySchemaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
