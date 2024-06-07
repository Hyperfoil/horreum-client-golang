package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// DatasetItemPreviewlabelPreviewLabelRequestBuilder builds and executes requests for operations under \api\dataset\{dataset-id}\previewLabel
type DatasetItemPreviewlabelPreviewLabelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DatasetItemPreviewlabelPreviewLabelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DatasetItemPreviewlabelPreviewLabelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDatasetItemPreviewlabelPreviewLabelRequestBuilderInternal instantiates a new DatasetItemPreviewlabelPreviewLabelRequestBuilder and sets the default values.
func NewDatasetItemPreviewlabelPreviewLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetItemPreviewlabelPreviewLabelRequestBuilder) {
    m := &DatasetItemPreviewlabelPreviewLabelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/dataset/{dataset%2Did}/previewLabel", pathParameters),
    }
    return m
}
// NewDatasetItemPreviewlabelPreviewLabelRequestBuilder instantiates a new DatasetItemPreviewlabelPreviewLabelRequestBuilder and sets the default values.
func NewDatasetItemPreviewlabelPreviewLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetItemPreviewlabelPreviewLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasetItemPreviewlabelPreviewLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a LabelPreviewable when successful
func (m *DatasetItemPreviewlabelPreviewLabelRequestBuilder) Post(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Labelable, requestConfiguration *DatasetItemPreviewlabelPreviewLabelRequestBuilderPostRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelPreviewable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateLabelPreviewFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelPreviewable), nil
}
// returns a *RequestInformation when successful
func (m *DatasetItemPreviewlabelPreviewLabelRequestBuilder) ToPostRequestInformation(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Labelable, requestConfiguration *DatasetItemPreviewlabelPreviewLabelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DatasetItemPreviewlabelPreviewLabelRequestBuilder when successful
func (m *DatasetItemPreviewlabelPreviewLabelRequestBuilder) WithUrl(rawUrl string)(*DatasetItemPreviewlabelPreviewLabelRequestBuilder) {
    return NewDatasetItemPreviewlabelPreviewLabelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
