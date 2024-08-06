package api

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExperimentRequestBuilder builds and executes requests for operations under \api\experiment
type ExperimentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTestId gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.experiment.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ExperimentWithTestItemRequestBuilder when successful
func (m *ExperimentRequestBuilder) ByTestId(testId string)(*ExperimentWithTestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if testId != "" {
        urlTplParams["testId"] = testId
    }
    return NewExperimentWithTestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByTestIdInteger gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.experiment.item collection
// returns a *ExperimentWithTestItemRequestBuilder when successful
func (m *ExperimentRequestBuilder) ByTestIdInteger(testId int32)(*ExperimentWithTestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["testId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(testId), 10)
    return NewExperimentWithTestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewExperimentRequestBuilderInternal instantiates a new ExperimentRequestBuilder and sets the default values.
func NewExperimentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentRequestBuilder) {
    m := &ExperimentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/experiment", pathParameters),
    }
    return m
}
// NewExperimentRequestBuilder instantiates a new ExperimentRequestBuilder and sets the default values.
func NewExperimentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExperimentRequestBuilderInternal(urlParams, requestAdapter)
}
// Models the models property
// returns a *ExperimentModelsRequestsModelsRequestBuilder when successful
func (m *ExperimentRequestBuilder) Models()(*ExperimentModelsRequestsModelsRequestBuilder) {
    return NewExperimentModelsRequestsModelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Run the run property
// returns a *ExperimentRunRequestBuilder when successful
func (m *ExperimentRequestBuilder) Run()(*ExperimentRunRequestBuilder) {
    return NewExperimentRunRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
