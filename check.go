package gotype

// Equal Reports whether the type type is equal
func Equal(t0 Type, t1 Type) bool {
	if t0.PkgPath() != t1.PkgPath() {
		return false
	}
	n0 := t0.Name()
	n1 := t1.Name()
	if n0 != n1 {
		return false
	}
	if n0 == "" {
		return Identical(t0, t1)
	}
	return true
}

// Implements reports whether type inter implements interface t.
func Implements(t Type, inter Type) bool {
	if inter.Kind() == Declaration {
		inter = inter.Declaration()
	}

	if inter.Kind() != Interface {
		return false
	}

	if t.Kind() == Declaration {
		t = t.Declaration()
	}

	lenMet0 := inter.NumField()
	lenMet1 := t.NumMethod()
	if lenMet1 < lenMet0 {
		return false
	}
	for i := 0; i != lenMet0; i++ {
		m0 := inter.Field(i)

		if m0.Kind() == Declaration {
			name := m0.Name()
			m1, ok := t.MethodByName(name)
			if !ok {
				return false
			}

			if !Identical(m0, m1) {
				return false
			}
		} else if !Implements(t, m0) {
			return false
		}
	}
	return true
}

// Identical reports whether t0 and t1 are identical types.
func Identical(t0, t1 Type) bool {
	return identical(t0, t1, true)
}

// IdenticalIgnoreTags reports whether t0 and t1 are identical types if tags are ignored.
func IdenticalIgnoreTags(t0, t1 Type) bool {
	return identical(t0, t1, false)
}

func identical(t0, t1 Type, cmpTags bool) bool {
	if t0 == t1 {
		return true
	}
	k0 := t0.Kind()
	k1 := t1.Kind()
	if k0 != k1 {
		return false
	}

	switch k0 {
	case Bool,
		Int, Int8, Int16, Int32, Int64,
		Uint, Uint8, Uint16, Uint32, Uint64,
		Uintptr,
		Float32, Float64,
		Complex64, Complex128,
		String, Byte, Rune, Error:
		return true
	case Slice, Ptr:
		return identical(t0.Elem(), t1.Elem(), cmpTags)
	case Array:
		return t0.Len() == t1.Len() && identical(t0.Elem(), t1.Elem(), cmpTags)
	case Chan:
		return t0.ChanDir() == t1.ChanDir() && t0.Len() == t1.Len() && identical(t0.Elem(), t1.Elem(), cmpTags)
	case Map:
		return identical(t0.Key(), t1.Key(), cmpTags) && identical(t0.Elem(), t1.Elem(), cmpTags)
	case Field:
		if cmpTags {
			return t0.Tag() == t1.Tag() && identical(t0.Elem(), t1.Elem(), cmpTags)
		}
		return identical(t0.Elem(), t1.Elem(), cmpTags)
	case Scope:
		return t0.PkgPath() == t1.PkgPath()
	case Declaration:
		return identical(t0.Declaration(), t1.Declaration(), cmpTags)
	case Func:
		numIn0 := t0.NumIn()
		numIn1 := t1.NumIn()
		if numIn0 != numIn1 {
			return false
		}
		for i := 0; i != numIn0; i++ {
			if !identical(t0.In(i), t1.In(i), cmpTags) {
				return false
			}
		}

		numOut0 := t0.NumOut()
		numOut1 := t1.NumOut()
		if numOut0 != numOut1 {
			return false
		}
		for i := 0; i != numOut0; i++ {
			if !identical(t0.Out(i), t1.Out(i), cmpTags) {
				return false
			}
		}
		return true
	case Interface:
		numMethod0 := t0.NumMethod()
		numMethod1 := t1.NumMethod()
		if numMethod0 != numMethod1 {
			return false
		}
		for i := 0; i != numMethod0; i++ {
			if !identical(t0.Method(i), t1.Method(i), cmpTags) {
				return false
			}
		}
		return true
	case Struct:
		numField0 := t0.NumField()
		numField1 := t1.NumField()
		if numField0 != numField1 {
			return false
		}
		for i := 0; i != numField0; i++ {
			if !identical(t0.Field(i), t1.Field(i), cmpTags) {
				return false
			}
		}
		return true
	}
	return false
}
