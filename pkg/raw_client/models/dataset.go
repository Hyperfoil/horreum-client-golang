package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Dataset a dataset is the JSON document used as the basis for all comparisons and reporting
type Dataset struct {
    ProtectedTimeType
    // Data payload
    data *string
    // Run description
    description *string
    // Dataset Unique ID
    id *int32
    // Dataset ordinal for ordered list of Datasets derived from a Run
    ordinal *int32
    // Run ID that Dataset relates to
    runId *int32
    // Test ID that Dataset relates to
    testid *int32
    // List of Validation Errors
    validationErrors []ValidationErrorable
}
// NewDataset instantiates a new Dataset and sets the default values.
func NewDataset()(*Dataset) {
    m := &Dataset{
        ProtectedTimeType: *NewProtectedTimeType(),
    }
    return m
}
// CreateDatasetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatasetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataset(), nil
}
// GetData gets the data property value. Data payload
// returns a *string when successful
func (m *Dataset) GetData()(*string) {
    return m.data
}
// GetDescription gets the description property value. Run description
// returns a *string when successful
func (m *Dataset) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Dataset) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectedTimeType.GetFieldDeserializers()
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val)
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
    res["ordinal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrdinal(val)
        }
        return nil
    }
    res["runId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunId(val)
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
// GetId gets the id property value. Dataset Unique ID
// returns a *int32 when successful
func (m *Dataset) GetId()(*int32) {
    return m.id
}
// GetOrdinal gets the ordinal property value. Dataset ordinal for ordered list of Datasets derived from a Run
// returns a *int32 when successful
func (m *Dataset) GetOrdinal()(*int32) {
    return m.ordinal
}
// GetRunId gets the runId property value. Run ID that Dataset relates to
// returns a *int32 when successful
func (m *Dataset) GetRunId()(*int32) {
    return m.runId
}
// GetTestid gets the testid property value. Test ID that Dataset relates to
// returns a *int32 when successful
func (m *Dataset) GetTestid()(*int32) {
    return m.testid
}
// GetValidationErrors gets the validationErrors property value. List of Validation Errors
// returns a []ValidationErrorable when successful
func (m *Dataset) GetValidationErrors()([]ValidationErrorable) {
    return m.validationErrors
}
// Serialize serializes information the current object
func (m *Dataset) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectedTimeType.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("data", m.GetData())
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
        err = writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ordinal", m.GetOrdinal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("runId", m.GetRunId())
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
// SetData sets the data property value. Data payload
func (m *Dataset) SetData(value *string)() {
    m.data = value
}
// SetDescription sets the description property value. Run description
func (m *Dataset) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. Dataset Unique ID
func (m *Dataset) SetId(value *int32)() {
    m.id = value
}
// SetOrdinal sets the ordinal property value. Dataset ordinal for ordered list of Datasets derived from a Run
func (m *Dataset) SetOrdinal(value *int32)() {
    m.ordinal = value
}
// SetRunId sets the runId property value. Run ID that Dataset relates to
func (m *Dataset) SetRunId(value *int32)() {
    m.runId = value
}
// SetTestid sets the testid property value. Test ID that Dataset relates to
func (m *Dataset) SetTestid(value *int32)() {
    m.testid = value
}
// SetValidationErrors sets the validationErrors property value. List of Validation Errors
func (m *Dataset) SetValidationErrors(value []ValidationErrorable)() {
    m.validationErrors = value
}
type Datasetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectedTimeTypeable
    GetData()(*string)
    GetDescription()(*string)
    GetId()(*int32)
    GetOrdinal()(*int32)
    GetRunId()(*int32)
    GetTestid()(*int32)
    GetValidationErrors()([]ValidationErrorable)
    SetData(value *string)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetOrdinal(value *int32)()
    SetRunId(value *int32)()
    SetTestid(value *int32)()
    SetValidationErrors(value []ValidationErrorable)()
}
