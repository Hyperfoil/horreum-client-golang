package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TestExport_subscriptions watcher object associated with test
type TestExport_subscriptions struct {
    Watch
}
// NewTestExport_subscriptions instantiates a new TestExport_subscriptions and sets the default values.
func NewTestExport_subscriptions()(*TestExport_subscriptions) {
    m := &TestExport_subscriptions{
        Watch: *NewWatch(),
    }
    return m
}
// CreateTestExport_subscriptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTestExport_subscriptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTestExport_subscriptions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TestExport_subscriptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Watch.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TestExport_subscriptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Watch.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type TestExport_subscriptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Watchable
}
