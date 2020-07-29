package ncserial

import (
	"errors"
	"sync"
)

type decoder struct {
	sync.Mutex
	*scanner
}

//UnmarshalD data to interface through directive
func UnmarshalD(data []byte, v interface{}) error {
	return errors.New("not implements yet")
}
