package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// DatasetDatasetItemRequestBuilder builds and executes requests for operations under \api\dataset\{dataset-id}
type DatasetDatasetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DatasetDatasetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DatasetDatasetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDatasetDatasetItemRequestBuilderInternal instantiates a new DatasetDatasetItemRequestBuilder and sets the default values.
func NewDatasetDatasetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetDatasetItemRequestBuilder) {
    m := &DatasetDatasetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/dataset/{dataset%2Did}", pathParameters),
    }
    return m
}
// NewDatasetDatasetItemRequestBuilder instantiates a new DatasetDatasetItemRequestBuilder and sets the default values.
func NewDatasetDatasetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetDatasetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasetDatasetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve Dataset by ID
// returns a Datasetable when successful
func (m *DatasetDatasetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DatasetDatasetItemRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Datasetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateDatasetFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Datasetable), nil
}
// LabelValues the labelValues property
// returns a *DatasetItemLabelValuesRequestBuilder when successful
func (m *DatasetDatasetItemRequestBuilder) LabelValues()(*DatasetItemLabelValuesRequestBuilder) {
    return NewDatasetItemLabelValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PreviewLabel the previewLabel property
// returns a *DatasetItemPreviewLabelRequestBuilder when successful
func (m *DatasetDatasetItemRequestBuilder) PreviewLabel()(*DatasetItemPreviewLabelRequestBuilder) {
    return NewDatasetItemPreviewLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Summary the summary property
// returns a *DatasetItemSummaryRequestBuilder when successful
func (m *DatasetDatasetItemRequestBuilder) Summary()(*DatasetItemSummaryRequestBuilder) {
    return NewDatasetItemSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieve Dataset by ID
// returns a *RequestInformation when successful
func (m *DatasetDatasetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DatasetDatasetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DatasetDatasetItemRequestBuilder when successful
func (m *DatasetDatasetItemRequestBuilder) WithUrl(rawUrl string)(*DatasetDatasetItemRequestBuilder) {
    return NewDatasetDatasetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
