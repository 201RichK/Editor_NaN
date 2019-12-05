package query

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func (qu *Query) Parse(entities map[string]reflect.Type) {
	if etype, ok := entities[qu.Entity]; ok {
		qu.Errors = append(qu.Errors, parseFields(qu.Fields, etype)...)
		v, e := parseFilter(qu.Filter.(url.Values), etype)
		qu.Errors = append(qu.Errors, e...)
		b, err := json.Marshal(v)
		if err != nil {
			qu.Errors = append(qu.Errors, err.Error())
		}
		var filter map[string]interface{}
		json.Unmarshal(b, &filter)
		qu.Filter = filter
	} else {
		qu.Errors = append(qu.Errors, fmt.Sprintf("unknow entity '%v'", qu.Entity))
	}
}

//// Parse Fields
func parseFields(fields []string, etype reflect.Type) []string {
	errors := []string{}
	for _, name := range fields {
		errors = append(errors, verifyNames(name, etype)...)
	}
	return errors
}

func verifyNames(s string, etype reflect.Type) []string {
	errors := []string{}
	names := strings.Split(s, "__")
	switch etype.Kind() {
	case reflect.Slice:
		etype = etype.Elem()
		names = append(names[1:], names[2:]...)
		fallthrough
	case reflect.Ptr:
		etype = etype.Elem()
	}
	if f, ok := etype.FieldByName(names[0]); ok {
		if others := strings.Join(names[1:], "__"); len(others) > 0 {
			errors = append(errors, verifyNames(others, f.Type)...) //f.Type.Elem()
			return errors
		}
		return nil
	}
	errors = append(errors, fmt.Sprintf("unknow fields's name '%v' in 'fields'", s))
	return errors
}

/// Parse Filter
// func parseFilter(filter url.Values, etype reflect.Type) (interface{}, []string) {
// 	errors := []string{}
// 	mp := make(map[string]interface{})
// 	for k, v := range filter {
// 		v, e := assignValue(etype, k, v)
// 		errors = append(errors, e...)
// 		mp[k] = v.Interface()
// 	}
// 	return mp, errors
// }

func parseFilter(filter url.Values, etype reflect.Type) (interface{}, []string) {
	errors := []string{}
	mp := make(map[string]interface{})
	for k, v := range filter {
		v, e := assignFields(k, etype, v)
		errors = append(errors, e...)
		mp[k] = v.Interface()
	}
	return mp, errors
}

func assignType(s string, etype reflect.Type, values []string) (reflect.Value, []string) {
	errors := []string{}
	vz := reflect.Zero(etype)
	switch etype.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, e := assignInt(etype, values[0])
		vz = v
		errors = append(errors, e...)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, e := assignUint(etype, values[0])
		vz = v
		errors = append(errors, e...)
	case reflect.Bool:
		v, e := assignBool(etype, values[0])
		vz = v
		errors = append(errors, e...)
	case reflect.String:
		vz = reflect.ValueOf(values[0])
	case reflect.Ptr:
		v, e := assignType(s, etype.Elem(), values)
		vz = v
		errors = append(errors, e...)
	case reflect.Struct:
		v, e := assignFields(s, etype, values)
		vz = v
		errors = append(errors, e...)
	case reflect.Slice:
		fmt.Println(etype.Elem().Name())
		chaines := strings.Split(s, "__")
		if e := etype.Elem(); e.Kind() == reflect.Ptr && e.Elem().Name() == chaines[0] {
			fmt.Println("OK")
			if others := strings.Join(chaines[1:], "__"); len(others) > 0 {
				v, e := assignType(others, etype.Elem(), values)
				errors = append(errors, e...)
				vz = v
			}
		}
	}
	return vz, errors
}

func assignFields(s string, etype reflect.Type, values []string) (reflect.Value, []string) {
	errors := []string{}
	chaines := strings.Split(s, "__")
	if f, ok := etype.FieldByName(chaines[0]); ok {
		if others := strings.Join(chaines[1:], "__"); len(others) > 0 {
			return assignType(others, f.Type, values)
		}
		return assignType("", f.Type, values)
	}
	errors = append(errors, fmt.Sprintf("unknow fields's name '%v' in '%v'", s, etype.Name()))
	return reflect.Zero(etype), errors
}

func assignValue(etype reflect.Type, names string, values []string) (reflect.Value, []string) {
	errors := []string{}
	vz := reflect.Zero(etype)
	switch etype.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, e := assignInt(etype, values[0])
		errors = append(errors, e...)
		vz = v
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, e := assignUint(etype, values[0])
		errors = append(errors, e...)
		vz = v
	case reflect.Bool:
		v, e := assignBool(etype, values[0])
		errors = append(errors, e...)
		vz = v
	case reflect.String:
		vz = reflect.ValueOf(values[0])
	case reflect.Struct:
		v, e := assignStruct(etype, names, values)
		errors = append(errors, e...)
		vz = v
	case reflect.Ptr:
		v, e := assignValue(etype.Elem(), names, values)
		evalue := reflect.New(etype.Elem())
		evalue.Elem().Set(v)
		vz = evalue
		errors = append(errors, e...)
	case reflect.Slice:
		slice := reflect.MakeSlice(etype, 0, len(values))
		child := etype.Elem()
		for _, value := range values {
			v, e := assignValue(child, names, []string{value})
			errors = append(errors, e...)
			slice = reflect.Append(slice, v)
		}
		vz = slice
	}
	return vz, errors
}

func assignStruct(etype reflect.Type, field string, value []string) (reflect.Value, []string) {
	errors := []string{}
	vz := reflect.Zero(etype)
	names := strings.Split(field, "__")
	others := strings.Join(names[1:], "__")
	if f, ok := etype.FieldByName(names[0]); ok {
		v, e := assignValue(f.Type, others, value)
		errors = append(errors, e...)
		evalue := reflect.New(etype).Elem()
		evalue.FieldByName(names[0]).Set(v)
		vz = evalue
		return vz, errors
	}
	errors = append(errors, fmt.Sprintf("unknow fields's name '%v' in '%v'", names[0], etype.Name()))
	return vz, errors
}

func assignInt(etype reflect.Type, value string) (reflect.Value, []string) {
	errors := []string{}
	intvalue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		errors = append(errors, fmt.Sprintf("Invalid syntax: Cannot convert '%v' in Int", value))
		return reflect.Zero(etype), errors
	}
	evalue := reflect.New(etype).Elem()
	evalue.SetInt(intvalue)
	return evalue, errors
}

func assignUint(etype reflect.Type, value string) (reflect.Value, []string) {
	errors := []string{}

	uintvalue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		errors = append(errors, fmt.Sprintf("Invalid syntax: Cannot convert '%v' in Uint", value))
		return reflect.Zero(etype), errors
	}
	evalue := reflect.New(etype).Elem()
	evalue.SetUint(uintvalue)
	return evalue, errors
}

func assignBool(etype reflect.Type, value string) (reflect.Value, []string) {
	errors := []string{}
	boolvalue, err := strconv.ParseBool(value)
	if err != nil {
		errors = append(errors, fmt.Sprintf("Invalid syntax: Cannot convert '%v' in boolean", value))
		return reflect.Zero(etype), errors
	}
	evalue := reflect.New(etype).Elem()
	evalue.SetBool(boolvalue)
	return evalue, errors
}
