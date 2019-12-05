package entities

type Message struct {
	ID           int           `json:",omitempty"`
	Content      string        `json:",omitempty"`
	Student      *Student      `json:",omitempty" orm:"rel(fk);null"`
	Organisation *Organisation `json:",omitempty" orm:"rel(fk);null"`
	Team         *Team         `json:",omitempty" orm:"rel(fk)"`
}
