package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UsernamePassAuth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Password
    password *string
    // type
    typeEscaped *string
    // Username
    username *string
}
// NewUsernamePassAuth instantiates a new UsernamePassAuth and sets the default values.
func NewUsernamePassAuth()(*UsernamePassAuth) {
    m := &UsernamePassAuth{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUsernamePassAuthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsernamePassAuthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsernamePassAuth(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UsernamePassAuth) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UsernamePassAuth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. Password
// returns a *string when successful
func (m *UsernamePassAuth) GetPassword()(*string) {
    return m.password
}
// GetTypeEscaped gets the type property value. type
// returns a *string when successful
func (m *UsernamePassAuth) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUsername gets the username property value. Username
// returns a *string when successful
func (m *UsernamePassAuth) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *UsernamePassAuth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("username", m.GetUsername())
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
func (m *UsernamePassAuth) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPassword sets the password property value. Password
func (m *UsernamePassAuth) SetPassword(value *string)() {
    m.password = value
}
// SetTypeEscaped sets the type property value. type
func (m *UsernamePassAuth) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUsername sets the username property value. Username
func (m *UsernamePassAuth) SetUsername(value *string)() {
    m.username = value
}
type UsernamePassAuthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPassword()(*string)
    GetTypeEscaped()(*string)
    GetUsername()(*string)
    SetPassword(value *string)()
    SetTypeEscaped(value *string)()
    SetUsername(value *string)()
}
