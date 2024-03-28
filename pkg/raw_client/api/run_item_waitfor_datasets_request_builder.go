package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemWaitforDatasetsRequestBuilder builds and executes requests for operations under \api\run\{id}\waitforDatasets
type RunItemWaitforDatasetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemWaitforDatasetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemWaitforDatasetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRunItemWaitforDatasetsRequestBuilderInternal instantiates a new RunItemWaitforDatasetsRequestBuilder and sets the default values.
func NewRunItemWaitforDatasetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemWaitforDatasetsRequestBuilder) {
    m := &RunItemWaitforDatasetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/waitforDatasets", pathParameters),
    }
    return m
}
// NewRunItemWaitforDatasetsRequestBuilder instantiates a new RunItemWaitforDatasetsRequestBuilder and sets the default values.
func NewRunItemWaitforDatasetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemWaitforDatasetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemWaitforDatasetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get blocking call, waiting for datasets to be produced
func (m *RunItemWaitforDatasetsRequestBuilder) Get(ctx context.Context, requestConfiguration *RunItemWaitforDatasetsRequestBuilderGetRequestConfiguration)(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation blocking call, waiting for datasets to be produced
// returns a *RequestInformation when successful
func (m *RunItemWaitforDatasetsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunItemWaitforDatasetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemWaitforDatasetsRequestBuilder when successful
func (m *RunItemWaitforDatasetsRequestBuilder) WithUrl(rawUrl string)(*RunItemWaitforDatasetsRequestBuilder) {
    return NewRunItemWaitforDatasetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
