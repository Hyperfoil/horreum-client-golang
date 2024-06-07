package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// SchemaAlllabelsAllLabelsRequestBuilder builds and executes requests for operations under \api\schema\allLabels
type SchemaAlllabelsAllLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaAlllabelsAllLabelsRequestBuilderGetQueryParameters retrieve list of Labels for ny name. Allows users to retrieve all Label Definitions that have the same name
type SchemaAlllabelsAllLabelsRequestBuilderGetQueryParameters struct {
    // Label name
    Name *string `uriparametername:"name"`
}
// SchemaAlllabelsAllLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaAlllabelsAllLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchemaAlllabelsAllLabelsRequestBuilderGetQueryParameters
}
// NewSchemaAlllabelsAllLabelsRequestBuilderInternal instantiates a new SchemaAlllabelsAllLabelsRequestBuilder and sets the default values.
func NewSchemaAlllabelsAllLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaAlllabelsAllLabelsRequestBuilder) {
    m := &SchemaAlllabelsAllLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/allLabels{?name*}", pathParameters),
    }
    return m
}
// NewSchemaAlllabelsAllLabelsRequestBuilder instantiates a new SchemaAlllabelsAllLabelsRequestBuilder and sets the default values.
func NewSchemaAlllabelsAllLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaAlllabelsAllLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaAlllabelsAllLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve list of Labels for ny name. Allows users to retrieve all Label Definitions that have the same name
// returns a []LabelInfoable when successful
func (m *SchemaAlllabelsAllLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaAlllabelsAllLabelsRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateLabelInfoFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelInfoable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.LabelInfoable)
        }
    }
    return val, nil
}
// ToGetRequestInformation retrieve list of Labels for ny name. Allows users to retrieve all Label Definitions that have the same name
// returns a *RequestInformation when successful
func (m *SchemaAlllabelsAllLabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaAlllabelsAllLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SchemaAlllabelsAllLabelsRequestBuilder when successful
func (m *SchemaAlllabelsAllLabelsRequestBuilder) WithUrl(rawUrl string)(*SchemaAlllabelsAllLabelsRequestBuilder) {
    return NewSchemaAlllabelsAllLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
