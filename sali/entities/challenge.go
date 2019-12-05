package entities

type Challenge struct {
	ID        int         `json:",omitempty"`
	Expired   bool        `json:",omitempty"`
	Admin     uint64      `json:",omitempty"`
	Winner    uint64      `json:",omitempty"`
	Exercises []*Exercise `json:",omitempty" orm:"rel(m2m)"`
	Students  []*Student  `json:",omitempty" orm:"rel(m2m)"`
}
