package models

import (
    ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6 "strings"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ChangeDetection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The config property
    config ChangeDetection_ChangeDetection_configable
    // The id property
    id *int32
    // The model property
    model *string
}
// ChangeDetection_ChangeDetection_config composed type wrapper for classes EDivisiveDetectionConfigable, FixedThresholdDetectionConfigable, RelativeDifferenceDetectionConfigable
type ChangeDetection_ChangeDetection_config struct {
    // Composed type representation for type EDivisiveDetectionConfigable
    eDivisiveDetectionConfig EDivisiveDetectionConfigable
    // Composed type representation for type FixedThresholdDetectionConfigable
    fixedThresholdDetectionConfig FixedThresholdDetectionConfigable
    // Composed type representation for type RelativeDifferenceDetectionConfigable
    relativeDifferenceDetectionConfig RelativeDifferenceDetectionConfigable
}
// NewChangeDetection_ChangeDetection_config instantiates a new ChangeDetection_ChangeDetection_config and sets the default values.
func NewChangeDetection_ChangeDetection_config()(*ChangeDetection_ChangeDetection_config) {
    m := &ChangeDetection_ChangeDetection_config{
    }
    return m
}
// CreateChangeDetection_ChangeDetection_configFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChangeDetection_ChangeDetection_configFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewChangeDetection_ChangeDetection_config()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("model")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "eDivisive") {
                    result.SetEDivisiveDetectionConfig(NewEDivisiveDetectionConfig())
                } else if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "fixedThreshold") {
                    result.SetFixedThresholdDetectionConfig(NewFixedThresholdDetectionConfig())
                } else if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "relativeDifference") {
                    result.SetRelativeDifferenceDetectionConfig(NewRelativeDifferenceDetectionConfig())
                }
            }
        }
    }
    return result, nil
}
// GetEDivisiveDetectionConfig gets the EDivisiveDetectionConfig property value. Composed type representation for type EDivisiveDetectionConfigable
// returns a EDivisiveDetectionConfigable when successful
func (m *ChangeDetection_ChangeDetection_config) GetEDivisiveDetectionConfig()(EDivisiveDetectionConfigable) {
    return m.eDivisiveDetectionConfig
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChangeDetection_ChangeDetection_config) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFixedThresholdDetectionConfig gets the FixedThresholdDetectionConfig property value. Composed type representation for type FixedThresholdDetectionConfigable
// returns a FixedThresholdDetectionConfigable when successful
func (m *ChangeDetection_ChangeDetection_config) GetFixedThresholdDetectionConfig()(FixedThresholdDetectionConfigable) {
    return m.fixedThresholdDetectionConfig
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ChangeDetection_ChangeDetection_config) GetIsComposedType()(bool) {
    return true
}
// GetRelativeDifferenceDetectionConfig gets the RelativeDifferenceDetectionConfig property value. Composed type representation for type RelativeDifferenceDetectionConfigable
// returns a RelativeDifferenceDetectionConfigable when successful
func (m *ChangeDetection_ChangeDetection_config) GetRelativeDifferenceDetectionConfig()(RelativeDifferenceDetectionConfigable) {
    return m.relativeDifferenceDetectionConfig
}
// Serialize serializes information the current object
func (m *ChangeDetection_ChangeDetection_config) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEDivisiveDetectionConfig() != nil {
        err := writer.WriteObjectValue("", m.GetEDivisiveDetectionConfig())
        if err != nil {
            return err
        }
    } else if m.GetFixedThresholdDetectionConfig() != nil {
        err := writer.WriteObjectValue("", m.GetFixedThresholdDetectionConfig())
        if err != nil {
            return err
        }
    } else if m.GetRelativeDifferenceDetectionConfig() != nil {
        err := writer.WriteObjectValue("", m.GetRelativeDifferenceDetectionConfig())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEDivisiveDetectionConfig sets the EDivisiveDetectionConfig property value. Composed type representation for type EDivisiveDetectionConfigable
func (m *ChangeDetection_ChangeDetection_config) SetEDivisiveDetectionConfig(value EDivisiveDetectionConfigable)() {
    m.eDivisiveDetectionConfig = value
}
// SetFixedThresholdDetectionConfig sets the FixedThresholdDetectionConfig property value. Composed type representation for type FixedThresholdDetectionConfigable
func (m *ChangeDetection_ChangeDetection_config) SetFixedThresholdDetectionConfig(value FixedThresholdDetectionConfigable)() {
    m.fixedThresholdDetectionConfig = value
}
// SetRelativeDifferenceDetectionConfig sets the RelativeDifferenceDetectionConfig property value. Composed type representation for type RelativeDifferenceDetectionConfigable
func (m *ChangeDetection_ChangeDetection_config) SetRelativeDifferenceDetectionConfig(value RelativeDifferenceDetectionConfigable)() {
    m.relativeDifferenceDetectionConfig = value
}
type ChangeDetection_ChangeDetection_configable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEDivisiveDetectionConfig()(EDivisiveDetectionConfigable)
    GetFixedThresholdDetectionConfig()(FixedThresholdDetectionConfigable)
    GetRelativeDifferenceDetectionConfig()(RelativeDifferenceDetectionConfigable)
    SetEDivisiveDetectionConfig(value EDivisiveDetectionConfigable)()
    SetFixedThresholdDetectionConfig(value FixedThresholdDetectionConfigable)()
    SetRelativeDifferenceDetectionConfig(value RelativeDifferenceDetectionConfigable)()
}
// NewChangeDetection instantiates a new ChangeDetection and sets the default values.
func NewChangeDetection()(*ChangeDetection) {
    m := &ChangeDetection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChangeDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChangeDetectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChangeDetection(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ChangeDetection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. The config property
// returns a ChangeDetection_ChangeDetection_configable when successful
func (m *ChangeDetection) GetConfig()(ChangeDetection_ChangeDetection_configable) {
    return m.config
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChangeDetection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChangeDetection_ChangeDetection_configFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(ChangeDetection_ChangeDetection_configable))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *ChangeDetection) GetId()(*int32) {
    return m.id
}
// GetModel gets the model property value. The model property
// returns a *string when successful
func (m *ChangeDetection) GetModel()(*string) {
    return m.model
}
// Serialize serializes information the current object
func (m *ChangeDetection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("config", m.GetConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
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
func (m *ChangeDetection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. The config property
func (m *ChangeDetection) SetConfig(value ChangeDetection_ChangeDetection_configable)() {
    m.config = value
}
// SetId sets the id property value. The id property
func (m *ChangeDetection) SetId(value *int32)() {
    m.id = value
}
// SetModel sets the model property value. The model property
func (m *ChangeDetection) SetModel(value *string)() {
    m.model = value
}
type ChangeDetectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfig()(ChangeDetection_ChangeDetection_configable)
    GetId()(*int32)
    GetModel()(*string)
    SetConfig(value ChangeDetection_ChangeDetection_configable)()
    SetId(value *int32)()
    SetModel(value *string)()
}
