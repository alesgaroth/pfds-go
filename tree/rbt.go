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

func EmptyRbTree[T Ordered[T]]() *RbTree[T] {
	return nil
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

func (t *RbTree[T]) Insert(elem T) interfaces.Set[T] {
	if t == nil {
		return &RbTree[T]{Black, nil, nil, elem}
	}
	t2 := t.ins(elem)
	return &RbTree[T]{Black, t2.left, t2.right, t2.data}
}

type balancer[T Ordered[T]] func(colour Colour, left *RbTree[T], right *RbTree[T], data T) *RbTree[T]

func (t *RbTree[T]) ins(elem T) (retval *RbTree[T]) {
	if t == nil {
		return &RbTree[T]{Red, nil, nil, elem}
	}
	if elem.Lt(t.data) {
		newleft, balance := t.left.lIns(elem)
		return balance(t.colour, newleft, t.right, t.data)
	} else if t.data.Lt(elem) {
		newright, balance := t.right.rIns(elem)
		return balance(t.colour, t.left, newright, t.data)
	} else {
		return t // aleady here
	}
}

func (t *RbTree[T]) lIns(elem T) (*RbTree[T], balancer[T]) {
	if t == nil {
		return &RbTree[T]{Red, nil, nil, elem}, nobalance
	}
	if elem.Lt(t.data) {
		newleft, balance := t.left.lIns(elem)
		return balance(t.colour, newleft, t.right, t.data), llbalance
	} else if t.data.Lt(elem) {
		newright, balance := t.right.rIns(elem)
		return balance(t.colour, t.left, newright, t.data), lrbalance
	} else {
		return t, nobalance // aleady here
	}
}

func (t *RbTree[T]) rIns(elem T) (*RbTree[T], balancer[T]) {
	if t == nil {
		return &RbTree[T]{Red, nil, nil, elem}, nobalance
	}
	if elem.Lt(t.data) {
		newleft, balance := t.left.lIns(elem)
		return balance(t.colour, newleft, t.right, t.data), rlbalance
	} else if t.data.Lt(elem) {
		newright, balance := t.right.rIns(elem)
		return balance(t.colour, t.left, newright, t.data), rrbalance
	} else {
		return t, nobalance // aleady here
	}
}

func nobalance[T Ordered[T]](colour Colour, left *RbTree[T], right *RbTree[T], data T) *RbTree[T] {
	return &RbTree[T]{colour, left, right, data}
}
func llbalance[T Ordered[T]](colour Colour, left *RbTree[T], right *RbTree[T], data T) *RbTree[T] {
	if colour == Black && left.colour == Red && left.left.colour == Red {
		var a, x, b, y, c, z, d = left.left.left, left.left.data, left.left.right, left.data, left.right, data, right
		return &RbTree[T]{Red, &RbTree[T]{Black, a, b, x}, &RbTree[T]{Black, c, d, z}, y}
	}
	return nobalance(colour, left, right, data)
}
func lrbalance[T Ordered[T]](colour Colour, left *RbTree[T], right *RbTree[T], data T) *RbTree[T] {
	if colour == Black && left.colour == Red && left.right.colour == Red {
		var a, x, b, y, c, z, d = left.left, left.data, left.right.left, left.right.data, left.right.right, data, right
		return &RbTree[T]{Red, &RbTree[T]{Black, a, b, x}, &RbTree[T]{Black, c, d, z}, y}
	}
	return nobalance(colour, left, right, data)
}
func rlbalance[T Ordered[T]](colour Colour, left *RbTree[T], right *RbTree[T], data T) *RbTree[T] {
	if colour == Black && right.colour == Red && right.left.colour == Red {
		var a, x, b, y, c, z, d = left, data, right.left.left, right.left.data, right.left.right, right.data, right.right
		return &RbTree[T]{Red, &RbTree[T]{Black, a, b, x}, &RbTree[T]{Black, c, d, z}, y}
	}
	return nobalance(colour, left, right, data)
}
func rrbalance[T Ordered[T]](colour Colour, left *RbTree[T], right *RbTree[T], data T) *RbTree[T] {
	if colour == Black && right.right != nil && right.colour == Red && right.right.colour == Red {
		var a, x, b, y, c, z, d = left, data, right.left, right.data, right.right.left, right.right.data, right.right.right
		return &RbTree[T]{Red, &RbTree[T]{Black, a, b, x}, &RbTree[T]{Black, c, d, z}, y}
	}
	return nobalance(colour, left, right, data)
}
