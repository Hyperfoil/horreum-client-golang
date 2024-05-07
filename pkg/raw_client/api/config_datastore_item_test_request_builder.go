package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// ConfigDatastoreItemTestRequestBuilder builds and executes requests for operations under \api\config\datastore\{-id}\test
type ConfigDatastoreItemTestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigDatastoreItemTestRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigDatastoreItemTestRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigDatastoreItemTestRequestBuilderInternal instantiates a new ConfigDatastoreItemTestRequestBuilder and sets the default values.
func NewConfigDatastoreItemTestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigDatastoreItemTestRequestBuilder) {
    m := &ConfigDatastoreItemTestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/config/datastore/{%2Did}/test", pathParameters),
    }
    return m
}
// NewConfigDatastoreItemTestRequestBuilder instantiates a new ConfigDatastoreItemTestRequestBuilder and sets the default values.
func NewConfigDatastoreItemTestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigDatastoreItemTestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigDatastoreItemTestRequestBuilderInternal(urlParams, requestAdapter)
}
// Get test a Datastore connection
// returns a DatastoreTestResponseable when successful
func (m *ConfigDatastoreItemTestRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigDatastoreItemTestRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.DatastoreTestResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateDatastoreTestResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.DatastoreTestResponseable), nil
}
// ToGetRequestInformation test a Datastore connection
// returns a *RequestInformation when successful
func (m *ConfigDatastoreItemTestRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigDatastoreItemTestRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConfigDatastoreItemTestRequestBuilder when successful
func (m *ConfigDatastoreItemTestRequestBuilder) WithUrl(rawUrl string)(*ConfigDatastoreItemTestRequestBuilder) {
    return NewConfigDatastoreItemTestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
