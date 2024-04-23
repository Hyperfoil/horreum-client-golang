package models

import (
    ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6 "strings"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Action struct {
    // The active property
    active *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The config property
    config Action_Action_configable
    // The event property
    event *string
    // The id property
    id *int32
    // The runAlways property
    runAlways *bool
    // The secrets property
    secrets Action_secretsable
    // The testId property
    testId *int32
    // The type property
    typeEscaped *string
}
// Action_Action_config composed type wrapper for classes GithubIssueCommentActionConfigable, GithubIssueCreateActionConfigable, HttpActionConfigable, SlackChannelMessageActionConfigable
type Action_Action_config struct {
    // Composed type representation for type GithubIssueCommentActionConfigable
    githubIssueCommentActionConfig GithubIssueCommentActionConfigable
    // Composed type representation for type GithubIssueCreateActionConfigable
    githubIssueCreateActionConfig GithubIssueCreateActionConfigable
    // Composed type representation for type HttpActionConfigable
    httpActionConfig HttpActionConfigable
    // Composed type representation for type SlackChannelMessageActionConfigable
    slackChannelMessageActionConfig SlackChannelMessageActionConfigable
}
// NewAction_Action_config instantiates a new Action_Action_config and sets the default values.
func NewAction_Action_config()(*Action_Action_config) {
    m := &Action_Action_config{
    }
    return m
}
// CreateAction_Action_configFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAction_Action_configFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAction_Action_config()
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
                if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "github-issue-comment") {
                    result.SetGithubIssueCommentActionConfig(NewGithubIssueCommentActionConfig())
                } else if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "github-issue-create") {
                    result.SetGithubIssueCreateActionConfig(NewGithubIssueCreateActionConfig())
                } else if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "http") {
                    result.SetHttpActionConfig(NewHttpActionConfig())
                } else if ie967d16dae74a49b5e0e051225c5dac0d76e5e38f13dd1628028cbce108c25b6.EqualFold(*mappingValue, "slack-channel-message") {
                    result.SetSlackChannelMessageActionConfig(NewSlackChannelMessageActionConfig())
                }
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Action_Action_config) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetGithubIssueCommentActionConfig gets the GithubIssueCommentActionConfig property value. Composed type representation for type GithubIssueCommentActionConfigable
// returns a GithubIssueCommentActionConfigable when successful
func (m *Action_Action_config) GetGithubIssueCommentActionConfig()(GithubIssueCommentActionConfigable) {
    return m.githubIssueCommentActionConfig
}
// GetGithubIssueCreateActionConfig gets the GithubIssueCreateActionConfig property value. Composed type representation for type GithubIssueCreateActionConfigable
// returns a GithubIssueCreateActionConfigable when successful
func (m *Action_Action_config) GetGithubIssueCreateActionConfig()(GithubIssueCreateActionConfigable) {
    return m.githubIssueCreateActionConfig
}
// GetHttpActionConfig gets the HttpActionConfig property value. Composed type representation for type HttpActionConfigable
// returns a HttpActionConfigable when successful
func (m *Action_Action_config) GetHttpActionConfig()(HttpActionConfigable) {
    return m.httpActionConfig
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *Action_Action_config) GetIsComposedType()(bool) {
    return true
}
// GetSlackChannelMessageActionConfig gets the SlackChannelMessageActionConfig property value. Composed type representation for type SlackChannelMessageActionConfigable
// returns a SlackChannelMessageActionConfigable when successful
func (m *Action_Action_config) GetSlackChannelMessageActionConfig()(SlackChannelMessageActionConfigable) {
    return m.slackChannelMessageActionConfig
}
// Serialize serializes information the current object
func (m *Action_Action_config) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGithubIssueCommentActionConfig() != nil {
        err := writer.WriteObjectValue("", m.GetGithubIssueCommentActionConfig())
        if err != nil {
            return err
        }
    } else if m.GetGithubIssueCreateActionConfig() != nil {
        err := writer.WriteObjectValue("", m.GetGithubIssueCreateActionConfig())
        if err != nil {
            return err
        }
    } else if m.GetHttpActionConfig() != nil {
        err := writer.WriteObjectValue("", m.GetHttpActionConfig())
        if err != nil {
            return err
        }
    } else if m.GetSlackChannelMessageActionConfig() != nil {
        err := writer.WriteObjectValue("", m.GetSlackChannelMessageActionConfig())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGithubIssueCommentActionConfig sets the GithubIssueCommentActionConfig property value. Composed type representation for type GithubIssueCommentActionConfigable
func (m *Action_Action_config) SetGithubIssueCommentActionConfig(value GithubIssueCommentActionConfigable)() {
    m.githubIssueCommentActionConfig = value
}
// SetGithubIssueCreateActionConfig sets the GithubIssueCreateActionConfig property value. Composed type representation for type GithubIssueCreateActionConfigable
func (m *Action_Action_config) SetGithubIssueCreateActionConfig(value GithubIssueCreateActionConfigable)() {
    m.githubIssueCreateActionConfig = value
}
// SetHttpActionConfig sets the HttpActionConfig property value. Composed type representation for type HttpActionConfigable
func (m *Action_Action_config) SetHttpActionConfig(value HttpActionConfigable)() {
    m.httpActionConfig = value
}
// SetSlackChannelMessageActionConfig sets the SlackChannelMessageActionConfig property value. Composed type representation for type SlackChannelMessageActionConfigable
func (m *Action_Action_config) SetSlackChannelMessageActionConfig(value SlackChannelMessageActionConfigable)() {
    m.slackChannelMessageActionConfig = value
}
type Action_Action_configable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGithubIssueCommentActionConfig()(GithubIssueCommentActionConfigable)
    GetGithubIssueCreateActionConfig()(GithubIssueCreateActionConfigable)
    GetHttpActionConfig()(HttpActionConfigable)
    GetSlackChannelMessageActionConfig()(SlackChannelMessageActionConfigable)
    SetGithubIssueCommentActionConfig(value GithubIssueCommentActionConfigable)()
    SetGithubIssueCreateActionConfig(value GithubIssueCreateActionConfigable)()
    SetHttpActionConfig(value HttpActionConfigable)()
    SetSlackChannelMessageActionConfig(value SlackChannelMessageActionConfigable)()
}
// NewAction instantiates a new Action and sets the default values.
func NewAction()(*Action) {
    m := &Action{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAction(), nil
}
// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *Action) GetActive()(*bool) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Action) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. The config property
// returns a Action_Action_configable when successful
func (m *Action) GetConfig()(Action_Action_configable) {
    return m.config
}
// GetEvent gets the event property value. The event property
// returns a *string when successful
func (m *Action) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Action) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAction_Action_configFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(Action_Action_configable))
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val)
        }
        return nil
    }
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
    res["runAlways"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAlways(val)
        }
        return nil
    }
    res["secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAction_secretsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecrets(val.(Action_secretsable))
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Action) GetId()(*int32) {
    return m.id
}
// GetRunAlways gets the runAlways property value. The runAlways property
// returns a *bool when successful
func (m *Action) GetRunAlways()(*bool) {
    return m.runAlways
}
// GetSecrets gets the secrets property value. The secrets property
// returns a Action_secretsable when successful
func (m *Action) GetSecrets()(Action_secretsable) {
    return m.secrets
}
// GetTestId gets the testId property value. The testId property
// returns a *int32 when successful
func (m *Action) GetTestId()(*int32) {
    return m.testId
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *Action) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Action) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("config", m.GetConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event", m.GetEvent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("runAlways", m.GetRunAlways())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secrets", m.GetSecrets())
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
// SetActive sets the active property value. The active property
func (m *Action) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Action) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. The config property
func (m *Action) SetConfig(value Action_Action_configable)() {
    m.config = value
}
// SetEvent sets the event property value. The event property
func (m *Action) SetEvent(value *string)() {
    m.event = value
}
// SetId sets the id property value. The id property
func (m *Action) SetId(value *int32)() {
    m.id = value
}
// SetRunAlways sets the runAlways property value. The runAlways property
func (m *Action) SetRunAlways(value *bool)() {
    m.runAlways = value
}
// SetSecrets sets the secrets property value. The secrets property
func (m *Action) SetSecrets(value Action_secretsable)() {
    m.secrets = value
}
// SetTestId sets the testId property value. The testId property
func (m *Action) SetTestId(value *int32)() {
    m.testId = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Action) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type Actionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetConfig()(Action_Action_configable)
    GetEvent()(*string)
    GetId()(*int32)
    GetRunAlways()(*bool)
    GetSecrets()(Action_secretsable)
    GetTestId()(*int32)
    GetTypeEscaped()(*string)
    SetActive(value *bool)()
    SetConfig(value Action_Action_configable)()
    SetEvent(value *string)()
    SetId(value *int32)()
    SetRunAlways(value *bool)()
    SetSecrets(value Action_secretsable)()
    SetTestId(value *int32)()
    SetTypeEscaped(value *string)()
}
