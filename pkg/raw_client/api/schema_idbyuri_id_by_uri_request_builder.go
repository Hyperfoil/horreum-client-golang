package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaIdbyuriIdByUriRequestBuilder builds and executes requests for operations under \api\schema\idByUri
type SchemaIdbyuriIdByUriRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUri gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.schema.idByUri.item collection
// returns a *SchemaIdbyuriWithUriItemRequestBuilder when successful
func (m *SchemaIdbyuriIdByUriRequestBuilder) ByUri(uri string)(*SchemaIdbyuriWithUriItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if uri != "" {
        urlTplParams["uri"] = uri
    }
    return NewSchemaIdbyuriWithUriItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSchemaIdbyuriIdByUriRequestBuilderInternal instantiates a new SchemaIdbyuriIdByUriRequestBuilder and sets the default values.
func NewSchemaIdbyuriIdByUriRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdbyuriIdByUriRequestBuilder) {
    m := &SchemaIdbyuriIdByUriRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/idByUri", pathParameters),
    }
    return m
}
// NewSchemaIdbyuriIdByUriRequestBuilder instantiates a new SchemaIdbyuriIdByUriRequestBuilder and sets the default values.
func NewSchemaIdbyuriIdByUriRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaIdbyuriIdByUriRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaIdbyuriIdByUriRequestBuilderInternal(urlParams, requestAdapter)
}
