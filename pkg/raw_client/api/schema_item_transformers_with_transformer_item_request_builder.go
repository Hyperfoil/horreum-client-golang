package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaItemTransformersWithTransformerItemRequestBuilder builds and executes requests for operations under \api\schema\{-id}\transformers\{transformerId}
type SchemaItemTransformersWithTransformerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaItemTransformersWithTransformerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaItemTransformersWithTransformerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaItemTransformersWithTransformerItemRequestBuilderInternal instantiates a new SchemaItemTransformersWithTransformerItemRequestBuilder and sets the default values.
func NewSchemaItemTransformersWithTransformerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemTransformersWithTransformerItemRequestBuilder) {
    m := &SchemaItemTransformersWithTransformerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/{%2Did}/transformers/{transformerId}", pathParameters),
    }
    return m
}
// NewSchemaItemTransformersWithTransformerItemRequestBuilder instantiates a new SchemaItemTransformersWithTransformerItemRequestBuilder and sets the default values.
func NewSchemaItemTransformersWithTransformerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemTransformersWithTransformerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaItemTransformersWithTransformerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a Transformer defined for a Schema
func (m *SchemaItemTransformersWithTransformerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchemaItemTransformersWithTransformerItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation delete a Transformer defined for a Schema
// returns a *RequestInformation when successful
func (m *SchemaItemTransformersWithTransformerItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SchemaItemTransformersWithTransformerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaItemTransformersWithTransformerItemRequestBuilder when successful
func (m *SchemaItemTransformersWithTransformerItemRequestBuilder) WithUrl(rawUrl string)(*SchemaItemTransformersWithTransformerItemRequestBuilder) {
    return NewSchemaItemTransformersWithTransformerItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
