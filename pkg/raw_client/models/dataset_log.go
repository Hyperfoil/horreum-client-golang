package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DatasetLog dataset Log
type DatasetLog struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The datasetId property
    datasetId *int32
    // The datasetOrdinal property
    datasetOrdinal *int32
    // The id property
    id *int64
    // The level property
    level *int32
    // The message property
    message *string
    // The runId property
    runId *int32
    // The source property
    source *string
    // The testId property
    testId *int32
    // The timestamp property
    timestamp *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewDatasetLog instantiates a new DatasetLog and sets the default values.
func NewDatasetLog()(*DatasetLog) {
    m := &DatasetLog{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDatasetLogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatasetLogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatasetLog(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DatasetLog) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatasetId gets the datasetId property value. The datasetId property
// returns a *int32 when successful
func (m *DatasetLog) GetDatasetId()(*int32) {
    return m.datasetId
}
// GetDatasetOrdinal gets the datasetOrdinal property value. The datasetOrdinal property
// returns a *int32 when successful
func (m *DatasetLog) GetDatasetOrdinal()(*int32) {
    return m.datasetOrdinal
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DatasetLog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["datasetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatasetId(val)
        }
        return nil
    }
    res["datasetOrdinal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatasetOrdinal(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
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
    res["timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *DatasetLog) GetId()(*int64) {
    return m.id
}
// GetLevel gets the level property value. The level property
// returns a *int32 when successful
func (m *DatasetLog) GetLevel()(*int32) {
    return m.level
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *DatasetLog) GetMessage()(*string) {
    return m.message
}
// GetRunId gets the runId property value. The runId property
// returns a *int32 when successful
func (m *DatasetLog) GetRunId()(*int32) {
    return m.runId
}
// GetSource gets the source property value. The source property
// returns a *string when successful
func (m *DatasetLog) GetSource()(*string) {
    return m.source
}
// GetTestId gets the testId property value. The testId property
// returns a *int32 when successful
func (m *DatasetLog) GetTestId()(*int32) {
    return m.testId
}
// GetTimestamp gets the timestamp property value. The timestamp property
// returns a *Time when successful
func (m *DatasetLog) GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.timestamp
}
// Serialize serializes information the current object
func (m *DatasetLog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("datasetId", m.GetDatasetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("datasetOrdinal", m.GetDatasetOrdinal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("level", m.GetLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
        err := writer.WriteStringValue("source", m.GetSource())
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
        err := writer.WriteTimeValue("timestamp", m.GetTimestamp())
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
func (m *DatasetLog) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatasetId sets the datasetId property value. The datasetId property
func (m *DatasetLog) SetDatasetId(value *int32)() {
    m.datasetId = value
}
// SetDatasetOrdinal sets the datasetOrdinal property value. The datasetOrdinal property
func (m *DatasetLog) SetDatasetOrdinal(value *int32)() {
    m.datasetOrdinal = value
}
// SetId sets the id property value. The id property
func (m *DatasetLog) SetId(value *int64)() {
    m.id = value
}
// SetLevel sets the level property value. The level property
func (m *DatasetLog) SetLevel(value *int32)() {
    m.level = value
}
// SetMessage sets the message property value. The message property
func (m *DatasetLog) SetMessage(value *string)() {
    m.message = value
}
// SetRunId sets the runId property value. The runId property
func (m *DatasetLog) SetRunId(value *int32)() {
    m.runId = value
}
// SetSource sets the source property value. The source property
func (m *DatasetLog) SetSource(value *string)() {
    m.source = value
}
// SetTestId sets the testId property value. The testId property
func (m *DatasetLog) SetTestId(value *int32)() {
    m.testId = value
}
// SetTimestamp sets the timestamp property value. The timestamp property
func (m *DatasetLog) SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timestamp = value
}
type DatasetLogable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatasetId()(*int32)
    GetDatasetOrdinal()(*int32)
    GetId()(*int64)
    GetLevel()(*int32)
    GetMessage()(*string)
    GetRunId()(*int32)
    GetSource()(*string)
    GetTestId()(*int32)
    GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDatasetId(value *int32)()
    SetDatasetOrdinal(value *int32)()
    SetId(value *int64)()
    SetLevel(value *int32)()
    SetMessage(value *string)()
    SetRunId(value *int32)()
    SetSource(value *string)()
    SetTestId(value *int32)()
    SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
