package db

import (
	"fmt"

	"sake/db/fieldtype"
	"sake/scheme"
	"sake/tools/nbutils"
)

type postgres struct{}

// var pgOperators = map[operator.Operator]string{
// 	operator.Equals:         "= ?",
// 	operator.NotEquals:      "!= ?",
// 	operator.Contains:       "LIKE ?",
// 	operator.NotContains:    "NOT LIKE ?",
// 	operator.Like:           "LIKE ?",
// 	operator.IContains:      "ILIKE ?",
// 	operator.NotIContains:   "NOT ILIKE ?",
// 	operator.ILike:          "ILIKE ?",
// 	operator.In:             "IN (?)",
// 	operator.NotIn:          "NOT IN (?)",
// 	operator.Lower:          "< ?",
// 	operator.LowerOrEqual:   "<= ?",
// 	operator.Greater:        "> ?",
// 	operator.GreaterOrEqual: ">= ?",
// }

// pg中还有的变量
var pgTypes = map[fieldtype.Type]string{
	fieldtype.Boolean:   "boolean",
	fieldtype.Char:      "character varying",
	fieldtype.Text:      "text",
	fieldtype.Date:      "date",
	fieldtype.DateTime:  "timestamp without time zone",
	fieldtype.Integer:   "integer",
	fieldtype.Float:     "numeric",
	fieldtype.HTML:      "text",
	fieldtype.Binary:    "bytea",
	fieldtype.Selection: "character varying",
	fieldtype.Many2One:  "integer",
	fieldtype.One2One:   "integer",
}

// connectionString returns the connection string for the given parameters
func (p *postgres) connectionString(params ConnectionParams) string {
	connectString := fmt.Sprintf("dbname=%s", params.DBName)
	if params.SSLMode != "" {
		connectString += fmt.Sprintf(" sslmode=%s", params.SSLMode)
	}
	if params.SSLCert != "" {
		connectString += fmt.Sprintf(" sslcert=%s", params.SSLCert)
	}
	if params.SSLKey != "" {
		connectString += fmt.Sprintf(" sslkey=%s", params.SSLKey)
	}
	if params.SSLCA != "" {
		connectString += fmt.Sprintf(" sslrootcert=%s", params.SSLCA)
	}
	if params.User != "" {
		connectString += fmt.Sprintf(" user=%s", params.User)
	}
	if params.Password != "" {
		connectString += fmt.Sprintf(" password=%s", params.Password)
	}
	if params.Host != "" {
		connectString += fmt.Sprintf(" host=%s", params.Host)
	}
	if params.Port != "" && params.Port != "5432" {
		connectString += fmt.Sprintf(" port=%s", params.Port)
	}
	return connectString
}

// 如何让operator转换成sql

// 获取field的类型
// typeSQL returns the sql type string for the given Field
func (p *postgres) typeSQL(f *scheme.Field) string {
	typ, _ := pgTypes[f.FieldType]
	return typ
}

// 获取field在数据库重的定义
func (p *postgres) columnSQLDefinition(f *scheme.Field, null bool) string {
	var res string
	typ, _ := pgTypes[f.FieldType]
	// typ, ok := pgTypes[f.FieldType]
	res = typ
	// if !ok {
	// 	log.Panic("Unknown column type", "type", f.fieldType, "model", f.model.name, "field", f.name)
	// }
	switch f.FieldType {
	case fieldtype.Char:
		if f.Size > 0 {
			res = fmt.Sprintf("%s(%d)", res, f.Size)
		}
	case fieldtype.Float:
		emptyD := nbutils.Digits{}
		if f.Digits != emptyD {
			res = fmt.Sprintf("numeric(%d, %d)", f.Digits.Precision, f.Digits.Scale)
		}
	}
	if p.fieldIsNotNull(f) && !null {
		res += " NOT NULL"
	}

	if f.Unique || f.FieldType == fieldtype.One2One {
		res += " UNIQUE"
	}
	return res
}

// 这里有个判断是否是一个 正确 的field是否是一个 null
// 为什么要这么去判断
// fieldIsNull returns true if the given Field results in a
// NOT NULL column in database.
func (p *postgres) fieldIsNotNull(f *scheme.Field) bool {
	if f.NotNull {
		return true
	}
	return false
}

// 关于获取table 的属性
// 现有获取所有tables name
// tables returns a map of table names of the database
func (p *postgres) tables() map[string]bool {
	var resList []string
	query := "SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema NOT IN ('pg_catalog', 'information_schema')"
	if err := db.Select(&resList, query); err != nil {
		log.Panic("Unable to get list of tables from database", "error", err)
	}
	res := make(map[string]bool, len(resList))
	for _, tableName := range resList {
		res[tableName] = true
	}
	return res
}

// quoteTableName returns the given table name with sql quotes
func (p *postgres) quoteTableName(tableName string) string {
	return fmt.Sprintf(`"%s"`, tableName)
}

// 获取所有columns
// columns returns a list of ColumnData for the given tableName
func (p *postgres) columns(tableName string) map[string]ColumnData {
	query := fmt.Sprintf(`
		SELECT column_name, data_type, is_nullable, column_default
		FROM information_schema.columns
		WHERE table_schema NOT IN ('pg_catalog', 'information_schema') AND table_name = '%s'
	`, tableName)
	var colData []ColumnData
	if err := db.Select(&colData, query); err != nil {
		log.Panic("Unable to get list of columns for table", "table", tableName, "error", err)
	}
	res := make(map[string]ColumnData, len(colData))
	for _, col := range colData {
		res[col.ColumnName] = col
	}
	return res
}

// 判断是否含有索引
// indexExists returns true if an index with the given name exists in the given table
func (p *postgres) indexExists(table string, name string) bool {
	query := fmt.Sprintf("SELECT COUNT(*) FROM pg_indexes WHERE tablename = '%s' AND indexname = '%s'", table, name)
	var cnt int
	dbGetNoTx(&cnt, query)
	return cnt > 0
}

// // 约束
// // constraintExists returns true if a constraint with the given name exists in the given table
// func (p *postgres) constraintExists(name string) bool {
// 	query := fmt.Sprintf("SELECT COUNT(*) FROM pg_constraint WHERE conname = '%s'", name)
// 	var cnt int
// 	dbGetNoTx(&cnt, query)
// 	return cnt > 0
// }
