package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ValidationError_error validation Error Details
type ValidationError_error struct {
    ErrorDetails
}
// NewValidationError_error instantiates a new ValidationError_error and sets the default values.
func NewValidationError_error()(*ValidationError_error) {
    m := &ValidationError_error{
        ErrorDetails: *NewErrorDetails(),
    }
    return m
}
// CreateValidationError_errorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateValidationError_errorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidationError_error(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ValidationError_error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ErrorDetails.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ValidationError_error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ErrorDetails.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type ValidationError_errorable interface {
    ErrorDetailsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
