package migration

type Schema string

const (
	Public   Schema = "public"
	Academic Schema = "academic"
	Administration Schema = "administration"
)

func GenerateTableName(schema Schema, tableName string) string {
	return string(schema) + "." + tableName
}
