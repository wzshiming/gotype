package gotype

import (
	"bytes"
	"fmt"
)

func newTypeValuePair(key, value Type) Type {
	return &typeValuePair{
		key:   key,
		value: value,
	}
}

type typeValuePair struct {
	typeBase
	key, value Type
}

func (t *typeValuePair) String() string {
	return fmt.Sprintf("%v: %v", t.key, t.value)
}

func (t *typeValuePair) Name() string {
	return t.key.Name()
}

func (t *typeValuePair) Value() string {
	return t.value.Value()
}

type typeValuePairs struct {
	typeBase
	li types
}

func (t *typeValuePairs) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("{")
	for i, v := range t.li {
		if i != 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(v.String())
	}
	buf.WriteByte('}')
	return buf.String()
}
