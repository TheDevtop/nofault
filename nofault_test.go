package nofault

import (
	"testing"
)

func TestNew(t *testing.T) {
	// Setup
	f := func() uint {
		return 0
	}

	// Execution
	nf := New[uint](f)

	// Validation
	if nf.retfn() != f() {
		t.Fail()
	}
}

func TestHandle(t *testing.T) {
	// Setup
	f := func() uint {
		return 10
	}
	s := "Faulty data"
	var d uint = 10

	// Execution
	nf := New[uint](f)
	v := nf.Handle(s)

	// Validation
	if v != d {
		t.Fail()
	}
}
