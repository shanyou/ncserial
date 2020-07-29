package ncserial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyValueOption(t *testing.T) {
	name := "foo"
	value := "bar"
	k := NewKVOption(name, value)
	assert.Equal(t, k.Name(), "foo")
	assert.Equal(t, k.Value(), "bar")

	k = NewKVOption(name, &value)
	assert.Equal(t, k.Value(), "bar")

	k = NewKVOption(name, false)
	assert.Equal(t, k.Value(), "off")

	k = NewKVOption(name, 5)
	assert.Equal(t, k.Value(), "5")

	k = NewKVOption(name, 0.5)
	assert.Equal(t, k.Value(), "0.5")

	k = NewKVOption(name, []string{"a", "b", "c"})
	assert.Equal(t, k.Value(), "a b c")
	assert.Equal(t, k.String(), "\nfoo a b c;")

	k = NewKVOption(name, nil)
	assert.Equal(t, k.String(), "\nfoo ;")

	k = NewOption([]string{"foo", "bar"})
	assert.Equal(t, k.Name(), "foo")
	assert.Equal(t, k.Value(), "bar")
	assert.Equal(t, k.String(), "\nfoo bar;")

	k = NewOption([]string{"foo"})
	assert.Equal(t, k.Name(), "foo")
	assert.Equal(t, k.Value(), "")
	assert.Equal(t, k.String(), "\nfoo ;")

	k = NewOption([]string{"foo", "bar", "baz"})
	assert.Equal(t, k.Name(), "foo")
	assert.Equal(t, k.String(), "\nfoo bar baz;")
}
