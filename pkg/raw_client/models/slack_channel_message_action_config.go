package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SlackChannelMessageActionConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Slack channel
    channel *string
    // Object markdown formatter
    formatter *string
    // Action type
    typeEscaped *string
}
// NewSlackChannelMessageActionConfig instantiates a new SlackChannelMessageActionConfig and sets the default values.
func NewSlackChannelMessageActionConfig()(*SlackChannelMessageActionConfig) {
    m := &SlackChannelMessageActionConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSlackChannelMessageActionConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSlackChannelMessageActionConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSlackChannelMessageActionConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SlackChannelMessageActionConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChannel gets the channel property value. Slack channel
// returns a *string when successful
func (m *SlackChannelMessageActionConfig) GetChannel()(*string) {
    return m.channel
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SlackChannelMessageActionConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["channel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannel(val)
        }
        return nil
    }
    res["formatter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormatter(val)
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
// GetFormatter gets the formatter property value. Object markdown formatter
// returns a *string when successful
func (m *SlackChannelMessageActionConfig) GetFormatter()(*string) {
    return m.formatter
}
// GetTypeEscaped gets the type property value. Action type
// returns a *string when successful
func (m *SlackChannelMessageActionConfig) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SlackChannelMessageActionConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("channel", m.GetChannel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("formatter", m.GetFormatter())
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
func (m *SlackChannelMessageActionConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChannel sets the channel property value. Slack channel
func (m *SlackChannelMessageActionConfig) SetChannel(value *string)() {
    m.channel = value
}
// SetFormatter sets the formatter property value. Object markdown formatter
func (m *SlackChannelMessageActionConfig) SetFormatter(value *string)() {
    m.formatter = value
}
// SetTypeEscaped sets the type property value. Action type
func (m *SlackChannelMessageActionConfig) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type SlackChannelMessageActionConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannel()(*string)
    GetFormatter()(*string)
    GetTypeEscaped()(*string)
    SetChannel(value *string)()
    SetFormatter(value *string)()
    SetTypeEscaped(value *string)()
}
