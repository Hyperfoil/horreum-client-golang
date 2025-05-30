// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FixThresholdConfig lower bound for acceptable datapoint values
type FixThresholdConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Threshold enabled/disabled
    enabled *bool
    // Is threshold inclusive of defined value?
    inclusive *bool
    // Threshold Value
    value *int32
}
// NewFixThresholdConfig instantiates a new FixThresholdConfig and sets the default values.
func NewFixThresholdConfig()(*FixThresholdConfig) {
    m := &FixThresholdConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFixThresholdConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFixThresholdConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFixThresholdConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FixThresholdConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. Threshold enabled/disabled
// returns a *bool when successful
func (m *FixThresholdConfig) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FixThresholdConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["inclusive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInclusive(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetInclusive gets the inclusive property value. Is threshold inclusive of defined value?
// returns a *bool when successful
func (m *FixThresholdConfig) GetInclusive()(*bool) {
    return m.inclusive
}
// GetValue gets the value property value. Threshold Value
// returns a *int32 when successful
func (m *FixThresholdConfig) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *FixThresholdConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inclusive", m.GetInclusive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("value", m.GetValue())
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
func (m *FixThresholdConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. Threshold enabled/disabled
func (m *FixThresholdConfig) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetInclusive sets the inclusive property value. Is threshold inclusive of defined value?
func (m *FixThresholdConfig) SetInclusive(value *bool)() {
    m.inclusive = value
}
// SetValue sets the value property value. Threshold Value
func (m *FixThresholdConfig) SetValue(value *int32)() {
    m.value = value
}
type FixThresholdConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetInclusive()(*bool)
    GetValue()(*int32)
    SetEnabled(value *bool)()
    SetInclusive(value *bool)()
    SetValue(value *int32)()
}
