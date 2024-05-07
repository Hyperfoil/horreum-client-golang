package models
import (
    "errors"
)
// Access rights for the test. This defines the visibility of the Test in the UI
type Dataset_access int

const (
    PUBLIC_DATASET_ACCESS Dataset_access = iota
    PROTECTED_DATASET_ACCESS
    PRIVATE_DATASET_ACCESS
)

func (i Dataset_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseDataset_access(v string) (any, error) {
    result := PUBLIC_DATASET_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_DATASET_ACCESS
        case "PROTECTED":
            result = PROTECTED_DATASET_ACCESS
        case "PRIVATE":
            result = PRIVATE_DATASET_ACCESS
        default:
            return 0, errors.New("Unknown Dataset_access value: " + v)
    }
    return &result, nil
}
func SerializeDataset_access(values []Dataset_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Dataset_access) isMultiValue() bool {
    return false
}
