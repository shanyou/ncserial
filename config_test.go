package ncserial

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultRestyConfig(t *testing.T) {
	prefix := "/usr/local/nginx"
	logpath := "logs"
	libPath := "/usr/local/nginx/lib"
	conf := NewDefaultRestyConfig(prefix, logpath, libPath)
	emptyBlk := NewEmptyBlock()
	emptyBlk.AddInterface(conf)
	assert.True(t, strings.Contains(emptyBlk.String(), "application/zip zip;"))
	assert.True(t, strings.Contains(emptyBlk.String(), "pid logs/nginx.pid;"))
	assert.True(t, strings.Contains(emptyBlk.String(), "use epoll;"))
	assert.True(t, strings.Contains(emptyBlk.String(), "lua_package_cpath '/usr/local/nginx/lib/?.so;;';"))
}
