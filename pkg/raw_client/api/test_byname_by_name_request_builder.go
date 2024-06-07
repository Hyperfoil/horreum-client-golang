package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TestBynameByNameRequestBuilder builds and executes requests for operations under \api\test\byName
type TestBynameByNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByName gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.test.byName.item collection
// returns a *TestBynameWithNameItemRequestBuilder when successful
func (m *TestBynameByNameRequestBuilder) ByName(name string)(*TestBynameWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewTestBynameWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTestBynameByNameRequestBuilderInternal instantiates a new TestBynameByNameRequestBuilder and sets the default values.
func NewTestBynameByNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestBynameByNameRequestBuilder) {
    m := &TestBynameByNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/byName", pathParameters),
    }
    return m
}
// NewTestBynameByNameRequestBuilder instantiates a new TestBynameByNameRequestBuilder and sets the default values.
func NewTestBynameByNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestBynameByNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestBynameByNameRequestBuilderInternal(urlParams, requestAdapter)
}
