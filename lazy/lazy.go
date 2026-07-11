//go:build debug
package lazy

type Z[T any] interface {
	Force() T
}

type debuggable[T any] struct {
	thunk  func() T
	val    T
	forced bool
}


func MZ[T any](thunk func() T) Z[T] {
	return &debuggable[T]{thunk: thunk}
}

func (d *debuggable[T]) Force() T {
	if !d.forced {
		d.forced = true
		d.val = d.thunk()
	}
	return d.val
}
