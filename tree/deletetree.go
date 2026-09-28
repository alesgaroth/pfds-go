package tree

import (
	"github.com/alesgaroth/pfds-go/interfaces"
	"github.com/alesgaroth/pfds-go/list"
)

type DeleteTree[T Ordered[T]] struct {
	wrapped    interfaces.Set[T]
	tombstones interfaces.Set[T]
}

func (d *DeleteTree[T]) GetEmpty() interfaces.Set[T] {
	return EmptyRbTree[T]()
}

func (d *DeleteTree[T]) IsEmpty() bool {
	return nil == d
}

func (d *DeleteTree[T]) Insert(data T) interfaces.Set[T] {
	if d == nil {
		return d.GetEmpty().Insert(data)
	}
	return &DeleteTree[T]{
		d.wrapped.Insert(data),
		d.tombstones,
	}
}
func (d *DeleteTree[T]) Member(data T) bool {
	if d.tombstones.Member(data) {
		return false
	}
	return d.wrapped.Member(data)
}

func (d *DeleteTree[T]) Sequence() interfaces.Stack[T] {
	return seqDiff(d.wrapped.Sequence(), d.tombstones.Sequence())
}

func (d *DeleteTree[T]) Delete(v T) interfaces.Set[T] {
	return &DeleteTree[T]{
		d.wrapped,
		d.tombstones.Insert(v),
	}
}
func (d *DeleteTree[T]) Merge(interfaces.Set[T]) interfaces.Set[T] {
	panic("unimplemented")
}

type sDiff[T Ordered[T]] struct {
	positive, negative interfaces.Stack[T]
}

func skipToFirstPositive[T Ordered[T]](positive, negative interfaces.Stack[T]) (p, n interfaces.Stack[T]) {
	for {
		j := positive.Head()
		k := negative.Head()
		if j.Eq(k) {
			positive = positive.Tail()
		} else if j.Lt(k) { // negative is ahead of positive
			return positive, negative	
		} else { // positive is ahead of negative
			negative = negative.Tail()
		}
	}
}

func seqDiff[T Ordered[T]](positive, negative interfaces.Stack[T]) interfaces.Stack[T] {
	positive, negative = skipToFirstPositive(positive, negative)
	if negative.IsEmpty() {
		return positive
	} else if positive.IsEmpty()  {
		return positive.GetEmpty()
	} else {
		return &sDiff[T]{positive, negative}
	}
}
func (s*sDiff[T])GetEmpty() interfaces.Stack[T] {
	return list.EmptyList[T]()
}
func (s*sDiff[T])IsEmpty() bool {
	return s.positive.IsEmpty()
}
func (s*sDiff[T])Cons(t T) interfaces.Stack[T] {
	return list.Prepend[T](t, s)
}
func (s*sDiff[T])Head() T {
	return s.positive.Head()
}
func (s*sDiff[T])Tail() interfaces.Stack[T] {
	return seqDiff(s.positive.Tail(), s.negative)
}

/*
func eagerSeqDiff[T Ordered[T]]seqDiff(positive, negative interfaces.Stack[T]) interfaces.Stack[T] {
	k := negative.Head()
	for j := positive.Head(); !positive.IsEmpty(); positive = positive.Tail() {
		if j.Eq(k) {
			continue
		} else if j.Lt(k) {
			yield j
		} else {
			negative = negative.Tail()
		}
	}
}
*/
