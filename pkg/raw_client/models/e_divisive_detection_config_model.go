package models
import (
    "errors"
)
type EDivisiveDetectionConfig_model int

const (
    EDIVISIVE_EDIVISIVEDETECTIONCONFIG_MODEL EDivisiveDetectionConfig_model = iota
)

func (i EDivisiveDetectionConfig_model) String() string {
    return []string{"eDivisive"}[i]
}
func ParseEDivisiveDetectionConfig_model(v string) (any, error) {
    result := EDIVISIVE_EDIVISIVEDETECTIONCONFIG_MODEL
    switch v {
        case "eDivisive":
            result = EDIVISIVE_EDIVISIVEDETECTIONCONFIG_MODEL
        default:
            return 0, errors.New("Unknown EDivisiveDetectionConfig_model value: " + v)
    }
    return &result, nil
}
func SerializeEDivisiveDetectionConfig_model(values []EDivisiveDetectionConfig_model) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EDivisiveDetectionConfig_model) isMultiValue() bool {
    return false
}
