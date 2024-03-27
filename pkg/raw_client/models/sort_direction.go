package models
import (
    "errors"
)
type SortDirection int

const (
    ASCENDING_SORTDIRECTION SortDirection = iota
    DESCENDING_SORTDIRECTION
)

func (i SortDirection) String() string {
    return []string{"Ascending", "Descending"}[i]
}
func ParseSortDirection(v string) (any, error) {
    result := ASCENDING_SORTDIRECTION
    switch v {
        case "Ascending":
            result = ASCENDING_SORTDIRECTION
        case "Descending":
            result = DESCENDING_SORTDIRECTION
        default:
            return 0, errors.New("Unknown SortDirection value: " + v)
    }
    return &result, nil
}
func SerializeSortDirection(values []SortDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SortDirection) isMultiValue() bool {
    return false
}
