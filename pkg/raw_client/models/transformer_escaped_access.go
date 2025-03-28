package models
// Access rights for the test. This defines the visibility of the Test in the UI
type Transformer_access int

const (
    PUBLIC_TRANSFORMER_ACCESS Transformer_access = iota
    PROTECTED_TRANSFORMER_ACCESS
    PRIVATE_TRANSFORMER_ACCESS
)

func (i Transformer_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseTransformer_access(v string) (any, error) {
    result := PUBLIC_TRANSFORMER_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_TRANSFORMER_ACCESS
        case "PROTECTED":
            result = PROTECTED_TRANSFORMER_ACCESS
        case "PRIVATE":
            result = PRIVATE_TRANSFORMER_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransformer_access(values []Transformer_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Transformer_access) isMultiValue() bool {
    return false
}
