package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Dataset a dataset is the JSON document used as the basis for all comparisons and reporting
type Dataset struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *Dataset_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Data payload
    data *string
    // Run description
    description *string
    // Dataset Unique ID
    id *int32
    // Dataset ordinal for ordered list of Datasets derived from a Run
    ordinal *int32
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // Run ID that Dataset relates to
    runId *int32
    // Run Start timestamp
    start *int64
    // Run Stop timestamp
    stop *int64
    // Test ID that Dataset relates to
    testid *int32
    // List of Validation Errors
    validationErrors []ValidationErrorable
}
// NewDataset instantiates a new Dataset and sets the default values.
func NewDataset()(*Dataset) {
    m := &Dataset{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDatasetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatasetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataset(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *Dataset_access when successful
func (m *Dataset) GetAccess()(*Dataset_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Dataset) GetAdditionalData()(map[string]any) {
    return m.additionalData
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
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataset_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*Dataset_access))
        }
        return nil
    }
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
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    res["stop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStop(val)
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
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *Dataset) GetOwner()(*string) {
    return m.owner
}
// GetRunId gets the runId property value. Run ID that Dataset relates to
// returns a *int32 when successful
func (m *Dataset) GetRunId()(*int32) {
    return m.runId
}
// GetStart gets the start property value. Run Start timestamp
// returns a *int64 when successful
func (m *Dataset) GetStart()(*int64) {
    return m.start
}
// GetStop gets the stop property value. Run Stop timestamp
// returns a *int64 when successful
func (m *Dataset) GetStop()(*int64) {
    return m.stop
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
    if m.GetAccess() != nil {
        cast := (*m.GetAccess()).String()
        err := writer.WriteStringValue("access", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("data", m.GetData())
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
        err := writer.WriteInt32Value("ordinal", m.GetOrdinal())
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
        err := writer.WriteInt32Value("runId", m.GetRunId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("stop", m.GetStop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("testid", m.GetTestid())
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
        err := writer.WriteCollectionOfObjectValues("validationErrors", cast)
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
func (m *Dataset) SetAccess(value *Dataset_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Dataset) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
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
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *Dataset) SetOwner(value *string)() {
    m.owner = value
}
// SetRunId sets the runId property value. Run ID that Dataset relates to
func (m *Dataset) SetRunId(value *int32)() {
    m.runId = value
}
// SetStart sets the start property value. Run Start timestamp
func (m *Dataset) SetStart(value *int64)() {
    m.start = value
}
// SetStop sets the stop property value. Run Stop timestamp
func (m *Dataset) SetStop(value *int64)() {
    m.stop = value
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
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*Dataset_access)
    GetData()(*string)
    GetDescription()(*string)
    GetId()(*int32)
    GetOrdinal()(*int32)
    GetOwner()(*string)
    GetRunId()(*int32)
    GetStart()(*int64)
    GetStop()(*int64)
    GetTestid()(*int32)
    GetValidationErrors()([]ValidationErrorable)
    SetAccess(value *Dataset_access)()
    SetData(value *string)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetOrdinal(value *int32)()
    SetOwner(value *string)()
    SetRunId(value *int32)()
    SetStart(value *int64)()
    SetStop(value *int64)()
    SetTestid(value *int32)()
    SetValidationErrors(value []ValidationErrorable)()
}
