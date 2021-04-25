package atomix

import (
	"strconv"
	"sync/atomic"
)

const CacheLine = 64

// AlignedInt64 is an atomic wrapper around an int64.
type AlignedInt64 struct {
	_     [CacheLine - 8]byte
	value int64
	_     [CacheLine]byte
}

// NewInt64 creates an Int64.
func NewAlignedInt64(i int64) *AlignedInt64 {
	return &AlignedInt64{value: i}
}

func (a *AlignedInt64) String() string {
	return strconv.FormatInt(a.Load(), 10)
}

// Load atomically the value.
func (a *AlignedInt64) Load() int64 {
	return atomic.LoadInt64(&a.value)
}

// Store atomically the given value.
func (a *AlignedInt64) Store(n int64) {
	atomic.StoreInt64(&a.value, n)
}

// Swap atomically and return the old value.
func (a *AlignedInt64) Swap(n int64) int64 {
	return atomic.SwapInt64(&a.value, n)
}

// Add atomically and return the new value.
func (a *AlignedInt64) Add(n int64) int64 {
	return atomic.AddInt64(&a.value, n)
}

// Sub atomically and return the new value.
func (a *AlignedInt64) Sub(n int64) int64 {
	return atomic.AddInt64(&a.value, -n)
}

// Inc atomically and return the new value.
func (a *AlignedInt64) Inc() int64 {
	return a.Add(1)
}

// Dec atomically and return the new value.
func (a *AlignedInt64) Dec() int64 {
	return a.Sub(1)
}

// CAS is an atomic Compare-And-Swap operation.
func (a *AlignedInt64) CAS(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&a.value, old, new)
}
