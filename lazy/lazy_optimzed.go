//go:build !debug
package lazy

type Z[T any] interface {
	Force() T
}

type optimized[T any] struct {
	thunk func() T
}

func MZ[T any](thunk func() T) Z[T] {
	thing := &optimized[T]{}
	thing.thunk = func() T {
		val := thunk()
		thing.thunk = func() T {
			return val
		}
		return thing.thunk()
	}
	return thing
}

func (d *optimized[T]) Force() T {
	return d.thunk()
}
