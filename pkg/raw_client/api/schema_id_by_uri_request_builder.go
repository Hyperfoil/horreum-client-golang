package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaIdByUriRequestBuilder builds and executes requests for operations under \api\schema\idByUri
type SchemaIdByUriRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUri gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.schema.idByUri.item collection
// returns a *SchemaIdByUriWithUriItemRequestBuilder when successful
func (m *SchemaIdByUriRequestBuilder) ByUri(uri string)(*SchemaIdByUriWithUriItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if uri != "" {
        urlTplParams["uri"] = uri
    }
    return NewSchemaIdByUriWithUriItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSchemaIdByUriRequestBuilderInternal instantiates a new SchemaIdByUriRequestBuilder and sets the default values.
func NewSchemaIdByUriRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdByUriRequestBuilder) {
    m := &SchemaIdByUriRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/idByUri", pathParameters),
    }
    return m
}
// NewSchemaIdByUriRequestBuilder instantiates a new SchemaIdByUriRequestBuilder and sets the default values.
func NewSchemaIdByUriRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdByUriRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaIdByUriRequestBuilderInternal(urlParams, requestAdapter)
}
