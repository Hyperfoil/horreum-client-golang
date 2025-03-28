package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TestSummary struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *TestSummary_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Total number of Datasets for the Test
    datasets *float64
    // Datastore id
    datastoreId *int32
    // Description of the test
    description *string
    // Name of folder that the test is stored in. Folders allow tests to be organised in the UI
    folder *string
    // ID of tests
    id *int32
    // Test name
    name *string
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // Total number of Runs for the Test
    runs *float64
    // Subscriptions for each test for authenticated user
    watching []string
}
// NewTestSummary instantiates a new TestSummary and sets the default values.
func NewTestSummary()(*TestSummary) {
    m := &TestSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTestSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTestSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTestSummary(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *TestSummary_access when successful
func (m *TestSummary) GetAccess()(*TestSummary_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TestSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatasets gets the datasets property value. Total number of Datasets for the Test
// returns a *float64 when successful
func (m *TestSummary) GetDatasets()(*float64) {
    return m.datasets
}
// GetDatastoreId gets the datastoreId property value. Datastore id
// returns a *int32 when successful
func (m *TestSummary) GetDatastoreId()(*int32) {
    return m.datastoreId
}
// GetDescription gets the description property value. Description of the test
// returns a *string when successful
func (m *TestSummary) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TestSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTestSummary_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*TestSummary_access))
        }
        return nil
    }
    res["datasets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatasets(val)
        }
        return nil
    }
    res["datastoreId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatastoreId(val)
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
    res["folder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolder(val)
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
    res["runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuns(val)
        }
        return nil
    }
    res["watching"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWatching(res)
        }
        return nil
    }
    return res
}
// GetFolder gets the folder property value. Name of folder that the test is stored in. Folders allow tests to be organised in the UI
// returns a *string when successful
func (m *TestSummary) GetFolder()(*string) {
    return m.folder
}
// GetId gets the id property value. ID of tests
// returns a *int32 when successful
func (m *TestSummary) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Test name
// returns a *string when successful
func (m *TestSummary) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *TestSummary) GetOwner()(*string) {
    return m.owner
}
// GetRuns gets the runs property value. Total number of Runs for the Test
// returns a *float64 when successful
func (m *TestSummary) GetRuns()(*float64) {
    return m.runs
}
// GetWatching gets the watching property value. Subscriptions for each test for authenticated user
// returns a []string when successful
func (m *TestSummary) GetWatching()([]string) {
    return m.watching
}
// Serialize serializes information the current object
func (m *TestSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccess() != nil {
        cast := (*m.GetAccess()).String()
        err := writer.WriteStringValue("access", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("datasets", m.GetDatasets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("datastoreId", m.GetDatastoreId())
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
        err := writer.WriteStringValue("folder", m.GetFolder())
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
        err := writer.WriteFloat64Value("runs", m.GetRuns())
        if err != nil {
            return err
        }
    }
    if m.GetWatching() != nil {
        err := writer.WriteCollectionOfStringValues("watching", m.GetWatching())
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
func (m *TestSummary) SetAccess(value *TestSummary_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TestSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatasets sets the datasets property value. Total number of Datasets for the Test
func (m *TestSummary) SetDatasets(value *float64)() {
    m.datasets = value
}
// SetDatastoreId sets the datastoreId property value. Datastore id
func (m *TestSummary) SetDatastoreId(value *int32)() {
    m.datastoreId = value
}
// SetDescription sets the description property value. Description of the test
func (m *TestSummary) SetDescription(value *string)() {
    m.description = value
}
// SetFolder sets the folder property value. Name of folder that the test is stored in. Folders allow tests to be organised in the UI
func (m *TestSummary) SetFolder(value *string)() {
    m.folder = value
}
// SetId sets the id property value. ID of tests
func (m *TestSummary) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Test name
func (m *TestSummary) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *TestSummary) SetOwner(value *string)() {
    m.owner = value
}
// SetRuns sets the runs property value. Total number of Runs for the Test
func (m *TestSummary) SetRuns(value *float64)() {
    m.runs = value
}
// SetWatching sets the watching property value. Subscriptions for each test for authenticated user
func (m *TestSummary) SetWatching(value []string)() {
    m.watching = value
}
type TestSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*TestSummary_access)
    GetDatasets()(*float64)
    GetDatastoreId()(*int32)
    GetDescription()(*string)
    GetFolder()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetOwner()(*string)
    GetRuns()(*float64)
    GetWatching()([]string)
    SetAccess(value *TestSummary_access)()
    SetDatasets(value *float64)()
    SetDatastoreId(value *int32)()
    SetDescription(value *string)()
    SetFolder(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetOwner(value *string)()
    SetRuns(value *float64)()
    SetWatching(value []string)()
}
