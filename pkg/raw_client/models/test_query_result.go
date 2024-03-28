package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TestQueryResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Count of available tests. This is a count of tests that the current user has access to
    count *int64
    // Array of Tests
    tests []Testable
}
// NewTestQueryResult instantiates a new TestQueryResult and sets the default values.
func NewTestQueryResult()(*TestQueryResult) {
    m := &TestQueryResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTestQueryResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTestQueryResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTestQueryResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TestQueryResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. Count of available tests. This is a count of tests that the current user has access to
// returns a *int64 when successful
func (m *TestQueryResult) GetCount()(*int64) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TestQueryResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateTestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Testable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Testable)
                }
            }
            m.SetTests(res)
        }
        return nil
    }
    return res
}
// GetTests gets the tests property value. Array of Tests
// returns a []Testable when successful
func (m *TestQueryResult) GetTests()([]Testable) {
    return m.tests
}
// Serialize serializes information the current object
func (m *TestQueryResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TestQueryResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. Count of available tests. This is a count of tests that the current user has access to
func (m *TestQueryResult) SetCount(value *int64)() {
    m.count = value
}
// SetTests sets the tests property value. Array of Tests
func (m *TestQueryResult) SetTests(value []Testable)() {
    m.tests = value
}
type TestQueryResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int64)
    GetTests()([]Testable)
    SetCount(value *int64)()
    SetTests(value []Testable)()
}
