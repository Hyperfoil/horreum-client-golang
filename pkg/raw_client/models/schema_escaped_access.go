package models
// Access rights for the test. This defines the visibility of the Test in the UI
type Schema_access int

const (
    PUBLIC_SCHEMA_ACCESS Schema_access = iota
    PROTECTED_SCHEMA_ACCESS
    PRIVATE_SCHEMA_ACCESS
)

func (i Schema_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseSchema_access(v string) (any, error) {
    result := PUBLIC_SCHEMA_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_SCHEMA_ACCESS
        case "PROTECTED":
            result = PROTECTED_SCHEMA_ACCESS
        case "PRIVATE":
            result = PRIVATE_SCHEMA_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSchema_access(values []Schema_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Schema_access) isMultiValue() bool {
    return false
}
