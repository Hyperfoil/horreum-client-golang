package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExportedLabelValues a map of label names to label values with the associated datasetId and runId
type ExportedLabelValues struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // the unique dataset id
    datasetId *int32
    // the run id that created the dataset
    runId *int32
    // Start timestamp
    start *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Stop timestamp
    stop *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // a map of label name to value
    values LabelValueMapable
}
// NewExportedLabelValues instantiates a new ExportedLabelValues and sets the default values.
func NewExportedLabelValues()(*ExportedLabelValues) {
    m := &ExportedLabelValues{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExportedLabelValuesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExportedLabelValuesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExportedLabelValues(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExportedLabelValues) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatasetId gets the datasetId property value. the unique dataset id
// returns a *int32 when successful
func (m *ExportedLabelValues) GetDatasetId()(*int32) {
    return m.datasetId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExportedLabelValues) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    res["stop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStop(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLabelValueMapFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(LabelValueMapable))
        }
        return nil
    }
    return res
}
// GetRunId gets the runId property value. the run id that created the dataset
// returns a *int32 when successful
func (m *ExportedLabelValues) GetRunId()(*int32) {
    return m.runId
}
// GetStart gets the start property value. Start timestamp
// returns a *Time when successful
func (m *ExportedLabelValues) GetStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.start
}
// GetStop gets the stop property value. Stop timestamp
// returns a *Time when successful
func (m *ExportedLabelValues) GetStop()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.stop
}
// GetValues gets the values property value. a map of label name to value
// returns a LabelValueMapable when successful
func (m *ExportedLabelValues) GetValues()(LabelValueMapable) {
    return m.values
}
// Serialize serializes information the current object
func (m *ExportedLabelValues) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("datasetId", m.GetDatasetId())
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
        err := writer.WriteTimeValue("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("stop", m.GetStop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("values", m.GetValues())
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
func (m *ExportedLabelValues) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatasetId sets the datasetId property value. the unique dataset id
func (m *ExportedLabelValues) SetDatasetId(value *int32)() {
    m.datasetId = value
}
// SetRunId sets the runId property value. the run id that created the dataset
func (m *ExportedLabelValues) SetRunId(value *int32)() {
    m.runId = value
}
// SetStart sets the start property value. Start timestamp
func (m *ExportedLabelValues) SetStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.start = value
}
// SetStop sets the stop property value. Stop timestamp
func (m *ExportedLabelValues) SetStop(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.stop = value
}
// SetValues sets the values property value. a map of label name to value
func (m *ExportedLabelValues) SetValues(value LabelValueMapable)() {
    m.values = value
}
type ExportedLabelValuesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatasetId()(*int32)
    GetRunId()(*int32)
    GetStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStop()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetValues()(LabelValueMapable)
    SetDatasetId(value *int32)()
    SetRunId(value *int32)()
    SetStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStop(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetValues(value LabelValueMapable)()
}
