package ncserial

import (
	"fmt"
	"reflect"
	"strconv"
)

//CovertToString convert interface to string
func CovertToString(value interface{}) string {
	if value == nil {
		return ""
	}
	var val string
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.String {
		val = v.String()
	} else if v.Type() == stringPtrType {
		val = v.Elem().String()
	} else if v.Kind() == reflect.Bool {
		if v.Bool() {
			val = "on"
		} else {
			val = "off"
		}
	} else if v.Kind() == reflect.Int {
		val = strconv.FormatInt(v.Int(), 10)
	} else if v.Kind() == reflect.Float64 {
		val = fmt.Sprintf("%g", v.Float())
	} else if v.Kind() == reflect.Slice {
		val = ""
		len := v.Len()
		for i := 0; i < len; i++ {
			if i == (len - 1) {
				val = val + v.Index(i).String()
			} else {
				val = val + v.Index(i).String() + " "
			}
		}
	}

	return val
}

//BuildDirective build directive with name , value
func BuildDirective(name string, value interface{}) Directive {
	t, ok := value.(*Block)
	if ok {
		return (Directive)(t)
	}

	return NewKVOption(name, value)
}

//MergeDirectives merge two directives
func MergeDirectives(d1, d2 Directives) Directives {
	for _, d := range d2 {
		d1 = append(d1, d)
	}

	return d1
}

//MergeOptions merge two options
func MergeOptions(s1, s2 Options) Options {
	for _, v := range s2 {
		s1 = append(s1, v)
	}

	return s1
}
