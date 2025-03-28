package models
// UI Component type
type ConditionComponent_type int

const (
    LOG_SLIDER_CONDITIONCOMPONENT_TYPE ConditionComponent_type = iota
    ENUM_CONDITIONCOMPONENT_TYPE
    NUMBER_BOUND_CONDITIONCOMPONENT_TYPE
    SWITCH_CONDITIONCOMPONENT_TYPE
)

func (i ConditionComponent_type) String() string {
    return []string{"LOG_SLIDER", "ENUM", "NUMBER_BOUND", "SWITCH"}[i]
}
func ParseConditionComponent_type(v string) (any, error) {
    result := LOG_SLIDER_CONDITIONCOMPONENT_TYPE
    switch v {
        case "LOG_SLIDER":
            result = LOG_SLIDER_CONDITIONCOMPONENT_TYPE
        case "ENUM":
            result = ENUM_CONDITIONCOMPONENT_TYPE
        case "NUMBER_BOUND":
            result = NUMBER_BOUND_CONDITIONCOMPONENT_TYPE
        case "SWITCH":
            result = SWITCH_CONDITIONCOMPONENT_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConditionComponent_type(values []ConditionComponent_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConditionComponent_type) isMultiValue() bool {
    return false
}
