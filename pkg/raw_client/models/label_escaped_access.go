package models
// Access rights for the test. This defines the visibility of the Test in the UI
type Label_access int

const (
    PUBLIC_LABEL_ACCESS Label_access = iota
    PROTECTED_LABEL_ACCESS
    PRIVATE_LABEL_ACCESS
)

func (i Label_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseLabel_access(v string) (any, error) {
    result := PUBLIC_LABEL_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_LABEL_ACCESS
        case "PROTECTED":
            result = PROTECTED_LABEL_ACCESS
        case "PRIVATE":
            result = PRIVATE_LABEL_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLabel_access(values []Label_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Label_access) isMultiValue() bool {
    return false
}
