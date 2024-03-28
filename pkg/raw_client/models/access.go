package models
import (
    "errors"
)
// Resources have different visibility within the UI. 'PUBLIC', 'PROTECTED' and 'PRIVATE'. Restricted resources are not visible to users who do not have the correct permissions
type Access int

const (
    PUBLIC_ACCESS Access = iota
    PROTECTED_ACCESS
    PRIVATE_ACCESS
)

func (i Access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseAccess(v string) (any, error) {
    result := PUBLIC_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_ACCESS
        case "PROTECTED":
            result = PROTECTED_ACCESS
        case "PRIVATE":
            result = PRIVATE_ACCESS
        default:
            return 0, errors.New("Unknown Access value: " + v)
    }
    return &result, nil
}
func SerializeAccess(values []Access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Access) isMultiValue() bool {
    return false
}
