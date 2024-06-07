package models
// Access rights for the test. This defines the visibility of the Test in the UI
type ProtectedType_access int

const (
    PUBLIC_PROTECTEDTYPE_ACCESS ProtectedType_access = iota
    PROTECTED_PROTECTEDTYPE_ACCESS
    PRIVATE_PROTECTEDTYPE_ACCESS
)

func (i ProtectedType_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseProtectedType_access(v string) (any, error) {
    result := PUBLIC_PROTECTEDTYPE_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_PROTECTEDTYPE_ACCESS
        case "PROTECTED":
            result = PROTECTED_PROTECTEDTYPE_ACCESS
        case "PRIVATE":
            result = PRIVATE_PROTECTEDTYPE_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProtectedType_access(values []ProtectedType_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProtectedType_access) isMultiValue() bool {
    return false
}
