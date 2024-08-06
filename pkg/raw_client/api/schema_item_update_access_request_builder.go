package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// SchemaItemUpdateAccessRequestBuilder builds and executes requests for operations under \api\schema\{-id}\updateAccess
type SchemaItemUpdateAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaItemUpdateAccessRequestBuilderPostQueryParameters update the Access configuration for a Schema
type SchemaItemUpdateAccessRequestBuilderPostQueryParameters struct {
    // New Access level
    // Deprecated: This property is deprecated, use AccessAsAccess instead
    Access *string `uriparametername:"access"`
    // New Access level
    AccessAsAccess *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Access `uriparametername:"access"`
    // Name of the new owner
    Owner *string `uriparametername:"owner"`
}
// SchemaItemUpdateAccessRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaItemUpdateAccessRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchemaItemUpdateAccessRequestBuilderPostQueryParameters
}
// NewSchemaItemUpdateAccessRequestBuilderInternal instantiates a new SchemaItemUpdateAccessRequestBuilder and sets the default values.
func NewSchemaItemUpdateAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemUpdateAccessRequestBuilder) {
    m := &SchemaItemUpdateAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/{%2Did}/updateAccess?access={access}&owner={owner}", pathParameters),
    }
    return m
}
// NewSchemaItemUpdateAccessRequestBuilder instantiates a new SchemaItemUpdateAccessRequestBuilder and sets the default values.
func NewSchemaItemUpdateAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemUpdateAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaItemUpdateAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the Access configuration for a Schema
func (m *SchemaItemUpdateAccessRequestBuilder) Post(ctx context.Context, requestConfiguration *SchemaItemUpdateAccessRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation update the Access configuration for a Schema
// returns a *RequestInformation when successful
func (m *SchemaItemUpdateAccessRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SchemaItemUpdateAccessRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaItemUpdateAccessRequestBuilder when successful
func (m *SchemaItemUpdateAccessRequestBuilder) WithUrl(rawUrl string)(*SchemaItemUpdateAccessRequestBuilder) {
    return NewSchemaItemUpdateAccessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
