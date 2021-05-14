package atomix

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	ae := NewError(nil)

	Equal(t, "<nil>", ae.String(), "Load wrong value")
	Equal(t, false, ae.HasError(), "Bad value")
	Equal(t, nil, ae.Load(), "Load wrong value")

	err := errors.New("ouch")
	ae = NewError(err)
	Equal(t, true, ae.HasError(), "Bad value")
	Equal(t, err, ae.Load(), "Load wrong value")
	Equal(t, "ouch", ae.String(), "Load wrong value")

	err2 := errors.New("very-ouch")
	ae.Store(err2)
	Equal(t, err2, ae.Load(), "Store wrong value")

	ae.Store(nil)
	Equal(t, err2, ae.Load(), "Store wrong value")
}
