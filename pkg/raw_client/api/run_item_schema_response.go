package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use RunItemSchemaPostResponseable instead.
type RunItemSchemaResponse struct {
    RunItemSchemaPostResponse
}
// NewRunItemSchemaResponse instantiates a new RunItemSchemaResponse and sets the default values.
func NewRunItemSchemaResponse()(*RunItemSchemaResponse) {
    m := &RunItemSchemaResponse{
        RunItemSchemaPostResponse: *NewRunItemSchemaPostResponse(),
    }
    return m
}
// CreateRunItemSchemaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunItemSchemaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunItemSchemaResponse(), nil
}
// Deprecated: This class is obsolete. Use RunItemSchemaPostResponseable instead.
type RunItemSchemaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RunItemSchemaPostResponseable
}
