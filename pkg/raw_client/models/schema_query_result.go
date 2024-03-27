package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SchemaQueryResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Count of available Schemas. This is a count of Schemas that the current user has access to
    count *int64
    // Array of Schemas
    schemas []Schemaable
}
// NewSchemaQueryResult instantiates a new SchemaQueryResult and sets the default values.
func NewSchemaQueryResult()(*SchemaQueryResult) {
    m := &SchemaQueryResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSchemaQueryResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemaQueryResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchemaQueryResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SchemaQueryResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. Count of available Schemas. This is a count of Schemas that the current user has access to
// returns a *int64 when successful
func (m *SchemaQueryResult) GetCount()(*int64) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SchemaQueryResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Schemaable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Schemaable)
                }
            }
            m.SetSchemas(res)
        }
        return nil
    }
    return res
}
// GetSchemas gets the schemas property value. Array of Schemas
// returns a []Schemaable when successful
func (m *SchemaQueryResult) GetSchemas()([]Schemaable) {
    return m.schemas
}
// Serialize serializes information the current object
func (m *SchemaQueryResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    if m.GetSchemas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchemas()))
        for i, v := range m.GetSchemas() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("schemas", cast)
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
func (m *SchemaQueryResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. Count of available Schemas. This is a count of Schemas that the current user has access to
func (m *SchemaQueryResult) SetCount(value *int64)() {
    m.count = value
}
// SetSchemas sets the schemas property value. Array of Schemas
func (m *SchemaQueryResult) SetSchemas(value []Schemaable)() {
    m.schemas = value
}
type SchemaQueryResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int64)
    GetSchemas()([]Schemaable)
    SetCount(value *int64)()
    SetSchemas(value []Schemaable)()
}
