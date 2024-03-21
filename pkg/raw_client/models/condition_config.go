package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionConfig a configuration object for Change detection models
type ConditionConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A dictionary of UI default configuration items for dynamically building the UI components
    defaults ConditionConfig_defaultsable
    // Change detection model description
    description *string
    // Name of Change detection model
    name *string
    // UI name for change detection model
    title *string
    // A list of UI components for dynamically building the UI components
    ui []ConditionComponentable
}
// NewConditionConfig instantiates a new ConditionConfig and sets the default values.
func NewConditionConfig()(*ConditionConfig) {
    m := &ConditionConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConditionConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaults gets the defaults property value. A dictionary of UI default configuration items for dynamically building the UI components
// returns a ConditionConfig_defaultsable when successful
func (m *ConditionConfig) GetDefaults()(ConditionConfig_defaultsable) {
    return m.defaults
}
// GetDescription gets the description property value. Change detection model description
// returns a *string when successful
func (m *ConditionConfig) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionConfig_defaultsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaults(val.(ConditionConfig_defaultsable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["ui"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionComponentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionComponentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConditionComponentable)
                }
            }
            m.SetUi(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of Change detection model
// returns a *string when successful
func (m *ConditionConfig) GetName()(*string) {
    return m.name
}
// GetTitle gets the title property value. UI name for change detection model
// returns a *string when successful
func (m *ConditionConfig) GetTitle()(*string) {
    return m.title
}
// GetUi gets the ui property value. A list of UI components for dynamically building the UI components
// returns a []ConditionComponentable when successful
func (m *ConditionConfig) GetUi()([]ConditionComponentable) {
    return m.ui
}
// Serialize serializes information the current object
func (m *ConditionConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("defaults", m.GetDefaults())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    if m.GetUi() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUi()))
        for i, v := range m.GetUi() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("ui", cast)
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
func (m *ConditionConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaults sets the defaults property value. A dictionary of UI default configuration items for dynamically building the UI components
func (m *ConditionConfig) SetDefaults(value ConditionConfig_defaultsable)() {
    m.defaults = value
}
// SetDescription sets the description property value. Change detection model description
func (m *ConditionConfig) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. Name of Change detection model
func (m *ConditionConfig) SetName(value *string)() {
    m.name = value
}
// SetTitle sets the title property value. UI name for change detection model
func (m *ConditionConfig) SetTitle(value *string)() {
    m.title = value
}
// SetUi sets the ui property value. A list of UI components for dynamically building the UI components
func (m *ConditionConfig) SetUi(value []ConditionComponentable)() {
    m.ui = value
}
type ConditionConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaults()(ConditionConfig_defaultsable)
    GetDescription()(*string)
    GetName()(*string)
    GetTitle()(*string)
    GetUi()([]ConditionComponentable)
    SetDefaults(value ConditionConfig_defaultsable)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetTitle(value *string)()
    SetUi(value []ConditionComponentable)()
}
