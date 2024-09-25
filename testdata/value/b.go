package a

import (
	"fmt"
	"time"
	t "time"
)

var (

	// To:"Declaration" Kind:"String"
	S0, S1 = "", ""

	// Name:"TNow"
	// To:"Declaration,Declaration" Name:"Time" Kind:"Struct"
	TNow = t.Now()

	// Name:"MyNow"
	// To:"Declaration" String:"t.Now"
	// To:"Declaration,Declaration" String:"func() (_)"
	MyNow = t.Now

	// Name:"MyPrintf"
	// To:"Declaration" String:"fmt.Printf"
	// To:"Declaration,Declaration" Kind:"Func" String:"func(format, a...) (n, err)"
	MyPrintf = fmt.Printf

	// Name:"ts" Value:""
	// To:"Declaration" String:"[]time.Time"
	ts []time.Time
)
