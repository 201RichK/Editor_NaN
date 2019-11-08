package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:password@localhost/editornan?sslmode=disable")
	orm.RegisterModel(new(User), new(Exercice), new(Composition), new(Language))
}

type User struct {
	Id           int64
	Email        string         `orm:"size(128)" valid:"Required;Email"`
	Password     string         `orm:"size(128)"  valid:"Required"`
	Languages    []*Language    `orm:"reverse(many)"`
	Compositions []*Composition `orm:"reverse(many)"`
}

type Exercice struct {
	Id         int64
	Sujet      string
	RunCode    int64
	UserCode   string
	Program    string
	ExecTime   *time.Time
	Success    bool
	Compostion *Composition `orm:"rel(fk)"`
}

type Composition struct {
	Id       int64
	User     *User       `orm:"rel(fk)"`
	Exercies []*Exercice `orm:"reverse(many)"`
}

type Language struct {
	Id         int64
	Name       string
	LanguageId int64
	User       *User `orm:"rel(fk)"`
}
