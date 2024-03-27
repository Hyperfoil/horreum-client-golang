package models
import (
    "errors"
)
// Type of backend datastore
type Datastore_type int

const (
    POSTGRES_DATASTORE_TYPE Datastore_type = iota
    ELASTICSEARCH_DATASTORE_TYPE
)

func (i Datastore_type) String() string {
    return []string{"POSTGRES", "ELASTICSEARCH"}[i]
}
func ParseDatastore_type(v string) (any, error) {
    result := POSTGRES_DATASTORE_TYPE
    switch v {
        case "POSTGRES":
            result = POSTGRES_DATASTORE_TYPE
        case "ELASTICSEARCH":
            result = ELASTICSEARCH_DATASTORE_TYPE
        default:
            return 0, errors.New("Unknown Datastore_type value: " + v)
    }
    return &result, nil
}
func SerializeDatastore_type(values []Datastore_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Datastore_type) isMultiValue() bool {
    return false
}
