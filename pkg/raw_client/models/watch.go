package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Watch watcher object associated with test
type Watch struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *int32
    // The optout property
    optout []string
    // The teams property
    teams []string
    // The testId property
    testId *int32
    // The users property
    users []string
}
// NewWatch instantiates a new Watch and sets the default values.
func NewWatch()(*Watch) {
    m := &Watch{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWatchFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWatchFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWatch(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Watch) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Watch) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["optout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetOptout(res)
        }
        return nil
    }
    res["teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTeams(res)
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
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Watch) GetId()(*int32) {
    return m.id
}
// GetOptout gets the optout property value. The optout property
// returns a []string when successful
func (m *Watch) GetOptout()([]string) {
    return m.optout
}
// GetTeams gets the teams property value. The teams property
// returns a []string when successful
func (m *Watch) GetTeams()([]string) {
    return m.teams
}
// GetTestId gets the testId property value. The testId property
// returns a *int32 when successful
func (m *Watch) GetTestId()(*int32) {
    return m.testId
}
// GetUsers gets the users property value. The users property
// returns a []string when successful
func (m *Watch) GetUsers()([]string) {
    return m.users
}
// Serialize serializes information the current object
func (m *Watch) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetOptout() != nil {
        err := writer.WriteCollectionOfStringValues("optout", m.GetOptout())
        if err != nil {
            return err
        }
    }
    if m.GetTeams() != nil {
        err := writer.WriteCollectionOfStringValues("teams", m.GetTeams())
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
    if m.GetUsers() != nil {
        err := writer.WriteCollectionOfStringValues("users", m.GetUsers())
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
func (m *Watch) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *Watch) SetId(value *int32)() {
    m.id = value
}
// SetOptout sets the optout property value. The optout property
func (m *Watch) SetOptout(value []string)() {
    m.optout = value
}
// SetTeams sets the teams property value. The teams property
func (m *Watch) SetTeams(value []string)() {
    m.teams = value
}
// SetTestId sets the testId property value. The testId property
func (m *Watch) SetTestId(value *int32)() {
    m.testId = value
}
// SetUsers sets the users property value. The users property
func (m *Watch) SetUsers(value []string)() {
    m.users = value
}
type Watchable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetOptout()([]string)
    GetTeams()([]string)
    GetTestId()(*int32)
    GetUsers()([]string)
    SetId(value *int32)()
    SetOptout(value []string)()
    SetTeams(value []string)()
    SetTestId(value *int32)()
    SetUsers(value []string)()
}
