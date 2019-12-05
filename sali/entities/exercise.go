package entities

type Exercise struct {
	ID         int          `json:",omitempty"`
	Kind       string       `json:",omitempty"`
	Level      string       `json:",omitempty"`
	Result     string       `json:",omitempty"`
	Subject    string       `json:",omitempty"`
	Example    string       `json:",omitempty"`
	Challenges []*Challenge `json:",omitempty" orm:"reverse(many)"`
	Runs       []*Run       `json:",omitempty" orm:"reverse(many)"`
}
