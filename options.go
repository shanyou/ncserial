package ncserial

import "fmt"

//Option for nginx command
type Option []string

//MarshalD implememt Marshaler
func (o Option) MarshalD() ([]Directive, error) {
	return []Directive{
		NewOption(o),
	}, nil
}

//Options represent list of words as kvoption
type Options []Option

//MarshalD implememt Marshaler
func (o Options) MarshalD() ([]Directive, error) {
	var ds []Directive
	for _, opt := range o {
		ds = append(ds, NewOption(opt))
	}

	return ds, nil
}

//KeyValueOption A key/value directive. implements Directive
type KeyValueOption struct {
	Base
	name  string
	value string
}

//NewOption with given strings
func NewOption(words []string) *KeyValueOption {
	if len(words) == 0 {
		return nil // empty Option
	} else if len(words) == 1 {
		return NewKVOption(words[0], nil)
	} else if len(words) == 2 {
		return NewKVOption(words[0], words[1])
	} else {
		return NewKVOption(words[0], words[1:])
	}
}

//NewKVOption init keyvalue option with key and give value
func NewKVOption(name string, value interface{}) *KeyValueOption {
	return &KeyValueOption{
		Base:  NewDefaultBase(),
		name:  name,
		value: CovertToString(value),
	}
}

func (k KeyValueOption) String() string {
	return k.Base.Render(fmt.Sprintf("%s %s;", k.name, k.value))
}

//Value implements Directive
func (k KeyValueOption) Value() interface{} {
	return k.value
}

//Name implements Directive
func (k KeyValueOption) Name() string {
	return k.name
}
