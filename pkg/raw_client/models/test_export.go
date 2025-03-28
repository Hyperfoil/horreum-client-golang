package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TestExport represents a Test with all associated data used for export/import operations.
type TestExport struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *TestExport_access
    // Array of Actions associated with test
    actions []Actionable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // URL to external service that can be called to compare runs.  This is typically an external reporting/visulization service
    compareUrl *string
    // Datastore associated with test
    datastore TestExport_datastoreable
    // backend ID for backing datastore
    datastoreId *int32
    // Description of the test
    description *string
    // Array of ExperimentProfiles associated with test
    experiments []ExperimentProfileable
    // Filter function to filter out datasets that are comparable for the purpose of change detection
    fingerprintFilter *string
    // Array of Label names that are used to create a fingerprint 
    fingerprintLabels []string
    // Name of folder that the test is stored in. Folders allow tests to be organised in the UI
    folder *string
    // Unique Test id
    id *int32
    // Array of MissingDataRules associated with test
    missingDataRules []MissingDataRuleable
    // Test name
    name *string
    // Are notifications enabled for the test
    notificationsEnabled *bool
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // Watcher object associated with test
    subscriptions TestExport_subscriptionsable
    // Label function to modify timeline labels to a produce a value used for ordering datapoints
    timelineFunction *string
    // List of label names that are used for determining metric to use as the time series
    timelineLabels []string
    // Array for transformers defined for the Test
    transformers []Transformerable
    // Array of Variables associated with test
    variables []Variableable
}
// NewTestExport instantiates a new TestExport and sets the default values.
func NewTestExport()(*TestExport) {
    m := &TestExport{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTestExportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTestExportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTestExport(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *TestExport_access when successful
func (m *TestExport) GetAccess()(*TestExport_access) {
    return m.access
}
// GetActions gets the actions property value. Array of Actions associated with test
// returns a []Actionable when successful
func (m *TestExport) GetActions()([]Actionable) {
    return m.actions
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TestExport) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompareUrl gets the compareUrl property value. URL to external service that can be called to compare runs.  This is typically an external reporting/visulization service
// returns a *string when successful
func (m *TestExport) GetCompareUrl()(*string) {
    return m.compareUrl
}
// GetDatastore gets the datastore property value. Datastore associated with test
// returns a TestExport_datastoreable when successful
func (m *TestExport) GetDatastore()(TestExport_datastoreable) {
    return m.datastore
}
// GetDatastoreId gets the datastoreId property value. backend ID for backing datastore
// returns a *int32 when successful
func (m *TestExport) GetDatastoreId()(*int32) {
    return m.datastoreId
}
// GetDescription gets the description property value. Description of the test
// returns a *string when successful
func (m *TestExport) GetDescription()(*string) {
    return m.description
}
// GetExperiments gets the experiments property value. Array of ExperimentProfiles associated with test
// returns a []ExperimentProfileable when successful
func (m *TestExport) GetExperiments()([]ExperimentProfileable) {
    return m.experiments
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TestExport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTestExport_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*TestExport_access))
        }
        return nil
    }
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Actionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Actionable)
                }
            }
            m.SetActions(res)
        }
        return nil
    }
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
    res["datastore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTestExport_datastoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatastore(val.(TestExport_datastoreable))
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
    res["experiments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExperimentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExperimentProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExperimentProfileable)
                }
            }
            m.SetExperiments(res)
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
    res["missingDataRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMissingDataRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MissingDataRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MissingDataRuleable)
                }
            }
            m.SetMissingDataRules(res)
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
    res["subscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTestExport_subscriptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptions(val.(TestExport_subscriptionsable))
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
    res["variables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVariableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Variableable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Variableable)
                }
            }
            m.SetVariables(res)
        }
        return nil
    }
    return res
}
// GetFingerprintFilter gets the fingerprintFilter property value. Filter function to filter out datasets that are comparable for the purpose of change detection
// returns a *string when successful
func (m *TestExport) GetFingerprintFilter()(*string) {
    return m.fingerprintFilter
}
// GetFingerprintLabels gets the fingerprintLabels property value. Array of Label names that are used to create a fingerprint 
// returns a []string when successful
func (m *TestExport) GetFingerprintLabels()([]string) {
    return m.fingerprintLabels
}
// GetFolder gets the folder property value. Name of folder that the test is stored in. Folders allow tests to be organised in the UI
// returns a *string when successful
func (m *TestExport) GetFolder()(*string) {
    return m.folder
}
// GetId gets the id property value. Unique Test id
// returns a *int32 when successful
func (m *TestExport) GetId()(*int32) {
    return m.id
}
// GetMissingDataRules gets the missingDataRules property value. Array of MissingDataRules associated with test
// returns a []MissingDataRuleable when successful
func (m *TestExport) GetMissingDataRules()([]MissingDataRuleable) {
    return m.missingDataRules
}
// GetName gets the name property value. Test name
// returns a *string when successful
func (m *TestExport) GetName()(*string) {
    return m.name
}
// GetNotificationsEnabled gets the notificationsEnabled property value. Are notifications enabled for the test
// returns a *bool when successful
func (m *TestExport) GetNotificationsEnabled()(*bool) {
    return m.notificationsEnabled
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *TestExport) GetOwner()(*string) {
    return m.owner
}
// GetSubscriptions gets the subscriptions property value. Watcher object associated with test
// returns a TestExport_subscriptionsable when successful
func (m *TestExport) GetSubscriptions()(TestExport_subscriptionsable) {
    return m.subscriptions
}
// GetTimelineFunction gets the timelineFunction property value. Label function to modify timeline labels to a produce a value used for ordering datapoints
// returns a *string when successful
func (m *TestExport) GetTimelineFunction()(*string) {
    return m.timelineFunction
}
// GetTimelineLabels gets the timelineLabels property value. List of label names that are used for determining metric to use as the time series
// returns a []string when successful
func (m *TestExport) GetTimelineLabels()([]string) {
    return m.timelineLabels
}
// GetTransformers gets the transformers property value. Array for transformers defined for the Test
// returns a []Transformerable when successful
func (m *TestExport) GetTransformers()([]Transformerable) {
    return m.transformers
}
// GetVariables gets the variables property value. Array of Variables associated with test
// returns a []Variableable when successful
func (m *TestExport) GetVariables()([]Variableable) {
    return m.variables
}
// Serialize serializes information the current object
func (m *TestExport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccess() != nil {
        cast := (*m.GetAccess()).String()
        err := writer.WriteStringValue("access", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActions()))
        for i, v := range m.GetActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("actions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("compareUrl", m.GetCompareUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("datastore", m.GetDatastore())
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
    if m.GetExperiments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExperiments()))
        for i, v := range m.GetExperiments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("experiments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fingerprintFilter", m.GetFingerprintFilter())
        if err != nil {
            return err
        }
    }
    if m.GetFingerprintLabels() != nil {
        err := writer.WriteCollectionOfStringValues("fingerprintLabels", m.GetFingerprintLabels())
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
    if m.GetMissingDataRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMissingDataRules()))
        for i, v := range m.GetMissingDataRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("missingDataRules", cast)
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
        err := writer.WriteBoolValue("notificationsEnabled", m.GetNotificationsEnabled())
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
        err := writer.WriteObjectValue("subscriptions", m.GetSubscriptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timelineFunction", m.GetTimelineFunction())
        if err != nil {
            return err
        }
    }
    if m.GetTimelineLabels() != nil {
        err := writer.WriteCollectionOfStringValues("timelineLabels", m.GetTimelineLabels())
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
    if m.GetVariables() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVariables()))
        for i, v := range m.GetVariables() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("variables", cast)
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
func (m *TestExport) SetAccess(value *TestExport_access)() {
    m.access = value
}
// SetActions sets the actions property value. Array of Actions associated with test
func (m *TestExport) SetActions(value []Actionable)() {
    m.actions = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TestExport) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompareUrl sets the compareUrl property value. URL to external service that can be called to compare runs.  This is typically an external reporting/visulization service
func (m *TestExport) SetCompareUrl(value *string)() {
    m.compareUrl = value
}
// SetDatastore sets the datastore property value. Datastore associated with test
func (m *TestExport) SetDatastore(value TestExport_datastoreable)() {
    m.datastore = value
}
// SetDatastoreId sets the datastoreId property value. backend ID for backing datastore
func (m *TestExport) SetDatastoreId(value *int32)() {
    m.datastoreId = value
}
// SetDescription sets the description property value. Description of the test
func (m *TestExport) SetDescription(value *string)() {
    m.description = value
}
// SetExperiments sets the experiments property value. Array of ExperimentProfiles associated with test
func (m *TestExport) SetExperiments(value []ExperimentProfileable)() {
    m.experiments = value
}
// SetFingerprintFilter sets the fingerprintFilter property value. Filter function to filter out datasets that are comparable for the purpose of change detection
func (m *TestExport) SetFingerprintFilter(value *string)() {
    m.fingerprintFilter = value
}
// SetFingerprintLabels sets the fingerprintLabels property value. Array of Label names that are used to create a fingerprint 
func (m *TestExport) SetFingerprintLabels(value []string)() {
    m.fingerprintLabels = value
}
// SetFolder sets the folder property value. Name of folder that the test is stored in. Folders allow tests to be organised in the UI
func (m *TestExport) SetFolder(value *string)() {
    m.folder = value
}
// SetId sets the id property value. Unique Test id
func (m *TestExport) SetId(value *int32)() {
    m.id = value
}
// SetMissingDataRules sets the missingDataRules property value. Array of MissingDataRules associated with test
func (m *TestExport) SetMissingDataRules(value []MissingDataRuleable)() {
    m.missingDataRules = value
}
// SetName sets the name property value. Test name
func (m *TestExport) SetName(value *string)() {
    m.name = value
}
// SetNotificationsEnabled sets the notificationsEnabled property value. Are notifications enabled for the test
func (m *TestExport) SetNotificationsEnabled(value *bool)() {
    m.notificationsEnabled = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *TestExport) SetOwner(value *string)() {
    m.owner = value
}
// SetSubscriptions sets the subscriptions property value. Watcher object associated with test
func (m *TestExport) SetSubscriptions(value TestExport_subscriptionsable)() {
    m.subscriptions = value
}
// SetTimelineFunction sets the timelineFunction property value. Label function to modify timeline labels to a produce a value used for ordering datapoints
func (m *TestExport) SetTimelineFunction(value *string)() {
    m.timelineFunction = value
}
// SetTimelineLabels sets the timelineLabels property value. List of label names that are used for determining metric to use as the time series
func (m *TestExport) SetTimelineLabels(value []string)() {
    m.timelineLabels = value
}
// SetTransformers sets the transformers property value. Array for transformers defined for the Test
func (m *TestExport) SetTransformers(value []Transformerable)() {
    m.transformers = value
}
// SetVariables sets the variables property value. Array of Variables associated with test
func (m *TestExport) SetVariables(value []Variableable)() {
    m.variables = value
}
type TestExportable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*TestExport_access)
    GetActions()([]Actionable)
    GetCompareUrl()(*string)
    GetDatastore()(TestExport_datastoreable)
    GetDatastoreId()(*int32)
    GetDescription()(*string)
    GetExperiments()([]ExperimentProfileable)
    GetFingerprintFilter()(*string)
    GetFingerprintLabels()([]string)
    GetFolder()(*string)
    GetId()(*int32)
    GetMissingDataRules()([]MissingDataRuleable)
    GetName()(*string)
    GetNotificationsEnabled()(*bool)
    GetOwner()(*string)
    GetSubscriptions()(TestExport_subscriptionsable)
    GetTimelineFunction()(*string)
    GetTimelineLabels()([]string)
    GetTransformers()([]Transformerable)
    GetVariables()([]Variableable)
    SetAccess(value *TestExport_access)()
    SetActions(value []Actionable)()
    SetCompareUrl(value *string)()
    SetDatastore(value TestExport_datastoreable)()
    SetDatastoreId(value *int32)()
    SetDescription(value *string)()
    SetExperiments(value []ExperimentProfileable)()
    SetFingerprintFilter(value *string)()
    SetFingerprintLabels(value []string)()
    SetFolder(value *string)()
    SetId(value *int32)()
    SetMissingDataRules(value []MissingDataRuleable)()
    SetName(value *string)()
    SetNotificationsEnabled(value *bool)()
    SetOwner(value *string)()
    SetSubscriptions(value TestExport_subscriptionsable)()
    SetTimelineFunction(value *string)()
    SetTimelineLabels(value []string)()
    SetTransformers(value []Transformerable)()
    SetVariables(value []Variableable)()
}
