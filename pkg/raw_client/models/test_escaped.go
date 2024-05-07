package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Test represents a Test. Tests are typically equivalent to a particular benchmark
type Test struct {
    ProtectedType
    // URL to external service that can be called to compare runs.  This is typically an external reporting/visulization service
    compareUrl *string
    // backend ID for backing datastore
    datastoreId *int32
    // Description of the test
    description *string
    // Filter function to filter out datasets that are comparable for the purpose of change detection
    fingerprintFilter *string
    // Array of Label names that are used to create a fingerprint 
    fingerprintLabels []string
    // Name of folder that the test is stored in. Folders allow tests to be organised in the UI
    folder *string
    // Unique Test id
    id *int32
    // Test name
    name *string
    // Are notifications enabled for the test
    notificationsEnabled *bool
    // Label function to modify timeline labels to a produce a value used for ordering datapoints
    timelineFunction *string
    // List of label names that are used for determining metric to use as the time series
    timelineLabels []string
    // Array of API tokens associated with test
    tokens []TestTokenable
    // Array for transformers defined for the Test
    transformers []Transformerable
}
// NewTest instantiates a new Test and sets the default values.
func NewTest()(*Test) {
    m := &Test{
        ProtectedType: *NewProtectedType(),
    }
    return m
}
// CreateTestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTest(), nil
}
// GetCompareUrl gets the compareUrl property value. URL to external service that can be called to compare runs.  This is typically an external reporting/visulization service
// returns a *string when successful
func (m *Test) GetCompareUrl()(*string) {
    return m.compareUrl
}
// GetDatastoreId gets the datastoreId property value. backend ID for backing datastore
// returns a *int32 when successful
func (m *Test) GetDatastoreId()(*int32) {
    return m.datastoreId
}
// GetDescription gets the description property value. Description of the test
// returns a *string when successful
func (m *Test) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Test) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectedType.GetFieldDeserializers()
    res["compareUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompareUrl(val)
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
    res["fingerprintFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprintFilter(val)
        }
        return nil
    }
    res["fingerprintLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetFingerprintLabels(res)
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
    res["notificationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationsEnabled(val)
        }
        return nil
    }
    res["timelineFunction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimelineFunction(val)
        }
        return nil
    }
    res["timelineLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTimelineLabels(res)
        }
        return nil
    }
    res["tokens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTestTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TestTokenable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TestTokenable)
                }
            }
            m.SetTokens(res)
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
    return res
}
// GetFingerprintFilter gets the fingerprintFilter property value. Filter function to filter out datasets that are comparable for the purpose of change detection
// returns a *string when successful
func (m *Test) GetFingerprintFilter()(*string) {
    return m.fingerprintFilter
}
// GetFingerprintLabels gets the fingerprintLabels property value. Array of Label names that are used to create a fingerprint 
// returns a []string when successful
func (m *Test) GetFingerprintLabels()([]string) {
    return m.fingerprintLabels
}
// GetFolder gets the folder property value. Name of folder that the test is stored in. Folders allow tests to be organised in the UI
// returns a *string when successful
func (m *Test) GetFolder()(*string) {
    return m.folder
}
// GetId gets the id property value. Unique Test id
// returns a *int32 when successful
func (m *Test) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Test name
// returns a *string when successful
func (m *Test) GetName()(*string) {
    return m.name
}
// GetNotificationsEnabled gets the notificationsEnabled property value. Are notifications enabled for the test
// returns a *bool when successful
func (m *Test) GetNotificationsEnabled()(*bool) {
    return m.notificationsEnabled
}
// GetTimelineFunction gets the timelineFunction property value. Label function to modify timeline labels to a produce a value used for ordering datapoints
// returns a *string when successful
func (m *Test) GetTimelineFunction()(*string) {
    return m.timelineFunction
}
// GetTimelineLabels gets the timelineLabels property value. List of label names that are used for determining metric to use as the time series
// returns a []string when successful
func (m *Test) GetTimelineLabels()([]string) {
    return m.timelineLabels
}
// GetTokens gets the tokens property value. Array of API tokens associated with test
// returns a []TestTokenable when successful
func (m *Test) GetTokens()([]TestTokenable) {
    return m.tokens
}
// GetTransformers gets the transformers property value. Array for transformers defined for the Test
// returns a []Transformerable when successful
func (m *Test) GetTransformers()([]Transformerable) {
    return m.transformers
}
// Serialize serializes information the current object
func (m *Test) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectedType.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("compareUrl", m.GetCompareUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("datastoreId", m.GetDatastoreId())
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
        err = writer.WriteStringValue("fingerprintFilter", m.GetFingerprintFilter())
        if err != nil {
            return err
        }
    }
    if m.GetFingerprintLabels() != nil {
        err = writer.WriteCollectionOfStringValues("fingerprintLabels", m.GetFingerprintLabels())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("folder", m.GetFolder())
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("notificationsEnabled", m.GetNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timelineFunction", m.GetTimelineFunction())
        if err != nil {
            return err
        }
    }
    if m.GetTimelineLabels() != nil {
        err = writer.WriteCollectionOfStringValues("timelineLabels", m.GetTimelineLabels())
        if err != nil {
            return err
        }
    }
    if m.GetTokens() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokens()))
        for i, v := range m.GetTokens() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tokens", cast)
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
        err = writer.WriteCollectionOfObjectValues("transformers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompareUrl sets the compareUrl property value. URL to external service that can be called to compare runs.  This is typically an external reporting/visulization service
func (m *Test) SetCompareUrl(value *string)() {
    m.compareUrl = value
}
// SetDatastoreId sets the datastoreId property value. backend ID for backing datastore
func (m *Test) SetDatastoreId(value *int32)() {
    m.datastoreId = value
}
// SetDescription sets the description property value. Description of the test
func (m *Test) SetDescription(value *string)() {
    m.description = value
}
// SetFingerprintFilter sets the fingerprintFilter property value. Filter function to filter out datasets that are comparable for the purpose of change detection
func (m *Test) SetFingerprintFilter(value *string)() {
    m.fingerprintFilter = value
}
// SetFingerprintLabels sets the fingerprintLabels property value. Array of Label names that are used to create a fingerprint 
func (m *Test) SetFingerprintLabels(value []string)() {
    m.fingerprintLabels = value
}
// SetFolder sets the folder property value. Name of folder that the test is stored in. Folders allow tests to be organised in the UI
func (m *Test) SetFolder(value *string)() {
    m.folder = value
}
// SetId sets the id property value. Unique Test id
func (m *Test) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Test name
func (m *Test) SetName(value *string)() {
    m.name = value
}
// SetNotificationsEnabled sets the notificationsEnabled property value. Are notifications enabled for the test
func (m *Test) SetNotificationsEnabled(value *bool)() {
    m.notificationsEnabled = value
}
// SetTimelineFunction sets the timelineFunction property value. Label function to modify timeline labels to a produce a value used for ordering datapoints
func (m *Test) SetTimelineFunction(value *string)() {
    m.timelineFunction = value
}
// SetTimelineLabels sets the timelineLabels property value. List of label names that are used for determining metric to use as the time series
func (m *Test) SetTimelineLabels(value []string)() {
    m.timelineLabels = value
}
// SetTokens sets the tokens property value. Array of API tokens associated with test
func (m *Test) SetTokens(value []TestTokenable)() {
    m.tokens = value
}
// SetTransformers sets the transformers property value. Array for transformers defined for the Test
func (m *Test) SetTransformers(value []Transformerable)() {
    m.transformers = value
}
type Testable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectedTypeable
    GetCompareUrl()(*string)
    GetDatastoreId()(*int32)
    GetDescription()(*string)
    GetFingerprintFilter()(*string)
    GetFingerprintLabels()([]string)
    GetFolder()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNotificationsEnabled()(*bool)
    GetTimelineFunction()(*string)
    GetTimelineLabels()([]string)
    GetTokens()([]TestTokenable)
    GetTransformers()([]Transformerable)
    SetCompareUrl(value *string)()
    SetDatastoreId(value *int32)()
    SetDescription(value *string)()
    SetFingerprintFilter(value *string)()
    SetFingerprintLabels(value []string)()
    SetFolder(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNotificationsEnabled(value *bool)()
    SetTimelineFunction(value *string)()
    SetTimelineLabels(value []string)()
    SetTokens(value []TestTokenable)()
    SetTransformers(value []Transformerable)()
}
