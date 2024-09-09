package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FixedThresholdDetectionConfig_max upper bound for acceptable datapoint values
type FixedThresholdDetectionConfig_max struct {
    FixThresholdConfig
}
// NewFixedThresholdDetectionConfig_max instantiates a new FixedThresholdDetectionConfig_max and sets the default values.
func NewFixedThresholdDetectionConfig_max()(*FixedThresholdDetectionConfig_max) {
    m := &FixedThresholdDetectionConfig_max{
        FixThresholdConfig: *NewFixThresholdConfig(),
    }
    return m
}
// CreateFixedThresholdDetectionConfig_maxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFixedThresholdDetectionConfig_maxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFixedThresholdDetectionConfig_max(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FixedThresholdDetectionConfig_max) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FixThresholdConfig.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *FixedThresholdDetectionConfig_max) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FixThresholdConfig.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type FixedThresholdDetectionConfig_maxable interface {
    FixThresholdConfigable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
