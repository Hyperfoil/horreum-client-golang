package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FixedThresholdDetectionConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Built In
    builtIn *bool
    // Upper bound for acceptable datapoint values
    max FixThresholdConfigable
    // Lower bound for acceptable datapoint values
    min FixThresholdConfigable
    // The model property
    model *FixedThresholdDetectionConfig_model
}
// NewFixedThresholdDetectionConfig instantiates a new FixedThresholdDetectionConfig and sets the default values.
func NewFixedThresholdDetectionConfig()(*FixedThresholdDetectionConfig) {
    m := &FixedThresholdDetectionConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFixedThresholdDetectionConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFixedThresholdDetectionConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFixedThresholdDetectionConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FixedThresholdDetectionConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuiltIn gets the builtIn property value. Built In
// returns a *bool when successful
func (m *FixedThresholdDetectionConfig) GetBuiltIn()(*bool) {
    return m.builtIn
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FixedThresholdDetectionConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["max"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFixThresholdConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMax(val.(FixThresholdConfigable))
        }
        return nil
    }
    res["min"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFixThresholdConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMin(val.(FixThresholdConfigable))
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFixedThresholdDetectionConfig_model)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val.(*FixedThresholdDetectionConfig_model))
        }
        return nil
    }
    return res
}
// GetMax gets the max property value. Upper bound for acceptable datapoint values
// returns a FixThresholdConfigable when successful
func (m *FixedThresholdDetectionConfig) GetMax()(FixThresholdConfigable) {
    return m.max
}
// GetMin gets the min property value. Lower bound for acceptable datapoint values
// returns a FixThresholdConfigable when successful
func (m *FixedThresholdDetectionConfig) GetMin()(FixThresholdConfigable) {
    return m.min
}
// GetModel gets the model property value. The model property
// returns a *FixedThresholdDetectionConfig_model when successful
func (m *FixedThresholdDetectionConfig) GetModel()(*FixedThresholdDetectionConfig_model) {
    return m.model
}
// Serialize serializes information the current object
func (m *FixedThresholdDetectionConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("builtIn", m.GetBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("max", m.GetMax())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("min", m.GetMin())
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
func (m *FixedThresholdDetectionConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuiltIn sets the builtIn property value. Built In
func (m *FixedThresholdDetectionConfig) SetBuiltIn(value *bool)() {
    m.builtIn = value
}
// SetMax sets the max property value. Upper bound for acceptable datapoint values
func (m *FixedThresholdDetectionConfig) SetMax(value FixThresholdConfigable)() {
    m.max = value
}
// SetMin sets the min property value. Lower bound for acceptable datapoint values
func (m *FixedThresholdDetectionConfig) SetMin(value FixThresholdConfigable)() {
    m.min = value
}
// SetModel sets the model property value. The model property
func (m *FixedThresholdDetectionConfig) SetModel(value *FixedThresholdDetectionConfig_model)() {
    m.model = value
}
type FixedThresholdDetectionConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuiltIn()(*bool)
    GetMax()(FixThresholdConfigable)
    GetMin()(FixThresholdConfigable)
    GetModel()(*FixedThresholdDetectionConfig_model)
    SetBuiltIn(value *bool)()
    SetMax(value FixThresholdConfigable)()
    SetMin(value FixThresholdConfigable)()
    SetModel(value *FixedThresholdDetectionConfig_model)()
}
