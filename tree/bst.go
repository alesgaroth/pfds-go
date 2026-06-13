package tree

import (
	"github.com/alesgaroth/pfds-go/interfaces"
)

type Ordered [T Ordered[T]] interface {
	Eq(T) bool
	Lt(T) bool
	Leq(T) bool
}

type Tree[T Ordered[T]] struct {
	left, right *Tree[T]
	data        T
}

func (l *Tree[T]) GetEmpty() interfaces.Set[T] {
	return nil
}

func EmptyTree[T Ordered[T]]() *Tree[T] {
	return nil
}

func (t *Tree[T]) Member(elem T) bool {
	if t == nil {
		return false
	}
	if elem.Lt(t.data) {
		return t.left.Member(elem)
	} else if t.data.Lt(elem) {
		return t.right.Member(elem)
	} else {
		return true
	}
}

func (t *Tree[T]) Insert(elem T) *Tree[T] {
	if t == nil {
		return &Tree[T]{nil, nil, elem}
	}
	if elem.Lt(t.data) {
		return &Tree[T]{t.left.Insert(elem), t.right, t.data}
	} else if t.data.Lt(elem) {
		return &Tree[T]{t.left, t.right.Insert(elem), t.data}
	} else {
		return t
	}
}
