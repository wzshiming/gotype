package generics

// String:"Struct" NumField:"2" NumMethod:"1" NumParam:"2" Kind:"Struct"
// To:"FieldByName:Data" Name:"Data" String:"Data T" Kind:"Field"
// To:"FieldByName:Data,Elem" Name:"T" String:"T"  Kind:"Param"
type Struct[T any, P comparable] struct {
	Data  T
	Data2 P
}

func (s *Struct[T, P]) String() string {
	return ""
}

var (
	// Kind:"Declaration" Name:"a" String:"a"
	// To:"Declaration" Kind:"Struct" String:"Struct[T, P]" NumParam:"2"
	// To:"Declaration,FieldByName:Data" Kind:"Field" String:"Data T"
	// To:"Declaration,FieldByName:Data,Elem" Kind:"String" String:"T"
	a Struct[string, int]
)
