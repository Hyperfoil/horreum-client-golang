package api

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// UserApikeyRequestBuilder builds and executes requests for operations under \api\user\apikey
type UserApikeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserApikeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserApikeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserApikeyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserApikeyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.user.apikey.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *UserApikeyApikeyItemRequestBuilder when successful
func (m *UserApikeyRequestBuilder) ById(id string)(*UserApikeyApikeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewUserApikeyApikeyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdInteger gets an item from the github.com/hyperfoil/horreum-client-golang/pkg/raw_client.api.user.apikey.item collection
// returns a *UserApikeyApikeyItemRequestBuilder when successful
func (m *UserApikeyRequestBuilder) ByIdInteger(id int32)(*UserApikeyApikeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(id), 10)
    return NewUserApikeyApikeyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserApikeyRequestBuilderInternal instantiates a new UserApikeyRequestBuilder and sets the default values.
func NewUserApikeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserApikeyRequestBuilder) {
    m := &UserApikeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/apikey", pathParameters),
    }
    return m
}
// NewUserApikeyRequestBuilder instantiates a new UserApikeyRequestBuilder and sets the default values.
func NewUserApikeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserApikeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserApikeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list API keys.
// returns a []ApiKeyResponseable when successful
func (m *UserApikeyRequestBuilder) Get(ctx context.Context, requestConfiguration *UserApikeyRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ApiKeyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateApiKeyResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ApiKeyResponseable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ApiKeyResponseable)
        }
    }
    return val, nil
}
// Post create a new API key.
// returns a *string when successful
func (m *UserApikeyRequestBuilder) Post(ctx context.Context, body UserApikeyPostRequestBodyable, requestConfiguration *UserApikeyRequestBuilderPostRequestConfiguration)(*string, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*string), nil
}
// ToGetRequestInformation list API keys.
// returns a *RequestInformation when successful
func (m *UserApikeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserApikeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a new API key.
// returns a *RequestInformation when successful
func (m *UserApikeyRequestBuilder) ToPostRequestInformation(ctx context.Context, body UserApikeyPostRequestBodyable, requestConfiguration *UserApikeyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserApikeyRequestBuilder when successful
func (m *UserApikeyRequestBuilder) WithUrl(rawUrl string)(*UserApikeyRequestBuilder) {
    return NewUserApikeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
