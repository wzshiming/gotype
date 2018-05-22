package gotype

func (r *Parser) Search(name string) Type {
	return r.Nameds.Search(name)
}

func (r *Parser) Child(i int) Type {
	return r.Nameds.Index(i)
}

func (r *Parser) NumChild() int {
	return len(r.Nameds)
}
