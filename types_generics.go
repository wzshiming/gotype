package gotype

import (
	"bytes"
)

func newTypeGenerics(base Type, params, actual []Type) Type {
	return &typeGenerics{
		Type:   base,
		params: params,
		actual: actual,
	}
}

type typeGenerics struct {
	Type
	params []Type
	actual []Type
}

func (t *typeGenerics) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(t.Type.String())
	buf.WriteString("[")
	for i, ind := range t.params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(ind.String())
	}
	buf.WriteString("]")
	return buf.String()
}

func (t *typeGenerics) NumParam() int {
	return len(t.params)
}

func (t *typeGenerics) Param(i int) Type {
	return t.params[i]
}

func (t *typeGenerics) ParamByName(name string) (Type, bool) {
	for _, ind := range t.params {
		if ind.Name() == name {
			return ind, true
		}
	}
	return nil, false
}
