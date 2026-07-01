package queue

import (
	"github.com/alesgaroth/pfds-go/interfaces"
	"github.com/alesgaroth/pfds-go/list"
)

type proGress[T any] struct {
	input    interfaces.Stack[T]
	reversed interfaces.Stack[T]
	internal Queue[any]
}

type IncrementalQueue[T any] struct {
	enqueue  interfaces.Stack[T]
	dequeue  interfaces.Stack[T]
	num      int
	enSpace  int
	progress *proGress[T]
}

func EmptyIncrementalQueue[T any]() Queue[T] {
	return (*IncrementalQueue[T])(nil)
}

func (rtq *IncrementalQueue[T]) Cons(a T) Queue[T] {
	if rtq.IsEmpty() {
		return &IncrementalQueue[T]{
			list.EmptyList[T](),
			SingletonStack(a),
			1,
			2,
			nil,
		}
	}
	var newEnqueue interfaces.Stack[T]
	if rtq.enqueue.IsEmpty() {
		newEnqueue = SingletonStack(a)
	} else {
		newEnqueue = rtq.enqueue.Cons(a)
	}
	return rebal(newEnqueue, rtq.dequeue, rtq.num+1, rtq.enSpace-1, rtq.progress)
}

func (rtq *IncrementalQueue[T]) Head() T {
	if !rtq.IsEmpty() {
		return rtq.dequeue.Head()
	} else {
		var t T
		return t
	}
}
func (rtq *IncrementalQueue[T]) IsEmpty() bool {
	return rtq == nil || rtq.num == 0
}
func (rtq *IncrementalQueue[T]) Tail() Queue[T] {
	if rtq.IsEmpty() {
		return EmptyIncrementalQueue[T]()
	}
	var newEnqueue = rtq.enqueue
	var newDequeue = rtq.dequeue.Tail()
	var newProgress = rtq.progress
	var newEnspace = rtq.enSpace - 1
	if newDequeue.IsEmpty() {
		if newProgress != nil {
			if !newProgress.internal.IsEmpty() {
				newDequeue = newProgress.internal.Head().(interfaces.Stack[T])
				newProgress = &proGress[T]{newProgress.input, newProgress.reversed, newProgress.internal.Tail()}
			} else if newProgress.input.IsEmpty() {
				newDequeue = newProgress.reversed
				newProgress = nil
			} else {
				newDequeue = newProgress.reversed.Cons(newProgress.input.Head())
				newProgress = nil
			}
		} else if newEnqueue.IsEmpty() {
			return EmptyIncrementalQueue[T]()
		} else {
			len := list.Length(newEnqueue)
			if len == 1 {
				newProgress = nil
				newDequeue = newEnqueue
				newEnqueue = list.EmptyList[T]()
				newEnspace = 2
			} else {
				panic("HELP HELP, I'm stuck in a software factory")
			}
		}
	}
	return rebal(newEnqueue, newDequeue, rtq.num-1, newEnspace, newProgress)
}

func rebal[T any](newEnqueue, newDequeue interfaces.Stack[T], newNum, newEnSpace int, progress *proGress[T]) Queue[T] {
	if newEnSpace == 0 {
		var newInternal Queue[any] =  EmptyIncrementalQueue[any]()
		if progress != nil {
			newInternal = progress.internal
			if !progress.reversed.IsEmpty() {
				newInternal = newInternal.Cons(progress.reversed)
			}
		}
		return &IncrementalQueue[T]{
			list.EmptyList[T](), newDequeue, newNum, newNum + 1,
			&proGress[T]{newEnqueue.Tail(), SingletonStack(newEnqueue.Head()), newInternal},
		}
	} else {
		if progress != nil && !progress.input.IsEmpty() {
			progress = &proGress[T]{
				progress.input.Tail(),
				progress.reversed.Cons(progress.input.Head()),
				progress.internal,
			}
		}
		return &IncrementalQueue[T]{
			newEnqueue,
			newDequeue,
			newNum,
			newEnSpace,
			progress,
		}
	}
}

func SingletonStack[T any](a T) interfaces.Stack[T] {
	var ls interfaces.Stack[T] = list.EmptyList[T]()
	return ls.Cons(a)
}
