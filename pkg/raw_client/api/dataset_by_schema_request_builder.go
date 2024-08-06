package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// DatasetBySchemaRequestBuilder builds and executes requests for operations under \api\dataset\bySchema
type DatasetBySchemaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DatasetBySchemaRequestBuilderGetQueryParameters retrieve a paginated list of Datasets, with total count, by Schema
type DatasetBySchemaRequestBuilderGetQueryParameters struct {
    // Sort direction
    // Deprecated: This property is deprecated, use DirectionAsSortDirection instead
    Direction *string `uriparametername:"direction"`
    // Sort direction
    DirectionAsSortDirection *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SortDirection `uriparametername:"direction"`
    // limit the number of results
    Limit *int32 `uriparametername:"limit"`
    // filter by page number of a paginated list of Schemas
    Page *int32 `uriparametername:"page"`
    // Field name to sort results
    Sort *string `uriparametername:"sort"`
    // Schema URI
    Uri *string `uriparametername:"uri"`
}
// DatasetBySchemaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DatasetBySchemaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DatasetBySchemaRequestBuilderGetQueryParameters
}
// NewDatasetBySchemaRequestBuilderInternal instantiates a new DatasetBySchemaRequestBuilder and sets the default values.
func NewDatasetBySchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetBySchemaRequestBuilder) {
    m := &DatasetBySchemaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/dataset/bySchema?uri={uri}{&direction*,limit*,page*,sort*}", pathParameters),
    }
    return m
}
// NewDatasetBySchemaRequestBuilder instantiates a new DatasetBySchemaRequestBuilder and sets the default values.
func NewDatasetBySchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetBySchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasetBySchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a paginated list of Datasets, with total count, by Schema
// returns a DatasetListable when successful
func (m *DatasetBySchemaRequestBuilder) Get(ctx context.Context, requestConfiguration *DatasetBySchemaRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.DatasetListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateDatasetListFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.DatasetListable), nil
}
// ToGetRequestInformation retrieve a paginated list of Datasets, with total count, by Schema
// returns a *RequestInformation when successful
func (m *DatasetBySchemaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DatasetBySchemaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DatasetBySchemaRequestBuilder when successful
func (m *DatasetBySchemaRequestBuilder) WithUrl(rawUrl string)(*DatasetBySchemaRequestBuilder) {
    return NewDatasetBySchemaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
