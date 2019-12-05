package mutation

import (
	"errors"
	"sali/handler/request/orm"
)

func (mu *Mutation) delete() (int64, error) {
	o := orm.NewOrm()

	if mu.M2M != nil && mu.Reverse != nil {
		m2m := o.QueryM2M(mu.M2M, mu.Rel)
		if m2m.Exist(mu.Reverse) {
			id, err := m2m.Remove(mu.Reverse)
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
	id, err := o.QueryTable(mu.Entity).SetCond(cond).Delete()
	if err != nil {
		return id, err
	}
	return id, nil
}
