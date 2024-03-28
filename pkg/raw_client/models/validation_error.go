package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ValidationError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Validation Error Details
    error ValidationError_errorable
    // Schema ID that Validation Error relates to
    schemaId *int32
}
// NewValidationError instantiates a new ValidationError and sets the default values.
func NewValidationError()(*ValidationError) {
    m := &ValidationError{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateValidationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateValidationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidationError(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ValidationError) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetError gets the error property value. Validation Error Details
// returns a ValidationError_errorable when successful
func (m *ValidationError) GetError()(ValidationError_errorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ValidationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateValidationError_errorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(ValidationError_errorable))
        }
        return nil
    }
    res["schemaId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaId(val)
        }
        return nil
    }
    return res
}
// GetSchemaId gets the schemaId property value. Schema ID that Validation Error relates to
// returns a *int32 when successful
func (m *ValidationError) GetSchemaId()(*int32) {
    return m.schemaId
}
// Serialize serializes information the current object
func (m *ValidationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("schemaId", m.GetSchemaId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidationError) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetError sets the error property value. Validation Error Details
func (m *ValidationError) SetError(value ValidationError_errorable)() {
    m.error = value
}
// SetSchemaId sets the schemaId property value. Schema ID that Validation Error relates to
func (m *ValidationError) SetSchemaId(value *int32)() {
    m.schemaId = value
}
type ValidationErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(ValidationError_errorable)
    GetSchemaId()(*int32)
    SetError(value ValidationError_errorable)()
    SetSchemaId(value *int32)()
}
