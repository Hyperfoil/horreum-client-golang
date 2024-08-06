package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TestByNameRequestBuilder builds and executes requests for operations under \api\test\byName
type TestByNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByName gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.test.byName.item collection
// returns a *TestByNameWithNameItemRequestBuilder when successful
func (m *TestByNameRequestBuilder) ByName(name string)(*TestByNameWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewTestByNameWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTestByNameRequestBuilderInternal instantiates a new TestByNameRequestBuilder and sets the default values.
func NewTestByNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestByNameRequestBuilder) {
    m := &TestByNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/byName", pathParameters),
    }
    return m
}
// NewTestByNameRequestBuilder instantiates a new TestByNameRequestBuilder and sets the default values.
func NewTestByNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestByNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestByNameRequestBuilderInternal(urlParams, requestAdapter)
}
