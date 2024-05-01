package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RelativeDifferenceDetectionConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Built In
    builtIn *bool
    // Relative Difference Detection filter
    filter *string
    // Minimal number of preceding datapoints
    minPrevious *int32
    // The model property
    model *RelativeDifferenceDetectionConfig_model
    // Maximum difference between the aggregated value of last <window> datapoints and the mean of preceding values.
    threshold *float64
    // Number of most recent datapoints used for aggregating the value for comparison.
    window *int32
}
// NewRelativeDifferenceDetectionConfig instantiates a new RelativeDifferenceDetectionConfig and sets the default values.
func NewRelativeDifferenceDetectionConfig()(*RelativeDifferenceDetectionConfig) {
    m := &RelativeDifferenceDetectionConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRelativeDifferenceDetectionConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelativeDifferenceDetectionConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelativeDifferenceDetectionConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RelativeDifferenceDetectionConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuiltIn gets the builtIn property value. Built In
// returns a *bool when successful
func (m *RelativeDifferenceDetectionConfig) GetBuiltIn()(*bool) {
    return m.builtIn
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RelativeDifferenceDetectionConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["builtIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuiltIn(val)
        }
        return nil
    }
    res["filter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["minPrevious"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinPrevious(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRelativeDifferenceDetectionConfig_model)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val.(*RelativeDifferenceDetectionConfig_model))
        }
        return nil
    }
    res["threshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreshold(val)
        }
        return nil
    }
    res["window"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindow(val)
        }
        return nil
    }
    return res
}
// GetFilter gets the filter property value. Relative Difference Detection filter
// returns a *string when successful
func (m *RelativeDifferenceDetectionConfig) GetFilter()(*string) {
    return m.filter
}
// GetMinPrevious gets the minPrevious property value. Minimal number of preceding datapoints
// returns a *int32 when successful
func (m *RelativeDifferenceDetectionConfig) GetMinPrevious()(*int32) {
    return m.minPrevious
}
// GetModel gets the model property value. The model property
// returns a *RelativeDifferenceDetectionConfig_model when successful
func (m *RelativeDifferenceDetectionConfig) GetModel()(*RelativeDifferenceDetectionConfig_model) {
    return m.model
}
// GetThreshold gets the threshold property value. Maximum difference between the aggregated value of last <window> datapoints and the mean of preceding values.
// returns a *float64 when successful
func (m *RelativeDifferenceDetectionConfig) GetThreshold()(*float64) {
    return m.threshold
}
// GetWindow gets the window property value. Number of most recent datapoints used for aggregating the value for comparison.
// returns a *int32 when successful
func (m *RelativeDifferenceDetectionConfig) GetWindow()(*int32) {
    return m.window
}
// Serialize serializes information the current object
func (m *RelativeDifferenceDetectionConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("builtIn", m.GetBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minPrevious", m.GetMinPrevious())
        if err != nil {
            return err
        }
    }
    if m.GetModel() != nil {
        cast := (*m.GetModel()).String()
        err := writer.WriteStringValue("model", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("threshold", m.GetThreshold())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("window", m.GetWindow())
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
func (m *RelativeDifferenceDetectionConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuiltIn sets the builtIn property value. Built In
func (m *RelativeDifferenceDetectionConfig) SetBuiltIn(value *bool)() {
    m.builtIn = value
}
// SetFilter sets the filter property value. Relative Difference Detection filter
func (m *RelativeDifferenceDetectionConfig) SetFilter(value *string)() {
    m.filter = value
}
// SetMinPrevious sets the minPrevious property value. Minimal number of preceding datapoints
func (m *RelativeDifferenceDetectionConfig) SetMinPrevious(value *int32)() {
    m.minPrevious = value
}
// SetModel sets the model property value. The model property
func (m *RelativeDifferenceDetectionConfig) SetModel(value *RelativeDifferenceDetectionConfig_model)() {
    m.model = value
}
// SetThreshold sets the threshold property value. Maximum difference between the aggregated value of last <window> datapoints and the mean of preceding values.
func (m *RelativeDifferenceDetectionConfig) SetThreshold(value *float64)() {
    m.threshold = value
}
// SetWindow sets the window property value. Number of most recent datapoints used for aggregating the value for comparison.
func (m *RelativeDifferenceDetectionConfig) SetWindow(value *int32)() {
    m.window = value
}
type RelativeDifferenceDetectionConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuiltIn()(*bool)
    GetFilter()(*string)
    GetMinPrevious()(*int32)
    GetModel()(*RelativeDifferenceDetectionConfig_model)
    GetThreshold()(*float64)
    GetWindow()(*int32)
    SetBuiltIn(value *bool)()
    SetFilter(value *string)()
    SetMinPrevious(value *int32)()
    SetModel(value *RelativeDifferenceDetectionConfig_model)()
    SetThreshold(value *float64)()
    SetWindow(value *int32)()
}
