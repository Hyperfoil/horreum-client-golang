package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SchemaUsage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Does schema have a JSON validation schema defined?
    hasJsonSchema *bool
    // Schema unique ID
    id *int32
    // Ordinal position of schema usage in Run/Dataset
    key *string
    // Schema name
    name *string
    // Source of schema usage, 0 is data, 1 is metadata. DataSets always use 0
    source *int32
    // Location of Schema Usage, 0 for Run, 1 for Dataset
    typeEscaped *int32
    // Schema name
    uri *string
}
// NewSchemaUsage instantiates a new SchemaUsage and sets the default values.
func NewSchemaUsage()(*SchemaUsage) {
    m := &SchemaUsage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSchemaUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemaUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchemaUsage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SchemaUsage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SchemaUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hasJsonSchema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasJsonSchema(val)
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
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
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
// GetHasJsonSchema gets the hasJsonSchema property value. Does schema have a JSON validation schema defined?
// returns a *bool when successful
func (m *SchemaUsage) GetHasJsonSchema()(*bool) {
    return m.hasJsonSchema
}
// GetId gets the id property value. Schema unique ID
// returns a *int32 when successful
func (m *SchemaUsage) GetId()(*int32) {
    return m.id
}
// GetKey gets the key property value. Ordinal position of schema usage in Run/Dataset
// returns a *string when successful
func (m *SchemaUsage) GetKey()(*string) {
    return m.key
}
// GetName gets the name property value. Schema name
// returns a *string when successful
func (m *SchemaUsage) GetName()(*string) {
    return m.name
}
// GetSource gets the source property value. Source of schema usage, 0 is data, 1 is metadata. DataSets always use 0
// returns a *int32 when successful
func (m *SchemaUsage) GetSource()(*int32) {
    return m.source
}
// GetTypeEscaped gets the type property value. Location of Schema Usage, 0 for Run, 1 for Dataset
// returns a *int32 when successful
func (m *SchemaUsage) GetTypeEscaped()(*int32) {
    return m.typeEscaped
}
// GetUri gets the uri property value. Schema name
// returns a *string when successful
func (m *SchemaUsage) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *SchemaUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hasJsonSchema", m.GetHasJsonSchema())
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
        err := writer.WriteStringValue("key", m.GetKey())
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
        err := writer.WriteInt32Value("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("type", m.GetTypeEscaped())
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
func (m *SchemaUsage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHasJsonSchema sets the hasJsonSchema property value. Does schema have a JSON validation schema defined?
func (m *SchemaUsage) SetHasJsonSchema(value *bool)() {
    m.hasJsonSchema = value
}
// SetId sets the id property value. Schema unique ID
func (m *SchemaUsage) SetId(value *int32)() {
    m.id = value
}
// SetKey sets the key property value. Ordinal position of schema usage in Run/Dataset
func (m *SchemaUsage) SetKey(value *string)() {
    m.key = value
}
// SetName sets the name property value. Schema name
func (m *SchemaUsage) SetName(value *string)() {
    m.name = value
}
// SetSource sets the source property value. Source of schema usage, 0 is data, 1 is metadata. DataSets always use 0
func (m *SchemaUsage) SetSource(value *int32)() {
    m.source = value
}
// SetTypeEscaped sets the type property value. Location of Schema Usage, 0 for Run, 1 for Dataset
func (m *SchemaUsage) SetTypeEscaped(value *int32)() {
    m.typeEscaped = value
}
// SetUri sets the uri property value. Schema name
func (m *SchemaUsage) SetUri(value *string)() {
    m.uri = value
}
type SchemaUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHasJsonSchema()(*bool)
    GetId()(*int32)
    GetKey()(*string)
    GetName()(*string)
    GetSource()(*int32)
    GetTypeEscaped()(*int32)
    GetUri()(*string)
    SetHasJsonSchema(value *bool)()
    SetId(value *int32)()
    SetKey(value *string)()
    SetName(value *string)()
    SetSource(value *int32)()
    SetTypeEscaped(value *int32)()
    SetUri(value *string)()
}
