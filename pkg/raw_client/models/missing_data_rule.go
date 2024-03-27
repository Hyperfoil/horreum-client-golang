package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MissingDataRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The condition property
    condition *string
    // The id property
    id *int32
    // The labels property
    labels []string
    // The lastNotification property
    lastNotification *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The maxStaleness property
    maxStaleness *int64
    // The name property
    name *string
    // The testId property
    testId *int32
}
// NewMissingDataRule instantiates a new MissingDataRule and sets the default values.
func NewMissingDataRule()(*MissingDataRule) {
    m := &MissingDataRule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMissingDataRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMissingDataRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMissingDataRule(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MissingDataRule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCondition gets the condition property value. The condition property
// returns a *string when successful
func (m *MissingDataRule) GetCondition()(*string) {
    return m.condition
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MissingDataRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["condition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCondition(val)
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
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetLabels(res)
        }
        return nil
    }
    res["lastNotification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastNotification(val)
        }
        return nil
    }
    res["maxStaleness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxStaleness(val)
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
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *MissingDataRule) GetId()(*int32) {
    return m.id
}
// GetLabels gets the labels property value. The labels property
// returns a []string when successful
func (m *MissingDataRule) GetLabels()([]string) {
    return m.labels
}
// GetLastNotification gets the lastNotification property value. The lastNotification property
// returns a *Time when successful
func (m *MissingDataRule) GetLastNotification()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastNotification
}
// GetMaxStaleness gets the maxStaleness property value. The maxStaleness property
// returns a *int64 when successful
func (m *MissingDataRule) GetMaxStaleness()(*int64) {
    return m.maxStaleness
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *MissingDataRule) GetName()(*string) {
    return m.name
}
// GetTestId gets the testId property value. The testId property
// returns a *int32 when successful
func (m *MissingDataRule) GetTestId()(*int32) {
    return m.testId
}
// Serialize serializes information the current object
func (m *MissingDataRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("condition", m.GetCondition())
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
    if m.GetLabels() != nil {
        err := writer.WriteCollectionOfStringValues("labels", m.GetLabels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastNotification", m.GetLastNotification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("maxStaleness", m.GetMaxStaleness())
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
        err := writer.WriteInt32Value("testId", m.GetTestId())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MissingDataRule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCondition sets the condition property value. The condition property
func (m *MissingDataRule) SetCondition(value *string)() {
    m.condition = value
}
// SetId sets the id property value. The id property
func (m *MissingDataRule) SetId(value *int32)() {
    m.id = value
}
// SetLabels sets the labels property value. The labels property
func (m *MissingDataRule) SetLabels(value []string)() {
    m.labels = value
}
// SetLastNotification sets the lastNotification property value. The lastNotification property
func (m *MissingDataRule) SetLastNotification(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastNotification = value
}
// SetMaxStaleness sets the maxStaleness property value. The maxStaleness property
func (m *MissingDataRule) SetMaxStaleness(value *int64)() {
    m.maxStaleness = value
}
// SetName sets the name property value. The name property
func (m *MissingDataRule) SetName(value *string)() {
    m.name = value
}
// SetTestId sets the testId property value. The testId property
func (m *MissingDataRule) SetTestId(value *int32)() {
    m.testId = value
}
type MissingDataRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCondition()(*string)
    GetId()(*int32)
    GetLabels()([]string)
    GetLastNotification()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMaxStaleness()(*int64)
    GetName()(*string)
    GetTestId()(*int32)
    SetCondition(value *string)()
    SetId(value *int32)()
    SetLabels(value []string)()
    SetLastNotification(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMaxStaleness(value *int64)()
    SetName(value *string)()
    SetTestId(value *int32)()
}
