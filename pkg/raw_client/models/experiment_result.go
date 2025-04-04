package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExperimentResult result of running an Experiment
type ExperimentResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A list of Dataset Info for experiment baseline(s)
    baseline []DatasetInfoable
    // Dataset Info about dataset used for experiment
    datasetInfo DatasetInfoable
    // The extraLabels property
    extraLabels *string
    // A list of log statements recorded while Experiment was evaluated
    logs []DatasetLogable
    // The notify property
    notify *bool
    // Experiment profile that results relates to
    profile ExperimentProfileable
    // A Map of all comparisons and results evaluated during an Experiment
    results ExperimentResult_resultsable
}
// NewExperimentResult instantiates a new ExperimentResult and sets the default values.
func NewExperimentResult()(*ExperimentResult) {
    m := &ExperimentResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExperimentResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExperimentResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExperimentResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExperimentResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBaseline gets the baseline property value. A list of Dataset Info for experiment baseline(s)
// returns a []DatasetInfoable when successful
func (m *ExperimentResult) GetBaseline()([]DatasetInfoable) {
    return m.baseline
}
// GetDatasetInfo gets the datasetInfo property value. Dataset Info about dataset used for experiment
// returns a DatasetInfoable when successful
func (m *ExperimentResult) GetDatasetInfo()(DatasetInfoable) {
    return m.datasetInfo
}
// GetExtraLabels gets the extraLabels property value. The extraLabels property
// returns a *string when successful
func (m *ExperimentResult) GetExtraLabels()(*string) {
    return m.extraLabels
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExperimentResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["baseline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDatasetInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DatasetInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DatasetInfoable)
                }
            }
            m.SetBaseline(res)
        }
        return nil
    }
    res["datasetInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDatasetInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatasetInfo(val.(DatasetInfoable))
        }
        return nil
    }
    res["extraLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtraLabels(val)
        }
        return nil
    }
    res["logs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDatasetLogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DatasetLogable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DatasetLogable)
                }
            }
            m.SetLogs(res)
        }
        return nil
    }
    res["notify"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotify(val)
        }
        return nil
    }
    res["profile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExperimentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfile(val.(ExperimentProfileable))
        }
        return nil
    }
    res["results"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExperimentResult_resultsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResults(val.(ExperimentResult_resultsable))
        }
        return nil
    }
    return res
}
// GetLogs gets the logs property value. A list of log statements recorded while Experiment was evaluated
// returns a []DatasetLogable when successful
func (m *ExperimentResult) GetLogs()([]DatasetLogable) {
    return m.logs
}
// GetNotify gets the notify property value. The notify property
// returns a *bool when successful
func (m *ExperimentResult) GetNotify()(*bool) {
    return m.notify
}
// GetProfile gets the profile property value. Experiment profile that results relates to
// returns a ExperimentProfileable when successful
func (m *ExperimentResult) GetProfile()(ExperimentProfileable) {
    return m.profile
}
// GetResults gets the results property value. A Map of all comparisons and results evaluated during an Experiment
// returns a ExperimentResult_resultsable when successful
func (m *ExperimentResult) GetResults()(ExperimentResult_resultsable) {
    return m.results
}
// Serialize serializes information the current object
func (m *ExperimentResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBaseline() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBaseline()))
        for i, v := range m.GetBaseline() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("baseline", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("datasetInfo", m.GetDatasetInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extraLabels", m.GetExtraLabels())
        if err != nil {
            return err
        }
    }
    if m.GetLogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLogs()))
        for i, v := range m.GetLogs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("logs", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("notify", m.GetNotify())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("profile", m.GetProfile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("results", m.GetResults())
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
func (m *ExperimentResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBaseline sets the baseline property value. A list of Dataset Info for experiment baseline(s)
func (m *ExperimentResult) SetBaseline(value []DatasetInfoable)() {
    m.baseline = value
}
// SetDatasetInfo sets the datasetInfo property value. Dataset Info about dataset used for experiment
func (m *ExperimentResult) SetDatasetInfo(value DatasetInfoable)() {
    m.datasetInfo = value
}
// SetExtraLabels sets the extraLabels property value. The extraLabels property
func (m *ExperimentResult) SetExtraLabels(value *string)() {
    m.extraLabels = value
}
// SetLogs sets the logs property value. A list of log statements recorded while Experiment was evaluated
func (m *ExperimentResult) SetLogs(value []DatasetLogable)() {
    m.logs = value
}
// SetNotify sets the notify property value. The notify property
func (m *ExperimentResult) SetNotify(value *bool)() {
    m.notify = value
}
// SetProfile sets the profile property value. Experiment profile that results relates to
func (m *ExperimentResult) SetProfile(value ExperimentProfileable)() {
    m.profile = value
}
// SetResults sets the results property value. A Map of all comparisons and results evaluated during an Experiment
func (m *ExperimentResult) SetResults(value ExperimentResult_resultsable)() {
    m.results = value
}
type ExperimentResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseline()([]DatasetInfoable)
    GetDatasetInfo()(DatasetInfoable)
    GetExtraLabels()(*string)
    GetLogs()([]DatasetLogable)
    GetNotify()(*bool)
    GetProfile()(ExperimentProfileable)
    GetResults()(ExperimentResult_resultsable)
    SetBaseline(value []DatasetInfoable)()
    SetDatasetInfo(value DatasetInfoable)()
    SetExtraLabels(value *string)()
    SetLogs(value []DatasetLogable)()
    SetNotify(value *bool)()
    SetProfile(value ExperimentProfileable)()
    SetResults(value ExperimentResult_resultsable)()
}
