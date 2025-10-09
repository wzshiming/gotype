package generics

// Generic type with single type parameter
// String:"Box[T any]"
// NumTypeParam:"1"
// To:"TypeParam:0" Name:"T" Kind:"TypeParam" String:"T any"
// To:"TypeParam:0,Constraint" Name:"any"
type Box[T any] struct {
	Value T
}

// Generic type with constraint
// String:"Comparable[T comparable]"
// NumTypeParam:"1"
// To:"TypeParam:0" Name:"T" Kind:"TypeParam" String:"T comparable"
// To:"TypeParam:0,Constraint" Name:"comparable"
type Comparable[T comparable] struct {
	Value T
}

// Generic type with multiple type parameters
// String:"Pair[K comparable, V any]"
// NumTypeParam:"2"
// To:"TypeParam:0" Name:"K" Kind:"TypeParam" String:"K comparable"
// To:"TypeParam:0,Constraint" Name:"comparable"
// To:"TypeParam:1" Name:"V" Kind:"TypeParam" String:"V any"
// To:"TypeParam:1,Constraint" Name:"any"
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Generic function
// String:"Max[T comparable]"
// To:"Declaration" NumTypeParam:"1" String:"func[T comparable](a, b) (_)"
// To:"Declaration,TypeParam:0" Name:"T" Kind:"TypeParam" String:"T comparable"
// To:"Declaration,TypeParam:0,Constraint" Name:"comparable"
func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Generic function with multiple type parameters
// String:"Map[K comparable, V any]"
// To:"Declaration" NumTypeParam:"2" String:"func[K comparable, V any](m, f) (_)"
// To:"Declaration,TypeParam:0" Name:"K" Kind:"TypeParam" String:"K comparable"
// To:"Declaration,TypeParam:0,Constraint" Name:"comparable"
// To:"Declaration,TypeParam:1" Name:"V" Kind:"TypeParam" String:"V any"
// To:"Declaration,TypeParam:1,Constraint" Name:"any"
func Map[K comparable, V any](m map[K]V, f func(V) V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = f(v)
	}
	return result
}

// Custom interface constraint
type Stringer interface {
	String() string
}

// Generic type with custom interface constraint
// String:"StringableBox[T Stringer]"
// NumTypeParam:"1"
// To:"TypeParam:0" Name:"T" Kind:"TypeParam" String:"T Stringer"
// To:"TypeParam:0,Constraint" Name:"Stringer" Kind:"Interface"
type StringableBox[T Stringer] struct {
	Value T
}

// Generic function with custom interface constraint
// String:"Stringify[T Stringer]"
// To:"Declaration" NumTypeParam:"1" String:"func[T Stringer](value) (_)"
// To:"Declaration,TypeParam:0" Name:"T" Kind:"TypeParam" String:"T Stringer"
// To:"Declaration,TypeParam:0,Constraint" Name:"Stringer" Kind:"Interface"
func Stringify[T Stringer](value T) string {
	return value.String()
}
