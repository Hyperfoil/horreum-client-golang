package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GithubIssueCommentActionConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Object markdown formatter
    formatter *string
    // GitHub issue number
    issue *string
    // GitHub issue URL
    issueUrl *string
    // GitHub repo owner
    owner *string
    // GitHub repo name
    repo *string
    // Action type
    typeEscaped *string
}
// NewGithubIssueCommentActionConfig instantiates a new GithubIssueCommentActionConfig and sets the default values.
func NewGithubIssueCommentActionConfig()(*GithubIssueCommentActionConfig) {
    m := &GithubIssueCommentActionConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGithubIssueCommentActionConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGithubIssueCommentActionConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGithubIssueCommentActionConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GithubIssueCommentActionConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GithubIssueCommentActionConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["issue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssue(val)
        }
        return nil
    }
    res["issueUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueUrl(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["repo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepo(val)
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
func (m *GithubIssueCommentActionConfig) GetFormatter()(*string) {
    return m.formatter
}
// GetIssue gets the issue property value. GitHub issue number
// returns a *string when successful
func (m *GithubIssueCommentActionConfig) GetIssue()(*string) {
    return m.issue
}
// GetIssueUrl gets the issueUrl property value. GitHub issue URL
// returns a *string when successful
func (m *GithubIssueCommentActionConfig) GetIssueUrl()(*string) {
    return m.issueUrl
}
// GetOwner gets the owner property value. GitHub repo owner
// returns a *string when successful
func (m *GithubIssueCommentActionConfig) GetOwner()(*string) {
    return m.owner
}
// GetRepo gets the repo property value. GitHub repo name
// returns a *string when successful
func (m *GithubIssueCommentActionConfig) GetRepo()(*string) {
    return m.repo
}
// GetTypeEscaped gets the type property value. Action type
// returns a *string when successful
func (m *GithubIssueCommentActionConfig) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *GithubIssueCommentActionConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("formatter", m.GetFormatter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issue", m.GetIssue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issueUrl", m.GetIssueUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repo", m.GetRepo())
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
func (m *GithubIssueCommentActionConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFormatter sets the formatter property value. Object markdown formatter
func (m *GithubIssueCommentActionConfig) SetFormatter(value *string)() {
    m.formatter = value
}
// SetIssue sets the issue property value. GitHub issue number
func (m *GithubIssueCommentActionConfig) SetIssue(value *string)() {
    m.issue = value
}
// SetIssueUrl sets the issueUrl property value. GitHub issue URL
func (m *GithubIssueCommentActionConfig) SetIssueUrl(value *string)() {
    m.issueUrl = value
}
// SetOwner sets the owner property value. GitHub repo owner
func (m *GithubIssueCommentActionConfig) SetOwner(value *string)() {
    m.owner = value
}
// SetRepo sets the repo property value. GitHub repo name
func (m *GithubIssueCommentActionConfig) SetRepo(value *string)() {
    m.repo = value
}
// SetTypeEscaped sets the type property value. Action type
func (m *GithubIssueCommentActionConfig) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type GithubIssueCommentActionConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormatter()(*string)
    GetIssue()(*string)
    GetIssueUrl()(*string)
    GetOwner()(*string)
    GetRepo()(*string)
    GetTypeEscaped()(*string)
    SetFormatter(value *string)()
    SetIssue(value *string)()
    SetIssueUrl(value *string)()
    SetOwner(value *string)()
    SetRepo(value *string)()
    SetTypeEscaped(value *string)()
}
