// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SchemaDescriptor summary description of Schema
type SchemaDescriptor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Schema unique ID
    id *int32
    // Schema name
    name *string
    // Schema name
    uri *string
}
// NewSchemaDescriptor instantiates a new SchemaDescriptor and sets the default values.
func NewSchemaDescriptor()(*SchemaDescriptor) {
    m := &SchemaDescriptor{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSchemaDescriptorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemaDescriptorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchemaDescriptor(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SchemaDescriptor) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SchemaDescriptor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Schema unique ID
// returns a *int32 when successful
func (m *SchemaDescriptor) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Schema name
// returns a *string when successful
func (m *SchemaDescriptor) GetName()(*string) {
    return m.name
}
// GetUri gets the uri property value. Schema name
// returns a *string when successful
func (m *SchemaDescriptor) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *SchemaDescriptor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("uri", m.GetUri())
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
func (m *SchemaDescriptor) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Schema unique ID
func (m *SchemaDescriptor) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Schema name
func (m *SchemaDescriptor) SetName(value *string)() {
    m.name = value
}
// SetUri sets the uri property value. Schema name
func (m *SchemaDescriptor) SetUri(value *string)() {
    m.uri = value
}
type SchemaDescriptorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetName()(*string)
    GetUri()(*string)
    SetId(value *int32)()
    SetName(value *string)()
    SetUri(value *string)()
}
