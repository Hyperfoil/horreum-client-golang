package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1 "github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
)

type UserApikeyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
    // The type property
    typeEscaped *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.KeyType
}
// NewUserApikeyPostRequestBody instantiates a new UserApikeyPostRequestBody and sets the default values.
func NewUserApikeyPostRequestBody()(*UserApikeyPostRequestBody) {
    m := &UserApikeyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserApikeyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserApikeyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserApikeyPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserApikeyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserApikeyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.ParseKeyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.KeyType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *UserApikeyPostRequestBody) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the type property value. The type property
// returns a *KeyType when successful
func (m *UserApikeyPostRequestBody) GetTypeEscaped()(*i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.KeyType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *UserApikeyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *UserApikeyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *UserApikeyPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *UserApikeyPostRequestBody) SetTypeEscaped(value *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.KeyType)() {
    m.typeEscaped = value
}
type UserApikeyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetTypeEscaped()(*i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.KeyType)
    SetName(value *string)()
    SetTypeEscaped(value *i24479a9d05b05b7c1efaeda9ae24aee51c8acc6f59ee3190ae7f0941a410c8a1.KeyType)()
}
