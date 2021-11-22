package scheme

import (
	"sake/db/fieldtype"
	"sake/scheme/convert"
	"sake/tools/nbutils"
)

// Field is column
type Field struct {
	Name     string
	DBName   string
	ColumnID string

	Scheme *Scheme

	Type      string
	FieldType fieldtype.Type
	Size      int
	Digits    nbutils.Digits

	PrimaryKey    bool
	AutoIncrement bool

	HasDefaultValue bool
	DefaultValue    string

	Index   bool
	NotNull bool // required 为什么用这个参数
	Unique  bool

	Check   string
	Comment string

	Tag    string
	TagMap map[string]string
}

// ParseField is xxx
func ParseField(fieldInfo *convert.Field) *Field {

	field := &Field{
		Name:   fieldInfo.Name,
		DBName: SnakeString(fieldInfo.Name),
		Type:   fieldInfo.Type,
		Tag:    fieldInfo.Tag,
		TagMap: ParseTagSetting(fieldInfo.Tag, ";"),
	}

	// if dbName, ok := field.TagSettings["COLUMN"]; ok {
	// 	field.DBName = dbName
	// }

	if val, ok := field.TagMap["PRIMARYKEY"]; ok && CheckTruth(val) {
		field.PrimaryKey = true
	} else if val, ok := field.TagMap["PRIMARY_KEY"]; ok && CheckTruth(val) {
		field.PrimaryKey = true
	}

	if val, ok := field.TagMap["AUTOINCREMENT"]; ok && CheckTruth(val) {
		field.AutoIncrement = true
		field.HasDefaultValue = true
	}

	// if num, ok := field.TagSettings["AUTOINCREMENTINCREMENT"]; ok {
	// 	field.AutoIncrementIncrement, _ = strconv.ParseInt(num, 10, 64)
	// }

	if v, ok := field.TagMap["DEFAULT"]; ok {
		field.HasDefaultValue = true
		field.DefaultValue = v
	}

	if val, ok := field.TagMap["NOT NULL"]; ok && CheckTruth(val) {
		field.NotNull = true
	} else if val, ok := field.TagMap["NOTNULL"]; ok && CheckTruth(val) {
		field.NotNull = true
	}

	if val, ok := field.TagMap["UNIQUE"]; ok && CheckTruth(val) {
		field.Unique = true
	}

	if val, ok := field.TagMap["COMMENT"]; ok {
		field.Comment = val
	}

	return field
}
