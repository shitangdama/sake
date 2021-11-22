package scheme

// Scheme represents a table of database
type Scheme struct {
	Name string

	DBName  string
	TableID string

	PrioritizedPrimaryField *Field

	PrimaryFields       []*Field
	PrimaryFieldDBNames []string

	// Model  map[string]string
	Schema string

	Fields       []*Field
	FieldNames   []string
	FieldDBNames []string
	fieldMap     map[string]*Field
	fieldDBMap   map[string]*Field

	// Relationship map[string]*Relationship
}
