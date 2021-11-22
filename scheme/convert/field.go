package convert

// Field represents a column of database
type Field struct {
	Name string `yaml:"name" json:"name"`
	Type string `yaml:"type" json:"type"`
	Tag  string `yaml:"tag"  json:"tag"`
	// 这里可以给其他信息
}
