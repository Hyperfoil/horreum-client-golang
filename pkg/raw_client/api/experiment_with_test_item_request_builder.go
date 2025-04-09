// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExperimentWithTestItemRequestBuilder builds and executes requests for operations under \api\experiment\{testId}
type ExperimentWithTestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewExperimentWithTestItemRequestBuilderInternal instantiates a new ExperimentWithTestItemRequestBuilder and sets the default values.
func NewExperimentWithTestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentWithTestItemRequestBuilder) {
    m := &ExperimentWithTestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/experiment/{testId}", pathParameters),
    }
    return m
}
// NewExperimentWithTestItemRequestBuilder instantiates a new ExperimentWithTestItemRequestBuilder and sets the default values.
func NewExperimentWithTestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExperimentWithTestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExperimentWithTestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Profiles the profiles property
// returns a *ExperimentItemProfilesRequestBuilder when successful
func (m *ExperimentWithTestItemRequestBuilder) Profiles()(*ExperimentItemProfilesRequestBuilder) {
    return NewExperimentItemProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
