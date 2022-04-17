package atomix

import (
	"strconv"
	"sync/atomic"
)

// Uint32 is an atomic wrapper around an uint32.
type Uint32 struct {
	atomicType
	value uint32
}

// NewUint32 creates an Uint32.
func NewUint32(i uint32) *Uint32 {
	return &Uint32{value: i}
}

func (i *Uint32) String() string {
	return strconv.FormatUint(uint64(i.Load()), 10)
}

// Load atomically the value.
func (i *Uint32) Load() uint32 {
	return atomic.LoadUint32(&i.value)
}

// Store atomically the given value.
func (i *Uint32) Store(n uint32) {
	atomic.StoreUint32(&i.value, n)
}

// Swap atomically and return the old value.
func (i *Uint32) Swap(n uint32) uint32 {
	return atomic.SwapUint32(&i.value, n)
}

// Add atomically and return the new value.
func (i *Uint32) Add(n uint32) uint32 {
	return atomic.AddUint32(&i.value, n)
}

// Sub atomically and return the new value.
func (i *Uint32) Sub(n uint32) uint32 {
	return atomic.AddUint32(&i.value, ^(n - 1))
}

// Inc atomically and return the new value.
func (i *Uint32) Inc() uint32 {
	return i.Add(1)
}

// Dec atomically and return the new value.
func (i *Uint32) Dec() uint32 {
	return i.Sub(1)
}

// CAS is an atomic Compare-And-Swap operation.
func (i *Uint32) CAS(old, new uint32) bool {
	return atomic.CompareAndSwapUint32(&i.value, old, new)
}

// SwapGreater value atomically, returns old and swap result.
func (i *Uint32) SwapGreater(new uint32) (old uint32, swapped bool) {
	for {
		old := i.Load()
		if new <= old {
			return old, false
		}
		if i.CAS(old, new) {
			return old, true
		}
	}
}

// SwapLess value atomically, returns old and swap result.
func (i *Uint32) SwapLess(new uint32) (old uint32, swapped bool) {
	for {
		old := i.Load()
		if new >= old {
			return old, false
		}
		if i.CAS(old, new) {
			return old, true
		}
	}
}
