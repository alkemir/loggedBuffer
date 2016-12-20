package Buffer

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

// Buffer wraps a bytes.Buffer with logging capabilities.
type Buffer struct {
	log []string
	*bytes.Buffer
}

func (b *Buffer) logCall(call string, ret interface{}) {
	_, file, line, _ := runtime.Caller(3)
	b.log = append(b.log, fmt.Sprintf("%s|%s:%d> %s=%v\n", time.Now().Format(timeFormat), file, line, call, ret))
}

// Logs returns the logs up to this point.
func (b *Buffer) Logs() []string {
	ret := make([]string, len(b.log))
	for i := 0; i < len(ret); i++ {
		ret[i] = b.log[i]
	}
	return ret
}

// NewBuffer wraps bytes.NewBuffer with logging capabilities.
func NewBuffer(buf []byte) *Buffer {
	ret := &Buffer{Buffer: bytes.NewBuffer(buf)}
	ret.logCall("NewBuffer", nil)
	return ret
}

// NewBufferString wraps bytes.NewBufferString with logging capabilities.
func NewBufferString(s string) *Buffer {
	ret := &Buffer{Buffer: bytes.NewBufferString(s)}
	ret.logCall("NewBufferString", nil)
	return ret
}

// Bytes wraps bytes.Bytes with logging capabilities.
func (b *Buffer) Bytes() []byte {
	if b == nil {
		return (&bytes.Buffer{}).Bytes()
	}
	ret := b.Buffer.Bytes()
	b.logCall("Bytes", ret)
	return ret
}

// Cap wraps bytes.Cap with logging capabilities.
func (b *Buffer) Cap() int {
	ret := b.Buffer.Cap()
	b.logCall("Cap", ret)
	return ret
}

// Grow wraps bytes.Grow with logging capabilities.
func (b *Buffer) Grow(n int) {
	b.Buffer.Grow(n)
	b.logCall("Grow", nil)
}

// Len  wraps bytes.Len with logging capabilities.
func (b *Buffer) Len() int {
	ret := b.Buffer.Len()
	b.logCall("Len", ret)
	return ret
}

// Next wraps bytes.Next with logging capabilities.
func (b *Buffer) Next(n int) []byte {
	ret := b.Buffer.Next(n)
	b.logCall("Next", nil)
	return ret
}

// Read wraps bytes.Read with logging capabilities.
func (b *Buffer) Read(p []byte) (n int, err error) {
	n, err = b.Buffer.Read(p)
	b.logCall("Read", n)
	return n, err
}

// ReadByte wraps bytes.ReadByte with logging capabilities.
func (b *Buffer) ReadByte() (byte, error) {
	ret, err := b.Buffer.ReadByte()
	b.logCall("ReadByte", err)
	return ret, err
}

// ReadBytes wraps bytes.ReadBytes with logging capabilities.
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
	line, err = b.Buffer.ReadBytes(delim)
	b.logCall("ReadBytes", err)
	return line, err
}

// ReadFrom wraps bytes.ReadFrom with logging capabilities.
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	n, err = b.Buffer.ReadFrom(r)
	b.logCall("ReadFrom", n)
	return n, err
}

// ReadRune wraps bytes.ReadRune with logging capabilities.
func (b *Buffer) ReadRune() (r rune, size int, err error) {
	r, size, err = b.Buffer.ReadRune()
	b.logCall("ReadRune", err)
	return r, size, err
}

// ReadString wraps bytes.ReadString with logging capabilities.
func (b *Buffer) ReadString(delim byte) (line string, err error) {
	line, err = b.Buffer.ReadString(delim)
	b.logCall("ReadString", err)
	return line, err
}

// Reset wraps bytes.Reset with logging capabilities.
func (b *Buffer) Reset() {
	b.Buffer.Reset()
	b.logCall("Reset", nil)
}

// String wraps bytes.String with logging capabilities.
func (b *Buffer) String() string {
	if b == nil {
		return "<nil>"
	}
	ret := b.Buffer.String()
	b.logCall("String", nil)
	return ret
}

// Truncate wraps bytes.Truncate with logging capabilities.
func (b *Buffer) Truncate(n int) {
	b.Buffer.Truncate(n)
	b.logCall("Truncate", nil)
}

// UnreadByte wraps bytes.UnreadByte with logging capabilities.
func (b *Buffer) UnreadByte() error {
	err := b.Buffer.UnreadByte()
	b.logCall("UnreadByte", err)
	return err
}

// UnreadRune wraps bytes.UnreadRune with logging capabilities.
func (b *Buffer) UnreadRune() error {
	err := b.Buffer.UnreadRune()
	b.logCall("UnreadRune", err)
	return err
}

// Write wraps bytes.Write with logging capabilities.
func (b *Buffer) Write(p []byte) (n int, err error) {
	n, err = b.Buffer.Write(p)
	b.logCall("Write", n)
	return n, err
}

// WriteByte wraps bytes.WriteByte with logging capabilities.
func (b *Buffer) WriteByte(c byte) error {
	err := b.Buffer.WriteByte(c)
	b.logCall("WriteByte", err)
	return err
}

// WriteRune wraps bytes.WriteRune with logging capabilities.
func (b *Buffer) WriteRune(r rune) (n int, err error) {
	n, err = b.Buffer.WriteRune(r)
	b.logCall("WriteRune", n)
	return n, err
}

// WriteString wraps bytes.WriteString with logging capabilities.
func (b *Buffer) WriteString(s string) (n int, err error) {
	n, err = b.Buffer.WriteString(s)
	b.logCall("WriteString", n)
	return n, err
}

// WriteTo wraps bytes.WroteTo with logging capabilities.
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	n, err = b.Buffer.WriteTo(w)
	b.logCall("WriteTo", n)
	return n, err
}
