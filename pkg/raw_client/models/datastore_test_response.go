package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DatastoreTestResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The msg property
    msg *string
    // The success property
    success *bool
}
// NewDatastoreTestResponse instantiates a new DatastoreTestResponse and sets the default values.
func NewDatastoreTestResponse()(*DatastoreTestResponse) {
    m := &DatastoreTestResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDatastoreTestResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatastoreTestResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatastoreTestResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DatastoreTestResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DatastoreTestResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["msg"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMsg(val)
        }
        return nil
    }
    res["success"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccess(val)
        }
        return nil
    }
    return res
}
// GetMsg gets the msg property value. The msg property
// returns a *string when successful
func (m *DatastoreTestResponse) GetMsg()(*string) {
    return m.msg
}
// GetSuccess gets the success property value. The success property
// returns a *bool when successful
func (m *DatastoreTestResponse) GetSuccess()(*bool) {
    return m.success
}
// Serialize serializes information the current object
func (m *DatastoreTestResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("msg", m.GetMsg())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("success", m.GetSuccess())
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
func (m *DatastoreTestResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMsg sets the msg property value. The msg property
func (m *DatastoreTestResponse) SetMsg(value *string)() {
    m.msg = value
}
// SetSuccess sets the success property value. The success property
func (m *DatastoreTestResponse) SetSuccess(value *bool)() {
    m.success = value
}
type DatastoreTestResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMsg()(*string)
    GetSuccess()(*bool)
    SetMsg(value *string)()
    SetSuccess(value *bool)()
}
