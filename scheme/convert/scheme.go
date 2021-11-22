package convert

// Scheme calling this Struct
type Scheme struct {
	Name   string   `yaml:"name" json:"name"`
	DBName string   `yaml:"db_name" json:"db_name"`
	Schema string   `yaml:"schema" json:"schema"`
	Type   string   `yaml:"type" json:"type"`
	Fields []*Field `yaml:"fields" json:"fields"`
	// Relationships []*RelationshipInfo `yaml:"relationships" json:"relationships"`
	// Mixins        []MixinStruct       `yaml:"mixins" json:"mixins"`
}
