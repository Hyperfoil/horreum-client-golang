package models
// Access rights for the test. This defines the visibility of the Test in the UI
type Test_access int

const (
    PUBLIC_TEST_ACCESS Test_access = iota
    PROTECTED_TEST_ACCESS
    PRIVATE_TEST_ACCESS
)

func (i Test_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseTest_access(v string) (any, error) {
    result := PUBLIC_TEST_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_TEST_ACCESS
        case "PROTECTED":
            result = PROTECTED_TEST_ACCESS
        case "PRIVATE":
            result = PRIVATE_TEST_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTest_access(values []Test_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Test_access) isMultiValue() bool {
    return false
}
