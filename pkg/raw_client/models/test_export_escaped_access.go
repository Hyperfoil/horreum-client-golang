package models
// Access rights for the test. This defines the visibility of the Test in the UI
type TestExport_access int

const (
    PUBLIC_TESTEXPORT_ACCESS TestExport_access = iota
    PROTECTED_TESTEXPORT_ACCESS
    PRIVATE_TESTEXPORT_ACCESS
)

func (i TestExport_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseTestExport_access(v string) (any, error) {
    result := PUBLIC_TESTEXPORT_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_TESTEXPORT_ACCESS
        case "PROTECTED":
            result = PROTECTED_TESTEXPORT_ACCESS
        case "PRIVATE":
            result = PRIVATE_TESTEXPORT_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTestExport_access(values []TestExport_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TestExport_access) isMultiValue() bool {
    return false
}
