package limit_writer

import (
	"fmt"
	"log"
	"testing"
)

type trace struct {
	i int
}

func (i *trace) Write(p []byte) (n int, err error) {
	i.i++
	log.Println(i.i, string(p))
	return len(p), nil
}

func TestWrite(t *testing.T) {
	w := New(&trace{}, 5)
	defer w.Flush()
	w.Write([]byte("123"))
	w.Write([]byte("4567"))

	fmt.Println(w.Write([]byte("333333")))

	w.Write([]byte("11111"))
	w.Write([]byte("2"))
	w.Write([]byte("33333"))

	// out
	// 2026/01/12 15:13:34 1 123
	// 0 write err: data len overflow
	// 2026/01/12 15:13:35 2 4567
	// 2026/01/12 15:13:35 3 11111
	// 2026/01/12 15:13:35 4 2
	// 2026/01/12 15:13:35 5 33333
}
