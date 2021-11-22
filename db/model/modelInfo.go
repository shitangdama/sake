package model

type (
	// Column defines column in table
	Column struct {
		*DbColumn
		Index *DbIndex
	}

	// 这里多了几个参数
	// DbColumn defines column info of columns
	DbColumn struct {
		Name            string      `db:"COLUMN_NAME"`
		DataType        string      `db:"DATA_TYPE"`
		Extra           string      `db:"EXTRA"`
		Comment         string      `db:"COLUMN_COMMENT"`
		ColumnDefault   interface{} `db:"COLUMN_DEFAULT"`
		IsNullAble      string      `db:"IS_NULLABLE"`
		OrdinalPosition int         `db:"ORDINAL_POSITION"`
	}
	// DbIndex defines index of columns in information_schema.statistic
	DbIndex struct {
		IndexName  string `db:"INDEX_NAME"`
		NonUnique  int    `db:"NON_UNIQUE"`
		SeqInIndex int    `db:"SEQ_IN_INDEX"`
	}
)
