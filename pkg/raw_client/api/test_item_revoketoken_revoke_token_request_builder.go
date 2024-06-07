package api

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TestItemRevoketokenRevokeTokenRequestBuilder builds and executes requests for operations under \api\test\{id}\revokeToken
type TestItemRevoketokenRevokeTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTokenId gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.test.item.revokeToken.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *TestItemRevoketokenWithTokenItemRequestBuilder when successful
func (m *TestItemRevoketokenRevokeTokenRequestBuilder) ByTokenId(tokenId string)(*TestItemRevoketokenWithTokenItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tokenId != "" {
        urlTplParams["tokenId"] = tokenId
    }
    return NewTestItemRevoketokenWithTokenItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByTokenIdInteger gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.test.item.revokeToken.item collection
// returns a *TestItemRevoketokenWithTokenItemRequestBuilder when successful
func (m *TestItemRevoketokenRevokeTokenRequestBuilder) ByTokenIdInteger(tokenId int32)(*TestItemRevoketokenWithTokenItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["tokenId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(tokenId), 10)
    return NewTestItemRevoketokenWithTokenItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTestItemRevoketokenRevokeTokenRequestBuilderInternal instantiates a new TestItemRevoketokenRevokeTokenRequestBuilder and sets the default values.
func NewTestItemRevoketokenRevokeTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemRevoketokenRevokeTokenRequestBuilder) {
    m := &TestItemRevoketokenRevokeTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/revokeToken", pathParameters),
    }
    return m
}
// NewTestItemRevoketokenRevokeTokenRequestBuilder instantiates a new TestItemRevoketokenRevokeTokenRequestBuilder and sets the default values.
func NewTestItemRevoketokenRevokeTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemRevoketokenRevokeTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemRevoketokenRevokeTokenRequestBuilderInternal(urlParams, requestAdapter)
}
