package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExperimentResult_profile experiment profile that results relates to
type ExperimentResult_profile struct {
    ExperimentProfile
}
// NewExperimentResult_profile instantiates a new ExperimentResult_profile and sets the default values.
func NewExperimentResult_profile()(*ExperimentResult_profile) {
    m := &ExperimentResult_profile{
        ExperimentProfile: *NewExperimentProfile(),
    }
    return m
}
// CreateExperimentResult_profileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExperimentResult_profileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExperimentResult_profile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExperimentResult_profile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExperimentProfile.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ExperimentResult_profile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExperimentProfile.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type ExperimentResult_profileable interface {
    ExperimentProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
