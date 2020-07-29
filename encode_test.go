package ncserial

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testMapWithTag struct {
	HTTP Directives `kv:"http"`
}

type testMapWithOutTag struct {
	HTTP Directives
}

func TestMapWithTag(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	test := testMapWithTag{
		HTTP: Directives{
			NewOption([]string{"abc", "bcd"}),
		},
	}
	ds, err := MarshalD(test)
	if err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		emptyBlk.AddDirective(d)
	}
	str := emptyBlk.String()
	assert.True(t, strings.Contains(strings.ToLower(str), "http"))
	assert.True(t, strings.Contains(strings.ToLower(str), "abc bcd;"))
}

func TestMapWithOutTag(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	test := testMapWithOutTag{
		HTTP: Directives{
			NewOption([]string{"abc", "bcd"}),
		},
	}
	ds, err := MarshalD(test)
	if err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		emptyBlk.AddDirective(d)
	}

	str := emptyBlk.String()
	assert.False(t, strings.Contains(strings.ToLower(str), "http"))
	assert.True(t, strings.Contains(strings.ToLower(str), "abc bcd;"))
}
