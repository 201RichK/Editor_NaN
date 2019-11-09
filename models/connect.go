package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:password@localhost/editornan?sslmode=disable")
	orm.RegisterModel(new(User), new(Exercice))
}

type User struct {
	Id           int64
	Email        string         `orm:"size(128)" valid:"Required;Email"`
	Password     string         `orm:"size(128)"  valid:"Required"`
	Exercice []*Exercice `orm:"reverse(many)"`
}

type Exercice struct {
	Id         int64
	Sujet      string
	Run    int64
	UserCode   string
	Program    string
	ExecTime   *time.Duration
	Success    bool
	User *User `orm:"rel(fk)"`
}


/*
INSERT INTO public.exercice (sujet, run, user_code, program, exec_time, success, user_id) VALUES ('somme de deux nombre a, b', 0, '', 'func (){\n\t\n}', CURRENT_TIMESTAMP, false, 1);
INSERT INTO public.user (email, password) VALUES ('john@mail.com', 'john');

 */