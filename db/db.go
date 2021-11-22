package db

import (
	"database/sql"
	"fmt"
	"sake/scheme"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// 旧的文件是在session里面， 这里转出来，

var (
	db         *sqlx.DB
	connParams ConnectionParams
	adapters   map[string]Dialect
)

// ConnectionParams are the database agnostic parameters to connect to the database
type ConnectionParams struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	SSLCert  string
	SSLKey   string
	SSLCA    string
}

// ConnectionString returns the connection string for these connection params
func (cp ConnectionParams) ConnectionString() string {
	adapter := adapters[cp.Driver]
	return adapter.connectionString(cp)
}

// A ColumnData holds information from the db schema about one column
type ColumnData struct {
	ColumnName    string         `db:"column_name"`
	DataType      string         `db:"data_type"`
	IsNullable    string         `db:"is_nullable"`
	ColumnDefault sql.NullString `db:"column_default"`
}

// 这里还是考虑下加入schema
// Dialect xxxx
type Dialect interface {
	// connectionString returns the connection string for the given parameters
	connectionString(ConnectionParams) string

	// // operatorSQL returns the sql string and placeholders for the given DomainOperator
	// operatorSQL(operator.Operator, interface{}) (string, interface{})

	// typeSQL returns the SQL type string, including columns constraints if any
	typeSQL(fi *scheme.Field) string
	// columnSQLDefinition returns the SQL type string, including columns constraints if any
	// If null is true, then the column will be nullable, whatever the field defines
	columnSQLDefinition(fi *scheme.Field, null bool) string
	// tables returns a map of table names of the database
	tables() map[string]bool
	// columns returns a list of ColumnData for the given tableName
	columns(tableName string) map[string]ColumnData
	// fieldIsNull returns true if the given Field results in a
	// NOT NULL column in database.
	// fieldIsNotNull(fi *Field) bool
	// // quoteTableName returns the given table name with sql quotes
	// quoteTableName(string) string
	// indexExists returns true if an index with the given name exists in the given table
	indexExists(table string, name string) bool
	// // constraintExists returns true if a constraint with the given name exists
	// constraintExists(name string) bool
	// // constraints returns a list of all constraints matching the given SQL pattern
	// constraints(pattern string) []string
	// // setTransactionIsolation returns the SQL string to set the transaction isolation
	// // level to serializable
	// setTransactionIsolation() string
	// // createSequence creates a DB sequence with the given name
	// createSequence(name string, increment, start int64)
	// // dropSequence drop the DB sequence with the given name
	// dropSequence(name string)
	// // alterSequence modifies the DB sequence given by name
	// alterSequence(name string, increment, restart int64)
	// // nextSequenceValue returns the next value of the given given sequence
	// nextSequenceValue(name string) int64
	// // sequences returns a list of all sequences matching the given SQL pattern
	// sequences(pattern string) []seqData
	// // childrenIdsQuery returns a query that finds all descendant of the given
	// // a record from table including itself. The query has a placeholder for the
	// // record's ID
	// childrenIdsQuery(table string) string
	// // substituteErrorMessage substitutes the given error's message by newMsg
	// substituteErrorMessage(err error, newMsg string) error
	// // isSerializationError returns true if the given error is a serialization error
	// // and that the failed transaction should be retried.
	// isSerializationError(err error) bool
}

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

// RegisterDialect adds a adapter to the adapters registry
// name of the adapter should match the database/sql driver name
func RegisterDialect(name string, dialect Dialect) {
	adapters[name] = dialect
}

// GetDialect is xxx
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = adapters[name]
	return
}

// Cursor is a wrapper around a database transaction
type Cursor struct {
	tx *sqlx.Tx
}

// Execute a query without returning any rows. It panics in case of error.
// The args are for any placeholder parameters in the query.
// func (c *Cursor) Execute(query string, args ...interface{}) sql.Result {
// 	return dbExecute(c.tx, query, args...)
// }

// // Get queries a row into the database and maps the result into dest.
// // The query must return only one row. Get panics on errors
// func (c *Cursor) Get(dest interface{}, query string, args ...interface{}) {
// 	dbGet(c.tx, dest, query, args...)
// }

// // Select queries multiple rows and map the result into dest which must be a slice.
// // Select panics on errors.
// func (c *Cursor) Select(dest interface{}, query string, args ...interface{}) {
// 	dbSelect(c.tx, dest, query, args...)
// }

// //
// // newCursor returns a new db cursor on the given database
// func newCursor(db *sqlx.DB) *Cursor {
// 	adapter := adapters[db.DriverName()]
// 	tx := db.MustBegin()
// 	dbExecute(tx, adapter.setTransactionIsolation())
// 	return &Cursor{
// 		tx: tx,
// 	}
// }

// 对外显示链接参数
// DBParams returns the DB connection parameters currently in use
func DBParams() ConnectionParams {
	return connParams
}

// DBConnect connects to a database using the given driver and arguments.
func DBConnect(params ConnectionParams) {
	connParams = params
	connStr := DBParams().ConnectionString()
	fmt.Println(33333)
	fmt.Println(connStr)
	db = sqlx.MustConnect(params.Driver, connStr)
	fmt.Println(33333)
	fmt.Println(db)

	log.Info("Connected to database", "driver", params.Driver, "connStr", connStr)
}

// DBClose is a wrapper around sqlx.Close
// It closes the connection to the database
func DBClose() {
	err := db.Close()
	log.Info("Closed database", "error", err)
}

// // 带tx和不带tx
// // 带tx
// // dbExecute is a wrapper around sqlx.MustExec
// // It executes a query that returns no row
// func dbExecute(cr *sqlx.Tx, query string, args ...interface{}) sql.Result {
// 	query, args = sanitizeQuery(query, args...)
// 	t := time.Now()
// 	res, err := cr.Exec(query, args...)
// 	logSQLResult(err, t, query, args...)
// 	return res
// }

// 不带tx
// dbGetNoTx is a wrapper around sqlx.Get outside a transaction
// It gets the value of a single row found by the
// given query and arguments
func dbGetNoTx(dest interface{}, query string, args ...interface{}) {
	query, args = sanitizeQuery(query, args...)
	// t := time.Now()
	_ = db.Get(dest, query, args...)
	logSQLResult(query, args)
}

// 我现在需要一个最简单的能查询 query的查询
// 用于生成一个简单的逻辑
// sanitizeQuery calls 'In' expansion and 'Rebind' on the given query and
// returns the new values to use. It panics in case of error
func sanitizeQuery(query string, args ...interface{}) (string, []interface{}) {
	originalArgs := args
	q, args, err := sqlx.In(query, args...)
	if err != nil {
		log.Panic("Unable to expand 'IN' statement", "error", err, "query", query, "args", originalArgs)
	}
	q = sqlx.Rebind(sqlx.BindType(db.DriverName()), q)
	return q, args
}

// 日志输出
// Log the result of the given sql query started at start time with the
// given args, and error. This function panics after logging if error is not nil.
func logSQLResult(query string, args ...interface{}) {

	log.Info(query, args)
}
