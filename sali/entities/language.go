package entities

type Language struct {
	ID       uint64     `json:",omitempty"`
	Name     string     `json:",omitempty" orm:"size(50);unique"`
	Index    uint8      `json:",omitempty" orm:"unique"`
	Team     *Team      `json:",omitempty" orm:"null;rel(one);on_delete(set_null)"`
	Students []*Student `json:",omitempty" orm:"rel(m2m)"`
}
