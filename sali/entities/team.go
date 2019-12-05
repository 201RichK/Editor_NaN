package entities

type Team struct {
	ID            int             `json:",omitempty"`
	Symbol        string          `json:",omitempty"`
	Name          string          `json:",omitempty" orm:"unique"`
	OrgAdmin      uint64          `json:",omitempty"`
	StudentAdmin  uint64          `json:",omitempty"`
	Organisations []*Organisation `json:",omitempty" orm:"rel(m2m);null"`
	Students      []*Student      `json:",omitempty" orm:"rel(m2m);null"`
	Language      *Language       `json:",omitempty" orm:"reverse(one);null"`
	Messages      []*Message      `json:",omitempty" orm:"reverse(many);null"`
}
