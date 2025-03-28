package models
// Access rights for the test. This defines the visibility of the Test in the UI
type RunSummary_access int

const (
    PUBLIC_RUNSUMMARY_ACCESS RunSummary_access = iota
    PROTECTED_RUNSUMMARY_ACCESS
    PRIVATE_RUNSUMMARY_ACCESS
)

func (i RunSummary_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseRunSummary_access(v string) (any, error) {
    result := PUBLIC_RUNSUMMARY_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_RUNSUMMARY_ACCESS
        case "PROTECTED":
            result = PROTECTED_RUNSUMMARY_ACCESS
        case "PRIVATE":
            result = PRIVATE_RUNSUMMARY_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRunSummary_access(values []RunSummary_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RunSummary_access) isMultiValue() bool {
    return false
}
