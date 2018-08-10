package a

// Kind:"Bool" Name:"TBool" String:"TBool"
type TBool bool

// Kind:"Int" Name:"TInt" String:"TInt"
type TInt int

// Kind:"Int8" Name:"TInt8" String:"TInt8"
type TInt8 int8

// Kind:"Int16" Name:"TInt16" String:"TInt16"
type TInt16 int16

// Kind:"Int32" Name:"TInt32" String:"TInt32"
type TInt32 int32

// Kind:"Int64" Name:"TInt64" String:"TInt64"
type TInt64 int64

// Kind:"Uint" Name:"TUint" String:"TUint"
type TUint uint

// Kind:"Uint8" Name:"TUint8" String:"TUint8"
type TUint8 uint8

// Kind:"Uint16" Name:"TUint16" String:"TUint16"
type TUint16 uint16

// Kind:"Uint32" Name:"TUint32" String:"TUint32"
type TUint32 uint32

// Kind:"Uint64" Name:"TUint64" String:"TUint64"
type TUint64 uint64

// Kind:"Uintptr" Name:"TUintptr" String:"TUintptr"
type TUintptr uintptr

// Kind:"Float32" Name:"TFloat32" String:"TFloat32"
type TFloat32 float32

// Kind:"Float64" Name:"TFloat64" String:"TFloat64"
type TFloat64 float64

// Kind:"Complex64" Name:"TComplex64" String:"TComplex64"
type TComplex64 complex64

// Kind:"Complex128" Name:"TComplex128" String:"TComplex128"
type TComplex128 complex128

// Kind:"String" Name:"TString" String:"TString"
type TString string

// Kind:"Byte" Name:"TByte" String:"TByte"
type TByte byte

// Kind:"Rune" Name:"TRune" String:"TRune"
type TRune rune

// Kind:"Error" Name:"TError" String:"TError"
type TError error

// Kind:"Array" Name:"TArray" String:"TArray"
type TArray [1]struct{}

// Kind:"Chan" Name:"TChan" String:"TChan"
type TChan chan struct{}

// Kind:"Func" Name:"TFunc" String:"TFunc"
type TFunc func()

// Kind:"Interface" Name:"TInterface" String:"TInterface"
type TInterface interface{}

// Kind:"Map" Name:"TMap" String:"TMap"
type TMap map[string]struct{}

// Kind:"Ptr" Name:"TPtr" String:"TPtr"
type TPtr *int

// Kind:"Slice" Name:"TSlice" String:"TSlice"
type TSlice []struct{}
