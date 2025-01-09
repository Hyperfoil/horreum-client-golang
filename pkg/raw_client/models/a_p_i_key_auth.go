package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type APIKeyAuth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Api key
    apiKey *string
    // type
    typeEscaped *string
}
// NewAPIKeyAuth instantiates a new APIKeyAuth and sets the default values.
func NewAPIKeyAuth()(*APIKeyAuth) {
    m := &APIKeyAuth{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAPIKeyAuthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAPIKeyAuthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAPIKeyAuth(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *APIKeyAuth) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApiKey gets the apiKey property value. Api key
// returns a *string when successful
func (m *APIKeyAuth) GetApiKey()(*string) {
    return m.apiKey
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *APIKeyAuth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apiKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiKey(val)
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
    return res
}
// GetTypeEscaped gets the type property value. type
// returns a *string when successful
func (m *APIKeyAuth) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *APIKeyAuth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("apiKey", m.GetApiKey())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *APIKeyAuth) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApiKey sets the apiKey property value. Api key
func (m *APIKeyAuth) SetApiKey(value *string)() {
    m.apiKey = value
}
// SetTypeEscaped sets the type property value. type
func (m *APIKeyAuth) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type APIKeyAuthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiKey()(*string)
    GetTypeEscaped()(*string)
    SetApiKey(value *string)()
    SetTypeEscaped(value *string)()
}
