package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunRunItemRequestBuilder builds and executes requests for operations under \api\run\{id}
type RunRunItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunRunItemRequestBuilderGetQueryParameters get extended Run information by Run ID
type RunRunItemRequestBuilderGetQueryParameters struct {
    // Run API token
    Token *string `uriparametername:"token"`
}
// RunRunItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunRunItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunRunItemRequestBuilderGetQueryParameters
}
// NewRunRunItemRequestBuilderInternal instantiates a new RunRunItemRequestBuilder and sets the default values.
func NewRunRunItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunRunItemRequestBuilder) {
    m := &RunRunItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}{?token*}", pathParameters),
    }
    return m
}
// NewRunRunItemRequestBuilder instantiates a new RunRunItemRequestBuilder and sets the default values.
func NewRunRunItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunRunItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunRunItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Data the data property
// returns a *RunItemDataRequestBuilder when successful
func (m *RunRunItemRequestBuilder) Data()(*RunItemDataRequestBuilder) {
    return NewRunItemDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Description the description property
// returns a *RunItemDescriptionRequestBuilder when successful
func (m *RunRunItemRequestBuilder) Description()(*RunItemDescriptionRequestBuilder) {
    return NewRunItemDescriptionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DropToken the dropToken property
// returns a *RunItemDropTokenRequestBuilder when successful
func (m *RunRunItemRequestBuilder) DropToken()(*RunItemDropTokenRequestBuilder) {
    return NewRunItemDropTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get extended Run information by Run ID
// returns a RunExtendedable when successful
func (m *RunRunItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RunRunItemRequestBuilderGetRequestConfiguration)(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunExtendedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateRunExtendedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.RunExtendedable), nil
}
// LabelValues the labelValues property
// returns a *RunItemLabelValuesRequestBuilder when successful
func (m *RunRunItemRequestBuilder) LabelValues()(*RunItemLabelValuesRequestBuilder) {
    return NewRunItemLabelValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Metadata the metadata property
// returns a *RunItemMetadataRequestBuilder when successful
func (m *RunRunItemRequestBuilder) Metadata()(*RunItemMetadataRequestBuilder) {
    return NewRunItemMetadataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Recalculate the recalculate property
// returns a *RunItemRecalculateRequestBuilder when successful
func (m *RunRunItemRequestBuilder) Recalculate()(*RunItemRecalculateRequestBuilder) {
    return NewRunItemRecalculateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetToken the resetToken property
// returns a *RunItemResetTokenRequestBuilder when successful
func (m *RunRunItemRequestBuilder) ResetToken()(*RunItemResetTokenRequestBuilder) {
    return NewRunItemResetTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schema the schema property
// returns a *RunItemSchemaRequestBuilder when successful
func (m *RunRunItemRequestBuilder) Schema()(*RunItemSchemaRequestBuilder) {
    return NewRunItemSchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Summary the summary property
// returns a *RunItemSummaryRequestBuilder when successful
func (m *RunRunItemRequestBuilder) Summary()(*RunItemSummaryRequestBuilder) {
    return NewRunItemSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get extended Run information by Run ID
// returns a *RequestInformation when successful
func (m *RunRunItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RunRunItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Trash the trash property
// returns a *RunItemTrashRequestBuilder when successful
func (m *RunRunItemRequestBuilder) Trash()(*RunItemTrashRequestBuilder) {
    return NewRunItemTrashRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateAccess the updateAccess property
// returns a *RunItemUpdateAccessRequestBuilder when successful
func (m *RunRunItemRequestBuilder) UpdateAccess()(*RunItemUpdateAccessRequestBuilder) {
    return NewRunItemUpdateAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunRunItemRequestBuilder when successful
func (m *RunRunItemRequestBuilder) WithUrl(rawUrl string)(*RunRunItemRequestBuilder) {
    return NewRunRunItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
