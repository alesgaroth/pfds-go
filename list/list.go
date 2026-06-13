package list

import (
	"github.com/alesgaroth/pfds-go/interfaces"
)

type List[T any] struct {
	data T
	next *List[T]
}


func EmptyList[T any]()  *List[T] {
	return nil
}

func (l *List[T])GetEmpty() interfaces.Stack[T] {
	return nil
}

func (l*List[T]) IsEmpty() bool {
	return l == nil
}
func (l*List[T]) Cons(t T) interfaces.Stack[T] {
	return &List[T]{t, l}
}
func (l*List[T]) Head() T {
	if l == nil {
		var zero T
		return zero
	} else {
		return l.data
	}
}
func (l*List[T]) Tail() interfaces.Stack[T] {
	if l == nil {
		return nil
	} else {
		return l.next
	}
}
