package mutation

import (
	"encoding/json"
	"sali/handler/request/orm"
)

func (mu *Mutation) update() (int64, error) {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	for k, v := range mu.Filter {
		cond.And(k, v)
	}
	var fields map[string]interface{}
	b, _ := json.Marshal(mu.Fields)
	json.Unmarshal(b, &fields)
	id, err := o.QueryTable(mu.Entity).SetCond(cond).Update(fields)
	if err != nil {
		return id, err
	}
	return id, nil
}
