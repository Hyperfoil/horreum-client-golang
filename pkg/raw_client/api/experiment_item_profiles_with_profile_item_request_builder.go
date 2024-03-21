package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExperimentItemProfilesWithProfileItemRequestBuilder builds and executes requests for operations under \api\experiment\{testId}\profiles\{profileId}
type ExperimentItemProfilesWithProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExperimentItemProfilesWithProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExperimentItemProfilesWithProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExperimentItemProfilesWithProfileItemRequestBuilderInternal instantiates a new ExperimentItemProfilesWithProfileItemRequestBuilder and sets the default values.
func NewExperimentItemProfilesWithProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentItemProfilesWithProfileItemRequestBuilder) {
    m := &ExperimentItemProfilesWithProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/experiment/{testId}/profiles/{profileId}", pathParameters),
    }
    return m
}
// NewExperimentItemProfilesWithProfileItemRequestBuilder instantiates a new ExperimentItemProfilesWithProfileItemRequestBuilder and sets the default values.
func NewExperimentItemProfilesWithProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentItemProfilesWithProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExperimentItemProfilesWithProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an Experiment Profiles for a Test
func (m *ExperimentItemProfilesWithProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExperimentItemProfilesWithProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation delete an Experiment Profiles for a Test
// returns a *RequestInformation when successful
func (m *ExperimentItemProfilesWithProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExperimentItemProfilesWithProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ExperimentItemProfilesWithProfileItemRequestBuilder when successful
func (m *ExperimentItemProfilesWithProfileItemRequestBuilder) WithUrl(rawUrl string)(*ExperimentItemProfilesWithProfileItemRequestBuilder) {
    return NewExperimentItemProfilesWithProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
