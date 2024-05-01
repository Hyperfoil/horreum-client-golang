package models
import (
    "errors"
)
type RelativeDifferenceDetectionConfig_model int

const (
    RELATIVEDIFFERENCE_RELATIVEDIFFERENCEDETECTIONCONFIG_MODEL RelativeDifferenceDetectionConfig_model = iota
)

func (i RelativeDifferenceDetectionConfig_model) String() string {
    return []string{"relativeDifference"}[i]
}
func ParseRelativeDifferenceDetectionConfig_model(v string) (any, error) {
    result := RELATIVEDIFFERENCE_RELATIVEDIFFERENCEDETECTIONCONFIG_MODEL
    switch v {
        case "relativeDifference":
            result = RELATIVEDIFFERENCE_RELATIVEDIFFERENCEDETECTIONCONFIG_MODEL
        default:
            return 0, errors.New("Unknown RelativeDifferenceDetectionConfig_model value: " + v)
    }
    return &result, nil
}
func SerializeRelativeDifferenceDetectionConfig_model(values []RelativeDifferenceDetectionConfig_model) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RelativeDifferenceDetectionConfig_model) isMultiValue() bool {
    return false
}
