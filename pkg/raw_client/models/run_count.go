package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RunCount struct {
    // Total count of active Runs visible
    active *int64
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Total count of Runs visible
    total *int64
    // Total count of trashed Runs
    trashed *int64
}
// NewRunCount instantiates a new RunCount and sets the default values.
func NewRunCount()(*RunCount) {
    m := &RunCount{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRunCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunCount(), nil
}
// GetActive gets the active property value. Total count of active Runs visible
// returns a *int64 when successful
func (m *RunCount) GetActive()(*int64) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RunCount) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RunCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    res["trashed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrashed(val)
        }
        return nil
    }
    return res
}
// GetTotal gets the total property value. Total count of Runs visible
// returns a *int64 when successful
func (m *RunCount) GetTotal()(*int64) {
    return m.total
}
// GetTrashed gets the trashed property value. Total count of trashed Runs
// returns a *int64 when successful
func (m *RunCount) GetTrashed()(*int64) {
    return m.trashed
}
// Serialize serializes information the current object
func (m *RunCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("trashed", m.GetTrashed())
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
// SetActive sets the active property value. Total count of active Runs visible
func (m *RunCount) SetActive(value *int64)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RunCount) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTotal sets the total property value. Total count of Runs visible
func (m *RunCount) SetTotal(value *int64)() {
    m.total = value
}
// SetTrashed sets the trashed property value. Total count of trashed Runs
func (m *RunCount) SetTrashed(value *int64)() {
    m.trashed = value
}
type RunCountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*int64)
    GetTotal()(*int64)
    GetTrashed()(*int64)
    SetActive(value *int64)()
    SetTotal(value *int64)()
    SetTrashed(value *int64)()
}
