package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TestExport_datastore datastore associated with test
type TestExport_datastore struct {
    Datastore
}
// NewTestExport_datastore instantiates a new TestExport_datastore and sets the default values.
func NewTestExport_datastore()(*TestExport_datastore) {
    m := &TestExport_datastore{
        Datastore: *NewDatastore(),
    }
    return m
}
// CreateTestExport_datastoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTestExport_datastoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTestExport_datastore(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TestExport_datastore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Datastore.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TestExport_datastore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Datastore.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type TestExport_datastoreable interface {
    Datastoreable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
