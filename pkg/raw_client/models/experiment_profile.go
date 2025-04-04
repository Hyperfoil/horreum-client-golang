package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExperimentProfile experiment profile that results relates to
type ExperimentProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Selector filter to apply to Baseline label values
    baselineFilter *string
    // Array of selector labels for comparison Baseline
    baselineLabels []string
    // Collection of Experiment Comparisons to run during an Experiment evaluation
    comparisons []ExperimentComparisonable
    // These labels are not used by Horreum but are added to the result event and therefore can be used e.g. when firing an Action.
    extraLabels []string
    // Experiment Profile unique ID
    id *int32
    // Name of Experiment Profile
    name *string
    // Selector filter to apply to Selector label values
    selectorFilter *string
    // Array of selector labels
    selectorLabels []string
    // Test ID that Experiment Profile relates to
    testId *int32
}
// NewExperimentProfile instantiates a new ExperimentProfile and sets the default values.
func NewExperimentProfile()(*ExperimentProfile) {
    m := &ExperimentProfile{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExperimentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExperimentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExperimentProfile(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExperimentProfile) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBaselineFilter gets the baselineFilter property value. Selector filter to apply to Baseline label values
// returns a *string when successful
func (m *ExperimentProfile) GetBaselineFilter()(*string) {
    return m.baselineFilter
}
// GetBaselineLabels gets the baselineLabels property value. Array of selector labels for comparison Baseline
// returns a []string when successful
func (m *ExperimentProfile) GetBaselineLabels()([]string) {
    return m.baselineLabels
}
// GetComparisons gets the comparisons property value. Collection of Experiment Comparisons to run during an Experiment evaluation
// returns a []ExperimentComparisonable when successful
func (m *ExperimentProfile) GetComparisons()([]ExperimentComparisonable) {
    return m.comparisons
}
// GetExtraLabels gets the extraLabels property value. These labels are not used by Horreum but are added to the result event and therefore can be used e.g. when firing an Action.
// returns a []string when successful
func (m *ExperimentProfile) GetExtraLabels()([]string) {
    return m.extraLabels
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExperimentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["baselineFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaselineFilter(val)
        }
        return nil
    }
    res["baselineLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBaselineLabels(res)
        }
        return nil
    }
    res["comparisons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExperimentComparisonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExperimentComparisonable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExperimentComparisonable)
                }
            }
            m.SetComparisons(res)
        }
        return nil
    }
    res["extraLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetExtraLabels(res)
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
    res["selectorFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelectorFilter(val)
        }
        return nil
    }
    res["selectorLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSelectorLabels(res)
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
// GetId gets the id property value. Experiment Profile unique ID
// returns a *int32 when successful
func (m *ExperimentProfile) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Name of Experiment Profile
// returns a *string when successful
func (m *ExperimentProfile) GetName()(*string) {
    return m.name
}
// GetSelectorFilter gets the selectorFilter property value. Selector filter to apply to Selector label values
// returns a *string when successful
func (m *ExperimentProfile) GetSelectorFilter()(*string) {
    return m.selectorFilter
}
// GetSelectorLabels gets the selectorLabels property value. Array of selector labels
// returns a []string when successful
func (m *ExperimentProfile) GetSelectorLabels()([]string) {
    return m.selectorLabels
}
// GetTestId gets the testId property value. Test ID that Experiment Profile relates to
// returns a *int32 when successful
func (m *ExperimentProfile) GetTestId()(*int32) {
    return m.testId
}
// Serialize serializes information the current object
func (m *ExperimentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("baselineFilter", m.GetBaselineFilter())
        if err != nil {
            return err
        }
    }
    if m.GetBaselineLabels() != nil {
        err := writer.WriteCollectionOfStringValues("baselineLabels", m.GetBaselineLabels())
        if err != nil {
            return err
        }
    }
    if m.GetComparisons() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComparisons()))
        for i, v := range m.GetComparisons() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("comparisons", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExtraLabels() != nil {
        err := writer.WriteCollectionOfStringValues("extraLabels", m.GetExtraLabels())
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
        err := writer.WriteStringValue("selectorFilter", m.GetSelectorFilter())
        if err != nil {
            return err
        }
    }
    if m.GetSelectorLabels() != nil {
        err := writer.WriteCollectionOfStringValues("selectorLabels", m.GetSelectorLabels())
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
func (m *ExperimentProfile) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBaselineFilter sets the baselineFilter property value. Selector filter to apply to Baseline label values
func (m *ExperimentProfile) SetBaselineFilter(value *string)() {
    m.baselineFilter = value
}
// SetBaselineLabels sets the baselineLabels property value. Array of selector labels for comparison Baseline
func (m *ExperimentProfile) SetBaselineLabels(value []string)() {
    m.baselineLabels = value
}
// SetComparisons sets the comparisons property value. Collection of Experiment Comparisons to run during an Experiment evaluation
func (m *ExperimentProfile) SetComparisons(value []ExperimentComparisonable)() {
    m.comparisons = value
}
// SetExtraLabels sets the extraLabels property value. These labels are not used by Horreum but are added to the result event and therefore can be used e.g. when firing an Action.
func (m *ExperimentProfile) SetExtraLabels(value []string)() {
    m.extraLabels = value
}
// SetId sets the id property value. Experiment Profile unique ID
func (m *ExperimentProfile) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Name of Experiment Profile
func (m *ExperimentProfile) SetName(value *string)() {
    m.name = value
}
// SetSelectorFilter sets the selectorFilter property value. Selector filter to apply to Selector label values
func (m *ExperimentProfile) SetSelectorFilter(value *string)() {
    m.selectorFilter = value
}
// SetSelectorLabels sets the selectorLabels property value. Array of selector labels
func (m *ExperimentProfile) SetSelectorLabels(value []string)() {
    m.selectorLabels = value
}
// SetTestId sets the testId property value. Test ID that Experiment Profile relates to
func (m *ExperimentProfile) SetTestId(value *int32)() {
    m.testId = value
}
type ExperimentProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaselineFilter()(*string)
    GetBaselineLabels()([]string)
    GetComparisons()([]ExperimentComparisonable)
    GetExtraLabels()([]string)
    GetId()(*int32)
    GetName()(*string)
    GetSelectorFilter()(*string)
    GetSelectorLabels()([]string)
    GetTestId()(*int32)
    SetBaselineFilter(value *string)()
    SetBaselineLabels(value []string)()
    SetComparisons(value []ExperimentComparisonable)()
    SetExtraLabels(value []string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetSelectorFilter(value *string)()
    SetSelectorLabels(value []string)()
    SetTestId(value *int32)()
}
