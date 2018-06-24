package gotype

func newTypeValueBind(typ, val Type, info *info) Type {
	switch typ.Kind() {
	case Struct:
		if v, ok := val.(*typeValuePairs); ok {
			nt := &typeStruct{}
			fl := typ.NumField()
			for i := 0; i != fl; i++ {
				f := typ.Field(i)
				name := f.Name()
				b := f
				if val, ok := v.li.Search(name); ok {
					b = newTypeValueBind(f, val, info)
				}
				nt.fields = append(nt.fields, b)
			}
			al := typ.NumAnonymo()
			for i := 0; i != al; i++ {
				f := typ.Field(i)
				name := f.Name()
				b := f
				if val, ok := v.li.Search(name); ok {
					b = newTypeValueBind(f, val, info)
				}
				nt.anonymo = append(nt.anonymo, b)
			}
			return nt
		}
	case Var:
		name := typ.Name()
		t := newTypeValueBind(typ.Elem(), val, info)
		if typ.Origin() != nil {
			t = newTypeOrigin(t, typ.Origin(), info, typ.Doc(), typ.Comment())
		}
		return newTypeVar(name, t)
	}

	return newValueBind(typ, val.Value)
}

func newValueBind(typ Type, valfunc func() string) *typeValueBind {
	if tv, ok := typ.(*typeValueBind); ok {
		return newValueBind(tv.Type, tv.Value)
	}
	return &typeValueBind{
		Type:    typ,
		valfunc: valfunc,
	}
}

func newEvalBind(val Type, index int64, info *info) Type {
	ori := val.Origin()
	val = newValueBind(val, func() string {
		v, err := constantEval(ori, int64(index), info)
		if err != nil {
			return ""
		}
		return v.ExactString()
	})
	return val
}

type typeValueBind struct {
	Type
	valfunc func() string
	val     string
}

func (t *typeValueBind) Value() string {
	if t.val == "" {
		t.val = t.valfunc()
	}
	return t.val
}
