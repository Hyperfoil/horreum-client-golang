package models
// Access rights for the test. This defines the visibility of the Test in the UI
type RunExtended_access int

const (
    PUBLIC_RUNEXTENDED_ACCESS RunExtended_access = iota
    PROTECTED_RUNEXTENDED_ACCESS
    PRIVATE_RUNEXTENDED_ACCESS
)

func (i RunExtended_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseRunExtended_access(v string) (any, error) {
    result := PUBLIC_RUNEXTENDED_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_RUNEXTENDED_ACCESS
        case "PROTECTED":
            result = PROTECTED_RUNEXTENDED_ACCESS
        case "PRIVATE":
            result = PRIVATE_RUNEXTENDED_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRunExtended_access(values []RunExtended_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RunExtended_access) isMultiValue() bool {
    return false
}
