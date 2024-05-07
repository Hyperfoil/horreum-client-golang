package api

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// SchemaRequestBuilder builds and executes requests for operations under \api\schema
type SchemaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaRequestBuilderGetQueryParameters retrieve a paginated list of Schemas with available count
type SchemaRequestBuilderGetQueryParameters struct {
    // Sort direction
    // Deprecated: This property is deprecated, use DirectionAsSortDirection instead
    Direction *string `uriparametername:"direction"`
    // Sort direction
    DirectionAsSortDirection *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SortDirection `uriparametername:"direction"`
    // limit the number of results
    Limit *int32 `uriparametername:"limit"`
    // filter by page number of a paginated list of Schemas
    Page *int32 `uriparametername:"page"`
    // __my, __all or a comma delimited  list of roles
    Roles *string `uriparametername:"roles"`
    // Field name to sort results
    Sort *string `uriparametername:"sort"`
}
// SchemaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchemaRequestBuilderGetQueryParameters
}
// SchemaRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllLabels the allLabels property
// returns a *SchemaAllLabelsRequestBuilder when successful
func (m *SchemaRequestBuilder) AllLabels()(*SchemaAllLabelsRequestBuilder) {
    return NewSchemaAllLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AllTransformers the allTransformers property
// returns a *SchemaAllTransformersRequestBuilder when successful
func (m *SchemaRequestBuilder) AllTransformers()(*SchemaAllTransformersRequestBuilder) {
    return NewSchemaAllTransformersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ById gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.schema.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *SchemaItemRequestBuilder when successful
func (m *SchemaRequestBuilder) ById(id string)(*SchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["%2Did"] = id
    }
    return NewSchemaItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdInteger gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.schema.item collection
// returns a *SchemaItemRequestBuilder when successful
func (m *SchemaRequestBuilder) ByIdInteger(id int32)(*SchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["%2Did"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(id), 10)
    return NewSchemaItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSchemaRequestBuilderInternal instantiates a new SchemaRequestBuilder and sets the default values.
func NewSchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaRequestBuilder) {
    m := &SchemaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema{?direction*,limit*,page*,roles*,sort*}", pathParameters),
    }
    return m
}
// NewSchemaRequestBuilder instantiates a new SchemaRequestBuilder and sets the default values.
func NewSchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// Descriptors the descriptors property
// returns a *SchemaDescriptorsRequestBuilder when successful
func (m *SchemaRequestBuilder) Descriptors()(*SchemaDescriptorsRequestBuilder) {
    return NewSchemaDescriptorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FindUsages the findUsages property
// returns a *SchemaFindUsagesRequestBuilder when successful
func (m *SchemaRequestBuilder) FindUsages()(*SchemaFindUsagesRequestBuilder) {
    return NewSchemaFindUsagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a paginated list of Schemas with available count
// returns a SchemaQueryResultable when successful
func (m *SchemaRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SchemaQueryResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateSchemaQueryResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.SchemaQueryResultable), nil
}
// IdByUri the idByUri property
// returns a *SchemaIdByUriRequestBuilder when successful
func (m *SchemaRequestBuilder) IdByUri()(*SchemaIdByUriRequestBuilder) {
    return NewSchemaIdByUriRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImportEscaped the import property
// returns a *SchemaImportRequestBuilder when successful
func (m *SchemaRequestBuilder) ImportEscaped()(*SchemaImportRequestBuilder) {
    return NewSchemaImportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post save a new Schema
// returns a *int32 when successful
func (m *SchemaRequestBuilder) Post(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Schemaable, requestConfiguration *SchemaRequestBuilderPostRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation retrieve a paginated list of Schemas with available count
// returns a *RequestInformation when successful
func (m *SchemaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation save a new Schema
// returns a *RequestInformation when successful
func (m *SchemaRequestBuilder) ToPostRequestInformation(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Schemaable, requestConfiguration *SchemaRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaRequestBuilder when successful
func (m *SchemaRequestBuilder) WithUrl(rawUrl string)(*SchemaRequestBuilder) {
    return NewSchemaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
