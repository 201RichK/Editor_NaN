package models

import (
	utils "Editor_NaN/EditorApi/utils/option"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:password@localhost/nanchallengedb?sslmode=disable")

	orm.RegisterModel(new(User), new(Demande), new(Challenge), new(Exercice), new(Enonce), new(UserChallenge), new(Language), new(Testeur))

	err := orm.RunSyncdb("default", false, true)
	utils.CheckError(err)
}
