package models
import (
    "errors"
)
// Access rights for the test. This defines the visibility of the Test in the UI
type Datastore_access int

const (
    PUBLIC_DATASTORE_ACCESS Datastore_access = iota
    PROTECTED_DATASTORE_ACCESS
    PRIVATE_DATASTORE_ACCESS
)

func (i Datastore_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseDatastore_access(v string) (any, error) {
    result := PUBLIC_DATASTORE_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_DATASTORE_ACCESS
        case "PROTECTED":
            result = PROTECTED_DATASTORE_ACCESS
        case "PRIVATE":
            result = PRIVATE_DATASTORE_ACCESS
        default:
            return 0, errors.New("Unknown Datastore_access value: " + v)
    }
    return &result, nil
}
func SerializeDatastore_access(values []Datastore_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Datastore_access) isMultiValue() bool {
    return false
}
