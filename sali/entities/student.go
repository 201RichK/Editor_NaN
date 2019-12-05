package entities

type Student struct {
	ID           int           `json:",omitempty"`
	State        int8          `json:",omitempty" orm:"default(-1)"`
	Name         string        `json:",omitempty" orm:"size(50);unique"`
	Email        string        `json:",omitempty" orm:"size(50);unique"`
	Password     string        `json:",omitempty"`
	Language     *Language     `json:",omitempty" orm:"rel(fk)"`
	Organisation *Organisation `json:",omitempty" orm:"rel(fk);null"`
	Teams        []*Team       `json:",omitempty" orm:"reverse(many)"`
	Scores       []*Score      `json:",omitempty" orm:"reverse(many)"`
	Messages     []*Message    `json:",omitempty" orm:"reverse(many)"`
	Challenges   []*Challenge  `json:",omitempty" orm:"reverse(many)"`
}
