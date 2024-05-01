package models
import (
    "errors"
)
type FixedThresholdDetectionConfig_model int

const (
    FIXEDTHRESHOLD_FIXEDTHRESHOLDDETECTIONCONFIG_MODEL FixedThresholdDetectionConfig_model = iota
)

func (i FixedThresholdDetectionConfig_model) String() string {
    return []string{"fixedThreshold"}[i]
}
func ParseFixedThresholdDetectionConfig_model(v string) (any, error) {
    result := FIXEDTHRESHOLD_FIXEDTHRESHOLDDETECTIONCONFIG_MODEL
    switch v {
        case "fixedThreshold":
            result = FIXEDTHRESHOLD_FIXEDTHRESHOLDDETECTIONCONFIG_MODEL
        default:
            return 0, errors.New("Unknown FixedThresholdDetectionConfig_model value: " + v)
    }
    return &result, nil
}
func SerializeFixedThresholdDetectionConfig_model(values []FixedThresholdDetectionConfig_model) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FixedThresholdDetectionConfig_model) isMultiValue() bool {
    return false
}
