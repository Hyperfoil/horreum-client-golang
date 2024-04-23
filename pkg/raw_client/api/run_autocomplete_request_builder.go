package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RunAutocompleteRequestBuilder builds and executes requests for operations under \api\run\autocomplete
type RunAutocompleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type RunAutocompleteRequestBuilderGetQueryParameters struct {
    // JSONPath to be autocompleted
    Query *string `uriparametername:"query"`
}
// RunAutocompleteRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunAutocompleteRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunAutocompleteRequestBuilderGetQueryParameters
}
// NewRunAutocompleteRequestBuilderInternal instantiates a new RunAutocompleteRequestBuilder and sets the default values.
func NewRunAutocompleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunAutocompleteRequestBuilder) {
    m := &RunAutocompleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/autocomplete?query={query}", pathParameters),
    }
    return m
}
// NewRunAutocompleteRequestBuilder instantiates a new RunAutocompleteRequestBuilder and sets the default values.
func NewRunAutocompleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunAutocompleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunAutocompleteRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []string when successful
func (m *RunAutocompleteRequestBuilder) Get(ctx context.Context, requestConfiguration *RunAutocompleteRequestBuilderGetRequestConfiguration)([]string, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", nil)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*string))
        }
    }
    return val, nil
}
// returns a *RequestInformation when successful
func (m *RunAutocompleteRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunAutocompleteRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RunAutocompleteRequestBuilder when successful
func (m *RunAutocompleteRequestBuilder) WithUrl(rawUrl string)(*RunAutocompleteRequestBuilder) {
    return NewRunAutocompleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
