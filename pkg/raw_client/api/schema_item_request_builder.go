package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// SchemaItemRequestBuilder builds and executes requests for operations under \api\schema\{-id}
type SchemaItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SchemaItemRequestBuilderGetQueryParameters retrieve Schema by ID
type SchemaItemRequestBuilderGetQueryParameters struct {
    // API token for authorization
    Token *string `uriparametername:"token"`
}
// SchemaItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchemaItemRequestBuilderGetQueryParameters
}
// NewSchemaItemRequestBuilderInternal instantiates a new SchemaItemRequestBuilder and sets the default values.
func NewSchemaItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemRequestBuilder) {
    m := &SchemaItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/{%2Did}{?token*}", pathParameters),
    }
    return m
}
// NewSchemaItemRequestBuilder instantiates a new SchemaItemRequestBuilder and sets the default values.
func NewSchemaItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a Schema by id
func (m *SchemaItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchemaItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DropToken the dropToken property
// returns a *SchemaItemDropTokenRequestBuilder when successful
func (m *SchemaItemRequestBuilder) DropToken()(*SchemaItemDropTokenRequestBuilder) {
    return NewSchemaItemDropTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Export the export property
// returns a *SchemaItemExportRequestBuilder when successful
func (m *SchemaItemRequestBuilder) Export()(*SchemaItemExportRequestBuilder) {
    return NewSchemaItemExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve Schema by ID
// returns a Schemaable when successful
func (m *SchemaItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaItemRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Schemaable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateSchemaFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Schemaable), nil
}
// Labels the labels property
// returns a *SchemaItemLabelsRequestBuilder when successful
func (m *SchemaItemRequestBuilder) Labels()(*SchemaItemLabelsRequestBuilder) {
    return NewSchemaItemLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetToken the resetToken property
// returns a *SchemaItemResetTokenRequestBuilder when successful
func (m *SchemaItemRequestBuilder) ResetToken()(*SchemaItemResetTokenRequestBuilder) {
    return NewSchemaItemResetTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a Schema by id
// returns a *RequestInformation when successful
func (m *SchemaItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SchemaItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve Schema by ID
// returns a *RequestInformation when successful
func (m *SchemaItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// Transformers the transformers property
// returns a *SchemaItemTransformersRequestBuilder when successful
func (m *SchemaItemRequestBuilder) Transformers()(*SchemaItemTransformersRequestBuilder) {
    return NewSchemaItemTransformersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateAccess the updateAccess property
// returns a *SchemaItemUpdateAccessRequestBuilder when successful
func (m *SchemaItemRequestBuilder) UpdateAccess()(*SchemaItemUpdateAccessRequestBuilder) {
    return NewSchemaItemUpdateAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaItemRequestBuilder when successful
func (m *SchemaItemRequestBuilder) WithUrl(rawUrl string)(*SchemaItemRequestBuilder) {
    return NewSchemaItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
