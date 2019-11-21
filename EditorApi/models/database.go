package models

import (
	utils "Editor_NaN/EditorApi/utils/option"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:admin@localhost/editordb?sslmode=disable")

	orm.RegisterModel(new(User), new(Language), new(Demande), new(Challenge), new(Vainqueur), new(ExoChallengeRand), new(Exercice), new(Testeur), new(UserChallenge))

	err := orm.RunSyncdb("default", false, true)
	utils.CheckError(err)
}
