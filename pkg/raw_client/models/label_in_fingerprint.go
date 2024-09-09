package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LabelInFingerprint struct {
    LabelLocation
}
// NewLabelInFingerprint instantiates a new LabelInFingerprint and sets the default values.
func NewLabelInFingerprint()(*LabelInFingerprint) {
    m := &LabelInFingerprint{
        LabelLocation: *NewLabelLocation(),
    }
    return m
}
// CreateLabelInFingerprintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelInFingerprintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelInFingerprint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelInFingerprint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelLocation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *LabelInFingerprint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelLocation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type LabelInFingerprintable interface {
    LabelLocationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
