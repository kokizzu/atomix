package atomix

import "sync/atomic"

// String is an atomic wrapper around a string
type String struct {
	noCopy
	value atomic.Value
}

// NewString creates a String
func NewString(str string) *String {
	s := &String{}
	s.Store(str)
	return s
}

// Load atomically the value
func (s *String) Load() string {
	v := s.value.Load()
	if v == nil {
		return ""
	}
	return v.(string)
}

// Store atomically the passed value.
func (s *String) Store(n string) {
	s.value.Store(n)
}
