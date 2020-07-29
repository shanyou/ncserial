package ncserial

import (
	"strings"
)

// Directive nginx config directive
type Directive interface {
	Name() string
	Value() interface{}
	String() string
	Parent() interface{}
	SetParent(parent interface{})
	SetIndentLevel(level int)
	GetIndentLevel() int
}

// Directives slice of directive
// usage:
// sort.Sort(Directives(sliceOfDirective))
type Directives []Directive

// Len implements sort.Interface
func (d Directives) Len() int { return len(d) }

// Less implments sort.Interface
func (d Directives) Less(i, j int) bool {
	siLower := strings.ToLower(d[i].Name())
	sjLower := strings.ToLower(d[j].Name())
	if siLower == sjLower {
		return d[i].Name() < d[j].Name()
	}
	return siLower < sjLower
}

// Swap implements sort.Interface
func (d Directives) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

//BlockDirective nginx block type directive with braces
type BlockDirective interface {
	Directive
	AddDirective(d Directive)
	AddInterface(i interface{})
	FindDirectiveByName(name string) (Directive, error)
}

// Marshaler is the interface implemented by types that
// can marshal themselves into valid Directive.
type Marshaler interface {
	MarshalD() ([]Directive, error)
}

// Unmarshaler is the interface implemented by types
// that can unmarshal a Nginx directive
type Unmarshaler interface {
	UnmarshalD([]byte) error
}
