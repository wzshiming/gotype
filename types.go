package gotype

type Type interface {
	Name() string
	Kind() Kind
	Key() Type
	Elem() Type // 指针 接口 别名 数组 map值 到下一级
	NumMethods() int
	Methods(int) *TypeMethod
	NumField() int
	Field(int) *TypeStructField
	Len() int     // 数组
	NumOut() int  // func
	Out(int) Type // func
	NumIn() int   // func
	In(int) Type  // func
}
