package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func GetExerciceById(id int64) (*[]Exercice, error) {
	o := orm.NewOrm()
	v := new([]Exercice)
	if _, err := o.QueryTable("Exercice").RelatedSel().All(v); err == nil {
		return v, nil
	} else {
		fmt.Println(v)
		return nil, err
	}
}