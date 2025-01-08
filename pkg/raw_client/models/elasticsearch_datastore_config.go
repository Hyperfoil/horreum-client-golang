package models

import (
    ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6 "strings"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ElasticsearchDatastoreConfig type of backend datastore
type ElasticsearchDatastoreConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The authentication property
    authentication ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationable
    // Built In
    builtIn *bool
    // Elasticsearch url
    url *string
}
// ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication composed type wrapper for classes APIKeyAuthable, NoAuthable, UsernamePassAuthable
type ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication struct {
    // Composed type representation for type APIKeyAuthable
    aPIKeyAuth APIKeyAuthable
    // Composed type representation for type NoAuthable
    noAuth NoAuthable
    // Composed type representation for type UsernamePassAuthable
    usernamePassAuth UsernamePassAuthable
}
// NewElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication instantiates a new ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication and sets the default values.
func NewElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication()(*ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) {
    m := &ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication{
    }
    return m
}
// CreateElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "api-key") {
                    result.SetAPIKeyAuth(NewAPIKeyAuth())
                } else if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "none") {
                    result.SetNoAuth(NewNoAuth())
                } else if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "username") {
                    result.SetUsernamePassAuth(NewUsernamePassAuth())
                }
            }
        }
    }
    return result, nil
}
// GetAPIKeyAuth gets the APIKeyAuth property value. Composed type representation for type APIKeyAuthable
// returns a APIKeyAuthable when successful
func (m *ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) GetAPIKeyAuth()(APIKeyAuthable) {
    return m.aPIKeyAuth
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    if m.GetAPIKeyAuth() != nil {
        return m.GetAPIKeyAuth().GetFieldDeserializers()
    } else if m.GetNoAuth() != nil {
        return m.GetNoAuth().GetFieldDeserializers()
    } else if m.GetUsernamePassAuth() != nil {
        return m.GetUsernamePassAuth().GetFieldDeserializers()
    }
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) GetIsComposedType()(bool) {
    return true
}
// GetNoAuth gets the NoAuth property value. Composed type representation for type NoAuthable
// returns a NoAuthable when successful
func (m *ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) GetNoAuth()(NoAuthable) {
    return m.noAuth
}
// GetUsernamePassAuth gets the UsernamePassAuth property value. Composed type representation for type UsernamePassAuthable
// returns a UsernamePassAuthable when successful
func (m *ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) GetUsernamePassAuth()(UsernamePassAuthable) {
    return m.usernamePassAuth
}
// Serialize serializes information the current object
func (m *ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAPIKeyAuth() != nil {
        err := writer.WriteObjectValue("", m.GetAPIKeyAuth())
        if err != nil {
            return err
        }
    } else if m.GetNoAuth() != nil {
        err := writer.WriteObjectValue("", m.GetNoAuth())
        if err != nil {
            return err
        }
    } else if m.GetUsernamePassAuth() != nil {
        err := writer.WriteObjectValue("", m.GetUsernamePassAuth())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAPIKeyAuth sets the APIKeyAuth property value. Composed type representation for type APIKeyAuthable
func (m *ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) SetAPIKeyAuth(value APIKeyAuthable)() {
    m.aPIKeyAuth = value
}
// SetNoAuth sets the NoAuth property value. Composed type representation for type NoAuthable
func (m *ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) SetNoAuth(value NoAuthable)() {
    m.noAuth = value
}
// SetUsernamePassAuth sets the UsernamePassAuth property value. Composed type representation for type UsernamePassAuthable
func (m *ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authentication) SetUsernamePassAuth(value UsernamePassAuthable)() {
    m.usernamePassAuth = value
}
type ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAPIKeyAuth()(APIKeyAuthable)
    GetNoAuth()(NoAuthable)
    GetUsernamePassAuth()(UsernamePassAuthable)
    SetAPIKeyAuth(value APIKeyAuthable)()
    SetNoAuth(value NoAuthable)()
    SetUsernamePassAuth(value UsernamePassAuthable)()
}
// NewElasticsearchDatastoreConfig instantiates a new ElasticsearchDatastoreConfig and sets the default values.
func NewElasticsearchDatastoreConfig()(*ElasticsearchDatastoreConfig) {
    m := &ElasticsearchDatastoreConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateElasticsearchDatastoreConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateElasticsearchDatastoreConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewElasticsearchDatastoreConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ElasticsearchDatastoreConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthentication gets the authentication property value. The authentication property
// returns a ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationable when successful
func (m *ElasticsearchDatastoreConfig) GetAuthentication()(ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationable) {
    return m.authentication
}
// GetBuiltIn gets the builtIn property value. Built In
// returns a *bool when successful
func (m *ElasticsearchDatastoreConfig) GetBuiltIn()(*bool) {
    return m.builtIn
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ElasticsearchDatastoreConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthentication(val.(ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationable))
        }
        return nil
    }
    res["builtIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuiltIn(val)
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
// GetUrl gets the url property value. Elasticsearch url
// returns a *string when successful
func (m *ElasticsearchDatastoreConfig) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ElasticsearchDatastoreConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authentication", m.GetAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("builtIn", m.GetBuiltIn())
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
func (m *ElasticsearchDatastoreConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthentication sets the authentication property value. The authentication property
func (m *ElasticsearchDatastoreConfig) SetAuthentication(value ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationable)() {
    m.authentication = value
}
// SetBuiltIn sets the builtIn property value. Built In
func (m *ElasticsearchDatastoreConfig) SetBuiltIn(value *bool)() {
    m.builtIn = value
}
// SetUrl sets the url property value. Elasticsearch url
func (m *ElasticsearchDatastoreConfig) SetUrl(value *string)() {
    m.url = value
}
type ElasticsearchDatastoreConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthentication()(ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationable)
    GetBuiltIn()(*bool)
    GetUrl()(*string)
    SetAuthentication(value ElasticsearchDatastoreConfig_ElasticsearchDatastoreConfig_authenticationable)()
    SetBuiltIn(value *bool)()
    SetUrl(value *string)()
}
