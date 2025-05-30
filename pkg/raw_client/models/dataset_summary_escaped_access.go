// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Access rights for the test. This defines the visibility of the Test in the UI
type DatasetSummary_access int

const (
    PUBLIC_DATASETSUMMARY_ACCESS DatasetSummary_access = iota
    PROTECTED_DATASETSUMMARY_ACCESS
    PRIVATE_DATASETSUMMARY_ACCESS
)

func (i DatasetSummary_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseDatasetSummary_access(v string) (any, error) {
    result := PUBLIC_DATASETSUMMARY_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_DATASETSUMMARY_ACCESS
        case "PROTECTED":
            result = PROTECTED_DATASETSUMMARY_ACCESS
        case "PRIVATE":
            result = PRIVATE_DATASETSUMMARY_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatasetSummary_access(values []DatasetSummary_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatasetSummary_access) isMultiValue() bool {
    return false
}
