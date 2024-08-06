package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// ExperimentModelsRequestsModelsRequestBuilder builds and executes requests for operations under \api\experiment\models
type ExperimentModelsRequestsModelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExperimentModelsRequestsModelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExperimentModelsRequestsModelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExperimentModelsRequestsModelsRequestBuilderInternal instantiates a new ExperimentModelsRequestsModelsRequestBuilder and sets the default values.
func NewExperimentModelsRequestsModelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentModelsRequestsModelsRequestBuilder) {
    m := &ExperimentModelsRequestsModelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/experiment/models", pathParameters),
    }
    return m
}
// NewExperimentModelsRequestsModelsRequestBuilder instantiates a new ExperimentModelsRequestsModelsRequestBuilder and sets the default values.
func NewExperimentModelsRequestsModelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentModelsRequestsModelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExperimentModelsRequestsModelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a list of Condition Config models
// returns a []ConditionConfigable when successful
func (m *ExperimentModelsRequestsModelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ExperimentModelsRequestsModelsRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ConditionConfigable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateConditionConfigFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ConditionConfigable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ConditionConfigable)
        }
    }
    return val, nil
}
// ToGetRequestInformation retrieve a list of Condition Config models
// returns a *RequestInformation when successful
func (m *ExperimentModelsRequestsModelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExperimentModelsRequestsModelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ExperimentModelsRequestsModelsRequestBuilder when successful
func (m *ExperimentModelsRequestsModelsRequestBuilder) WithUrl(rawUrl string)(*ExperimentModelsRequestsModelsRequestBuilder) {
    return NewExperimentModelsRequestsModelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
