package operator

// An Operator inside an SQL WHERE clause
type Operator string

// 这里要考虑怎么匹配json里面的数据
// 是不是进行匹配

// Operators
const (
	Equals         Operator = "="
	NotEquals      Operator = "!="
	Greater        Operator = ">"
	GreaterOrEqual Operator = ">="
	Lower          Operator = "<"
	LowerOrEqual   Operator = "<="
	Like           Operator = "=like"
	Contains       Operator = "like"
	NotContains    Operator = "not like"
	IContains      Operator = "ilike"
	NotIContains   Operator = "not ilike"
	ILike          Operator = "=ilike"
	In             Operator = "in"
	NotIn          Operator = "not in"
	ChildOf        Operator = "child_of"
)

// var allowedOperators = map[Operator]bool{
// 	Equals:         true,
// 	NotEquals:      true,
// 	Greater:        true,
// 	GreaterOrEqual: true,
// 	Lower:          true,
// 	LowerOrEqual:   true,
// 	Like:           true,
// 	Contains:       true,
// 	NotContains:    true,
// 	IContains:      true,
// 	NotIContains:   true,
// 	ILike:          true,
// 	In:             true,
// 	NotIn:          true,
// 	ChildOf:        true,
// }

// var negativeOperators = map[Operator]bool{
// 	NotEquals:    true,
// 	NotContains:  true,
// 	NotIContains: true,
// 	NotIn:        true,
// }

// var positiveOperators = map[Operator]bool{
// 	Equals:    true,
// 	IContains: true,
// 	ILike:     true,
// 	Contains:  true,
// 	Like:      true,
// 	In:        true,
// }

// var multiOperator = map[Operator]bool{
// 	In:    true,
// 	NotIn: true,
// }
