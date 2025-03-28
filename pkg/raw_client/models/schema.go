package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Schema data object that describes the schema definition for a test
type Schema struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *Schema_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Schema Description
    description *string
    // Unique Schema ID
    id *int32
    // Schema name
    name *string
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // JSON validation schema. Used to validate uploaded JSON documents
    schema *string
    // Unique, versioned schema URI
    uri *string
}
// NewSchema instantiates a new Schema and sets the default values.
func NewSchema()(*Schema) {
    m := &Schema{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchema(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *Schema_access when successful
func (m *Schema) GetAccess()(*Schema_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Schema) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Schema Description
// returns a *string when successful
func (m *Schema) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Schema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSchema_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*Schema_access))
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val)
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
// GetId gets the id property value. Unique Schema ID
// returns a *int32 when successful
func (m *Schema) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Schema name
// returns a *string when successful
func (m *Schema) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *Schema) GetOwner()(*string) {
    return m.owner
}
// GetSchema gets the schema property value. JSON validation schema. Used to validate uploaded JSON documents
// returns a *string when successful
func (m *Schema) GetSchema()(*string) {
    return m.schema
}
// GetUri gets the uri property value. Unique, versioned schema URI
// returns a *string when successful
func (m *Schema) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *Schema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccess() != nil {
        cast := (*m.GetAccess()).String()
        err := writer.WriteStringValue("access", &cast)
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
        err := writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schema", m.GetSchema())
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
// SetAccess sets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
func (m *Schema) SetAccess(value *Schema_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Schema) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Schema Description
func (m *Schema) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. Unique Schema ID
func (m *Schema) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Schema name
func (m *Schema) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *Schema) SetOwner(value *string)() {
    m.owner = value
}
// SetSchema sets the schema property value. JSON validation schema. Used to validate uploaded JSON documents
func (m *Schema) SetSchema(value *string)() {
    m.schema = value
}
// SetUri sets the uri property value. Unique, versioned schema URI
func (m *Schema) SetUri(value *string)() {
    m.uri = value
}
type Schemaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*Schema_access)
    GetDescription()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetOwner()(*string)
    GetSchema()(*string)
    GetUri()(*string)
    SetAccess(value *Schema_access)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetOwner(value *string)()
    SetSchema(value *string)()
    SetUri(value *string)()
}
