package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExperimentComparison struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Model JSON configuration
    config *string
    // Name of comparison model
    model *string
    // Variable ID to run experiment against
    variableId *int32
    // Variable Name to run experiment against
    variableName *string
}
// NewExperimentComparison instantiates a new ExperimentComparison and sets the default values.
func NewExperimentComparison()(*ExperimentComparison) {
    m := &ExperimentComparison{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExperimentComparisonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExperimentComparisonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExperimentComparison(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExperimentComparison) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. Model JSON configuration
// returns a *string when successful
func (m *ExperimentComparison) GetConfig()(*string) {
    return m.config
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExperimentComparison) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["variableId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVariableId(val)
        }
        return nil
    }
    res["variableName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVariableName(val)
        }
        return nil
    }
    return res
}
// GetModel gets the model property value. Name of comparison model
// returns a *string when successful
func (m *ExperimentComparison) GetModel()(*string) {
    return m.model
}
// GetVariableId gets the variableId property value. Variable ID to run experiment against
// returns a *int32 when successful
func (m *ExperimentComparison) GetVariableId()(*int32) {
    return m.variableId
}
// GetVariableName gets the variableName property value. Variable Name to run experiment against
// returns a *string when successful
func (m *ExperimentComparison) GetVariableName()(*string) {
    return m.variableName
}
// Serialize serializes information the current object
func (m *ExperimentComparison) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("config", m.GetConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("variableId", m.GetVariableId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("variableName", m.GetVariableName())
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
func (m *ExperimentComparison) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. Model JSON configuration
func (m *ExperimentComparison) SetConfig(value *string)() {
    m.config = value
}
// SetModel sets the model property value. Name of comparison model
func (m *ExperimentComparison) SetModel(value *string)() {
    m.model = value
}
// SetVariableId sets the variableId property value. Variable ID to run experiment against
func (m *ExperimentComparison) SetVariableId(value *int32)() {
    m.variableId = value
}
// SetVariableName sets the variableName property value. Variable Name to run experiment against
func (m *ExperimentComparison) SetVariableName(value *string)() {
    m.variableName = value
}
type ExperimentComparisonable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfig()(*string)
    GetModel()(*string)
    GetVariableId()(*int32)
    GetVariableName()(*string)
    SetConfig(value *string)()
    SetModel(value *string)()
    SetVariableId(value *int32)()
    SetVariableName(value *string)()
}
