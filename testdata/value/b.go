package a

import (
	"fmt"
	t "time"
)

var (

	// To:"Declaration" Kind:"String"
	S0, S1 = "", ""

	// Name:"TNow"
	// To:"Declaration,Declaration" Name:"Time" Kind:"Struct"
	TNow = t.Now()

	// Name:"MyNow"
	// To:"Declaration" String:"Now"
	// To:"Declaration,Declaration" String:"func() (_)"
	MyNow = t.Now

	// Name:"MyPrintf"
	// To:"Declaration" String:"Printf"
	// To:"Declaration,Declaration" Kind:"Func" String:"func(format, a...) (n, err)"
	// To:"Declaration,Declaration,In:1,Declaration" String:"[]interface{}"
	MyPrintf = fmt.Printf
)
