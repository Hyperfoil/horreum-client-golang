package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiKeyResponse struct {
    // The access property
    access *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The creation property
    creation *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The id property
    id *int64
    // The isRevoked property
    isRevoked *bool
    // The name property
    name *string
    // The toExpiration property
    toExpiration *int64
    // The type property
    typeEscaped *KeyType
}
// NewApiKeyResponse instantiates a new ApiKeyResponse and sets the default values.
func NewApiKeyResponse()(*ApiKeyResponse) {
    m := &ApiKeyResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiKeyResponse(), nil
}
// GetAccess gets the access property value. The access property
// returns a *Time when successful
func (m *ApiKeyResponse) GetAccess()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiKeyResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreation gets the creation property value. The creation property
// returns a *Time when successful
func (m *ApiKeyResponse) GetCreation()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.creation
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiKeyResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val)
        }
        return nil
    }
    res["creation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreation(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isRevoked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRevoked(val)
        }
        return nil
    }
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
    res["toExpiration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToExpiration(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*KeyType))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ApiKeyResponse) GetId()(*int64) {
    return m.id
}
// GetIsRevoked gets the isRevoked property value. The isRevoked property
// returns a *bool when successful
func (m *ApiKeyResponse) GetIsRevoked()(*bool) {
    return m.isRevoked
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ApiKeyResponse) GetName()(*string) {
    return m.name
}
// GetToExpiration gets the toExpiration property value. The toExpiration property
// returns a *int64 when successful
func (m *ApiKeyResponse) GetToExpiration()(*int64) {
    return m.toExpiration
}
// GetTypeEscaped gets the type property value. The type property
// returns a *KeyType when successful
func (m *ApiKeyResponse) GetTypeEscaped()(*KeyType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ApiKeyResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("access", m.GetAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("creation", m.GetCreation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRevoked", m.GetIsRevoked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("toExpiration", m.GetToExpiration())
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
// SetAccess sets the access property value. The access property
func (m *ApiKeyResponse) SetAccess(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiKeyResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreation sets the creation property value. The creation property
func (m *ApiKeyResponse) SetCreation(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creation = value
}
// SetId sets the id property value. The id property
func (m *ApiKeyResponse) SetId(value *int64)() {
    m.id = value
}
// SetIsRevoked sets the isRevoked property value. The isRevoked property
func (m *ApiKeyResponse) SetIsRevoked(value *bool)() {
    m.isRevoked = value
}
// SetName sets the name property value. The name property
func (m *ApiKeyResponse) SetName(value *string)() {
    m.name = value
}
// SetToExpiration sets the toExpiration property value. The toExpiration property
func (m *ApiKeyResponse) SetToExpiration(value *int64)() {
    m.toExpiration = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *ApiKeyResponse) SetTypeEscaped(value *KeyType)() {
    m.typeEscaped = value
}
type ApiKeyResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreation()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int64)
    GetIsRevoked()(*bool)
    GetName()(*string)
    GetToExpiration()(*int64)
    GetTypeEscaped()(*KeyType)
    SetAccess(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreation(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int64)()
    SetIsRevoked(value *bool)()
    SetName(value *string)()
    SetToExpiration(value *int64)()
    SetTypeEscaped(value *KeyType)()
}
