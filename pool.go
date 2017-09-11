package chaos

import (
	"bytes"
	"sync"
)

var bbFree = sync.Pool{}

func ByteBufferPoolGet() *bytes.Buffer {
	if buf := bbFree.Get(); buf != nil {
		return buf.(*bytes.Buffer)
	} else {
		return &bytes.Buffer{}
	}
}

func put(b *bytes.Buffer) { bbFree.Put(b) }

func BytesBufferPoolFree(b *bytes.Buffer) {
	b.Reset()
	put(b)
}
