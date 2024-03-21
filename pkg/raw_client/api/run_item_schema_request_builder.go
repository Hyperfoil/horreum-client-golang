package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunItemSchemaRequestBuilder builds and executes requests for operations under \api\run\{id}\schema
type RunItemSchemaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemSchemaRequestBuilderPostQueryParameters update Run schema for part of JSON data
type RunItemSchemaRequestBuilderPostQueryParameters struct {
    // JSON path expression to update schema
    Path *string `uriparametername:"path"`
}
// RunItemSchemaRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemSchemaRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunItemSchemaRequestBuilderPostQueryParameters
}
// NewRunItemSchemaRequestBuilderInternal instantiates a new RunItemSchemaRequestBuilder and sets the default values.
func NewRunItemSchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemSchemaRequestBuilder) {
    m := &RunItemSchemaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/schema{?path*}", pathParameters),
    }
    return m
}
// NewRunItemSchemaRequestBuilder instantiates a new RunItemSchemaRequestBuilder and sets the default values.
func NewRunItemSchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemSchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemSchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update Run schema for part of JSON data
// Deprecated: This method is obsolete. Use PostAsSchemaPostResponse instead.
// returns a RunItemSchemaResponseable when successful
func (m *RunItemSchemaRequestBuilder) Post(ctx context.Context, body *string, requestConfiguration *RunItemSchemaRequestBuilderPostRequestConfiguration)(RunItemSchemaResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRunItemSchemaResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RunItemSchemaResponseable), nil
}
// PostAsSchemaPostResponse update Run schema for part of JSON data
// returns a RunItemSchemaPostResponseable when successful
func (m *RunItemSchemaRequestBuilder) PostAsSchemaPostResponse(ctx context.Context, body *string, requestConfiguration *RunItemSchemaRequestBuilderPostRequestConfiguration)(RunItemSchemaPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRunItemSchemaPostResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RunItemSchemaPostResponseable), nil
}
// ToPostRequestInformation update Run schema for part of JSON data
// returns a *RequestInformation when successful
func (m *RunItemSchemaRequestBuilder) ToPostRequestInformation(ctx context.Context, body *string, requestConfiguration *RunItemSchemaRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetContentFromScalar(ctx, m.BaseRequestBuilder.RequestAdapter, "text/plain", body)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemSchemaRequestBuilder when successful
func (m *RunItemSchemaRequestBuilder) WithUrl(rawUrl string)(*RunItemSchemaRequestBuilder) {
    return NewRunItemSchemaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
