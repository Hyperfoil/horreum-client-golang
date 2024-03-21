package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransformerInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Schema ID
    schemaId *int32
    // Schema name
    schemaName *string
    // Schema uri
    schemaUri *string
    // Transformer ID
    transformerId *int32
    // Transformer name
    transformerName *string
}
// NewTransformerInfo instantiates a new TransformerInfo and sets the default values.
func NewTransformerInfo()(*TransformerInfo) {
    m := &TransformerInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransformerInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransformerInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransformerInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransformerInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransformerInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["schemaId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaId(val)
        }
        return nil
    }
    res["schemaName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaName(val)
        }
        return nil
    }
    res["schemaUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaUri(val)
        }
        return nil
    }
    res["transformerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransformerId(val)
        }
        return nil
    }
    res["transformerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransformerName(val)
        }
        return nil
    }
    return res
}
// GetSchemaId gets the schemaId property value. Schema ID
// returns a *int32 when successful
func (m *TransformerInfo) GetSchemaId()(*int32) {
    return m.schemaId
}
// GetSchemaName gets the schemaName property value. Schema name
// returns a *string when successful
func (m *TransformerInfo) GetSchemaName()(*string) {
    return m.schemaName
}
// GetSchemaUri gets the schemaUri property value. Schema uri
// returns a *string when successful
func (m *TransformerInfo) GetSchemaUri()(*string) {
    return m.schemaUri
}
// GetTransformerId gets the transformerId property value. Transformer ID
// returns a *int32 when successful
func (m *TransformerInfo) GetTransformerId()(*int32) {
    return m.transformerId
}
// GetTransformerName gets the transformerName property value. Transformer name
// returns a *string when successful
func (m *TransformerInfo) GetTransformerName()(*string) {
    return m.transformerName
}
// Serialize serializes information the current object
func (m *TransformerInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("schemaId", m.GetSchemaId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schemaName", m.GetSchemaName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schemaUri", m.GetSchemaUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("transformerId", m.GetTransformerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("transformerName", m.GetTransformerName())
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
func (m *TransformerInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSchemaId sets the schemaId property value. Schema ID
func (m *TransformerInfo) SetSchemaId(value *int32)() {
    m.schemaId = value
}
// SetSchemaName sets the schemaName property value. Schema name
func (m *TransformerInfo) SetSchemaName(value *string)() {
    m.schemaName = value
}
// SetSchemaUri sets the schemaUri property value. Schema uri
func (m *TransformerInfo) SetSchemaUri(value *string)() {
    m.schemaUri = value
}
// SetTransformerId sets the transformerId property value. Transformer ID
func (m *TransformerInfo) SetTransformerId(value *int32)() {
    m.transformerId = value
}
// SetTransformerName sets the transformerName property value. Transformer name
func (m *TransformerInfo) SetTransformerName(value *string)() {
    m.transformerName = value
}
type TransformerInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSchemaId()(*int32)
    GetSchemaName()(*string)
    GetSchemaUri()(*string)
    GetTransformerId()(*int32)
    GetTransformerName()(*string)
    SetSchemaId(value *int32)()
    SetSchemaName(value *string)()
    SetSchemaUri(value *string)()
    SetTransformerId(value *int32)()
    SetTransformerName(value *string)()
}
