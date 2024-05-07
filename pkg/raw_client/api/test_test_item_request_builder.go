package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// TestTestItemRequestBuilder builds and executes requests for operations under \api\test\{id}
type TestTestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TestTestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestTestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TestTestItemRequestBuilderGetQueryParameters retrieve a test by id
type TestTestItemRequestBuilderGetQueryParameters struct {
    Token *string `uriparametername:"token"`
}
// TestTestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TestTestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TestTestItemRequestBuilderGetQueryParameters
}
// AddToken the addToken property
// returns a *TestItemAddTokenRequestBuilder when successful
func (m *TestTestItemRequestBuilder) AddToken()(*TestItemAddTokenRequestBuilder) {
    return NewTestItemAddTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTestTestItemRequestBuilderInternal instantiates a new TestTestItemRequestBuilder and sets the default values.
func NewTestTestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestTestItemRequestBuilder) {
    m := &TestTestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/test/{id}{?token*}", pathParameters),
    }
    return m
}
// NewTestTestItemRequestBuilder instantiates a new TestTestItemRequestBuilder and sets the default values.
func NewTestTestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TestTestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTestTestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a Test by id
func (m *TestTestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TestTestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Export the export property
// returns a *TestItemExportRequestBuilder when successful
func (m *TestTestItemRequestBuilder) Export()(*TestItemExportRequestBuilder) {
    return NewTestItemExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Fingerprint the fingerprint property
// returns a *TestItemFingerprintRequestBuilder when successful
func (m *TestTestItemRequestBuilder) Fingerprint()(*TestItemFingerprintRequestBuilder) {
    return NewTestItemFingerprintRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a test by id
// returns a Testable when successful
func (m *TestTestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TestTestItemRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Testable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateTestFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Testable), nil
}
// LabelValues the labelValues property
// returns a *TestItemLabelValuesRequestBuilder when successful
func (m *TestTestItemRequestBuilder) LabelValues()(*TestItemLabelValuesRequestBuilder) {
    return NewTestItemLabelValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Move the move property
// returns a *TestItemMoveRequestBuilder when successful
func (m *TestTestItemRequestBuilder) Move()(*TestItemMoveRequestBuilder) {
    return NewTestItemMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
// returns a *TestItemNotificationsRequestBuilder when successful
func (m *TestTestItemRequestBuilder) Notifications()(*TestItemNotificationsRequestBuilder) {
    return NewTestItemNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Recalculate the recalculate property
// returns a *TestItemRecalculateRequestBuilder when successful
func (m *TestTestItemRequestBuilder) Recalculate()(*TestItemRecalculateRequestBuilder) {
    return NewTestItemRecalculateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeToken the revokeToken property
// returns a *TestItemRevokeTokenRequestBuilder when successful
func (m *TestTestItemRequestBuilder) RevokeToken()(*TestItemRevokeTokenRequestBuilder) {
    return NewTestItemRevokeTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a Test by id
// returns a *RequestInformation when successful
func (m *TestTestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TestTestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a test by id
// returns a *RequestInformation when successful
func (m *TestTestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TestTestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Tokens the tokens property
// returns a *TestItemTokensRequestBuilder when successful
func (m *TestTestItemRequestBuilder) Tokens()(*TestItemTokensRequestBuilder) {
    return NewTestItemTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Transformers the transformers property
// returns a *TestItemTransformersRequestBuilder when successful
func (m *TestTestItemRequestBuilder) Transformers()(*TestItemTransformersRequestBuilder) {
    return NewTestItemTransformersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateAccess the updateAccess property
// returns a *TestItemUpdateAccessRequestBuilder when successful
func (m *TestTestItemRequestBuilder) UpdateAccess()(*TestItemUpdateAccessRequestBuilder) {
    return NewTestItemUpdateAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TestTestItemRequestBuilder when successful
func (m *TestTestItemRequestBuilder) WithUrl(rawUrl string)(*TestTestItemRequestBuilder) {
    return NewTestTestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
