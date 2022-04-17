package atomix

import (
	"testing"
)

func TestInt32(t *testing.T) {
	a := NewInt32(10)

	Equal(t, "10", a.String(), "Wrong String")
	Equal(t, int32(10), a.Load(), "Load wrong value")

	Equal(t, int32(15), a.Add(5), "Add wrong value")
	Equal(t, int32(12), a.Sub(3), "Sub wrong value")

	Equal(t, int32(13), a.Inc(), "Inc wrong value")
	Equal(t, int32(12), a.Dec(), "Dec wrong value")

	OK(t, a.CAS(12, 0), "CAS should swap")
	Equal(t, int32(0), a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(13, 0), "CAS should not swap")

	Equal(t, int32(0), a.Swap(1), "Swap wrong value")
	Equal(t, int32(1), a.Load(), "Swap wrong value")

	a.Store(15)
	Equal(t, int32(15), a.Load(), "Store wrong value")
}

func TestInt32Compare(t *testing.T) {
	a := NewInt32(42)

	old, ok := a.SwapGreater(80)
	Equal(t, old, int32(42), "Store wrong value")
	Equal(t, true, ok, "Store wrong value")
	Equal(t, a.Load(), int32(80), "Store wrong value")

	old, ok = a.SwapGreater(40)
	Equal(t, old, int32(80), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")

	old, ok = a.SwapLess(-80)
	Equal(t, old, int32(80), "Store wrong value")
	Equal(t, ok, true, "Store wrong value")
	Equal(t, a.Load(), int32(-80), "Store wrong value")

	old, ok = a.SwapLess(-40)
	Equal(t, old, int32(-80), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")
}
