package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GithubIssueCreateActionConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Object markdown formatter
    formatter *string
    // GitHub repo owner
    owner *string
    // GitHub repo name
    repo *string
    // GitHub issue title
    title *string
    // Action type
    typeEscaped *string
}
// NewGithubIssueCreateActionConfig instantiates a new GithubIssueCreateActionConfig and sets the default values.
func NewGithubIssueCreateActionConfig()(*GithubIssueCreateActionConfig) {
    m := &GithubIssueCreateActionConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGithubIssueCreateActionConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGithubIssueCreateActionConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGithubIssueCreateActionConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GithubIssueCreateActionConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GithubIssueCreateActionConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
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
func (m *GithubIssueCreateActionConfig) GetFormatter()(*string) {
    return m.formatter
}
// GetOwner gets the owner property value. GitHub repo owner
// returns a *string when successful
func (m *GithubIssueCreateActionConfig) GetOwner()(*string) {
    return m.owner
}
// GetRepo gets the repo property value. GitHub repo name
// returns a *string when successful
func (m *GithubIssueCreateActionConfig) GetRepo()(*string) {
    return m.repo
}
// GetTitle gets the title property value. GitHub issue title
// returns a *string when successful
func (m *GithubIssueCreateActionConfig) GetTitle()(*string) {
    return m.title
}
// GetTypeEscaped gets the type property value. Action type
// returns a *string when successful
func (m *GithubIssueCreateActionConfig) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *GithubIssueCreateActionConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("formatter", m.GetFormatter())
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
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *GithubIssueCreateActionConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFormatter sets the formatter property value. Object markdown formatter
func (m *GithubIssueCreateActionConfig) SetFormatter(value *string)() {
    m.formatter = value
}
// SetOwner sets the owner property value. GitHub repo owner
func (m *GithubIssueCreateActionConfig) SetOwner(value *string)() {
    m.owner = value
}
// SetRepo sets the repo property value. GitHub repo name
func (m *GithubIssueCreateActionConfig) SetRepo(value *string)() {
    m.repo = value
}
// SetTitle sets the title property value. GitHub issue title
func (m *GithubIssueCreateActionConfig) SetTitle(value *string)() {
    m.title = value
}
// SetTypeEscaped sets the type property value. Action type
func (m *GithubIssueCreateActionConfig) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type GithubIssueCreateActionConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormatter()(*string)
    GetOwner()(*string)
    GetRepo()(*string)
    GetTitle()(*string)
    GetTypeEscaped()(*string)
    SetFormatter(value *string)()
    SetOwner(value *string)()
    SetRepo(value *string)()
    SetTitle(value *string)()
    SetTypeEscaped(value *string)()
}
