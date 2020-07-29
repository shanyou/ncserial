package ncserial

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ITRender interface {
	print()
}
type Test struct {
	foo int
}

func (t *Test) print() {
	fmt.Printf("%d/n", t.foo)
}

func TestBaseRender(t *testing.T) {
	b := NewBase(1, 4, 'a', nil)
	t.Log(b.Render("server"))
	assert.Equal(t, b.Render("server"), "\naaaaserver", "Base render fail")
}

func TestRefect(t *testing.T) {
	var x interface{} = []int{1, 2, 3}
	xType := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	t.Log(xType, xValue) // "[]int [1 2 3]"

	tt := &Test{5}
	tType := reflect.TypeOf(tt)
	testType := reflect.TypeOf((*Test)(nil)).Elem()
	renderType := reflect.TypeOf(new(ITRender)).Elem()
	t.Log(reflect.TypeOf(tt))

	assert.True(t, tType.Implements(renderType), "Test not implements IRender")
	assert.Equal(t, reflect.TypeOf(tt).Elem(), testType, "tt is not a type of Test")

}

func TestBaseGetIndent(t *testing.T) {
	b := NewBase(2, 4, 'a', nil)
	assert.Equal(t, b.GetIndent(), "aaaaaaaa", "Base getIndent fail")
}
