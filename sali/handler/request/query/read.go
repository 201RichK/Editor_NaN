package query

import (
	"fmt"
	"sali/handler/request/orm"
	"strings"
)

// func (qu *Query) read() (int64, interface{}, error) {
// 	slice := reflect.New(reflect.SliceOf(reflect.TypeOf(qu.Filter)))
// 	result := slice.Interface()
// 	fmt.Println(qu.Filter)
// 	db := db.DB.New()
// 	id := db.Set("gorm:auto_preload", true).Model(qu.Entity).Order(qu.Order).Offset(qu.Offset).Where(qu.Filter).Limit(qu.Limit).Find(result).RowsAffected
// 	er := db.Error
// 	if er != nil {
// 		return id, nil, er
// 	}
// 	if qu.Limit == 1 {
// 		if slice.Elem().Len() == 1 {
// 			return id, slice.Elem().Index(0).Interface(), nil
// 		}
// 		return id, struct{}{}, nil
// 	}
// 	return id, result, nil
// }

func (qu *Query) read() (int64, interface{}, error) {
	cond := orm.NewCondition()
	if filter, ok := qu.Filter.(map[string]interface{}); ok {
		for k, v := range filter {
			cond.And(k, v)
		}
	}
	o := orm.NewOrm()
	params := []orm.Params{}
	id, err := o.QueryTable(qu.Entity).SetCond(cond).Limit(qu.Limit).Offset(qu.Offset).OrderBy(qu.Order).Values(&params, qu.Fields)
	if err != nil {
		fmt.Println(err)
		return id, nil, err
	}

	if qu.Limit == 1 {
		if len(params) == 1 {
			return id, corrige(params[0]), nil
		}
		return id, struct{}{}, nil
	}

	var data []map[string]interface{}
	for _, v := range params {
		fmt.Println(v)
		data = append(data, corrige(v))
	}

	return id, data, nil
}

func corrige(data map[string]interface{}) map[string]interface{} {
	mp := make(map[string]interface{})
	for k, v := range data {
		s := strings.Split(k, "__")
		if first, others := s[0], strings.Join(s[1:], "__"); len(others) == 0 {
			mp[first] = v
		} else {
			m := make(map[string]interface{})
			m[others] = v
			if v, ok := mp[first]; ok {
				if v, ok := v.(map[string]interface{}); ok {
					for key, value := range corrige(m) {
						v[key] = value
					}
					mp[first] = v
				}
			} else {
				mp[first] = corrige(m)
			}
		}
	}
	return mp
}
