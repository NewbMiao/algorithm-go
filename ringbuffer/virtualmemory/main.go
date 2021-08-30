package ringbuffer

import (
	"errors"
)

type T interface{}

var ErrIsEmpty = errors.New("ringbuffer is empty")

// https://zh.wikipedia.org/wiki/%E7%92%B0%E5%BD%A2%E7%B7%A9%E8%A1%9D%E5%8D%80#%E9%95%9C%E5%83%8F%E6%8C%87%E7%A4%BA%E4%BD%8D
// RingBuffer is a ring buffer for common types.
// It never is full and always grows if it will be full.
// It is not thread-safe(goroutine-safe) so you must use the lock-like synchronization primitive to use it in multiple writers and multiple readers.
type RingBuffer struct {
	buf         []T
	initialSize int
	size        int
	r           int // read pointer
	w           int // write pointer
}

func NewRingBuffer(initialSize int) *RingBuffer {
	if initialSize <= 0 || initialSize&(initialSize-1) != 0 {
		panic("initial size must be power of 2")
	}

	return &RingBuffer{
		buf:         make([]T, initialSize),
		initialSize: initialSize,
		size:        initialSize,
	}
}

func (r *RingBuffer) incrCursor(p *int) {
	*p = (*p + 1) & (2*r.size - 1)
}
func (r *RingBuffer) Read() (T, error) {
	if r.IsEmpty() {
		return nil, ErrIsEmpty
	}

	v := r.buf[r.r&(r.size-1)]
	r.incrCursor(&r.r)

	return v, nil
}

func (r *RingBuffer) Pop() T {
	v, err := r.Read()
	if err == ErrIsEmpty { // Empty
		panic(ErrIsEmpty.Error())
	}

	return v
}

func (r *RingBuffer) Peek() T {
	if r.IsEmpty() { // Empty
		panic(ErrIsEmpty.Error())
	}

	v := r.buf[r.r&(r.size-1)]
	return v
}

func (r *RingBuffer) Write(v T) {
	r.buf[r.w&(r.size-1)] = v
	r.incrCursor(&r.w)

	if r.IsFull() { // full
		r.grow()
	}
}

func (r *RingBuffer) grow() {
	var size int
	if r.size < 1024 {
		size = r.size * 2
	} else {
		size = r.size + r.size/4
	}

	buf := make([]T, size)

	copy(buf[0:], r.buf[r.r:])
	copy(buf[r.size-r.r:], r.buf[0:r.r])

	r.r = 0
	r.w = r.size
	r.size = size
	r.buf = buf
}

func (r *RingBuffer) IsEmpty() bool {
	return r.r == r.w
}

func (r *RingBuffer) IsFull() bool {
	return (r.r ^ r.size) == r.w
}

// Capacity returns the size of the underlying buffer.
func (r *RingBuffer) Capacity() int {
	return r.size
}

func (r *RingBuffer) Len() int {
	if r.r == r.w {
		return 0
	}

	if r.w > r.r {
		return r.w - r.r
	}

	return r.size - r.r + r.w
}

func (r *RingBuffer) Reset() {
	r.r = 0
	r.w = 0
	r.size = r.initialSize
	r.buf = make([]T, r.initialSize)
}
