package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

// RunItemUpdateaccessUpdateAccessRequestBuilder builds and executes requests for operations under \api\run\{id}\updateAccess
type RunItemUpdateaccessUpdateAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RunItemUpdateaccessUpdateAccessRequestBuilderPostQueryParameters update the Access configuration for a Run
type RunItemUpdateaccessUpdateAccessRequestBuilderPostQueryParameters struct {
    // New Access level
    // Deprecated: This property is deprecated, use AccessAsAccess instead
    Access *string `uriparametername:"access"`
    // New Access level
    AccessAsAccess *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.Access `uriparametername:"access"`
    // Name of the new owner
    Owner *string `uriparametername:"owner"`
}
// RunItemUpdateaccessUpdateAccessRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemUpdateaccessUpdateAccessRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunItemUpdateaccessUpdateAccessRequestBuilderPostQueryParameters
}
// NewRunItemUpdateaccessUpdateAccessRequestBuilderInternal instantiates a new RunItemUpdateaccessUpdateAccessRequestBuilder and sets the default values.
func NewRunItemUpdateaccessUpdateAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemUpdateaccessUpdateAccessRequestBuilder) {
    m := &RunItemUpdateaccessUpdateAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/run/{id}/updateAccess?access={access}&owner={owner}", pathParameters),
    }
    return m
}
// NewRunItemUpdateaccessUpdateAccessRequestBuilder instantiates a new RunItemUpdateaccessUpdateAccessRequestBuilder and sets the default values.
func NewRunItemUpdateaccessUpdateAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemUpdateaccessUpdateAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemUpdateaccessUpdateAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the Access configuration for a Run
func (m *RunItemUpdateaccessUpdateAccessRequestBuilder) Post(ctx context.Context, requestConfiguration *RunItemUpdateaccessUpdateAccessRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation update the Access configuration for a Run
// returns a *RequestInformation when successful
func (m *RunItemUpdateaccessUpdateAccessRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RunItemUpdateaccessUpdateAccessRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RunItemUpdateaccessUpdateAccessRequestBuilder when successful
func (m *RunItemUpdateaccessUpdateAccessRequestBuilder) WithUrl(rawUrl string)(*RunItemUpdateaccessUpdateAccessRequestBuilder) {
    return NewRunItemUpdateaccessUpdateAccessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
