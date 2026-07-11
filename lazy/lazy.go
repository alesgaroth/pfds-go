package lazy

type Lazy[T any] interface {
	Force() T
}

type debuggable[T any] struct {
	thunk func() T
	val T
	forced bool
}

type optimized[T any] struct {
	thunk func() T
}

func Lzy[T any](thunk func() T) Lazy[T] {
	return &debuggable[T]{thunk: thunk}
}

func (d*debuggable[T])Force() T {
	if !d.forced  {
		d.forced = true
		d.val = d.thunk()
	}
	return d.val
}
