package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExperimentResult_datasetInfo dataset Info about dataset used for experiment
type ExperimentResult_datasetInfo struct {
    DatasetInfo
}
// NewExperimentResult_datasetInfo instantiates a new ExperimentResult_datasetInfo and sets the default values.
func NewExperimentResult_datasetInfo()(*ExperimentResult_datasetInfo) {
    m := &ExperimentResult_datasetInfo{
        DatasetInfo: *NewDatasetInfo(),
    }
    return m
}
// CreateExperimentResult_datasetInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExperimentResult_datasetInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExperimentResult_datasetInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExperimentResult_datasetInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DatasetInfo.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ExperimentResult_datasetInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DatasetInfo.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type ExperimentResult_datasetInfoable interface {
    DatasetInfoable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
