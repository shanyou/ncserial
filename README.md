ncserial
===
nginx config parser for the Go language

# Instruction
`ncserial` use interface `Directive` and `BlockDirective` to represent nginx config directive.
```go
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
//BlockDirective nginx block type directive with braces
type BlockDirective interface {
	Directive
	AddDirective(d Directive)
	AddInterface(i interface{})
	FindDirectiveByName(name string) (Directive, error)
}
```

use reflect method to marshal struct with tag to directive and output to config file
```go
// Events nginx events directive
type Events struct {
	WorkerConnections int    `kv:"worker_connections"`
	Use               string `kv:"use"`          //epoll
	MultiAccept       bool   `kv:"multi_accept"` //default on
}

//HTTP nginx http config section
type HTTP struct {
	DefalutType string  `kv:"default_type,omitempty"`
	LogFormat   string  `kv:"log_format"`
	MimeTypes   Options `kv:"types"`
	AccessLog   string  `kv:"access_log,omitempty"`
	ErrorLog    string  `kv:"error_log,omitempty"`
	SendFile    bool    `kv:"sendfile"`
	ExtConfig   Options
}

//Config represent nginx config
// follows https://www.nginx.com/resources/wiki/start/topics/examples/full/ to build nginx base config
type Config struct {
	User            string `kv:"user,omitempty"`
	WorkerProcesses string `kv:"worker_processes"`
	PId             string `kv:"pid"`
	ErrorLog        string `kv:"error_log,omitempty"`
	LimitNofile     int    `kv:"worker_rlimit_nofile"`
	Events          Events `kv:"events"`
	HTTP            HTTP   `kv:"http"`
	Extras          Options
}
```
# Example
```go
package main
import (
    "fmt"
    "github.com/shanyou/ncserial"
)
func main() {
    prefix := "/usr/local/nginx"
	logpath := "logs"
	libPath := "/usr/local/nginx/lib"
	conf := ncserial.NewDefaultRestyConfig(prefix, logpath, libPath)
	emptyBlk := ncserial.NewEmptyBlock()
    emptyBlk.AddInterface(conf)
    fmt.Println(emptyBlk)
}
```