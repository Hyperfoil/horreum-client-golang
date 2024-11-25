package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Datastore type of backend datastore
type Datastore struct {
    // Access rights for the test. This defines the visibility of the Test in the UI
    access *Datastore_access
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Is this a built-in datastore? Built-in datastores cannot be deleted or modified
    builtIn *bool
    // The config property
    config Datastore_Datastore_configable
    // Unique Datastore id
    id *int32
    // Name of the datastore, used to identify the datastore in the Test definition
    name *string
    // Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
    owner *string
    // Type of backend datastore
    typeEscaped *Datastore_type
}
// Datastore_Datastore_config composed type wrapper for classes ElasticsearchDatastoreConfigable, PostgresDatastoreConfigable
type Datastore_Datastore_config struct {
    // Composed type representation for type ElasticsearchDatastoreConfigable
    elasticsearchDatastoreConfig ElasticsearchDatastoreConfigable
    // Composed type representation for type PostgresDatastoreConfigable
    postgresDatastoreConfig PostgresDatastoreConfigable
}
// NewDatastore_Datastore_config instantiates a new Datastore_Datastore_config and sets the default values.
func NewDatastore_Datastore_config()(*Datastore_Datastore_config) {
    m := &Datastore_Datastore_config{
    }
    return m
}
// CreateDatastore_Datastore_configFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatastore_Datastore_configFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDatastore_Datastore_config()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetElasticsearchDatastoreConfig gets the ElasticsearchDatastoreConfig property value. Composed type representation for type ElasticsearchDatastoreConfigable
// returns a ElasticsearchDatastoreConfigable when successful
func (m *Datastore_Datastore_config) GetElasticsearchDatastoreConfig()(ElasticsearchDatastoreConfigable) {
    return m.elasticsearchDatastoreConfig
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Datastore_Datastore_config) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    if m.GetElasticsearchDatastoreConfig() != nil {
        return m.GetElasticsearchDatastoreConfig().GetFieldDeserializers()
    } else if m.GetPostgresDatastoreConfig() != nil {
        return m.GetPostgresDatastoreConfig().GetFieldDeserializers()
    }
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *Datastore_Datastore_config) GetIsComposedType()(bool) {
    return true
}
// GetPostgresDatastoreConfig gets the PostgresDatastoreConfig property value. Composed type representation for type PostgresDatastoreConfigable
// returns a PostgresDatastoreConfigable when successful
func (m *Datastore_Datastore_config) GetPostgresDatastoreConfig()(PostgresDatastoreConfigable) {
    return m.postgresDatastoreConfig
}
// Serialize serializes information the current object
func (m *Datastore_Datastore_config) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetElasticsearchDatastoreConfig() != nil {
        err := writer.WriteObjectValue("", m.GetElasticsearchDatastoreConfig())
        if err != nil {
            return err
        }
    } else if m.GetPostgresDatastoreConfig() != nil {
        err := writer.WriteObjectValue("", m.GetPostgresDatastoreConfig())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetElasticsearchDatastoreConfig sets the ElasticsearchDatastoreConfig property value. Composed type representation for type ElasticsearchDatastoreConfigable
func (m *Datastore_Datastore_config) SetElasticsearchDatastoreConfig(value ElasticsearchDatastoreConfigable)() {
    m.elasticsearchDatastoreConfig = value
}
// SetPostgresDatastoreConfig sets the PostgresDatastoreConfig property value. Composed type representation for type PostgresDatastoreConfigable
func (m *Datastore_Datastore_config) SetPostgresDatastoreConfig(value PostgresDatastoreConfigable)() {
    m.postgresDatastoreConfig = value
}
type Datastore_Datastore_configable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetElasticsearchDatastoreConfig()(ElasticsearchDatastoreConfigable)
    GetPostgresDatastoreConfig()(PostgresDatastoreConfigable)
    SetElasticsearchDatastoreConfig(value ElasticsearchDatastoreConfigable)()
    SetPostgresDatastoreConfig(value PostgresDatastoreConfigable)()
}
// NewDatastore instantiates a new Datastore and sets the default values.
func NewDatastore()(*Datastore) {
    m := &Datastore{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDatastoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatastoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatastore(), nil
}
// GetAccess gets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
// returns a *Datastore_access when successful
func (m *Datastore) GetAccess()(*Datastore_access) {
    return m.access
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Datastore) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuiltIn gets the builtIn property value. Is this a built-in datastore? Built-in datastores cannot be deleted or modified
// returns a *bool when successful
func (m *Datastore) GetBuiltIn()(*bool) {
    return m.builtIn
}
// GetConfig gets the config property value. The config property
// returns a Datastore_Datastore_configable when successful
func (m *Datastore) GetConfig()(Datastore_Datastore_configable) {
    return m.config
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Datastore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDatastore_access)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*Datastore_access))
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
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDatastore_Datastore_configFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(Datastore_Datastore_configable))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDatastore_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*Datastore_type))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique Datastore id
// returns a *int32 when successful
func (m *Datastore) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Name of the datastore, used to identify the datastore in the Test definition
// returns a *string when successful
func (m *Datastore) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
// returns a *string when successful
func (m *Datastore) GetOwner()(*string) {
    return m.owner
}
// GetTypeEscaped gets the type property value. Type of backend datastore
// returns a *Datastore_type when successful
func (m *Datastore) GetTypeEscaped()(*Datastore_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Datastore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccess() != nil {
        cast := (*m.GetAccess()).String()
        err := writer.WriteStringValue("access", &cast)
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
        err := writer.WriteObjectValue("config", m.GetConfig())
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
        err := writer.WriteStringValue("name", m.GetName())
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
// SetAccess sets the access property value. Access rights for the test. This defines the visibility of the Test in the UI
func (m *Datastore) SetAccess(value *Datastore_access)() {
    m.access = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Datastore) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuiltIn sets the builtIn property value. Is this a built-in datastore? Built-in datastores cannot be deleted or modified
func (m *Datastore) SetBuiltIn(value *bool)() {
    m.builtIn = value
}
// SetConfig sets the config property value. The config property
func (m *Datastore) SetConfig(value Datastore_Datastore_configable)() {
    m.config = value
}
// SetId sets the id property value. Unique Datastore id
func (m *Datastore) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Name of the datastore, used to identify the datastore in the Test definition
func (m *Datastore) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. Name of the team that owns the test. Users must belong to the team that owns a test to make modifications
func (m *Datastore) SetOwner(value *string)() {
    m.owner = value
}
// SetTypeEscaped sets the type property value. Type of backend datastore
func (m *Datastore) SetTypeEscaped(value *Datastore_type)() {
    m.typeEscaped = value
}
type Datastoreable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*Datastore_access)
    GetBuiltIn()(*bool)
    GetConfig()(Datastore_Datastore_configable)
    GetId()(*int32)
    GetName()(*string)
    GetOwner()(*string)
    GetTypeEscaped()(*Datastore_type)
    SetAccess(value *Datastore_access)()
    SetBuiltIn(value *bool)()
    SetConfig(value Datastore_Datastore_configable)()
    SetId(value *int32)()
    SetName(value *string)()
    SetOwner(value *string)()
    SetTypeEscaped(value *Datastore_type)()
}
