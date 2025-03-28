package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DatasetSummary struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *DatasetSummary_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Dataset description
    description *string
    // Unique Dataset ID
    id *int32
    // Ordinal position of Dataset Summary on returned List
    ordinal *int32
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // Run ID that Dataset relates to
    runId *int32
    // List of Schema usages
    schemas []SchemaUsageable
    // Run Start timestamp
    start *int64
    // Run Stop timestamp
    stop *int64
    // Test ID that Dataset relates to
    testId *int32
    // Test name that the Dataset relates to
    testname *string
    // List of Validation Errors
    validationErrors []ValidationErrorable
    // map of view component ids to the LabelValueMap to render the component for this dataset
    view DatasetSummary_viewable
}
// NewDatasetSummary instantiates a new DatasetSummary and sets the default values.
func NewDatasetSummary()(*DatasetSummary) {
    m := &DatasetSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDatasetSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatasetSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatasetSummary(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *DatasetSummary_access when successful
func (m *DatasetSummary) GetAccess()(*DatasetSummary_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DatasetSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Dataset description
// returns a *string when successful
func (m *DatasetSummary) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DatasetSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDatasetSummary_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*DatasetSummary_access))
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
    res["testId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestId(val)
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
    res["view"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDatasetSummary_viewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetView(val.(DatasetSummary_viewable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique Dataset ID
// returns a *int32 when successful
func (m *DatasetSummary) GetId()(*int32) {
    return m.id
}
// GetOrdinal gets the ordinal property value. Ordinal position of Dataset Summary on returned List
// returns a *int32 when successful
func (m *DatasetSummary) GetOrdinal()(*int32) {
    return m.ordinal
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *DatasetSummary) GetOwner()(*string) {
    return m.owner
}
// GetRunId gets the runId property value. Run ID that Dataset relates to
// returns a *int32 when successful
func (m *DatasetSummary) GetRunId()(*int32) {
    return m.runId
}
// GetSchemas gets the schemas property value. List of Schema usages
// returns a []SchemaUsageable when successful
func (m *DatasetSummary) GetSchemas()([]SchemaUsageable) {
    return m.schemas
}
// GetStart gets the start property value. Run Start timestamp
// returns a *int64 when successful
func (m *DatasetSummary) GetStart()(*int64) {
    return m.start
}
// GetStop gets the stop property value. Run Stop timestamp
// returns a *int64 when successful
func (m *DatasetSummary) GetStop()(*int64) {
    return m.stop
}
// GetTestId gets the testId property value. Test ID that Dataset relates to
// returns a *int32 when successful
func (m *DatasetSummary) GetTestId()(*int32) {
    return m.testId
}
// GetTestname gets the testname property value. Test name that the Dataset relates to
// returns a *string when successful
func (m *DatasetSummary) GetTestname()(*string) {
    return m.testname
}
// GetValidationErrors gets the validationErrors property value. List of Validation Errors
// returns a []ValidationErrorable when successful
func (m *DatasetSummary) GetValidationErrors()([]ValidationErrorable) {
    return m.validationErrors
}
// GetView gets the view property value. map of view component ids to the LabelValueMap to render the component for this dataset
// returns a DatasetSummary_viewable when successful
func (m *DatasetSummary) GetView()(DatasetSummary_viewable) {
    return m.view
}
// Serialize serializes information the current object
func (m *DatasetSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetSchemas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchemas()))
        for i, v := range m.GetSchemas() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("schemas", cast)
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
        err := writer.WriteInt32Value("testId", m.GetTestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("testname", m.GetTestname())
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
        err := writer.WriteObjectValue("view", m.GetView())
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
func (m *DatasetSummary) SetAccess(value *DatasetSummary_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DatasetSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Dataset description
func (m *DatasetSummary) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. Unique Dataset ID
func (m *DatasetSummary) SetId(value *int32)() {
    m.id = value
}
// SetOrdinal sets the ordinal property value. Ordinal position of Dataset Summary on returned List
func (m *DatasetSummary) SetOrdinal(value *int32)() {
    m.ordinal = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *DatasetSummary) SetOwner(value *string)() {
    m.owner = value
}
// SetRunId sets the runId property value. Run ID that Dataset relates to
func (m *DatasetSummary) SetRunId(value *int32)() {
    m.runId = value
}
// SetSchemas sets the schemas property value. List of Schema usages
func (m *DatasetSummary) SetSchemas(value []SchemaUsageable)() {
    m.schemas = value
}
// SetStart sets the start property value. Run Start timestamp
func (m *DatasetSummary) SetStart(value *int64)() {
    m.start = value
}
// SetStop sets the stop property value. Run Stop timestamp
func (m *DatasetSummary) SetStop(value *int64)() {
    m.stop = value
}
// SetTestId sets the testId property value. Test ID that Dataset relates to
func (m *DatasetSummary) SetTestId(value *int32)() {
    m.testId = value
}
// SetTestname sets the testname property value. Test name that the Dataset relates to
func (m *DatasetSummary) SetTestname(value *string)() {
    m.testname = value
}
// SetValidationErrors sets the validationErrors property value. List of Validation Errors
func (m *DatasetSummary) SetValidationErrors(value []ValidationErrorable)() {
    m.validationErrors = value
}
// SetView sets the view property value. map of view component ids to the LabelValueMap to render the component for this dataset
func (m *DatasetSummary) SetView(value DatasetSummary_viewable)() {
    m.view = value
}
type DatasetSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*DatasetSummary_access)
    GetDescription()(*string)
    GetId()(*int32)
    GetOrdinal()(*int32)
    GetOwner()(*string)
    GetRunId()(*int32)
    GetSchemas()([]SchemaUsageable)
    GetStart()(*int64)
    GetStop()(*int64)
    GetTestId()(*int32)
    GetTestname()(*string)
    GetValidationErrors()([]ValidationErrorable)
    GetView()(DatasetSummary_viewable)
    SetAccess(value *DatasetSummary_access)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetOrdinal(value *int32)()
    SetOwner(value *string)()
    SetRunId(value *int32)()
    SetSchemas(value []SchemaUsageable)()
    SetStart(value *int64)()
    SetStop(value *int64)()
    SetTestId(value *int32)()
    SetTestname(value *string)()
    SetValidationErrors(value []ValidationErrorable)()
    SetView(value DatasetSummary_viewable)()
}
