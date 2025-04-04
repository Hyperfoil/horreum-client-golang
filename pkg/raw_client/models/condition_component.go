package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConditionComponent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Change detection model component description
    description *string
    // Change detection model component name
    name *string
    // Map of properties for component
    properties ConditionComponent_propertiesable
    // Change detection model component title
    title *string
    // UI Component type
    typeEscaped ConditionComponent_ConditionComponent_typeable
}
// ConditionComponent_ConditionComponent_type composed type wrapper for classes string
type ConditionComponent_ConditionComponent_type struct {
    // Composed type representation for type string
    string *string
}
// NewConditionComponent_ConditionComponent_type instantiates a new ConditionComponent_ConditionComponent_type and sets the default values.
func NewConditionComponent_ConditionComponent_type()(*ConditionComponent_ConditionComponent_type) {
    m := &ConditionComponent_ConditionComponent_type{
    }
    return m
}
// CreateConditionComponent_ConditionComponent_typeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionComponent_ConditionComponent_typeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewConditionComponent_ConditionComponent_type()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionComponent_ConditionComponent_type) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ConditionComponent_ConditionComponent_type) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ConditionComponent_ConditionComponent_type) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ConditionComponent_ConditionComponent_type) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetString sets the string property value. Composed type representation for type string
func (m *ConditionComponent_ConditionComponent_type) SetString(value *string)() {
    m.string = value
}
type ConditionComponent_ConditionComponent_typeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    SetString(value *string)()
}
// NewConditionComponent instantiates a new ConditionComponent and sets the default values.
func NewConditionComponent()(*ConditionComponent) {
    m := &ConditionComponent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionComponentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionComponentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionComponent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConditionComponent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Change detection model component description
// returns a *string when successful
func (m *ConditionComponent) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionComponent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionComponent_propertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperties(val.(ConditionComponent_propertiesable))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionComponent_ConditionComponent_typeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(ConditionComponent_ConditionComponent_typeable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Change detection model component name
// returns a *string when successful
func (m *ConditionComponent) GetName()(*string) {
    return m.name
}
// GetProperties gets the properties property value. Map of properties for component
// returns a ConditionComponent_propertiesable when successful
func (m *ConditionComponent) GetProperties()(ConditionComponent_propertiesable) {
    return m.properties
}
// GetTitle gets the title property value. Change detection model component title
// returns a *string when successful
func (m *ConditionComponent) GetTitle()(*string) {
    return m.title
}
// GetTypeEscaped gets the type property value. UI Component type
// returns a ConditionComponent_ConditionComponent_typeable when successful
func (m *ConditionComponent) GetTypeEscaped()(ConditionComponent_ConditionComponent_typeable) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ConditionComponent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("properties", m.GetProperties())
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
    {
        err := writer.WriteObjectValue("type", m.GetTypeEscaped())
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
func (m *ConditionComponent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Change detection model component description
func (m *ConditionComponent) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. Change detection model component name
func (m *ConditionComponent) SetName(value *string)() {
    m.name = value
}
// SetProperties sets the properties property value. Map of properties for component
func (m *ConditionComponent) SetProperties(value ConditionComponent_propertiesable)() {
    m.properties = value
}
// SetTitle sets the title property value. Change detection model component title
func (m *ConditionComponent) SetTitle(value *string)() {
    m.title = value
}
// SetTypeEscaped sets the type property value. UI Component type
func (m *ConditionComponent) SetTypeEscaped(value ConditionComponent_ConditionComponent_typeable)() {
    m.typeEscaped = value
}
type ConditionComponentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetName()(*string)
    GetProperties()(ConditionComponent_propertiesable)
    GetTitle()(*string)
    GetTypeEscaped()(ConditionComponent_ConditionComponent_typeable)
    SetDescription(value *string)()
    SetName(value *string)()
    SetProperties(value ConditionComponent_propertiesable)()
    SetTitle(value *string)()
    SetTypeEscaped(value ConditionComponent_ConditionComponent_typeable)()
}
