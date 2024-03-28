package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunDataRequestBuilder builds and executes requests for operations under \api\run\data
type RunDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunDataRequestBuilderPostQueryParameters upload a new Run
type RunDataRequestBuilderPostQueryParameters struct {
    // New Access level
    // Deprecated: This property is deprecated, use AccessAsAccess instead
    Access *string `uriparametername:"access"`
    // New Access level
    AccessAsAccess *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Access `uriparametername:"access"`
    // Run description
    Description *string `uriparametername:"description"`
    // Name of the new owner
    Owner *string `uriparametername:"owner"`
    // Schema URI
    Schema *string `uriparametername:"schema"`
    // start timestamp of run, or json path expression
    Start *string `uriparametername:"start"`
    // stop timestamp of run, or json path expression
    Stop *string `uriparametername:"stop"`
    // test name of ID
    Test *string `uriparametername:"test"`
    // Horreum internal token. Incompatible with Keycloak
    Token *string `uriparametername:"token"`
}
// RunDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunDataRequestBuilderPostQueryParameters
}
// NewRunDataRequestBuilderInternal instantiates a new RunDataRequestBuilder and sets the default values.
func NewRunDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunDataRequestBuilder) {
    m := &RunDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/data?start={start}&stop={stop}&test={test}{&access*,description*,owner*,schema*,token*}", pathParameters),
    }
    return m
}
// NewRunDataRequestBuilder instantiates a new RunDataRequestBuilder and sets the default values.
func NewRunDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a new Run
// returns a *int32 when successful
func (m *RunDataRequestBuilder) Post(ctx context.Context, body *string, requestConfiguration *RunDataRequestBuilderPostRequestConfiguration)(*int32, error) {
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
// ToPostRequestInformation upload a new Run
// returns a *RequestInformation when successful
func (m *RunDataRequestBuilder) ToPostRequestInformation(ctx context.Context, body *string, requestConfiguration *RunDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetContentFromScalar(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunDataRequestBuilder when successful
func (m *RunDataRequestBuilder) WithUrl(rawUrl string)(*RunDataRequestBuilder) {
    return NewRunDataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
