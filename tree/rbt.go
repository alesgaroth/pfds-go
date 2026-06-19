package tree

import (
	"github.com/alesgaroth/pfds-go/interfaces"
)

type Colour int

const (
	Red Colour = iota
	Black
)

type RbTree[T Ordered[T]] struct {
	colour      Colour
	left, right *RbTree[T]
	data        T
}

func (t *RbTree[T]) GetEmpty() interfaces.Set[T] {
	return nil
}

func (t *RbTree[T]) IsEmpty() bool {
	return nil == t
}

func (t *RbTree[T]) Member(elem T) bool {
	if t == nil {
		return false
	}
	if elem.Lt(t.data) {
		return t.left.Member(elem)
	} else {
		return t.right.member(t, elem)
	}
}

func (t *RbTree[T]) member(last *RbTree[T], elem T) bool {
	if t == nil {
		return last.data.Eq(elem)
	}
	if elem.Lt(t.data) {
		return t.left.member(last, elem)
	} else {
		return t.right.member(t, elem)
	}
}

func (t *RbTree[T]) Insert(elem T) (retval interfaces.Set[T]) {
	t2 := t.ins(elem)
	return &RbTree[T]{Black, t2.left, t2.right, t2.data}
}

func (t *RbTree[T]) ins(elem T) (retval *RbTree[T]) {
	if t == nil {
		return &RbTree[T]{Red, nil, nil, elem}
	}
	if elem.Lt(t.data) {
		return balance(t.colour, t.left.ins(elem), t.right, t.data)
	} else if t.data.Lt(elem) {
		return balance(t.colour, t.left, t.right.ins(elem), t.data)
	} else {
		return t // aleady here
	}
}

func balance[T Ordered[T]](colour Colour, left *RbTree[T], right *RbTree[T], data T) *RbTree[T] {
	if colour == Black {
		var a, b, c, d *RbTree[T]
		var x, y, z T
		doIt := false
		if left != nil && left.left != nil && left.colour == Red && left.left.colour == Red {
			a, x, b, y, c, z, d = left.left.left, left.left.data, left.left.right, left.data, left.right, data, right
			doIt = true
		} else if left != nil && left.right != nil && left.colour == Red && left.right.colour == Red {
			a, x, b, y, c, z, d = left.left, left.data, left.right.left, left.right.data, left.right.right, data, right
			doIt = true
		} else if right != nil && right.left != nil && right.colour == Red && right.left.colour == Red {
			a, x, b, y, c, z, d = left, data, right.left.left, right.left.data, right.left.right, right.data, right.right
			doIt = true
		} else if right != nil && right.right != nil && right.colour == Red && right.right.colour == Red {
			a, x, b, y, c, z, d = left, data, right.left, right.data, right.right.left, right.right.data, right.right.right
			doIt = true
		}
		if doIt {
			return &RbTree[T]{Red, &RbTree[T]{Black, a, b, x}, &RbTree[T]{Black, c, d, z}, y}
		} // fall through
	}
	return &RbTree[T]{colour, left, right, data}
}
