package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// DatasetByschemaBySchemaRequestBuilder builds and executes requests for operations under \api\dataset\bySchema
type DatasetByschemaBySchemaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DatasetByschemaBySchemaRequestBuilderGetQueryParameters retrieve a paginated list of Datasets, with total count, by Schema
type DatasetByschemaBySchemaRequestBuilderGetQueryParameters struct {
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
// DatasetByschemaBySchemaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DatasetByschemaBySchemaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DatasetByschemaBySchemaRequestBuilderGetQueryParameters
}
// NewDatasetByschemaBySchemaRequestBuilderInternal instantiates a new DatasetByschemaBySchemaRequestBuilder and sets the default values.
func NewDatasetByschemaBySchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetByschemaBySchemaRequestBuilder) {
    m := &DatasetByschemaBySchemaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/dataset/bySchema?uri={uri}{&direction*,limit*,page*,sort*}", pathParameters),
    }
    return m
}
// NewDatasetByschemaBySchemaRequestBuilder instantiates a new DatasetByschemaBySchemaRequestBuilder and sets the default values.
func NewDatasetByschemaBySchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetByschemaBySchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasetByschemaBySchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a paginated list of Datasets, with total count, by Schema
// returns a DatasetListable when successful
func (m *DatasetByschemaBySchemaRequestBuilder) Get(ctx context.Context, requestConfiguration *DatasetByschemaBySchemaRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.DatasetListable, error) {
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
func (m *DatasetByschemaBySchemaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DatasetByschemaBySchemaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DatasetByschemaBySchemaRequestBuilder when successful
func (m *DatasetByschemaBySchemaRequestBuilder) WithUrl(rawUrl string)(*DatasetByschemaBySchemaRequestBuilder) {
    return NewDatasetByschemaBySchemaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
