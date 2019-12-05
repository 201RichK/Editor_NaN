package entities

type Organisation struct {
	ID       int        `json:",omitempty"`
	Name     string     `json:",omitempty"`
	Email    string     `json:",omitempty"`
	Password string     `json:",omitempty"`
	Students []*Student `json:",omitempty" orm:"reverse(many)"`
	Messages []*Message `json:",omitempty" orm:"reverse(many)"`
	Teams    []*Team    `json:",omitempty" orm:"reverse(many)"`
}
