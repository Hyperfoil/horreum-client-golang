package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// DatasetItemLabelvaluesLabelValuesRequestBuilder builds and executes requests for operations under \api\dataset\{dataset-id}\labelValues
type DatasetItemLabelvaluesLabelValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DatasetItemLabelvaluesLabelValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DatasetItemLabelvaluesLabelValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDatasetItemLabelvaluesLabelValuesRequestBuilderInternal instantiates a new DatasetItemLabelvaluesLabelValuesRequestBuilder and sets the default values.
func NewDatasetItemLabelvaluesLabelValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetItemLabelvaluesLabelValuesRequestBuilder) {
    m := &DatasetItemLabelvaluesLabelValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/dataset/{dataset%2Did}/labelValues", pathParameters),
    }
    return m
}
// NewDatasetItemLabelvaluesLabelValuesRequestBuilder instantiates a new DatasetItemLabelvaluesLabelValuesRequestBuilder and sets the default values.
func NewDatasetItemLabelvaluesLabelValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetItemLabelvaluesLabelValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasetItemLabelvaluesLabelValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []LabelValueable when successful
func (m *DatasetItemLabelvaluesLabelValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *DatasetItemLabelvaluesLabelValuesRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelValueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateLabelValueFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelValueable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelValueable)
        }
    }
    return val, nil
}
// returns a *RequestInformation when successful
func (m *DatasetItemLabelvaluesLabelValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DatasetItemLabelvaluesLabelValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DatasetItemLabelvaluesLabelValuesRequestBuilder when successful
func (m *DatasetItemLabelvaluesLabelValuesRequestBuilder) WithUrl(rawUrl string)(*DatasetItemLabelvaluesLabelValuesRequestBuilder) {
    return NewDatasetItemLabelvaluesLabelValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
