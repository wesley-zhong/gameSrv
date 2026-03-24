package ringbuffer

import (
	"errors"
	"unsafe"
)

var (
	ErrTooManyDataToWrite = errors.New("too many data to write")
	ErrIsFull             = errors.New("ringbuffer is full")
	ErrIsEmpty            = errors.New("ringbuffer is empty")
	ErrAccuqireLock       = errors.New("no lock to accquire")
)

// RingBuffer is a circular buffer that implement io.ReaderWriter interface.
type RingBuffer struct {
	buf    []byte
	size   int
	r      int // next position to read
	w      int // next position to write
	isFull bool
	mu     Mutex
}

// New returns a new RingBuffer whose buffer has the given size.
func NewRingBuff(size int) *RingBuffer {
	return &RingBuffer{
		buf:  make([]byte, size),
		size: size,
	}
}

// Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any error encountered. Even if Read returns n < len(p), it may use all of p as scratch space during the call. If some data is available but not len(p) bytes, Read conventionally returns what is available instead of waiting for more.
// When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it returns the number of bytes read. It may return the (non-nil) error from the same call or return the error (and n == 0) from a subsequent call.
// Callers should always process the n > 0 bytes returned before considering the error err. Doing so correctly handles I/O errors that happen after reading some bytes and also both of the allowed EOF behaviors.
func (r *RingBuffer) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.read(p)
}

// TryRead read up to len(p) bytes into p like Read but it is not blocking.
// If it has not succeeded to accquire the lock, it return 0 as n and ErrAccuqireLock.
func (r *RingBuffer) TryRead(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if !r.mu.TryLock() {
		return 0, ErrAccuqireLock
	}
	defer r.mu.Unlock()
	return r.read(p)
}

func (r *RingBuffer) read(p []byte) (n int, err error) {
	if r.isEmpty() {
		return 0, ErrIsEmpty
	}

	n = r.available()
	if n > len(p) {
		n = len(p)
	}

	r.readTo(p, n)

	r.isFull = false
	return n, nil
}

func (r *RingBuffer) isEmpty() bool {
	return r.w == r.r && !r.isFull
}

func (r *RingBuffer) available() int {
	if r.w > r.r {
		return r.w - r.r
	}
	return r.size - r.r + r.w
}

func (r *RingBuffer) readTo(p []byte, n int) {
	if r.r+n <= r.size {
		copy(p, r.buf[r.r:r.r+n])
		r.r = r.r + n
	} else {
		c1 := r.size - r.r
		copy(p, r.buf[r.r:])
		c2 := n - c1
		copy(p[c1:], r.buf[:c2])
		r.r = c2
	}
}

// ReadByte reads and returns the next byte from the input or ErrIsEmpty.
func (r *RingBuffer) ReadByte() (b byte, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.isEmpty() {
		return 0, ErrIsEmpty
	}
	b = r.buf[r.r]
	r.r++
	if r.r == r.size {
		r.r = 0
	}
	r.isFull = false
	return b, nil
}

// Write writes len(p) bytes from p to the underlying buf.
// It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early.
// Write returns a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
func (r *RingBuffer) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.write(p)
}

// TryWrite writes len(p) bytes from p to the underlying buf like Write, but it is not blocking.
// If it has not succeeded to accquire the lock, it return 0 as n and ErrAccuqireLock.
func (r *RingBuffer) TryWrite(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if !r.mu.TryLock() {
		return 0, ErrAccuqireLock
	}
	defer r.mu.Unlock()
	return r.write(p)
}

func (r *RingBuffer) write(p []byte) (n int, err error) {
	if r.isFull {
		return 0, ErrIsFull
	}

	avail := r.free()
	if len(p) > avail {
		err = ErrTooManyDataToWrite
		p = p[:avail]
	}
	n = len(p)

	r.writeFrom(p, n)
	return n, err
}

func (r *RingBuffer) free() int {
	if r.w < r.r {
		return r.r - r.w
	}
	return r.size - r.w + r.r
}

func (r *RingBuffer) writeFrom(p []byte, n int) {
	if r.w+n <= r.size {
		copy(r.buf[r.w:], p)
		r.w += n
	} else {
		c1 := r.size - r.w
		copy(r.buf[r.w:], p[:c1])
		c2 := n - c1
		copy(r.buf[:c2], p[c1:])
		r.w = c2
	}

	if r.w == r.size {
		r.w = 0
	}
	if r.w == r.r {
		r.isFull = true
	}
}

// WriteByte writes one byte into buffer, and returns ErrIsFull if buffer is full.
func (r *RingBuffer) WriteByte(c byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.writeByte(c)
}

// TryWriteByte writes one byte into buffer without blocking.
// If it has not succeeded to accquire the lock, it return ErrAccuqireLock.
func (r *RingBuffer) TryWriteByte(c byte) error {
	if !r.mu.TryLock() {
		return ErrAccuqireLock
	}
	defer r.mu.Unlock()
	return r.writeByte(c)
}

func (r *RingBuffer) writeByte(c byte) error {
	if r.w == r.r && r.isFull {
		return ErrIsFull
	}
	r.buf[r.w] = c
	r.w++
	if r.w == r.size {
		r.w = 0
	}
	if r.w == r.r {
		r.isFull = true
	}
	return nil
}

// Length return the length of available read bytes.
func (r *RingBuffer) Length() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.w == r.r {
		if r.isFull {
			return r.size
		}
		return 0
	}
	return r.available()
}

// Capacity returns the size of the underlying buffer.
func (r *RingBuffer) Capacity() int {
	return r.size
}

// Free returns the length of available bytes to write.
func (r *RingBuffer) Free() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.w == r.r {
		if r.isFull {
			return 0
		}
		return r.size
	}
	return r.free()
}

// WriteString writes the contents of the string s to buffer, which accepts a slice of bytes.
func (r *RingBuffer) WriteString(s string) (n int, err error) {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	buf := *(*[]byte)(unsafe.Pointer(&h))
	return r.Write(buf)
}

// Bytes returns all available read bytes. It does not move the read pointer and only copy the available data.
func (r *RingBuffer) Bytes() []byte {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.w == r.r {
		if r.isFull {
			buf := make([]byte, r.size)
			copy(buf, r.buf[r.r:])
			copy(buf[r.size-r.r:], r.buf[:r.w])
			return buf
		}
		return nil
	}

	if r.w > r.r {
		buf := make([]byte, r.w-r.r)
		copy(buf, r.buf[r.r:r.w])
		return buf
	}

	n := r.available()
	buf := make([]byte, n)
	r.copyTo(buf, n)
	return buf
}

func (r *RingBuffer) copyTo(buf []byte, n int) {
	if r.r+n <= r.size {
		copy(buf, r.buf[r.r:r.r+n])
	} else {
		c1 := r.size - r.r
		copy(buf, r.buf[r.r:])
		c2 := n - c1
		copy(buf[c1:], r.buf[:c2])
	}
}

// IsFull returns this ringbuffer is full.
func (r *RingBuffer) IsFull() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.isFull
}

// IsEmpty returns this ringbuffer is empty.
func (r *RingBuffer) IsEmpty() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.isEmpty()
}

// Reset the read pointer and writer pointer to zero.
func (r *RingBuffer) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.r = 0
	r.w = 0
	r.isFull = false
}