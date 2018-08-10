package a

const (
	// Value:"1" Name:"A1"
	A1 = 1 << iota
	// Value:"2" Name:"B1"
	B1
	// Value:"4" Name:"C1"
	C1
	// Value:"8" Name:"D1"
	D1
)

const (
	// Value:"0" Name:"A2"
	A2 = iota
	// Value:"1" Name:"B2"
	B2
	// Value:"2" Name:"C2"
	C2
	// Value:"3" Name:"D2"
	D2
)

// Value:"\"A\"" Name:"StrA"
const StrA = "A"

const (
	// Value:"\"CAB\"" Name:"StrB"
	StrB = StrC + "B"
	// Value:"\"CA\"" Name:"StrC"
	StrC = "C" + StrA
)

var (
	// Name:"Dir1"
	// To:"Declaration" String:"chan int"
	Dir1 chan int

	// Name:"Dir2"
	// To:"Declaration" String:"chan<- int"
	Dir2 chan<- int

	// Name:"Dir3"
	// To:"Declaration" String:"<-chan int"
	Dir3 <-chan int
)
