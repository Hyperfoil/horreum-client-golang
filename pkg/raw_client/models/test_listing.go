package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TestListing struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Number of tests when pagination is ignored
    count *int64
    // Array of Test Summaries
    tests []TestSummaryable
}
// NewTestListing instantiates a new TestListing and sets the default values.
func NewTestListing()(*TestListing) {
    m := &TestListing{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTestListingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTestListingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTestListing(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TestListing) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. Number of tests when pagination is ignored
// returns a *int64 when successful
func (m *TestListing) GetCount()(*int64) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TestListing) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["tests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTestSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TestSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TestSummaryable)
                }
            }
            m.SetTests(res)
        }
        return nil
    }
    return res
}
// GetTests gets the tests property value. Array of Test Summaries
// returns a []TestSummaryable when successful
func (m *TestListing) GetTests()([]TestSummaryable) {
    return m.tests
}
// Serialize serializes information the current object
func (m *TestListing) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    if m.GetTests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTests()))
        for i, v := range m.GetTests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("tests", cast)
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
func (m *TestListing) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. Number of tests when pagination is ignored
func (m *TestListing) SetCount(value *int64)() {
    m.count = value
}
// SetTests sets the tests property value. Array of Test Summaries
func (m *TestListing) SetTests(value []TestSummaryable)() {
    m.tests = value
}
type TestListingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int64)
    GetTests()([]TestSummaryable)
    SetCount(value *int64)()
    SetTests(value []TestSummaryable)()
}
