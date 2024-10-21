package models
type KeyType int

const (
    USER_KEYTYPE KeyType = iota
)

func (i KeyType) String() string {
    return []string{"USER"}[i]
}
func ParseKeyType(v string) (any, error) {
    result := USER_KEYTYPE
    switch v {
        case "USER":
            result = USER_KEYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeKeyType(values []KeyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i KeyType) isMultiValue() bool {
    return false
}
