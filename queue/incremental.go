package queue

import (
	"github.com/alesgaroth/pfds-go/interfaces"
	"github.com/alesgaroth/pfds-go/list"
)

type reverser[T any] struct {
	input    interfaces.Stack[T]
	reversed interfaces.Stack[T]
	internal Queue[any]
}

type IncrementalQueue[T any] struct {
	enqueue  interfaces.Stack[T]
	dequeue  interfaces.Stack[T]
	num      int
	enSpace  int
	progress *reverser[T]
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
	newEnqueue := rtq.enqueue.Cons(a)
	return rebal(newEnqueue, rtq.dequeue, rtq.num+1, rtq.enSpace-1, rtq.progress)
}

func (rtq *IncrementalQueue[T]) Head() T {
	if rtq != nil {
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
	if rtq == nil {
		return EmptyIncrementalQueue[T]()
	}
	var newDequeue = rtq.dequeue.Tail()
	if newDequeue.IsEmpty() {
		return rtq.emptyDequeue()
	} else {
		var newEnqueue = rtq.enqueue
		var newDequeue = rtq.dequeue.Tail()
		var newProgress = rtq.progress
		var newEnspace = rtq.enSpace - 1
		return rebal(newEnqueue, newDequeue, rtq.num-1, newEnspace, newProgress)
	}
}

func (rtq *IncrementalQueue[T]) emptyDequeue() Queue[T] {
	if rtq.progress != nil {
		if !rtq.progress.internal.IsEmpty() {
			var newEnspace = rtq.enSpace - 1
			var newDequeue = rtq.progress.internal.Head().(interfaces.Stack[T])
			var newProgress = &reverser[T]{rtq.progress.input, rtq.progress.reversed, rtq.progress.internal.Tail()}
			return rebal(rtq.enqueue, newDequeue, rtq.num-1, newEnspace, newProgress)
		} else if rtq.progress.input.IsEmpty() {
			if !rtq.progress.input.IsEmpty() {
				panic("Help me!  I'm stuck in a software factory")
			}
			var newEnspace = rtq.enSpace - 1
			var newDequeue = rtq.progress.reversed
			return rebal(rtq.enqueue, newDequeue, rtq.num-1, newEnspace, nil)
		} else {
			len := list.Length(rtq.progress.input)
			if len > 1 {
				panic("Help I'm stuck in a software factory")
			}
			var newEnspace = rtq.enSpace - 1
			var newDequeue = rtq.progress.reversed.Cons(rtq.progress.input.Head())
			return rebal(rtq.enqueue, newDequeue, rtq.num-1, newEnspace, nil)
		}
	} else if rtq.enqueue.IsEmpty() {
		return EmptyIncrementalQueue[T]()
	} else {
		len := list.Length(rtq.enqueue)
		if len == 1 {
			var newEnqueue = list.EmptyList[T]()
			return reverseStep(newEnqueue, rtq.enqueue, rtq.num-1, 2, nil)
		} else {
			panic("HELP HELP, I'm stuck in a software factory")
		}
	}
}

func rebal[T any](newEnqueue, newDequeue interfaces.Stack[T], newNum, newEnSpace int, progress *reverser[T]) Queue[T] {
	if newEnSpace == 0 {
		return noSpaceLeftInEnqueue(newEnqueue, newDequeue, newNum, newEnSpace, progress)
	} else {
		return reverseStep(newEnqueue, newDequeue, newNum, newEnSpace, progress)
	}
}

func noSpaceLeftInEnqueue[T any](newEnqueue, newDequeue interfaces.Stack[T], newNum, newEnSpace int, progress *reverser[T]) Queue[T] {
	// no space left in enqueue.  move it to the reverser
	var newInternal Queue[any] = EmptyIncrementalQueue[any]()
	if progress != nil {
		newInternal = progress.internal
		if !progress.input.IsEmpty() {
			panic("Help me! I'm stuck in a software factory")
		}
		if !progress.reversed.IsEmpty() {
			newInternal = newInternal.Cons(progress.reversed)
		}
	}
	// at this point progress.input and progress.reversed are empty
	return &IncrementalQueue[T]{
		list.EmptyList[T](), newDequeue, newNum, newNum + 1,
		&reverser[T]{newEnqueue.Tail(), SingletonStack(newEnqueue.Head()), newInternal},
	}
}

func reverseStep[T any](newEnqueue, newDequeue interfaces.Stack[T], newNum, newEnSpace int, progress *reverser[T]) Queue[T] {
	if progress != nil && !progress.input.IsEmpty() {
		progress = &reverser[T]{
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

func SingletonStack[T any](a T) interfaces.Stack[T] {
	var ls interfaces.Stack[T] = list.EmptyList[T]()
	return ls.Cons(a)
}
