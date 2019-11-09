package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:password@localhost/naneditor?sslmode=disable")
	orm.RegisterModel(&User{}, &Exercice{})
}

type User struct {
	Id       int64
	Email    string      `orm:"size(128)" valid:"Required;Email"`
	Password string      `orm:"size(128)"  valid:"Required"`
	Exercice []*Exercice `orm:"reverse(many)"`
}

type Exercice struct {
	Id       int64
	Sujet    string
	Run      int64
	UserCode string
	Program  string
	ExecTime *time.Duration
	Success  bool
	User     *User `orm:"rel(fk)"`
}
