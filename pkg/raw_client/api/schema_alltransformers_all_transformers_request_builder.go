package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// SchemaAlltransformersAllTransformersRequestBuilder builds and executes requests for operations under \api\schema\allTransformers
type SchemaAlltransformersAllTransformersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemaAlltransformersAllTransformersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemaAlltransformersAllTransformersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemaAlltransformersAllTransformersRequestBuilderInternal instantiates a new SchemaAlltransformersAllTransformersRequestBuilder and sets the default values.
func NewSchemaAlltransformersAllTransformersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaAlltransformersAllTransformersRequestBuilder) {
    m := &SchemaAlltransformersAllTransformersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/schema/allTransformers", pathParameters),
    }
    return m
}
// NewSchemaAlltransformersAllTransformersRequestBuilder instantiates a new SchemaAlltransformersAllTransformersRequestBuilder and sets the default values.
func NewSchemaAlltransformersAllTransformersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaAlltransformersAllTransformersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaAlltransformersAllTransformersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve all transformers
// returns a []TransformerInfoable when successful
func (m *SchemaAlltransformersAllTransformersRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemaAlltransformersAllTransformersRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.TransformerInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateTransformerInfoFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.TransformerInfoable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.TransformerInfoable)
        }
    }
    return val, nil
}
// ToGetRequestInformation retrieve all transformers
// returns a *RequestInformation when successful
func (m *SchemaAlltransformersAllTransformersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemaAlltransformersAllTransformersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemaAlltransformersAllTransformersRequestBuilder when successful
func (m *SchemaAlltransformersAllTransformersRequestBuilder) WithUrl(rawUrl string)(*SchemaAlltransformersAllTransformersRequestBuilder) {
    return NewSchemaAlltransformersAllTransformersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
