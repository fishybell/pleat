package pleat

import (
	"bytes"
	"sync"
)

type BufferPool interface {
	Put(*bytes.Buffer)
	Get() *bytes.Buffer
}

type defaultPool struct {
	pool *sync.Pool
}

func (p *defaultPool) Put(buf *bytes.Buffer) {
	panic("not implemented...yet...")
}

func (p *defaultPool) Get() *bytes.Buffer {
	panic("not implemented...yet...")
	return nil
}

// SetBufferPool allows to replace the default logrus buffer pool
// to better meets the specific needs of an application.
func SetBufferPool(bp BufferPool) {
	panic("not implemented...yet...")
}
