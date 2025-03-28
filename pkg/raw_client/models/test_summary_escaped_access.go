package models
// Access rights for the test. This defines the visibility of the Test in the UI
type TestSummary_access int

const (
    PUBLIC_TESTSUMMARY_ACCESS TestSummary_access = iota
    PROTECTED_TESTSUMMARY_ACCESS
    PRIVATE_TESTSUMMARY_ACCESS
)

func (i TestSummary_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseTestSummary_access(v string) (any, error) {
    result := PUBLIC_TESTSUMMARY_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_TESTSUMMARY_ACCESS
        case "PROTECTED":
            result = PROTECTED_TESTSUMMARY_ACCESS
        case "PRIVATE":
            result = PRIVATE_TESTSUMMARY_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTestSummary_access(values []TestSummary_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TestSummary_access) isMultiValue() bool {
    return false
}
