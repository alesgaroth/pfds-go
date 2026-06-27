package tree

import (
	"github.com/alesgaroth/pfds-go/interfaces"
	"github.com/alesgaroth/pfds-go/list"
)

type RbSeq[T Ordered[T]] struct {
	stack interfaces.Stack[*RbTree[T]]
}

func (t *RbTree[T]) Sequence() interfaces.Stack[T] {
	if t == nil {
		return nil //list.EmptyList[T]()
	}
	var stack interfaces.Stack[*RbTree[T]] = list.EmptyList[*RbTree[T]]()
	for ; t != nil; t = t.left {
		stack = stack.Cons(t)
	}
	return &RbSeq[T]{stack}
}

func (seq *RbSeq[T]) IsEmpty() bool {
	return seq.stack.IsEmpty()
}
func (seq *RbSeq[T]) Head() T {
	if seq.IsEmpty() {
		var t T
		return t
	}
	t := seq.stack.Head()
	return t.data // panic if we somehow got a nil
}
func (seq *RbSeq[T]) Tail() interfaces.Stack[T] {
	if seq.IsEmpty() {
		return nil //list.EmptyList[T]()
	}
	t := seq.stack.Head()
	stack := seq.stack.Tail()
	for t = t.right; t != nil; t = t.left {
		stack = stack.Cons(t)
	}
	return &RbSeq[T]{stack}
}

func (t *RbTree[T]) Count() int {
	if t == nil {
		return 0
	}
	return t.left.Count() + 1 + t.right.Count()
}

func (t *RbTree[T]) CountExplicitStack() int {
	if t == nil {
		return 0
	}
	var stack interfaces.Stack[*RbTree[T]] = list.EmptyList[*RbTree[T]]()
	for ; t != nil; t = t.left {
		stack = stack.Cons(t)
	}

	//  main loop
	count := 0
	for !stack.IsEmpty() {
		t = stack.Head()
		count += 1
		stack = stack.Tail()
		t = t.right
		for ; t != nil; t = t.left {
			stack = stack.Cons(t)
		}
	}

	return count
}

func (t *RbSeq[T]) GetEmpty() interfaces.Stack[T] {
	return nil
}
func (t *RbSeq[T]) Cons(value T) interfaces.Stack[T] {
	return list.Prepend[T](value, t)
}
