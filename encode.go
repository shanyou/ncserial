package ncserial

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

var (
	marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()
	stringType    = reflect.TypeOf("")
	stringPtrType = reflect.TypeOf((*string)(nil))
)

//MarshalD marshal struct to nginx directive
func MarshalD(i interface{}) ([]Directive, error) {
	//如果接口实现Marshaler则直接调用
	iv, ok := i.(Marshaler)
	if ok {
		return iv.MarshalD()
	}

	//is Directive return it self
	dv, ok := i.(Directive)
	if ok {
		return []Directive{dv}, nil
	}

	v := getValue(reflect.ValueOf(i))
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("type %s is not supported", t.Kind())
	}

	var directives []Directive
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// skip unexported fields. from godoc:
		// PkgPath is the package path that qualifies a lower case (unexported)
		// field name. It is empty for upper case (exported) field names.
		if f.PkgPath != "" {
			continue
		}

		fv := getValue(v.Field(i))
		if !fv.IsValid() {
			return nil, errors.New("invalid input value")
		}

		key, omit, findTag := readTag(f) //omit 忽略, "-"忽略
		if strings.Compare(key, "-") == 0 {
			// ignore
			continue
		}

		if f.Type.Implements(marshalerType) && !findTag {
			fv := v.Field(i)
			if fv.IsNil() {
				continue
			}
			m, _ := fv.Interface().(Marshaler)
			ds, err := m.MarshalD()
			if err == nil {
				for _, dd := range ds {
					directives = append(directives, dd)
				}
				continue
			}
		}
		var d Directive
		switch fv.Kind() {
		case reflect.Func, reflect.Chan:
			continue
		case reflect.Interface, reflect.Struct:
			b := NewBlock(key)
			dd, err := MarshalD(fv.Interface())
			if err != nil {
				continue
			}
			for _, sd := range dd {
				b.AddDirective(sd)
			}
			directives = append(directives, b)
		case reflect.Ptr, reflect.UnsafePointer:
			if fv.IsNil() {
				if !omit {
					d = BuildDirective(key, nil)
				} else {
					continue
				}
			} else {
				d = BuildDirective(key, fv.Interface())
			}
			directives = append(directives, d)
		case reflect.Slice:
			if fv.IsNil() {
				if !omit {
					d = BuildDirective(key, nil)
					directives = append(directives, d)
				} else {
					continue
				}
			} else {
				if findTag {
					b := NewBlock(key)
					for i := 0; i < fv.Len(); i++ {
						ifv := fv.Index(i)
						dd, err := MarshalD(ifv.Interface())
						if err != nil {
							continue
						}
						for _, sd := range dd {
							b.AddDirective(sd)
						}
					}
					directives = append(directives, b)
				} else {
					for i := 0; i < fv.Len(); i++ {
						ifv := fv.Index(i)
						dd, err := MarshalD(ifv.Interface())
						if err != nil {
							continue
						}
						for _, sd := range dd {
							directives = append(directives, sd)
						}
					}
				}
			}
		case reflect.Map:
			var blk *Block
			if findTag {
				blk = NewBlock(key)
			}
			for _, e := range fv.MapKeys() {
				if e.Kind() != reflect.String {
					continue
				}
				vv := fv.MapIndex(e)
				d = BuildDirective(e.String(), vv.Interface())
				if !findTag {
					// 如果没有标记则直接加入到当前struct的解析结果中
					directives = append(directives, d)
				} else {
					// 如果有标记则创建一个block加入到block中
					blk.AddDirective(d)
				}
			}

			if findTag {
				directives = append(directives, blk)
			}
		default:
			if fv.IsZero() && omit {
				continue
			}
			d = BuildDirective(key, fv.Interface())
			directives = append(directives, d)
		}
	}

	sort.Sort(Directives(directives))
	return directives, nil
}

//getValue 判断是否是Ptr如果是就调用Elm()
func getValue(value reflect.Value) reflect.Value {
	if !value.IsValid() {
		return reflect.ValueOf(nil)
	}
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return value
		}
		return getValue(value.Elem())
	default:
		return value
	}
}

// read tag like `kv:"email,omitempty"`
// if no tag last key return false
func readTag(f reflect.StructField) (string, bool, bool) {
	val, ok := f.Tag.Lookup("kv")
	if !ok {
		return f.Name, true, false
	}
	opts := strings.Split(val, ",")
	omit := false
	if len(opts) == 2 {
		omit = opts[1] == "omitempty"
	}
	return opts[0], omit, true
}
