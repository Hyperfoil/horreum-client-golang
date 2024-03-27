package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KeycloakConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Keycloak client ID in Horreum realm for User Interface
    clientId *string
    // Keycloak realm securing Horreum instance
    realm *string
    // URL of Keycloak instance securing Horreum
    url *string
}
// NewKeycloakConfig instantiates a new KeycloakConfig and sets the default values.
func NewKeycloakConfig()(*KeycloakConfig) {
    m := &KeycloakConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateKeycloakConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKeycloakConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeycloakConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *KeycloakConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the clientId property value. Keycloak client ID in Horreum realm for User Interface
// returns a *string when successful
func (m *KeycloakConfig) GetClientId()(*string) {
    return m.clientId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KeycloakConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["realm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealm(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetRealm gets the realm property value. Keycloak realm securing Horreum instance
// returns a *string when successful
func (m *KeycloakConfig) GetRealm()(*string) {
    return m.realm
}
// GetUrl gets the url property value. URL of Keycloak instance securing Horreum
// returns a *string when successful
func (m *KeycloakConfig) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *KeycloakConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("realm", m.GetRealm())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *KeycloakConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the clientId property value. Keycloak client ID in Horreum realm for User Interface
func (m *KeycloakConfig) SetClientId(value *string)() {
    m.clientId = value
}
// SetRealm sets the realm property value. Keycloak realm securing Horreum instance
func (m *KeycloakConfig) SetRealm(value *string)() {
    m.realm = value
}
// SetUrl sets the url property value. URL of Keycloak instance securing Horreum
func (m *KeycloakConfig) SetUrl(value *string)() {
    m.url = value
}
type KeycloakConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetRealm()(*string)
    GetUrl()(*string)
    SetClientId(value *string)()
    SetRealm(value *string)()
    SetUrl(value *string)()
}
