package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelValue label Value derived from Label definition and Dataset Data
type LabelValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique ID for Label Value
    id *int32
    // Label name
    name *string
    // Summary description of Schema
    schema SchemaDescriptorable
    // Value value extracted from Dataset. This can be a scalar, array or JSON object
    value *string
}
// NewLabelValue instantiates a new LabelValue and sets the default values.
func NewLabelValue()(*LabelValue) {
    m := &LabelValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LabelValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSchemaDescriptorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(SchemaDescriptorable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// GetId gets the id property value. Unique ID for Label Value
// returns a *int32 when successful
func (m *LabelValue) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Label name
// returns a *string when successful
func (m *LabelValue) GetName()(*string) {
    return m.name
}
// GetSchema gets the schema property value. Summary description of Schema
// returns a SchemaDescriptorable when successful
func (m *LabelValue) GetSchema()(SchemaDescriptorable) {
    return m.schema
}
// GetValue gets the value property value. Value value extracted from Dataset. This can be a scalar, array or JSON object
// returns a *string when successful
func (m *LabelValue) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *LabelValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
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
        err := writer.WriteObjectValue("schema", m.GetSchema())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *LabelValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Unique ID for Label Value
func (m *LabelValue) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Label name
func (m *LabelValue) SetName(value *string)() {
    m.name = value
}
// SetSchema sets the schema property value. Summary description of Schema
func (m *LabelValue) SetSchema(value SchemaDescriptorable)() {
    m.schema = value
}
// SetValue sets the value property value. Value value extracted from Dataset. This can be a scalar, array or JSON object
func (m *LabelValue) SetValue(value *string)() {
    m.value = value
}
type LabelValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetName()(*string)
    GetSchema()(SchemaDescriptorable)
    GetValue()(*string)
    SetId(value *int32)()
    SetName(value *string)()
    SetSchema(value SchemaDescriptorable)()
    SetValue(value *string)()
}
