package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// TestItemFingerprintRequestBuilder builds and executes requests for operations under \api\test\{id}\fingerprint
type TestItemFingerprintRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestItemFingerprintRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestItemFingerprintRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTestItemFingerprintRequestBuilderInternal instantiates a new TestItemFingerprintRequestBuilder and sets the default values.
func NewTestItemFingerprintRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemFingerprintRequestBuilder) {
    m := &TestItemFingerprintRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}/fingerprint", pathParameters),
    }
    return m
}
// NewTestItemFingerprintRequestBuilder instantiates a new TestItemFingerprintRequestBuilder and sets the default values.
func NewTestItemFingerprintRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestItemFingerprintRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestItemFingerprintRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all Fingerprints for a Test
// returns a []Fingerprintsable when successful
func (m *TestItemFingerprintRequestBuilder) Get(ctx context.Context, requestConfiguration *TestItemFingerprintRequestBuilderGetRequestConfiguration)([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Fingerprintsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateFingerprintsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Fingerprintsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Fingerprintsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list all Fingerprints for a Test
// returns a *RequestInformation when successful
func (m *TestItemFingerprintRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestItemFingerprintRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestItemFingerprintRequestBuilder when successful
func (m *TestItemFingerprintRequestBuilder) WithUrl(rawUrl string)(*TestItemFingerprintRequestBuilder) {
    return NewTestItemFingerprintRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
