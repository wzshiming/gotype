package gotype

func (r *Parser) ChildByName(name string) Type {
	v, ok := r.Values[name]
	if ok {
		return v
	}
	v, ok = r.Funcs[name]
	if ok {
		return v
	}
	v, ok = r.Types[name]
	if ok {
		return v
	}
	return nil
}

func (r *Parser) NumChild() int {
	return len(r.Values) + len(r.Funcs) + len(r.Types)
}

func (r *Parser) Range(f func(k string, v Type) bool) {
	r.RangeValues(f)
	r.RangeFuncs(f)
	r.RangeTypes(f)
}

func (r *Parser) RangeValues(f func(k string, v Type) bool) {
	for k, v := range r.Values {
		if !f(k, v) {
			return
		}
	}
}

func (r *Parser) RangeFuncs(f func(k string, v Type) bool) {
	for k, v := range r.Funcs {
		if !f(k, v) {
			return
		}
	}
}

func (r *Parser) RangeTypes(f func(k string, v Type) bool) {
	for k, v := range r.Types {
		if !f(k, v) {
			return
		}
	}
}
