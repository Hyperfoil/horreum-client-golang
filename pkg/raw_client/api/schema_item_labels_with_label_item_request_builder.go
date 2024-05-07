package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaItemLabelsWithLabelItemRequestBuilder builds and executes requests for operations under \api\schema\{-id}\labels\{labelId}
type SchemaItemLabelsWithLabelItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaItemLabelsWithLabelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaItemLabelsWithLabelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaItemLabelsWithLabelItemRequestBuilderInternal instantiates a new SchemaItemLabelsWithLabelItemRequestBuilder and sets the default values.
func NewSchemaItemLabelsWithLabelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemLabelsWithLabelItemRequestBuilder) {
    m := &SchemaItemLabelsWithLabelItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/{%2Did}/labels/{labelId}", pathParameters),
    }
    return m
}
// NewSchemaItemLabelsWithLabelItemRequestBuilder instantiates a new SchemaItemLabelsWithLabelItemRequestBuilder and sets the default values.
func NewSchemaItemLabelsWithLabelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemLabelsWithLabelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaItemLabelsWithLabelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete existing Label from a Schema
func (m *SchemaItemLabelsWithLabelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchemaItemLabelsWithLabelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation delete existing Label from a Schema
// returns a *RequestInformation when successful
func (m *SchemaItemLabelsWithLabelItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SchemaItemLabelsWithLabelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaItemLabelsWithLabelItemRequestBuilder when successful
func (m *SchemaItemLabelsWithLabelItemRequestBuilder) WithUrl(rawUrl string)(*SchemaItemLabelsWithLabelItemRequestBuilder) {
    return NewSchemaItemLabelsWithLabelItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
