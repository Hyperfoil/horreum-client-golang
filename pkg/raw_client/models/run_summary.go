package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RunSummary struct {
    ProtectedTimeType
    // Array of datasets ids
    datasets []int32
    // Run description
    description *string
    // does Run have metadata uploaded alongside Run data
    hasMetadata *bool
    // Run unique ID
    id *int32
    // List of all Schema Usages for Run
    schemas []SchemaUsageable
    // test ID run relates to
    testid *int32
    // test ID run relates to
    testname *string
    // has Run been trashed in the UI
    trashed *bool
    // Array of validation errors
    validationErrors []ValidationErrorable
}
// NewRunSummary instantiates a new RunSummary and sets the default values.
func NewRunSummary()(*RunSummary) {
    m := &RunSummary{
        ProtectedTimeType: *NewProtectedTimeType(),
    }
    return m
}
// CreateRunSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunSummary(), nil
}
// GetDatasets gets the datasets property value. Array of datasets ids
// returns a []int32 when successful
func (m *RunSummary) GetDatasets()([]int32) {
    return m.datasets
}
// GetDescription gets the description property value. Run description
// returns a *string when successful
func (m *RunSummary) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RunSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectedTimeType.GetFieldDeserializers()
    res["datasets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetDatasets(res)
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
    res["hasMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMetadata(val)
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
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSchemaUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SchemaUsageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SchemaUsageable)
                }
            }
            m.SetSchemas(res)
        }
        return nil
    }
    res["testid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestid(val)
        }
        return nil
    }
    res["testname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestname(val)
        }
        return nil
    }
    res["trashed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrashed(val)
        }
        return nil
    }
    res["validationErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateValidationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ValidationErrorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ValidationErrorable)
                }
            }
            m.SetValidationErrors(res)
        }
        return nil
    }
    return res
}
// GetHasMetadata gets the hasMetadata property value. does Run have metadata uploaded alongside Run data
// returns a *bool when successful
func (m *RunSummary) GetHasMetadata()(*bool) {
    return m.hasMetadata
}
// GetId gets the id property value. Run unique ID
// returns a *int32 when successful
func (m *RunSummary) GetId()(*int32) {
    return m.id
}
// GetSchemas gets the schemas property value. List of all Schema Usages for Run
// returns a []SchemaUsageable when successful
func (m *RunSummary) GetSchemas()([]SchemaUsageable) {
    return m.schemas
}
// GetTestid gets the testid property value. test ID run relates to
// returns a *int32 when successful
func (m *RunSummary) GetTestid()(*int32) {
    return m.testid
}
// GetTestname gets the testname property value. test ID run relates to
// returns a *string when successful
func (m *RunSummary) GetTestname()(*string) {
    return m.testname
}
// GetTrashed gets the trashed property value. has Run been trashed in the UI
// returns a *bool when successful
func (m *RunSummary) GetTrashed()(*bool) {
    return m.trashed
}
// GetValidationErrors gets the validationErrors property value. Array of validation errors
// returns a []ValidationErrorable when successful
func (m *RunSummary) GetValidationErrors()([]ValidationErrorable) {
    return m.validationErrors
}
// Serialize serializes information the current object
func (m *RunSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectedTimeType.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDatasets() != nil {
        err = writer.WriteCollectionOfInt32Values("datasets", m.GetDatasets())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasMetadata", m.GetHasMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetSchemas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchemas()))
        for i, v := range m.GetSchemas() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("schemas", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("testid", m.GetTestid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("testname", m.GetTestname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("trashed", m.GetTrashed())
        if err != nil {
            return err
        }
    }
    if m.GetValidationErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValidationErrors()))
        for i, v := range m.GetValidationErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("validationErrors", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDatasets sets the datasets property value. Array of datasets ids
func (m *RunSummary) SetDatasets(value []int32)() {
    m.datasets = value
}
// SetDescription sets the description property value. Run description
func (m *RunSummary) SetDescription(value *string)() {
    m.description = value
}
// SetHasMetadata sets the hasMetadata property value. does Run have metadata uploaded alongside Run data
func (m *RunSummary) SetHasMetadata(value *bool)() {
    m.hasMetadata = value
}
// SetId sets the id property value. Run unique ID
func (m *RunSummary) SetId(value *int32)() {
    m.id = value
}
// SetSchemas sets the schemas property value. List of all Schema Usages for Run
func (m *RunSummary) SetSchemas(value []SchemaUsageable)() {
    m.schemas = value
}
// SetTestid sets the testid property value. test ID run relates to
func (m *RunSummary) SetTestid(value *int32)() {
    m.testid = value
}
// SetTestname sets the testname property value. test ID run relates to
func (m *RunSummary) SetTestname(value *string)() {
    m.testname = value
}
// SetTrashed sets the trashed property value. has Run been trashed in the UI
func (m *RunSummary) SetTrashed(value *bool)() {
    m.trashed = value
}
// SetValidationErrors sets the validationErrors property value. Array of validation errors
func (m *RunSummary) SetValidationErrors(value []ValidationErrorable)() {
    m.validationErrors = value
}
type RunSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectedTimeTypeable
    GetDatasets()([]int32)
    GetDescription()(*string)
    GetHasMetadata()(*bool)
    GetId()(*int32)
    GetSchemas()([]SchemaUsageable)
    GetTestid()(*int32)
    GetTestname()(*string)
    GetTrashed()(*bool)
    GetValidationErrors()([]ValidationErrorable)
    SetDatasets(value []int32)()
    SetDescription(value *string)()
    SetHasMetadata(value *bool)()
    SetId(value *int32)()
    SetSchemas(value []SchemaUsageable)()
    SetTestid(value *int32)()
    SetTestname(value *string)()
    SetTrashed(value *bool)()
    SetValidationErrors(value []ValidationErrorable)()
}
