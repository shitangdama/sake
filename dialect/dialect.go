package dialect

// import "sake/scheme"

// var dialectsMap = map[string]Dialect{}

// // Dialect is interface
// type Dialect interface {
// 	// 这个地方改成由，map给的定义
// 	DataTypeOf(typ string) string
// 	TableExistSQL(tableName string) (string, []interface{})

// 	// Convert(sql string) string
// 	ParseFieldInfo(field *scheme.Field) string

// 	// CreateTable(scheme *scheme.Scheme) string
// 	// UpdateTable(scheme *scheme.Scheme) string
// 	// DropTable(name string) string
// 	// GetTable(scheme *scheme.Scheme) string

// 	// CreateColumn(field *scheme.Field) string
// 	// UpdateColumn(field *scheme.Field, c *ColumnUpdateInfo) string
// 	// DeleteColumn(field *scheme.Field) string
// 	// GetColumn(field *scheme.Field) string
// }

// // RegisterDialect is xx
// func RegisterDialect(name string, dialect Dialect) {
// 	dialectsMap[name] = dialect
// }

// // GetDialect is xxx
// func GetDialect(name string) (dialect Dialect, ok bool) {
// 	dialect, ok = dialectsMap[name]
// 	return
// }
