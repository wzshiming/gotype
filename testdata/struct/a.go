package a

// String:"Struct" NumField:"3" NumMethod:"1"
// To:"FieldByName:Name" Tag:"json:\"name\""
// To:"Field:0" Tag:"json:\"name\""
// To:"Field:0,Elem" Name:"string"
// To:"FieldByName:Age" Tag:"json:\"age\""
// To:"FieldByName:Msg,Elem,FieldByName:Msg2" Tag:"json:\"msg\""
// To:"FieldByName:Msg2" Tag:"json:\"msg\""
// To:"Field:1" Tag:"json:\"age\""
// To:"Field:1,Elem" Name:"int"
// To:"MethodByName:String" Name:"String"
// To:"Method:0" Name:"String"
// To:"MethodByName:Message1" Name:"Message1"
// To:"MethodByName:Message2" Name:"Message2"
type Struct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Msg
}

func (s *Struct) String() string {
	return ""
}

type Msg struct {
	Msg2 string `json:"msg"`
}

func (Msg) Message1() string {
	return ""
}
func (Msg) Message2() string {
	return ""
}
