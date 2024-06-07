package models
// Access rights for the test. This defines the visibility of the Test in the UI
type ProtectedTimeType_access int

const (
    PUBLIC_PROTECTEDTIMETYPE_ACCESS ProtectedTimeType_access = iota
    PROTECTED_PROTECTEDTIMETYPE_ACCESS
    PRIVATE_PROTECTEDTIMETYPE_ACCESS
)

func (i ProtectedTimeType_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseProtectedTimeType_access(v string) (any, error) {
    result := PUBLIC_PROTECTEDTIMETYPE_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_PROTECTEDTIMETYPE_ACCESS
        case "PROTECTED":
            result = PROTECTED_PROTECTEDTIMETYPE_ACCESS
        case "PRIVATE":
            result = PRIVATE_PROTECTEDTIMETYPE_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProtectedTimeType_access(values []ProtectedTimeType_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProtectedTimeType_access) isMultiValue() bool {
    return false
}
