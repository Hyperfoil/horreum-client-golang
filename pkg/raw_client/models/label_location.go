package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LabelLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique ID for location that references Schema
    testId *int32
    // Test name that references Schema
    testName *string
    // Location of Label usage
    typeEscaped *string
}
// NewLabelLocation instantiates a new LabelLocation and sets the default values.
func NewLabelLocation()(*LabelLocation) {
    m := &LabelLocation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelLocation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LabelLocation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["testName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetTestId gets the testId property value. Unique ID for location that references Schema
// returns a *int32 when successful
func (m *LabelLocation) GetTestId()(*int32) {
    return m.testId
}
// GetTestName gets the testName property value. Test name that references Schema
// returns a *string when successful
func (m *LabelLocation) GetTestName()(*string) {
    return m.testName
}
// GetTypeEscaped gets the type property value. Location of Label usage
// returns a *string when successful
func (m *LabelLocation) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *LabelLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("testId", m.GetTestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("testName", m.GetTestName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *LabelLocation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTestId sets the testId property value. Unique ID for location that references Schema
func (m *LabelLocation) SetTestId(value *int32)() {
    m.testId = value
}
// SetTestName sets the testName property value. Test name that references Schema
func (m *LabelLocation) SetTestName(value *string)() {
    m.testName = value
}
// SetTypeEscaped sets the type property value. Location of Label usage
func (m *LabelLocation) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type LabelLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTestId()(*int32)
    GetTestName()(*string)
    GetTypeEscaped()(*string)
    SetTestId(value *int32)()
    SetTestName(value *string)()
    SetTypeEscaped(value *string)()
}
