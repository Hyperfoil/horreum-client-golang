package models
// Access rights for the test. This defines the visibility of the Test in the UI
type Run_access int

const (
    PUBLIC_RUN_ACCESS Run_access = iota
    PROTECTED_RUN_ACCESS
    PRIVATE_RUN_ACCESS
)

func (i Run_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseRun_access(v string) (any, error) {
    result := PUBLIC_RUN_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_RUN_ACCESS
        case "PROTECTED":
            result = PROTECTED_RUN_ACCESS
        case "PRIVATE":
            result = PRIVATE_RUN_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRun_access(values []Run_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Run_access) isMultiValue() bool {
    return false
}
