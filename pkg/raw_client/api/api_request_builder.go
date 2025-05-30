// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApiRequestBuilder builds and executes requests for operations under \api
type ApiRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Config the config property
// returns a *ConfigRequestBuilder when successful
func (m *ApiRequestBuilder) Config()(*ConfigRequestBuilder) {
    return NewConfigRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiRequestBuilderInternal instantiates a new ApiRequestBuilder and sets the default values.
func NewApiRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiRequestBuilder) {
    m := &ApiRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api", pathParameters),
    }
    return m
}
// NewApiRequestBuilder instantiates a new ApiRequestBuilder and sets the default values.
func NewApiRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApiRequestBuilderInternal(urlParams, requestAdapter)
}
// Dataset the dataset property
// returns a *DatasetRequestBuilder when successful
func (m *ApiRequestBuilder) Dataset()(*DatasetRequestBuilder) {
    return NewDatasetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Experiment the experiment property
// returns a *ExperimentRequestBuilder when successful
func (m *ApiRequestBuilder) Experiment()(*ExperimentRequestBuilder) {
    return NewExperimentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Run the run property
// returns a *RunRequestBuilder when successful
func (m *ApiRequestBuilder) Run()(*RunRequestBuilder) {
    return NewRunRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schema the schema property
// returns a *SchemaRequestBuilder when successful
func (m *ApiRequestBuilder) Schema()(*SchemaRequestBuilder) {
    return NewSchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Test the test property
// returns a *TestRequestBuilder when successful
func (m *ApiRequestBuilder) Test()(*TestRequestBuilder) {
    return NewTestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// User the user property
// returns a *UserRequestBuilder when successful
func (m *ApiRequestBuilder) User()(*UserRequestBuilder) {
    return NewUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
