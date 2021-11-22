package dialect

// import (
// 	"fmt"
// 	"log"
// 	"sake/scheme"
// 	"strings"
// )

// type postgres struct{}

// // 这里考虑下这个函数的具体作用
// var _ Dialect = (*postgres)(nil)

// func init() {
// 	RegisterDialect("pgx", &postgres{})
// }

// func (p *postgres) TableExistSQL(tableName string) (string, []interface{}) {
// 	args := []interface{}{tableName}
// 	return "SELECT relname FROM pg_class where relname = '?'", args
// }

// func (p *postgres) DataTypeOf(typ string) string {
// 	switch typ {
// 	case "bool":
// 		return "bool"
// 	case "int":
// 		return "integer"
// 	case "int64", "uint64":
// 		return "bigint"
// 	case "float", "float64":
// 		return "real"
// 	case "string", "text":
// 		return "text"
// 	// case reflect.Array, reflect.Slice:
// 	// 	return "blob"
// 	case "datetime":
// 		return "datetime"
// 	}

// 	panic(fmt.Sprintf("invalid sql type %s", typ))
// }

// // columnSQLDefinition 使field字段转换成一个sql语句 FieldConvertSQL
// // If null is true, then the column will be nullable, whatever the field defines
// func (p *postgres) columnSQLDefinition(field *scheme.Field) string {
// 	var result []string
// 	result = append(result, field.DBName)

// 	if field.PrimaryKey == true {
// 		if field.AutoIncrement == true {
// 			result = append(result, "SERIAL")
// 		} else {
// 			result = append(result, p.DataTypeOf(field.Type))
// 		}
// 		result = append(result, "PRIMARY KEY")
// 	} else {
// 		result = append(result, p.DataTypeOf(field.Type))
// 	}

// 	if field.NotNull == true {
// 		result = append(result, "NOT NULL")
// 	}

// 	if field.Unique == true {
// 		result = append(result, field.DBName)
// 	}

// 	if field.Check != "" {
// 		result = append(result, fmt.Sprintf(`CHECK(%s)`, field.Check))
// 	}

// 	if field.Comment != "" {
// 		result = append(result, fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s"."%s" IS "%s"`, field.Scheme.Schema, field.Scheme.DBName, field.DBName, field.Comment))
// 	}

// 	return strings.Join(result, " ")
// }

// // tables returns a map of table names of the database
// func (p *postgres) tables() map[string]bool {
// 	var resList []string
// 	query := "SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema NOT IN ('pg_catalog', 'information_schema')"
// 	if err := db.Select(&resList, query); err != nil {
// 		log.Panic("Unable to get list of tables from database", "error", err)
// 	}
// 	res := make(map[string]bool, len(resList))
// 	for _, tableName := range resList {
// 		res[tableName] = true
// 	}
// 	return res
// }

// // quoteTableName returns the given table name with sql quotes
// func (d *postgresAdapter) quoteTableName(tableName string) string {
// 	return fmt.Sprintf(`"%s"`, tableName)
// }

// // columns returns a list of ColumnData for the given tableName
// func (d *postgresAdapter) columns(tableName string) map[string]ColumnData {
// 	query := fmt.Sprintf(`
// 		SELECT column_name, data_type, is_nullable, column_default
// 		FROM information_schema.columns
// 		WHERE table_schema NOT IN ('pg_catalog', 'information_schema') AND table_name = '%s'
// 	`, tableName)
// 	var colData []ColumnData
// 	if err := db.Select(&colData, query); err != nil {
// 		log.Panic("Unable to get list of columns for table", "table", tableName, "error", err)
// 	}
// 	res := make(map[string]ColumnData, len(colData))
// 	for _, col := range colData {
// 		res[col.ColumnName] = col
// 	}
// 	return res
// }
