package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelValue_schema summary description of Schema
type LabelValue_schema struct {
    SchemaDescriptor
}
// NewLabelValue_schema instantiates a new LabelValue_schema and sets the default values.
func NewLabelValue_schema()(*LabelValue_schema) {
    m := &LabelValue_schema{
        SchemaDescriptor: *NewSchemaDescriptor(),
    }
    return m
}
// CreateLabelValue_schemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelValue_schemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelValue_schema(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelValue_schema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SchemaDescriptor.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *LabelValue_schema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SchemaDescriptor.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type LabelValue_schemaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SchemaDescriptorable
}
