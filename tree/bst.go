package tree

import (
	"fmt"
	"github.com/alesgaroth/pfds-go/interfaces"
)

type Ordered[T Ordered[T]] interface {
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
	} else {
		return t.right.member(t, elem)
	}
}

func (t *Tree[T]) member(last *Tree[T], elem T) bool {
	if t == nil {
		return last.data.Eq(elem)
	}
	if elem.Lt(t.data) {
		return t.left.member(last, elem)
	} else {
		return t.right.member(t, elem)
	}
}

var alreadyThere = fmt.Errorf("That value is already in the set")

func (t *Tree[T]) Insert(elem T) (retval *Tree[T]) {
	defer func() {
		if r := recover(); r != nil {
			if r == alreadyThere {
				retval = t
			} else {
				panic(r)
			}
		}
	}()
	retval =  t.insert(elem)
	return retval
}


func (t *Tree[T]) insert(elem T) *Tree[T] {
	if t == nil {
		return &Tree[T]{nil, nil, elem}
	}
	if elem.Lt(t.data) {
		return &Tree[T]{t.left.insert(elem), t.right, t.data}
	} else if t.data.Lt(elem) {
		return &Tree[T]{t.left, t.right.insert(elem), t.data}
	} else {
		panic(alreadyThere)
	}
}
