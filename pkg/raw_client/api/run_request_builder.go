package api

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunRequestBuilder builds and executes requests for operations under \api\run
type RunRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Autocomplete the autocomplete property
// returns a *RunAutocompleteRequestBuilder when successful
func (m *RunRequestBuilder) Autocomplete()(*RunAutocompleteRequestBuilder) {
    return NewRunAutocompleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ById gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.run.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *RunRunItemRequestBuilder when successful
func (m *RunRequestBuilder) ById(id string)(*RunRunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewRunRunItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdInteger gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.run.item collection
// returns a *RunRunItemRequestBuilder when successful
func (m *RunRequestBuilder) ByIdInteger(id int32)(*RunRunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(id), 10)
    return NewRunRunItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// BySchema the bySchema property
// returns a *RunBySchemaRequestBuilder when successful
func (m *RunRequestBuilder) BySchema()(*RunBySchemaRequestBuilder) {
    return NewRunBySchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRunRequestBuilderInternal instantiates a new RunRequestBuilder and sets the default values.
func NewRunRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunRequestBuilder) {
    m := &RunRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run", pathParameters),
    }
    return m
}
// NewRunRequestBuilder instantiates a new RunRequestBuilder and sets the default values.
func NewRunRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
// returns a *RunCountRequestBuilder when successful
func (m *RunRequestBuilder) Count()(*RunCountRequestBuilder) {
    return NewRunCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Data the data property
// returns a *RunDataRequestBuilder when successful
func (m *RunRequestBuilder) Data()(*RunDataRequestBuilder) {
    return NewRunDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// List the list property
// returns a *RunListRequestBuilder when successful
func (m *RunRequestBuilder) List()(*RunListRequestBuilder) {
    return NewRunListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecalculateAll the recalculateAll property
// returns a *RunRecalculateAllRequestBuilder when successful
func (m *RunRequestBuilder) RecalculateAll()(*RunRecalculateAllRequestBuilder) {
    return NewRunRecalculateAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Test the test property
// returns a *RunTestRequestBuilder when successful
func (m *RunRequestBuilder) Test()(*RunTestRequestBuilder) {
    return NewRunTestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
