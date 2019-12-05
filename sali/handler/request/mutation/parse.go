package mutation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func (mu *Mutation) Parse(entities map[string]reflect.Type) {
	if mu.Entity != "" {
		parseEntity(mu, entities)
		return
	}
	parseManyToMany(mu, entities)
	fmt.Println(mu.Errors)
}

func parseEntity(mu *Mutation, entities map[string]reflect.Type) {
	if etype, ok := entities[mu.Entity]; ok {
		for id, v := range []interface{}{mu.Fields, mu.Filter} {
			entity := reflect.New(etype).Interface()
			b, err := json.Marshal(v)
			if err != nil {
				mu.Errors = append(mu.Errors, err.Error())
			}
			reader := bytes.NewBuffer(b)
			dec := json.NewDecoder(reader)
			dec.DisallowUnknownFields()
			err = dec.Decode(&entity)
			if err != nil {
				mu.Errors = append(mu.Errors, err.Error())
			}

			if id == 0 {
				mu.Fields = entity

			}
			// switch id {
			// case 0:
			// mu.Fields = entity
			// case 1:
			// 	mu.Filter = entity
			// }
		}
	} else {
		mu.Errors = append(mu.Errors, fmt.Sprintf("unknow entity '%v'", mu.Entity))
	}
}

func parseManyToMany(mu *Mutation, entities map[string]reflect.Type) {
	for id, v := range []string{mu.NM2M, mu.NReverse} {
		if etype, ok := entities[v]; ok {
			entity := reflect.New(etype).Interface()
			var b []byte
			var err error
			switch id {
			case 0:
				b, err = json.Marshal(mu.M2M)
			case 1:
				b, err = json.Marshal(mu.Reverse)
			}
			if err != nil {
				mu.Errors = append(mu.Errors, err.Error())
			}
			reader := bytes.NewBuffer(b)
			dec := json.NewDecoder(reader)
			dec.DisallowUnknownFields()
			err = dec.Decode(&entity)
			if err != nil {
				mu.Errors = append(mu.Errors, err.Error())
			}
			switch id {
			case 0:
				mu.M2M = entity
			case 1:
				mu.Reverse = entity
			}
		} else {
			mu.Errors = append(mu.Errors, fmt.Sprintf("unknow entity '%v'", v))
		}
	}
}
