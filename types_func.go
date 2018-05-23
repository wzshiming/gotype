package gotype

import "bytes"

type typeFunc struct {
	typeBase
	variadic bool
	params   Types
	results  Types
}

func (t *typeFunc) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("func(")
	for i, v := range t.params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(v.String())
	}
	buf.WriteString(")")

	if len(t.results) > 1 {
		buf.WriteString(" (")
	}
	for i, v := range t.results {
		if i != 0 {
			buf.WriteString(", ")
		}
		if t.variadic && i+1 == len(t.results) {
			buf.WriteString("...")
		}
		buf.WriteString(v.String())
	}
	if len(t.results) > 1 {
		buf.WriteString(")")
	}
	return buf.String()
}

func (t *typeFunc) Kind() Kind {
	return Func
}

func (t *typeFunc) NumOut() int {
	return t.results.Len()
}

func (t *typeFunc) Out(i int) Type {
	return t.results.Index(i)
}

func (t *typeFunc) NumIn() int {
	return t.params.Len()
}

func (t *typeFunc) In(i int) Type {
	return t.params.Index(i)
}

func (t *typeFunc) IsVariadic() bool {
	return t.variadic
}
