package mutation

import (
	"errors"
	"fmt"
	"sali/handler/request/orm"
)

func (mu *Mutation) create() (int64, error) {
	o := orm.NewOrm()
	if mu.M2M != nil && mu.Reverse != nil {
		m2m := o.QueryM2M(mu.M2M, mu.Rel)
		if !m2m.Exist(mu.Reverse) {
			id, err := m2m.Add(mu.Reverse)
			fmt.Println(id, err)
			if err != nil {
				return 0, err
			}
			return id, nil
		}
		return 0, errors.New("Record already exists")
	}
	cond := orm.NewCondition()
	for k, v := range mu.Filter {
		cond.And(k, v)
	}
	insert, err := o.QueryTable(mu.Entity).SetCond(cond).PrepareInsert()
	id, err := insert.Insert(mu.Fields)
	fmt.Println(mu.Fields, id)
	if err != nil {
		return id, err
	}
	return id, nil
}
