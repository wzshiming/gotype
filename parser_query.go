package gotype

func (r *astParser) Search(name string) Type {
	return r.nameds.Search(name)
}

func (r *astParser) Child(i int) Type {
	return r.nameds.Index(i)
}

func (r *astParser) NumChild() int {
	return len(r.nameds)
}
