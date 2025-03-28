package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Transformer a transformer extracts labals and applies a Function to convert a Run into one or more Datasets
type Transformer struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *Transformer_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Transformer description
    description *string
    // A collection of extractors to extract JSON values to create new Dataset JSON document
    extractors []Extractorable
    // The function property
    function *string
    // Unique Transformer id
    id *int32
    // Transformer name
    name *string
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // Schema ID that the transform is registered against
    schemaId *int32
    // Schema name that the transform is registered against
    schemaName *string
    // Schema Uri that the transform is registered against
    schemaUri *string
    // The schema associated with the calculated Datasets. Where a transformer creates a new JSON object with a new structure, this Schema is used to extafct values from the new Dataset JSON document
    targetSchemaUri *string
}
// NewTransformer instantiates a new Transformer and sets the default values.
func NewTransformer()(*Transformer) {
    m := &Transformer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransformerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransformerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransformer(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *Transformer_access when successful
func (m *Transformer) GetAccess()(*Transformer_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Transformer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Transformer description
// returns a *string when successful
func (m *Transformer) GetDescription()(*string) {
    return m.description
}
// GetExtractors gets the extractors property value. A collection of extractors to extract JSON values to create new Dataset JSON document
// returns a []Extractorable when successful
func (m *Transformer) GetExtractors()([]Extractorable) {
    return m.extractors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Transformer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransformer_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*Transformer_access))
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
    res["extractors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtractorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extractorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Extractorable)
                }
            }
            m.SetExtractors(res)
        }
        return nil
    }
    res["function"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunction(val)
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
    res["schemaId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaId(val)
        }
        return nil
    }
    res["schemaName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaName(val)
        }
        return nil
    }
    res["schemaUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaUri(val)
        }
        return nil
    }
    res["targetSchemaUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetSchemaUri(val)
        }
        return nil
    }
    return res
}
// GetFunction gets the function property value. The function property
// returns a *string when successful
func (m *Transformer) GetFunction()(*string) {
    return m.function
}
// GetId gets the id property value. Unique Transformer id
// returns a *int32 when successful
func (m *Transformer) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Transformer name
// returns a *string when successful
func (m *Transformer) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *Transformer) GetOwner()(*string) {
    return m.owner
}
// GetSchemaId gets the schemaId property value. Schema ID that the transform is registered against
// returns a *int32 when successful
func (m *Transformer) GetSchemaId()(*int32) {
    return m.schemaId
}
// GetSchemaName gets the schemaName property value. Schema name that the transform is registered against
// returns a *string when successful
func (m *Transformer) GetSchemaName()(*string) {
    return m.schemaName
}
// GetSchemaUri gets the schemaUri property value. Schema Uri that the transform is registered against
// returns a *string when successful
func (m *Transformer) GetSchemaUri()(*string) {
    return m.schemaUri
}
// GetTargetSchemaUri gets the targetSchemaUri property value. The schema associated with the calculated Datasets. Where a transformer creates a new JSON object with a new structure, this Schema is used to extafct values from the new Dataset JSON document
// returns a *string when successful
func (m *Transformer) GetTargetSchemaUri()(*string) {
    return m.targetSchemaUri
}
// Serialize serializes information the current object
func (m *Transformer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetExtractors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtractors()))
        for i, v := range m.GetExtractors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("extractors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("function", m.GetFunction())
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
        err := writer.WriteInt32Value("schemaId", m.GetSchemaId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schemaName", m.GetSchemaName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schemaUri", m.GetSchemaUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetSchemaUri", m.GetTargetSchemaUri())
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
func (m *Transformer) SetAccess(value *Transformer_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Transformer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Transformer description
func (m *Transformer) SetDescription(value *string)() {
    m.description = value
}
// SetExtractors sets the extractors property value. A collection of extractors to extract JSON values to create new Dataset JSON document
func (m *Transformer) SetExtractors(value []Extractorable)() {
    m.extractors = value
}
// SetFunction sets the function property value. The function property
func (m *Transformer) SetFunction(value *string)() {
    m.function = value
}
// SetId sets the id property value. Unique Transformer id
func (m *Transformer) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Transformer name
func (m *Transformer) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *Transformer) SetOwner(value *string)() {
    m.owner = value
}
// SetSchemaId sets the schemaId property value. Schema ID that the transform is registered against
func (m *Transformer) SetSchemaId(value *int32)() {
    m.schemaId = value
}
// SetSchemaName sets the schemaName property value. Schema name that the transform is registered against
func (m *Transformer) SetSchemaName(value *string)() {
    m.schemaName = value
}
// SetSchemaUri sets the schemaUri property value. Schema Uri that the transform is registered against
func (m *Transformer) SetSchemaUri(value *string)() {
    m.schemaUri = value
}
// SetTargetSchemaUri sets the targetSchemaUri property value. The schema associated with the calculated Datasets. Where a transformer creates a new JSON object with a new structure, this Schema is used to extafct values from the new Dataset JSON document
func (m *Transformer) SetTargetSchemaUri(value *string)() {
    m.targetSchemaUri = value
}
type Transformerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*Transformer_access)
    GetDescription()(*string)
    GetExtractors()([]Extractorable)
    GetFunction()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetOwner()(*string)
    GetSchemaId()(*int32)
    GetSchemaName()(*string)
    GetSchemaUri()(*string)
    GetTargetSchemaUri()(*string)
    SetAccess(value *Transformer_access)()
    SetDescription(value *string)()
    SetExtractors(value []Extractorable)()
    SetFunction(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetOwner(value *string)()
    SetSchemaId(value *int32)()
    SetSchemaName(value *string)()
    SetSchemaUri(value *string)()
    SetTargetSchemaUri(value *string)()
}
