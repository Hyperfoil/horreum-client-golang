// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SchemaExport represents a Schema with all associated data used for export/import operations.
type SchemaExport struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *SchemaExport_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Schema Description
    description *string
    // Unique Schema ID
    id *int32
    // Array of Labels associated with schema
    labels []Labelable
    // Schema name
    name *string
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // JSON validation schema. Used to validate uploaded JSON documents
    schema *string
    // Array of Transformers associated with schema
    transformers []Transformerable
    // Unique, versioned schema URI
    uri *string
}
// NewSchemaExport instantiates a new SchemaExport and sets the default values.
func NewSchemaExport()(*SchemaExport) {
    m := &SchemaExport{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSchemaExportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemaExportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchemaExport(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *SchemaExport_access when successful
func (m *SchemaExport) GetAccess()(*SchemaExport_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SchemaExport) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Schema Description
// returns a *string when successful
func (m *SchemaExport) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SchemaExport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSchemaExport_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*SchemaExport_access))
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
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Labelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Labelable)
                }
            }
            m.SetLabels(res)
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
    res["transformers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTransformerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Transformerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Transformerable)
                }
            }
            m.SetTransformers(res)
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
func (m *SchemaExport) GetId()(*int32) {
    return m.id
}
// GetLabels gets the labels property value. Array of Labels associated with schema
// returns a []Labelable when successful
func (m *SchemaExport) GetLabels()([]Labelable) {
    return m.labels
}
// GetName gets the name property value. Schema name
// returns a *string when successful
func (m *SchemaExport) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *SchemaExport) GetOwner()(*string) {
    return m.owner
}
// GetSchema gets the schema property value. JSON validation schema. Used to validate uploaded JSON documents
// returns a *string when successful
func (m *SchemaExport) GetSchema()(*string) {
    return m.schema
}
// GetTransformers gets the transformers property value. Array of Transformers associated with schema
// returns a []Transformerable when successful
func (m *SchemaExport) GetTransformers()([]Transformerable) {
    return m.transformers
}
// GetUri gets the uri property value. Unique, versioned schema URI
// returns a *string when successful
func (m *SchemaExport) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *SchemaExport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("labels", cast)
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
    if m.GetTransformers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransformers()))
        for i, v := range m.GetTransformers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("transformers", cast)
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
func (m *SchemaExport) SetAccess(value *SchemaExport_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SchemaExport) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Schema Description
func (m *SchemaExport) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. Unique Schema ID
func (m *SchemaExport) SetId(value *int32)() {
    m.id = value
}
// SetLabels sets the labels property value. Array of Labels associated with schema
func (m *SchemaExport) SetLabels(value []Labelable)() {
    m.labels = value
}
// SetName sets the name property value. Schema name
func (m *SchemaExport) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *SchemaExport) SetOwner(value *string)() {
    m.owner = value
}
// SetSchema sets the schema property value. JSON validation schema. Used to validate uploaded JSON documents
func (m *SchemaExport) SetSchema(value *string)() {
    m.schema = value
}
// SetTransformers sets the transformers property value. Array of Transformers associated with schema
func (m *SchemaExport) SetTransformers(value []Transformerable)() {
    m.transformers = value
}
// SetUri sets the uri property value. Unique, versioned schema URI
func (m *SchemaExport) SetUri(value *string)() {
    m.uri = value
}
type SchemaExportable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*SchemaExport_access)
    GetDescription()(*string)
    GetId()(*int32)
    GetLabels()([]Labelable)
    GetName()(*string)
    GetOwner()(*string)
    GetSchema()(*string)
    GetTransformers()([]Transformerable)
    GetUri()(*string)
    SetAccess(value *SchemaExport_access)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetLabels(value []Labelable)()
    SetName(value *string)()
    SetOwner(value *string)()
    SetSchema(value *string)()
    SetTransformers(value []Transformerable)()
    SetUri(value *string)()
}
