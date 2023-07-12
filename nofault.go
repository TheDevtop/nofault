package nofault

/*
 This Source Code Form is subject to the terms of the Mozilla Public
 License, v. 2.0. If a copy of the MPL was not distributed with this
 file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

// NoFault: A generic monad-like construct,
// for handling errors.
type NoFault[T any] struct {
	retfn func() T // Returns known-good value
}

// Create a new generic fault handler
func New[T any](fn func() T) *NoFault[T] {
	nf := new(NoFault[T])
	nf.retfn = fn
	return nf
}

// Handle the input => return input or known-good
func (nf *NoFault[T]) Handle(v any) T {
	if r, ok := v.(T); ok {
		return r
	} else {
		return nf.retfn()
	}
}
