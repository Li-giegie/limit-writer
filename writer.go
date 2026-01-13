package limit_writer

import (
	"errors"
	"io"
)

var ErrOverflow = errors.New("write err: data len overflow")

// New limit 限制写入大小和缓冲器容量
func New(w io.Writer, limit int) *Writer {
	return NewWriterSize(w, limit, limit)
}

// NewWriterSize limit 限制写入大小，size 缓冲器容量
func NewWriterSize(w io.Writer, limit, size int) *Writer {
	return NewWriterBuffer(w, limit, make([]byte, 0, size))
}

func NewWriterBuffer(w io.Writer, limit int, buf []byte) *Writer {
	return &Writer{
		w:     w,
		limit: limit,
		buf:   buf,
	}
}

// Writer 实现了限制向底层Writer写入一次数据的大小；
// 当待写入p字节数 > 缓冲区总容量，返回ErrOverflow；
// 当待写入p字节数 > 缓冲区剩余容量，先将缓冲区数据写入底层Writer，再写入p到缓存区
// 当待写入p字节数 <= 缓冲区剩余容量，copy到缓冲区中
type Writer struct {
	limit int
	buf   []byte
	w     io.Writer
}

func (w *Writer) Write(p []byte) (n int, err error) {
	if len(p) > w.limit {
		return 0, ErrOverflow
	}
	if len(w.buf)+len(p) > w.limit {
		if n, err = w.w.Write(w.buf); err != nil {
			return
		}
		w.buf = w.buf[:0]
	}
	w.buf = append(w.buf, p...)
	return len(p), nil
}

func (w *Writer) Flush() (err error) {
	_, err = w.w.Write(w.buf)
	return
}

func (w *Writer) Size() int {
	return len(w.buf)
}

func (w *Writer) Reset() {
	w.buf = w.buf[:0]
}
