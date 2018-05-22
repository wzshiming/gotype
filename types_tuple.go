package gotype

func newTypeTuple(all Types) Type {
	switch len(all) {
	case 0:
		return nil
	case 1:
		return all[0]
	default:
		return &typeTuple{
			Type: all[0],
			all:  all,
		}
	}

}

type typeTuple struct {
	Type       // [0]
	all  Types // [:]
}
