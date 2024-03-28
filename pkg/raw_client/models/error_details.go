package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ErrorDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The arguments property
    arguments []string
    // The code property
    code *string
    // The details property
    details *string
    // The evaluationPath property
    evaluationPath *string
    // The instanceLocation property
    instanceLocation *string
    // The message property
    message *string
    // The messageKey property
    messageKey *string
    // The path property
    path *string
    // The property property
    property *string
    // The schemaLocation property
    schemaLocation *string
    // The schemaPath property
    // Deprecated: 
    schemaPath *string
    // Validation Error type
    typeEscaped *string
    // The valid property
    valid *bool
}
// NewErrorDetails instantiates a new ErrorDetails and sets the default values.
func NewErrorDetails()(*ErrorDetails) {
    m := &ErrorDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewErrorDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ErrorDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArguments gets the arguments property value. The arguments property
// returns a []string when successful
func (m *ErrorDetails) GetArguments()([]string) {
    return m.arguments
}
// GetCode gets the code property value. The code property
// returns a *string when successful
func (m *ErrorDetails) GetCode()(*string) {
    return m.code
}
// GetDetails gets the details property value. The details property
// returns a *string when successful
func (m *ErrorDetails) GetDetails()(*string) {
    return m.details
}
// GetEvaluationPath gets the evaluationPath property value. The evaluationPath property
// returns a *string when successful
func (m *ErrorDetails) GetEvaluationPath()(*string) {
    return m.evaluationPath
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ErrorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["arguments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetArguments(res)
        }
        return nil
    }
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val)
        }
        return nil
    }
    res["evaluationPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationPath(val)
        }
        return nil
    }
    res["instanceLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceLocation(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["messageKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageKey(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["property"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperty(val)
        }
        return nil
    }
    res["schemaLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaLocation(val)
        }
        return nil
    }
    res["schemaPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaPath(val)
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
    res["valid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValid(val)
        }
        return nil
    }
    return res
}
// GetInstanceLocation gets the instanceLocation property value. The instanceLocation property
// returns a *string when successful
func (m *ErrorDetails) GetInstanceLocation()(*string) {
    return m.instanceLocation
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *ErrorDetails) GetMessage()(*string) {
    return m.message
}
// GetMessageKey gets the messageKey property value. The messageKey property
// returns a *string when successful
func (m *ErrorDetails) GetMessageKey()(*string) {
    return m.messageKey
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *ErrorDetails) GetPath()(*string) {
    return m.path
}
// GetProperty gets the property property value. The property property
// returns a *string when successful
func (m *ErrorDetails) GetProperty()(*string) {
    return m.property
}
// GetSchemaLocation gets the schemaLocation property value. The schemaLocation property
// returns a *string when successful
func (m *ErrorDetails) GetSchemaLocation()(*string) {
    return m.schemaLocation
}
// GetSchemaPath gets the schemaPath property value. The schemaPath property
// Deprecated: 
// returns a *string when successful
func (m *ErrorDetails) GetSchemaPath()(*string) {
    return m.schemaPath
}
// GetTypeEscaped gets the type property value. Validation Error type
// returns a *string when successful
func (m *ErrorDetails) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetValid gets the valid property value. The valid property
// returns a *bool when successful
func (m *ErrorDetails) GetValid()(*bool) {
    return m.valid
}
// Serialize serializes information the current object
func (m *ErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetArguments() != nil {
        err := writer.WriteCollectionOfStringValues("arguments", m.GetArguments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("evaluationPath", m.GetEvaluationPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("instanceLocation", m.GetInstanceLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("messageKey", m.GetMessageKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("property", m.GetProperty())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schemaLocation", m.GetSchemaLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schemaPath", m.GetSchemaPath())
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
        err := writer.WriteBoolValue("valid", m.GetValid())
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
func (m *ErrorDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArguments sets the arguments property value. The arguments property
func (m *ErrorDetails) SetArguments(value []string)() {
    m.arguments = value
}
// SetCode sets the code property value. The code property
func (m *ErrorDetails) SetCode(value *string)() {
    m.code = value
}
// SetDetails sets the details property value. The details property
func (m *ErrorDetails) SetDetails(value *string)() {
    m.details = value
}
// SetEvaluationPath sets the evaluationPath property value. The evaluationPath property
func (m *ErrorDetails) SetEvaluationPath(value *string)() {
    m.evaluationPath = value
}
// SetInstanceLocation sets the instanceLocation property value. The instanceLocation property
func (m *ErrorDetails) SetInstanceLocation(value *string)() {
    m.instanceLocation = value
}
// SetMessage sets the message property value. The message property
func (m *ErrorDetails) SetMessage(value *string)() {
    m.message = value
}
// SetMessageKey sets the messageKey property value. The messageKey property
func (m *ErrorDetails) SetMessageKey(value *string)() {
    m.messageKey = value
}
// SetPath sets the path property value. The path property
func (m *ErrorDetails) SetPath(value *string)() {
    m.path = value
}
// SetProperty sets the property property value. The property property
func (m *ErrorDetails) SetProperty(value *string)() {
    m.property = value
}
// SetSchemaLocation sets the schemaLocation property value. The schemaLocation property
func (m *ErrorDetails) SetSchemaLocation(value *string)() {
    m.schemaLocation = value
}
// SetSchemaPath sets the schemaPath property value. The schemaPath property
// Deprecated: 
func (m *ErrorDetails) SetSchemaPath(value *string)() {
    m.schemaPath = value
}
// SetTypeEscaped sets the type property value. Validation Error type
func (m *ErrorDetails) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetValid sets the valid property value. The valid property
func (m *ErrorDetails) SetValid(value *bool)() {
    m.valid = value
}
type ErrorDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArguments()([]string)
    GetCode()(*string)
    GetDetails()(*string)
    GetEvaluationPath()(*string)
    GetInstanceLocation()(*string)
    GetMessage()(*string)
    GetMessageKey()(*string)
    GetPath()(*string)
    GetProperty()(*string)
    GetSchemaLocation()(*string)
    GetSchemaPath()(*string)
    GetTypeEscaped()(*string)
    GetValid()(*bool)
    SetArguments(value []string)()
    SetCode(value *string)()
    SetDetails(value *string)()
    SetEvaluationPath(value *string)()
    SetInstanceLocation(value *string)()
    SetMessage(value *string)()
    SetMessageKey(value *string)()
    SetPath(value *string)()
    SetProperty(value *string)()
    SetSchemaLocation(value *string)()
    SetSchemaPath(value *string)()
    SetTypeEscaped(value *string)()
    SetValid(value *bool)()
}
