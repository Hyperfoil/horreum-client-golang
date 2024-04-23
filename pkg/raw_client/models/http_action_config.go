package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HttpActionConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Action type
    typeEscaped *string
    // HTTP address
    url *string
}
// NewHttpActionConfig instantiates a new HttpActionConfig and sets the default values.
func NewHttpActionConfig()(*HttpActionConfig) {
    m := &HttpActionConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHttpActionConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHttpActionConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHttpActionConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HttpActionConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HttpActionConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetTypeEscaped gets the type property value. Action type
// returns a *string when successful
func (m *HttpActionConfig) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUrl gets the url property value. HTTP address
// returns a *string when successful
func (m *HttpActionConfig) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *HttpActionConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *HttpActionConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. Action type
func (m *HttpActionConfig) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. HTTP address
func (m *HttpActionConfig) SetUrl(value *string)() {
    m.url = value
}
type HttpActionConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*string)
    GetUrl()(*string)
    SetTypeEscaped(value *string)()
    SetUrl(value *string)()
}
