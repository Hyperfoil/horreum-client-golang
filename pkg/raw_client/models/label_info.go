package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LabelInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Is label a filtering label?
    filtering *bool
    // Is label a metrics label?
    metrics *bool
    // Label name
    name *string
    // List of schemas where label is referenced
    schemas []SchemaDescriptorable
}
// NewLabelInfo instantiates a new LabelInfo and sets the default values.
func NewLabelInfo()(*LabelInfo) {
    m := &LabelInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LabelInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["filtering"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFiltering(val)
        }
        return nil
    }
    res["metrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetrics(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSchemaDescriptorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SchemaDescriptorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SchemaDescriptorable)
                }
            }
            m.SetSchemas(res)
        }
        return nil
    }
    return res
}
// GetFiltering gets the filtering property value. Is label a filtering label?
// returns a *bool when successful
func (m *LabelInfo) GetFiltering()(*bool) {
    return m.filtering
}
// GetMetrics gets the metrics property value. Is label a metrics label?
// returns a *bool when successful
func (m *LabelInfo) GetMetrics()(*bool) {
    return m.metrics
}
// GetName gets the name property value. Label name
// returns a *string when successful
func (m *LabelInfo) GetName()(*string) {
    return m.name
}
// GetSchemas gets the schemas property value. List of schemas where label is referenced
// returns a []SchemaDescriptorable when successful
func (m *LabelInfo) GetSchemas()([]SchemaDescriptorable) {
    return m.schemas
}
// Serialize serializes information the current object
func (m *LabelInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("filtering", m.GetFiltering())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("metrics", m.GetMetrics())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *LabelInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFiltering sets the filtering property value. Is label a filtering label?
func (m *LabelInfo) SetFiltering(value *bool)() {
    m.filtering = value
}
// SetMetrics sets the metrics property value. Is label a metrics label?
func (m *LabelInfo) SetMetrics(value *bool)() {
    m.metrics = value
}
// SetName sets the name property value. Label name
func (m *LabelInfo) SetName(value *string)() {
    m.name = value
}
// SetSchemas sets the schemas property value. List of schemas where label is referenced
func (m *LabelInfo) SetSchemas(value []SchemaDescriptorable)() {
    m.schemas = value
}
type LabelInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFiltering()(*bool)
    GetMetrics()(*bool)
    GetName()(*string)
    GetSchemas()([]SchemaDescriptorable)
    SetFiltering(value *bool)()
    SetMetrics(value *bool)()
    SetName(value *string)()
    SetSchemas(value []SchemaDescriptorable)()
}
