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
	panic("defaultPool.Put not implemented...yet...")
}

func (p *defaultPool) Get() *bytes.Buffer {
	panic("defaultPool.Get not implemented...yet...")
	return nil
}

// SetBufferPool ...
func SetBufferPool(bp BufferPool) {
	panic("BufferPool. not implemented...yet...")
}
