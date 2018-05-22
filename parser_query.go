package gotype

func (r *astParser) Search(name string) Type {
	return r.Nameds.Search(name)
}

func (r *astParser) Child(i int) Type {
	return r.Nameds.Index(i)
}

func (r *astParser) NumChild() int {
	return len(r.Nameds)
}
