package api

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DatasetListRequestBuilder builds and executes requests for operations under \api\dataset\list
type DatasetListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTestId gets an item from the github.com/hyperfoil/horreum/pkg/raw_client.api.dataset.list.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *DatasetListWithTestItemRequestBuilder when successful
func (m *DatasetListRequestBuilder) ByTestId(testId string)(*DatasetListWithTestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if testId != "" {
        urlTplParams["testId"] = testId
    }
    return NewDatasetListWithTestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByTestIdInteger gets an item from the github.com/hyperfoil/horreum/pkg/raw_client.api.dataset.list.item collection
// returns a *DatasetListWithTestItemRequestBuilder when successful
func (m *DatasetListRequestBuilder) ByTestIdInteger(testId int32)(*DatasetListWithTestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["testId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(testId), 10)
    return NewDatasetListWithTestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDatasetListRequestBuilderInternal instantiates a new DatasetListRequestBuilder and sets the default values.
func NewDatasetListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetListRequestBuilder) {
    m := &DatasetListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/dataset/list", pathParameters),
    }
    return m
}
// NewDatasetListRequestBuilder instantiates a new DatasetListRequestBuilder and sets the default values.
func NewDatasetListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasetListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasetListRequestBuilderInternal(urlParams, requestAdapter)
}
