package entities

type Score struct {
	ID        int      `json:",omitempty"`
	Points    uint8    `json:",omitempty"`
	Challenge uint64   `json:",omitempty"`
	Student   *Student `json:",omitempty" orm:"rel(fk)"`
}
