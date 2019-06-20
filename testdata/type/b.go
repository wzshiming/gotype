package a

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"time"
)

// Name:"MyStringer"
// Kind:"Interface"
type MyStringer fmt.Stringer

// Name:"a" Value:""
// To:"Declaration" Kind:"Interface" NumMethod:"1" NumField:"2" String:"interface{Stringer; MyString}"
// To:"Declaration,MethodByName:String,Declaration" String:"func() (_)" Value:""
var a = (interface {
	fmt.Stringer
	MyString() string
})(nil)

// Name:"b" Value:""
// To:"Declaration" Kind:"Struct" String:"struct{fmt.Stringer; String string `s:\"\"`}"
// To:"Declaration,MethodByName:String,Declaration" String:"func() (_)" Value:""
var b = struct {
	fmt.Stringer
	String string `s:""`
}{}

// Name:"c" Value:""
// To:"Declaration" Kind:"Map" Key:"String" Elme:"Int" String:"map[string]int" Value:""
var c = map[string]int{
	"S": 1,
}

// Name:"d" Value:""
// To:"Declaration" Kind:"Slice" Elme:"Int" String:"[]int" Value:""
var d = make([]int, 10)

// Name:"e" Value:""
// To:"Declaration" Kind:"Array" Elme:"Byte" Len:"8" String:"[8]byte" Value:""
var e = [8]byte{}

// Name:"s" Value:""
// Kind:"Interface" NumMethod:"1" String:"s" Value:""
type s fmt.Stringer

// Name:"t" Value:""
// Kind:"Struct" NumField:"3" String:"t" Value:""
type t time.Time

// Name:"v" Value:""
// Kind:"Map" String:"v" Key:"string" Elme:"[]string" Value:""
type v url.Values

// Name:"l" Value:""
// String:"l" Elme:"byte" Kind:"Slice" Value:""
type l json.RawMessage

// Name:"w" Value:""
// String:"w" Kind:"Func" Value:""
type w filepath.WalkFunc
