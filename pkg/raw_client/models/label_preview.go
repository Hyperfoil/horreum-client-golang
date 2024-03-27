package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelPreview preview a Label Value derived from a Dataset Data. A preview allows users to apply a Label to a dataset and preview the Label Value result and processing errors in the UI
type LabelPreview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Description of errors occurred attempting to generate Label Value Preview
    output *string
    // Value value extracted from Dataset. This can be a scalar, array or JSON object
    value *string
}
// NewLabelPreview instantiates a new LabelPreview and sets the default values.
func NewLabelPreview()(*LabelPreview) {
    m := &LabelPreview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelPreviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelPreviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelPreview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LabelPreview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelPreview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["output"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutput(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetOutput gets the output property value. Description of errors occurred attempting to generate Label Value Preview
// returns a *string when successful
func (m *LabelPreview) GetOutput()(*string) {
    return m.output
}
// GetValue gets the value property value. Value value extracted from Dataset. This can be a scalar, array or JSON object
// returns a *string when successful
func (m *LabelPreview) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *LabelPreview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("output", m.GetOutput())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *LabelPreview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOutput sets the output property value. Description of errors occurred attempting to generate Label Value Preview
func (m *LabelPreview) SetOutput(value *string)() {
    m.output = value
}
// SetValue sets the value property value. Value value extracted from Dataset. This can be a scalar, array or JSON object
func (m *LabelPreview) SetValue(value *string)() {
    m.value = value
}
type LabelPreviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOutput()(*string)
    GetValue()(*string)
    SetOutput(value *string)()
    SetValue(value *string)()
}
