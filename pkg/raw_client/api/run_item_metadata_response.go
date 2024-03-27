package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use RunItemMetadataGetResponseable instead.
type RunItemMetadataResponse struct {
    RunItemMetadataGetResponse
}
// NewRunItemMetadataResponse instantiates a new RunItemMetadataResponse and sets the default values.
func NewRunItemMetadataResponse()(*RunItemMetadataResponse) {
    m := &RunItemMetadataResponse{
        RunItemMetadataGetResponse: *NewRunItemMetadataGetResponse(),
    }
    return m
}
// CreateRunItemMetadataResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunItemMetadataResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunItemMetadataResponse(), nil
}
// Deprecated: This class is obsolete. Use RunItemMetadataGetResponseable instead.
type RunItemMetadataResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RunItemMetadataGetResponseable
}
