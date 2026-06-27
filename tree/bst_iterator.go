package tree

import (
	"github.com/alesgaroth/pfds-go/interfaces"
	"github.com/alesgaroth/pfds-go/list"
)

type BsSeq[T Ordered[T]] struct {
	stack interfaces.Stack[*Tree[T]]
}

func (t *Tree[T]) Sequence() interfaces.Stack[T] {
	if t == nil {
		return nil //list.EmptyList[T]()
	}
	var stack interfaces.Stack[*Tree[T]] = list.EmptyList[*Tree[T]]()
	for ; t != nil; t = t.left {
		stack = stack.Cons(t)
	}
	return &BsSeq[T]{stack}
}

func (seq *BsSeq[T]) IsEmpty() bool {
	return seq.stack.IsEmpty()
}
func (seq *BsSeq[T]) Head() T {
	if seq.IsEmpty() {
		var t T
		return t
	}
	t := seq.stack.Head()
	return t.data // panic if we somehow got a nil
}
func (seq *BsSeq[T]) Tail() interfaces.Stack[T] {
	if seq.IsEmpty() {
		return nil //list.EmptyList[T]()
	}
	t := seq.stack.Head()
	stack := seq.stack.Tail()
	for t = t.right; t != nil; t = t.left {
		stack = stack.Cons(t)
	}
	return &BsSeq[T]{stack}
}

func (t *Tree[T]) Count() int {
	if t == nil {
		return 0
	}
	return t.left.Count() + 1 + t.right.Count()
}

func (t *Tree[T]) CountExplicitStack() int {
	if t == nil {
		return 0
	}
	var stack interfaces.Stack[*Tree[T]] = list.EmptyList[*Tree[T]]()
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

func (t *BsSeq[T]) GetEmpty() interfaces.Stack[T] {
	return nil
}
func (t *BsSeq[T]) Cons(value T) interfaces.Stack[T] {
	return list.Prepend[T](value, t)
}
