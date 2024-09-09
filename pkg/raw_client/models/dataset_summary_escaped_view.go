package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DatasetSummary_view map of view component ids to the LabelValueMap to render the component for this dataset
type DatasetSummary_view struct {
    IndexedLabelValueMap
}
// NewDatasetSummary_view instantiates a new DatasetSummary_view and sets the default values.
func NewDatasetSummary_view()(*DatasetSummary_view) {
    m := &DatasetSummary_view{
        IndexedLabelValueMap: *NewIndexedLabelValueMap(),
    }
    return m
}
// CreateDatasetSummary_viewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatasetSummary_viewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatasetSummary_view(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DatasetSummary_view) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IndexedLabelValueMap.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DatasetSummary_view) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IndexedLabelValueMap.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type DatasetSummary_viewable interface {
    IndexedLabelValueMapable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
