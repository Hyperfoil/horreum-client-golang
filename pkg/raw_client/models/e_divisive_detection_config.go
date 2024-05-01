package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EDivisiveDetectionConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Built In
    builtIn *bool
    // The model property
    model *EDivisiveDetectionConfig_model
}
// NewEDivisiveDetectionConfig instantiates a new EDivisiveDetectionConfig and sets the default values.
func NewEDivisiveDetectionConfig()(*EDivisiveDetectionConfig) {
    m := &EDivisiveDetectionConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEDivisiveDetectionConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEDivisiveDetectionConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEDivisiveDetectionConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EDivisiveDetectionConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuiltIn gets the builtIn property value. Built In
// returns a *bool when successful
func (m *EDivisiveDetectionConfig) GetBuiltIn()(*bool) {
    return m.builtIn
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EDivisiveDetectionConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["builtIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuiltIn(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEDivisiveDetectionConfig_model)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val.(*EDivisiveDetectionConfig_model))
        }
        return nil
    }
    return res
}
// GetModel gets the model property value. The model property
// returns a *EDivisiveDetectionConfig_model when successful
func (m *EDivisiveDetectionConfig) GetModel()(*EDivisiveDetectionConfig_model) {
    return m.model
}
// Serialize serializes information the current object
func (m *EDivisiveDetectionConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("builtIn", m.GetBuiltIn())
        if err != nil {
            return err
        }
    }
    if m.GetModel() != nil {
        cast := (*m.GetModel()).String()
        err := writer.WriteStringValue("model", &cast)
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
func (m *EDivisiveDetectionConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuiltIn sets the builtIn property value. Built In
func (m *EDivisiveDetectionConfig) SetBuiltIn(value *bool)() {
    m.builtIn = value
}
// SetModel sets the model property value. The model property
func (m *EDivisiveDetectionConfig) SetModel(value *EDivisiveDetectionConfig_model)() {
    m.model = value
}
type EDivisiveDetectionConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuiltIn()(*bool)
    GetModel()(*EDivisiveDetectionConfig_model)
    SetBuiltIn(value *bool)()
    SetModel(value *EDivisiveDetectionConfig_model)()
}
