package models

import (
	"github.com/astaxie/beego/orm"
	log "github.com/sirupsen/logrus"
)

func SelectUser(email string) (interface{}, error) {
	db := orm.NewOrm()
	var user User
	err := db.QueryTable(new(User)).Filter("Email", email).One(&user)
	if err != nil {
		return nil, err
	}
	log.Info("user ======== ", user)
	return user, nil
}