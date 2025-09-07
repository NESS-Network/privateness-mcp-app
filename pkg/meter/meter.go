package meter

import (
	"sync/atomic"
	"time"
	"io"
)

type Meter struct {
	bytesIn  atomic.Uint64
	bytesOut atomic.Uint64
	start    time.Time
}

func New() *Meter { return &Meter{start: time.Now()} }

func (m *Meter) BytesIn() uint64  { return m.bytesIn.Load() }
func (m *Meter) BytesOut() uint64 { return m.bytesOut.Load() }
func (m *Meter) Duration() time.Duration { return time.Since(m.start) }

type countingRW struct {
	inner io.ReadWriteCloser
	m     *Meter
}

func (c *countingRW) Read(p []byte) (int, error) {
	n, err := c.inner.Read(p)
	c.m.bytesIn.Add(uint64(n))
	return n, err
}
func (c *countingRW) Write(p []byte) (int, error) {
	n, err := c.inner.Write(p)
	c.m.bytesOut.Add(uint64(n))
	return n, err
}
func (c *countingRW) Close() error { return c.inner.Close() }

func Wrap(inner io.ReadWriteCloser, m *Meter) io.ReadWriteCloser {
	return &countingRW{inner: inner, m: m}
}
