package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Variable struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The calculation property
    calculation *string
    // The changeDetection property
    changeDetection []ChangeDetectionable
    // The group property
    group *string
    // The id property
    id *int32
    // The labels property
    labels []string
    // The name property
    name *string
    // The order property
    order *int32
    // The testId property
    testId *int32
}
// NewVariable instantiates a new Variable and sets the default values.
func NewVariable()(*Variable) {
    m := &Variable{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVariableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVariableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVariable(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Variable) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCalculation gets the calculation property value. The calculation property
// returns a *string when successful
func (m *Variable) GetCalculation()(*string) {
    return m.calculation
}
// GetChangeDetection gets the changeDetection property value. The changeDetection property
// returns a []ChangeDetectionable when successful
func (m *Variable) GetChangeDetection()([]ChangeDetectionable) {
    return m.changeDetection
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Variable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["calculation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalculation(val)
        }
        return nil
    }
    res["changeDetection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChangeDetectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChangeDetectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ChangeDetectionable)
                }
            }
            m.SetChangeDetection(res)
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val)
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
    res["order"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrder(val)
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
// GetGroup gets the group property value. The group property
// returns a *string when successful
func (m *Variable) GetGroup()(*string) {
    return m.group
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Variable) GetId()(*int32) {
    return m.id
}
// GetLabels gets the labels property value. The labels property
// returns a []string when successful
func (m *Variable) GetLabels()([]string) {
    return m.labels
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Variable) GetName()(*string) {
    return m.name
}
// GetOrder gets the order property value. The order property
// returns a *int32 when successful
func (m *Variable) GetOrder()(*int32) {
    return m.order
}
// GetTestId gets the testId property value. The testId property
// returns a *int32 when successful
func (m *Variable) GetTestId()(*int32) {
    return m.testId
}
// Serialize serializes information the current object
func (m *Variable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("calculation", m.GetCalculation())
        if err != nil {
            return err
        }
    }
    if m.GetChangeDetection() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChangeDetection()))
        for i, v := range m.GetChangeDetection() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("changeDetection", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("group", m.GetGroup())
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
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("order", m.GetOrder())
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
func (m *Variable) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCalculation sets the calculation property value. The calculation property
func (m *Variable) SetCalculation(value *string)() {
    m.calculation = value
}
// SetChangeDetection sets the changeDetection property value. The changeDetection property
func (m *Variable) SetChangeDetection(value []ChangeDetectionable)() {
    m.changeDetection = value
}
// SetGroup sets the group property value. The group property
func (m *Variable) SetGroup(value *string)() {
    m.group = value
}
// SetId sets the id property value. The id property
func (m *Variable) SetId(value *int32)() {
    m.id = value
}
// SetLabels sets the labels property value. The labels property
func (m *Variable) SetLabels(value []string)() {
    m.labels = value
}
// SetName sets the name property value. The name property
func (m *Variable) SetName(value *string)() {
    m.name = value
}
// SetOrder sets the order property value. The order property
func (m *Variable) SetOrder(value *int32)() {
    m.order = value
}
// SetTestId sets the testId property value. The testId property
func (m *Variable) SetTestId(value *int32)() {
    m.testId = value
}
type Variableable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCalculation()(*string)
    GetChangeDetection()([]ChangeDetectionable)
    GetGroup()(*string)
    GetId()(*int32)
    GetLabels()([]string)
    GetName()(*string)
    GetOrder()(*int32)
    GetTestId()(*int32)
    SetCalculation(value *string)()
    SetChangeDetection(value []ChangeDetectionable)()
    SetGroup(value *string)()
    SetId(value *int32)()
    SetLabels(value []string)()
    SetName(value *string)()
    SetOrder(value *int32)()
    SetTestId(value *int32)()
}
