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
	_, err := t.Lookup(elem)
	return err != NotFound
}

func (t *RbTree[T]) Insert(elem T) interfaces.Set[T] {
	if t == nil {
		return &RbTree[T]{Black, nil, nil, elem}
	}
	t2, _ := t.ins(elem)
	return &RbTree[T]{Black, t2.left, t2.right, t2.data}
}

type balancer[T Ordered[T]] func(colour Colour, left *RbTree[T], right *RbTree[T], data T) *RbTree[T]

func (t *RbTree[T]) ins(elem T) (retval *RbTree[T], retbalancer balancer[T]) {
	if t == nil {
		return &RbTree[T]{Red, nil, nil, elem}, nobalance
	}
	if elem.Lt(t.data) {
		newleft, balance := t.left.ins(elem)
		return balance(t.colour, newleft, t.right, t.data), llbalance
	} else {
		return t.balancedRIns(elem, lrbalance)
	}
}

func (t *RbTree[T]) balancedRIns(elem T, upbalancer balancer[T]) (retval *RbTree[T], retbalancer balancer[T]) {
	defer func() {
		if r := recover(); r != nil {
			if r == alreadyThere {
				retval = &RbTree[T]{t.colour, t.left, t.right, elem}
				retbalancer = nobalance
			} else {
				panic(r)
			}
		}
	}()
	newright, balance := t.right.rIns(t, elem)
	return balance(t.colour, t.left, newright, t.data), upbalancer
}

func (t *RbTree[T]) lIns(last *RbTree[T], elem T) (retval *RbTree[T], retbalancer balancer[T]) {
	if t == nil {
		if last.data.Eq(elem) {
			panic(alreadyThere)
		}
		return &RbTree[T]{Red, nil, nil, elem}, nobalance
	}
	if elem.Lt(t.data) {
		newleft, balance := t.left.lIns(last, elem)
		return balance(t.colour, newleft, t.right, t.data), llbalance
	} else {
		return t.balancedRIns(elem, lrbalance)
	}
}

func (t *RbTree[T]) rIns(last *RbTree[T], elem T) (retval *RbTree[T], retbalancer balancer[T]) {
	if t == nil {
		if last.data.Eq(elem) {
			panic(alreadyThere)
		}
		return &RbTree[T]{Red, nil, nil, elem}, nobalance
	}
	if elem.Lt(t.data) {
		newleft, balance := t.left.lIns(last, elem)
		return balance(t.colour, newleft, t.right, t.data), rlbalance
	} else {
		return t.balancedRIns(elem, rrbalance)
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

// this is identical to what's in Bsm.go. (Aside from the casts)
func (t *RbTree[T]) Lookup(elem T) (T, error) {
	if t == nil {
		var v T
		return v, NotFound
	}
	if elem.Lt(t.data) {
		return t.left.Lookup(elem)
	} else {
		return t.right.lookup(t, elem)
	}
}

func (t *RbTree[T]) lookup(last *RbTree[T], elem T) (T, error) {
	if t == nil {
		if last.data.Eq(elem) {
			return last.data, nil
		} else {
			var v T
			return v, NotFound
		}
	}
	if elem.Lt(t.data) {
		return t.left.lookup(last, elem)
	} else {
		return t.right.lookup(t, elem)
	}
}
