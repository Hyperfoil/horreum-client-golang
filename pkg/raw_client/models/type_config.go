package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TypeConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The builtIn property
    builtIn *bool
    // The enumName property
    enumName *string
    // The label property
    label *string
    // The name property
    name *string
    // The supportedAuths property
    supportedAuths []string
}
// NewTypeConfig instantiates a new TypeConfig and sets the default values.
func NewTypeConfig()(*TypeConfig) {
    m := &TypeConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTypeConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTypeConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTypeConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TypeConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuiltIn gets the builtIn property value. The builtIn property
// returns a *bool when successful
func (m *TypeConfig) GetBuiltIn()(*bool) {
    return m.builtIn
}
// GetEnumName gets the enumName property value. The enumName property
// returns a *string when successful
func (m *TypeConfig) GetEnumName()(*string) {
    return m.enumName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TypeConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["enumName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnumName(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
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
    res["supportedAuths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSupportedAuths(res)
        }
        return nil
    }
    return res
}
// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *TypeConfig) GetLabel()(*string) {
    return m.label
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *TypeConfig) GetName()(*string) {
    return m.name
}
// GetSupportedAuths gets the supportedAuths property value. The supportedAuths property
// returns a []string when successful
func (m *TypeConfig) GetSupportedAuths()([]string) {
    return m.supportedAuths
}
// Serialize serializes information the current object
func (m *TypeConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("builtIn", m.GetBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("enumName", m.GetEnumName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("label", m.GetLabel())
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
    if m.GetSupportedAuths() != nil {
        err := writer.WriteCollectionOfStringValues("supportedAuths", m.GetSupportedAuths())
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
func (m *TypeConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuiltIn sets the builtIn property value. The builtIn property
func (m *TypeConfig) SetBuiltIn(value *bool)() {
    m.builtIn = value
}
// SetEnumName sets the enumName property value. The enumName property
func (m *TypeConfig) SetEnumName(value *string)() {
    m.enumName = value
}
// SetLabel sets the label property value. The label property
func (m *TypeConfig) SetLabel(value *string)() {
    m.label = value
}
// SetName sets the name property value. The name property
func (m *TypeConfig) SetName(value *string)() {
    m.name = value
}
// SetSupportedAuths sets the supportedAuths property value. The supportedAuths property
func (m *TypeConfig) SetSupportedAuths(value []string)() {
    m.supportedAuths = value
}
type TypeConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuiltIn()(*bool)
    GetEnumName()(*string)
    GetLabel()(*string)
    GetName()(*string)
    GetSupportedAuths()([]string)
    SetBuiltIn(value *bool)()
    SetEnumName(value *string)()
    SetLabel(value *string)()
    SetName(value *string)()
    SetSupportedAuths(value []string)()
}
