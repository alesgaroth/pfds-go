package lazy

type Stream[T any] interface {
	IsEmpty() bool
	Cons(t T) Stream[T]
	Head() T
	Tail() Stream[T]
	Append(Stream[T]) Stream[T]
	Take(num int) Stream[T]
	Drop(num int) Stream[T]
	Reverse() Stream[T]
}

type streem[T any] struct {
	cl Z[*cell[T]]
}

type cell[T any] struct {
	val  T
	next Stream[T]
}

func MakeStream[T any]() *streem[T] {
	return nil
}

func (s *streem[T]) IsEmpty() bool {
	return s == nil || s.cl.Force() == nil
}

func (s *streem[T]) Cons(t T) Stream[T] {
	return &streem[T]{MZ[*cell[T]](func() *cell[T] { return &cell[T]{t, s} })}
}
func (s *streem[T]) Head() T {
	if s == nil {
		panic("oops trying to head an empty list")
	}
	return s.cl.Force().val
}
func (s *streem[T]) Tail() Stream[T] {
	return s.cl.Force().next
}
func (s *streem[T]) Append(other Stream[T]) Stream[T] {
	if s == nil {
		return other
	}
	if s.cl == nil {
		return other
	}
	q := s.cl.Force()
	if q == nil {
		return other
	} else {
		return &streem[T]{MZ[*cell[T]](func() *cell[T] {
			return &cell[T]{
				q.val,
				q.next.Append(other),
			}
		})}
	}
}
func (s *streem[T]) Take(num int) Stream[T] {
	if num == 0 {
		return MakeStream[T]()
	} else if s == nil {
		return MakeStream[T]()
	} else {
		q := s.cl.Force()
		return &streem[T]{MZ[*cell[T]](func() *cell[T]{
				return &cell[T] {
					q.val,
					q.next.Take(num - 1),
				}
			})}
	}
}


func (s *streem[T]) Drop(num int ) Stream[T] {
	if num == 0 {
		return s
	} else if  s == nil {
		return MakeStream[T]()
	} else {
		q := s.cl.Force()
		return q.next.Drop(num - 1)
	}
}

func (s *streem[T]) Dropdoesntwork(num int) Stream[T] {
	if num == 0  || s == nil {
		return MakeStream[T]()
	}
	q := s.cl.Force()
	qq := q.next.(*streem[T])
	return &streem[T]{MZ[*cell[T]](func() *cell[T] {
		return &cell[T] {
			q.val,
			qq.Drop(num - 1),
		}
	})}
}

func (s *streem[T]) rever(s2 Stream[T]) Stream[T] {
	if s == nil || s.cl.Force() == nil {
		return s2
	}
	q := s.cl.Force()
	qq := q.next.(*streem[T])
	return qq.rever(s2.Cons(q.val))
}
func (s *streem[T]) Reverse() Stream[T] {
	return s.rever(MakeStream[T]())
}
