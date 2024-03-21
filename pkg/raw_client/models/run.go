package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Run struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *Run_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Run result payload
    data *string
    // Collection of Datasets derived from Run payload
    datasets []Datasetable
    // Run description
    description *string
    // Unique Run ID
    id *int32
    // JSON metadata related to run, can be tool configuration etc
    metadata *string
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // Run Start timestamp
    start *int64
    // Run Stop timestamp
    stop *int64
    // Test ID run relates to
    testid *int32
    // Has Run been deleted from UI
    trashed *bool
    // Collection of Validation Errors in Run payload
    validationErrors []ValidationErrorable
}
// NewRun instantiates a new Run and sets the default values.
func NewRun()(*Run) {
    m := &Run{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRunFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRun(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *Run_access when successful
func (m *Run) GetAccess()(*Run_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Run) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. Run result payload
// returns a *string when successful
func (m *Run) GetData()(*string) {
    return m.data
}
// GetDatasets gets the datasets property value. Collection of Datasets derived from Run payload
// returns a []Datasetable when successful
func (m *Run) GetDatasets()([]Datasetable) {
    return m.datasets
}
// GetDescription gets the description property value. Run description
// returns a *string when successful
func (m *Run) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Run) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRun_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*Run_access))
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
    res["datasets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDatasetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Datasetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Datasetable)
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
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val)
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
// GetId gets the id property value. Unique Run ID
// returns a *int32 when successful
func (m *Run) GetId()(*int32) {
    return m.id
}
// GetMetadata gets the metadata property value. JSON metadata related to run, can be tool configuration etc
// returns a *string when successful
func (m *Run) GetMetadata()(*string) {
    return m.metadata
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *Run) GetOwner()(*string) {
    return m.owner
}
// GetStart gets the start property value. Run Start timestamp
// returns a *int64 when successful
func (m *Run) GetStart()(*int64) {
    return m.start
}
// GetStop gets the stop property value. Run Stop timestamp
// returns a *int64 when successful
func (m *Run) GetStop()(*int64) {
    return m.stop
}
// GetTestid gets the testid property value. Test ID run relates to
// returns a *int32 when successful
func (m *Run) GetTestid()(*int32) {
    return m.testid
}
// GetTrashed gets the trashed property value. Has Run been deleted from UI
// returns a *bool when successful
func (m *Run) GetTrashed()(*bool) {
    return m.trashed
}
// GetValidationErrors gets the validationErrors property value. Collection of Validation Errors in Run payload
// returns a []ValidationErrorable when successful
func (m *Run) GetValidationErrors()([]ValidationErrorable) {
    return m.validationErrors
}
// Serialize serializes information the current object
func (m *Run) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetDatasets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDatasets()))
        for i, v := range m.GetDatasets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("datasets", cast)
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
        err := writer.WriteStringValue("metadata", m.GetMetadata())
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
    {
        err := writer.WriteBoolValue("trashed", m.GetTrashed())
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
func (m *Run) SetAccess(value *Run_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Run) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. Run result payload
func (m *Run) SetData(value *string)() {
    m.data = value
}
// SetDatasets sets the datasets property value. Collection of Datasets derived from Run payload
func (m *Run) SetDatasets(value []Datasetable)() {
    m.datasets = value
}
// SetDescription sets the description property value. Run description
func (m *Run) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. Unique Run ID
func (m *Run) SetId(value *int32)() {
    m.id = value
}
// SetMetadata sets the metadata property value. JSON metadata related to run, can be tool configuration etc
func (m *Run) SetMetadata(value *string)() {
    m.metadata = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *Run) SetOwner(value *string)() {
    m.owner = value
}
// SetStart sets the start property value. Run Start timestamp
func (m *Run) SetStart(value *int64)() {
    m.start = value
}
// SetStop sets the stop property value. Run Stop timestamp
func (m *Run) SetStop(value *int64)() {
    m.stop = value
}
// SetTestid sets the testid property value. Test ID run relates to
func (m *Run) SetTestid(value *int32)() {
    m.testid = value
}
// SetTrashed sets the trashed property value. Has Run been deleted from UI
func (m *Run) SetTrashed(value *bool)() {
    m.trashed = value
}
// SetValidationErrors sets the validationErrors property value. Collection of Validation Errors in Run payload
func (m *Run) SetValidationErrors(value []ValidationErrorable)() {
    m.validationErrors = value
}
type Runable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*Run_access)
    GetData()(*string)
    GetDatasets()([]Datasetable)
    GetDescription()(*string)
    GetId()(*int32)
    GetMetadata()(*string)
    GetOwner()(*string)
    GetStart()(*int64)
    GetStop()(*int64)
    GetTestid()(*int32)
    GetTrashed()(*bool)
    GetValidationErrors()([]ValidationErrorable)
    SetAccess(value *Run_access)()
    SetData(value *string)()
    SetDatasets(value []Datasetable)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetMetadata(value *string)()
    SetOwner(value *string)()
    SetStart(value *int64)()
    SetStop(value *int64)()
    SetTestid(value *int32)()
    SetTrashed(value *bool)()
    SetValidationErrors(value []ValidationErrorable)()
}
