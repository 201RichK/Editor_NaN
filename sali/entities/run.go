package entities

type Run struct {
	ID       uint64
	Program  string
	Language string
	Exercise *Exercise `orm:"rel(fk)"`
}
