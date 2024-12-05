package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunTestRequestBuilder builds and executes requests for operations under \api\run\test
type RunTestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunTestRequestBuilderPostQueryParameters upload a new Run
type RunTestRequestBuilderPostQueryParameters struct {
    // New Access level
    // Deprecated: This property is deprecated, use AccessAsAccess instead
    Access *string `uriparametername:"access"`
    // New Access level
    AccessAsAccess *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Access `uriparametername:"access"`
    // Name of the new owner
    Owner *string `uriparametername:"owner"`
    // test name of ID
    Test *string `uriparametername:"test"`
}
// RunTestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunTestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunTestRequestBuilderPostQueryParameters
}
// NewRunTestRequestBuilderInternal instantiates a new RunTestRequestBuilder and sets the default values.
func NewRunTestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunTestRequestBuilder) {
    m := &RunTestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/test{?access*,owner*,test*}", pathParameters),
    }
    return m
}
// NewRunTestRequestBuilder instantiates a new RunTestRequestBuilder and sets the default values.
func NewRunTestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunTestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunTestRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a new Run
// returns a []int32 when successful
func (m *RunTestRequestBuilder) Post(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Runable, requestConfiguration *RunTestRequestBuilderPostRequestConfiguration)([]int32, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    val := make([]int32, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*int32))
        }
    }
    return val, nil
}
// ToPostRequestInformation upload a new Run
// returns a *RequestInformation when successful
func (m *RunTestRequestBuilder) ToPostRequestInformation(ctx context.Context, body i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Runable, requestConfiguration *RunTestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
// returns a *RunTestRequestBuilder when successful
func (m *RunTestRequestBuilder) WithUrl(rawUrl string)(*RunTestRequestBuilder) {
    return NewRunTestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
