package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

type UserCreateUserPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The password property
    password *string
    // The roles property
    roles []string
    // The team property
    team *string
    // The user property
    user i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.UserDataable
}
// NewUserCreateUserPostRequestBody instantiates a new UserCreateUserPostRequestBody and sets the default values.
func NewUserCreateUserPostRequestBody()(*UserCreateUserPostRequestBody) {
    m := &UserCreateUserPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserCreateUserPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserCreateUserPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserCreateUserPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserCreateUserPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserCreateUserPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRoles(res)
        }
        return nil
    }
    res["team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.CreateUserDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.UserDataable))
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password property
// returns a *string when successful
func (m *UserCreateUserPostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetRoles gets the roles property value. The roles property
// returns a []string when successful
func (m *UserCreateUserPostRequestBody) GetRoles()([]string) {
    return m.roles
}
// GetTeam gets the team property value. The team property
// returns a *string when successful
func (m *UserCreateUserPostRequestBody) GetTeam()(*string) {
    return m.team
}
// GetUser gets the user property value. The user property
// returns a UserDataable when successful
func (m *UserCreateUserPostRequestBody) GetUser()(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.UserDataable) {
    return m.user
}
// Serialize serializes information the current object
func (m *UserCreateUserPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        err := writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("team", m.GetTeam())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *UserCreateUserPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPassword sets the password property value. The password property
func (m *UserCreateUserPostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetRoles sets the roles property value. The roles property
func (m *UserCreateUserPostRequestBody) SetRoles(value []string)() {
    m.roles = value
}
// SetTeam sets the team property value. The team property
func (m *UserCreateUserPostRequestBody) SetTeam(value *string)() {
    m.team = value
}
// SetUser sets the user property value. The user property
func (m *UserCreateUserPostRequestBody) SetUser(value i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.UserDataable)() {
    m.user = value
}
type UserCreateUserPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPassword()(*string)
    GetRoles()([]string)
    GetTeam()(*string)
    GetUser()(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.UserDataable)
    SetPassword(value *string)()
    SetRoles(value []string)()
    SetTeam(value *string)()
    SetUser(value i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.UserDataable)()
}
