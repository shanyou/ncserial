package ncserial

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyBlock(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	emptyBlk.AddKVOption("proxy_set_header", "Host $host")
	emptyBlk.AddKVOption("proxy_set_header", "X-Real-IP $remote_addr")
	blkStr := emptyBlk.String()
	assert.True(t, strings.Contains(blkStr, "proxy_set_header"))
	assert.True(t, strings.Contains(blkStr, "X-Real-IP $remote_addr"))
}

func TestFindBlock(t *testing.T) {
	emptyBlk := NewEmptyBlock()

	blk := NewBlock("test")
	blk.AddKVOption("foo", "bar")
	sblk := NewBlock("sub")
	sblk.AddKVOption("aa", "bb")
	blk.AddDirective(sblk)
	emptyBlk.AddDirective(blk)

	fb, err := emptyBlk.FindDirectiveByName("aa")
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, fb.Name() == "aa")

	fb, err = emptyBlk.FindDirectiveByName("foo")
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, fb.Name() == "foo")

	fb, err = emptyBlk.FindDirectiveByName("sub")
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, fb.Name() == "sub")
}

func TestFindBlockFail(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	blk := NewBlock("test")
	blk.AddKVOption("foo", "bar")
	sblk := NewBlock("sub")
	sblk.AddKVOption("aa", "bb")
	blk.AddDirective(sblk)
	emptyBlk.AddDirective(blk)
	fb, err := emptyBlk.FindDirectiveByName("abc")
	assert.NotNil(t, err)
	assert.Nil(t, fb)
}
