package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Label a Label is a core component of Horreum, defining which components of the JSON document are part of a KPI and how the metric values are calculated
type Label struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *Label_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A collection of Extractors, that will be combined in the Combination Function
    extractors []Extractorable
    // Is Label a filtering label? Filtering labels contains values that are used to filter datasets for comparison
    filtering *bool
    // A Combination Function that defines how values from Extractors are combined to produce a Label Value
    function *string
    // Unique ID for Label
    id *int32
    // Is Label a metrics label? Metrics labels are contain Metrics that are used for comparison
    metrics *bool
    // Name for label. NOTE: all Labels are considered to have the same semantic meaning throughout the entire system
    name *string
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // Schema ID that the Label relates to
    schemaId *int32
}
// NewLabel instantiates a new Label and sets the default values.
func NewLabel()(*Label) {
    m := &Label{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabel(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *Label_access when successful
func (m *Label) GetAccess()(*Label_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Label) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExtractors gets the extractors property value. A collection of Extractors, that will be combined in the Combination Function
// returns a []Extractorable when successful
func (m *Label) GetExtractors()([]Extractorable) {
    return m.extractors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Label) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLabel_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*Label_access))
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
    res["filtering"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFiltering(val)
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
    res["metrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetrics(val)
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
    return res
}
// GetFiltering gets the filtering property value. Is Label a filtering label? Filtering labels contains values that are used to filter datasets for comparison
// returns a *bool when successful
func (m *Label) GetFiltering()(*bool) {
    return m.filtering
}
// GetFunction gets the function property value. A Combination Function that defines how values from Extractors are combined to produce a Label Value
// returns a *string when successful
func (m *Label) GetFunction()(*string) {
    return m.function
}
// GetId gets the id property value. Unique ID for Label
// returns a *int32 when successful
func (m *Label) GetId()(*int32) {
    return m.id
}
// GetMetrics gets the metrics property value. Is Label a metrics label? Metrics labels are contain Metrics that are used for comparison
// returns a *bool when successful
func (m *Label) GetMetrics()(*bool) {
    return m.metrics
}
// GetName gets the name property value. Name for label. NOTE: all Labels are considered to have the same semantic meaning throughout the entire system
// returns a *string when successful
func (m *Label) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *Label) GetOwner()(*string) {
    return m.owner
}
// GetSchemaId gets the schemaId property value. Schema ID that the Label relates to
// returns a *int32 when successful
func (m *Label) GetSchemaId()(*int32) {
    return m.schemaId
}
// Serialize serializes information the current object
func (m *Label) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccess() != nil {
        cast := (*m.GetAccess()).String()
        err := writer.WriteStringValue("access", &cast)
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
        err := writer.WriteBoolValue("filtering", m.GetFiltering())
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
        err := writer.WriteBoolValue("metrics", m.GetMetrics())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccess sets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
func (m *Label) SetAccess(value *Label_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Label) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExtractors sets the extractors property value. A collection of Extractors, that will be combined in the Combination Function
func (m *Label) SetExtractors(value []Extractorable)() {
    m.extractors = value
}
// SetFiltering sets the filtering property value. Is Label a filtering label? Filtering labels contains values that are used to filter datasets for comparison
func (m *Label) SetFiltering(value *bool)() {
    m.filtering = value
}
// SetFunction sets the function property value. A Combination Function that defines how values from Extractors are combined to produce a Label Value
func (m *Label) SetFunction(value *string)() {
    m.function = value
}
// SetId sets the id property value. Unique ID for Label
func (m *Label) SetId(value *int32)() {
    m.id = value
}
// SetMetrics sets the metrics property value. Is Label a metrics label? Metrics labels are contain Metrics that are used for comparison
func (m *Label) SetMetrics(value *bool)() {
    m.metrics = value
}
// SetName sets the name property value. Name for label. NOTE: all Labels are considered to have the same semantic meaning throughout the entire system
func (m *Label) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *Label) SetOwner(value *string)() {
    m.owner = value
}
// SetSchemaId sets the schemaId property value. Schema ID that the Label relates to
func (m *Label) SetSchemaId(value *int32)() {
    m.schemaId = value
}
type Labelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*Label_access)
    GetExtractors()([]Extractorable)
    GetFiltering()(*bool)
    GetFunction()(*string)
    GetId()(*int32)
    GetMetrics()(*bool)
    GetName()(*string)
    GetOwner()(*string)
    GetSchemaId()(*int32)
    SetAccess(value *Label_access)()
    SetExtractors(value []Extractorable)()
    SetFiltering(value *bool)()
    SetFunction(value *string)()
    SetId(value *int32)()
    SetMetrics(value *bool)()
    SetName(value *string)()
    SetOwner(value *string)()
    SetSchemaId(value *int32)()
}
