package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// SchemaFindUsagesRequestBuilder builds and executes requests for operations under \api\schema\findUsages
type SchemaFindUsagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaFindUsagesRequestBuilderGetQueryParameters find all usages of a Schema by label name
type SchemaFindUsagesRequestBuilderGetQueryParameters struct {
    // Name of label to search for
    Label *string `uriparametername:"label"`
}
// SchemaFindUsagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaFindUsagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchemaFindUsagesRequestBuilderGetQueryParameters
}
// NewSchemaFindUsagesRequestBuilderInternal instantiates a new SchemaFindUsagesRequestBuilder and sets the default values.
func NewSchemaFindUsagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaFindUsagesRequestBuilder) {
    m := &SchemaFindUsagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/findUsages?label={label}", pathParameters),
    }
    return m
}
// NewSchemaFindUsagesRequestBuilder instantiates a new SchemaFindUsagesRequestBuilder and sets the default values.
func NewSchemaFindUsagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaFindUsagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaFindUsagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get find all usages of a Schema by label name
// returns a []LabelLocationable when successful
func (m *SchemaFindUsagesRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaFindUsagesRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelLocationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateLabelLocationFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelLocationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelLocationable)
        }
    }
    return val, nil
}
// ToGetRequestInformation find all usages of a Schema by label name
// returns a *RequestInformation when successful
func (m *SchemaFindUsagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaFindUsagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaFindUsagesRequestBuilder when successful
func (m *SchemaFindUsagesRequestBuilder) WithUrl(rawUrl string)(*SchemaFindUsagesRequestBuilder) {
    return NewSchemaFindUsagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
