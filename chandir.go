package gotype

//go:generate stringer -type ChanDir chandir.go

// ChanDir represents a channel type's direction.
type ChanDir int

// Define channel direction
const (
	RecvDir ChanDir = 1 << iota         // chan<-
	SendDir                             // <-chan
	BothDir ChanDir = RecvDir | SendDir // chan
)
