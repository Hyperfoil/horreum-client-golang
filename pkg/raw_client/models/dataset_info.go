package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DatasetInfo dataset Info about dataset used for experiment
type DatasetInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Dataset ID for Dataset
    id *int32
    // Ordinal position in ordered list
    ordinal *int32
    // Run ID that Dataset relates to
    runId *int32
    // Test ID that Dataset relates to
    testId *int32
}
// NewDatasetInfo instantiates a new DatasetInfo and sets the default values.
func NewDatasetInfo()(*DatasetInfo) {
    m := &DatasetInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDatasetInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatasetInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatasetInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DatasetInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DatasetInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["ordinal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrdinal(val)
        }
        return nil
    }
    res["runId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunId(val)
        }
        return nil
    }
    res["testId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Dataset ID for Dataset
// returns a *int32 when successful
func (m *DatasetInfo) GetId()(*int32) {
    return m.id
}
// GetOrdinal gets the ordinal property value. Ordinal position in ordered list
// returns a *int32 when successful
func (m *DatasetInfo) GetOrdinal()(*int32) {
    return m.ordinal
}
// GetRunId gets the runId property value. Run ID that Dataset relates to
// returns a *int32 when successful
func (m *DatasetInfo) GetRunId()(*int32) {
    return m.runId
}
// GetTestId gets the testId property value. Test ID that Dataset relates to
// returns a *int32 when successful
func (m *DatasetInfo) GetTestId()(*int32) {
    return m.testId
}
// Serialize serializes information the current object
func (m *DatasetInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ordinal", m.GetOrdinal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("runId", m.GetRunId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("testId", m.GetTestId())
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
func (m *DatasetInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Dataset ID for Dataset
func (m *DatasetInfo) SetId(value *int32)() {
    m.id = value
}
// SetOrdinal sets the ordinal property value. Ordinal position in ordered list
func (m *DatasetInfo) SetOrdinal(value *int32)() {
    m.ordinal = value
}
// SetRunId sets the runId property value. Run ID that Dataset relates to
func (m *DatasetInfo) SetRunId(value *int32)() {
    m.runId = value
}
// SetTestId sets the testId property value. Test ID that Dataset relates to
func (m *DatasetInfo) SetTestId(value *int32)() {
    m.testId = value
}
type DatasetInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetOrdinal()(*int32)
    GetRunId()(*int32)
    GetTestId()(*int32)
    SetId(value *int32)()
    SetOrdinal(value *int32)()
    SetRunId(value *int32)()
    SetTestId(value *int32)()
}
