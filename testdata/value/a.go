package a

const (
	// Value:"1" Name:"U1"
	U1 uint = 1 << iota
	// Value:"2" Name:"U2"
	U2
	// Value:"4" Name:"U3"
	U3
	// Value:"3" Name:"U4"
	U4 = U1 | U2
)

const (
	// Value:"0" Name:"I1"
	I1 int = iota + 0
	// Value:"1" Name:"I2"
	I2
	// Value:"2" Name:"I3"
	I3
	// Value:"3" Name:"I4"
	I4 = I2 + I3
	// Value:"4" Name:"I5"
	I5 = -I2 + I3 + I4
)

const (
	// Value:"0" Name:"F1"
	F1 = float64(iota / 10)
	// Value:"1/10" Name:"F2"
	F2
	// Value:"1/5" Name:"F3"
	F3
	// Value:"3" Name:"F4"
	F4 float64 = iota
)

const (
	// Value:"\"A\"" Name:"StrA"
	StrA = "A"
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

var (
	// To:"Declaration" String:"[3]int"
	Arr1 [I4]int
)
