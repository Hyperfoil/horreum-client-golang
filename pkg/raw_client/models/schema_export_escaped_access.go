package models
// Access rights for the test. This defines the visibility of the Test in the UI
type SchemaExport_access int

const (
    PUBLIC_SCHEMAEXPORT_ACCESS SchemaExport_access = iota
    PROTECTED_SCHEMAEXPORT_ACCESS
    PRIVATE_SCHEMAEXPORT_ACCESS
)

func (i SchemaExport_access) String() string {
    return []string{"PUBLIC", "PROTECTED", "PRIVATE"}[i]
}
func ParseSchemaExport_access(v string) (any, error) {
    result := PUBLIC_SCHEMAEXPORT_ACCESS
    switch v {
        case "PUBLIC":
            result = PUBLIC_SCHEMAEXPORT_ACCESS
        case "PROTECTED":
            result = PROTECTED_SCHEMAEXPORT_ACCESS
        case "PRIVATE":
            result = PRIVATE_SCHEMAEXPORT_ACCESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSchemaExport_access(values []SchemaExport_access) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SchemaExport_access) isMultiValue() bool {
    return false
}
