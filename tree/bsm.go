package tree

import (
	"fmt"
	"github.com/alesgaroth/pfds-go/interfaces"
)

type MapEntry[K Ordered[K], V comparable] struct {
	key K
	val V
}

func (me MapEntry[K, V]) Key() K {
	return me.key
}
func (me MapEntry[K, V]) Value() V {
	return me.val
}

type TreeMap[K Ordered[K], V comparable] struct {
	left, right *TreeMap[K, V]
	data        MapEntry[K, V]
}

func (l *TreeMap[K, V]) EmptyMap() interfaces.Map[K, V] {
	return nil
}

func EmptyTreeMap[K Ordered[K], V comparable]() *TreeMap[K, V] {
	return nil
}

var NotFound = fmt.Errorf("That value is not in the set")

func (t *TreeMap[K, V]) Lookup(elem K) (V, error) {
	if t == nil {
		var v V
		return v, NotFound
	}
	if elem.Lt(t.data.key) {
		return t.left.Lookup(elem)
	} else {
		return t.right.lookup(t, elem)
	}
}

func (t *TreeMap[K, V]) lookup(last *TreeMap[K, V], elem K) (V, error) {
	if t == nil {
		if last.data.key.Eq(elem) {
			return last.data.val, nil
		} else {
			var v V
			return v, NotFound
		}
	}
	if elem.Lt(t.data.key) {
		return t.left.lookup(last, elem)
	} else {
		return t.right.lookup(t, elem)
	}
}

func (t *TreeMap[K, V]) IsEmpty() bool {
	return t == nil
}

func (t *TreeMap[K, V]) Bind(key K, val V) (retval interfaces.Map[K, V]) {
	defer func() {
		if r := recover(); r != nil {
			if r == alreadyThere {
				retval = t
			} else {
				panic(r)
			}
		}
	}()
	retval = t.bind(key, val)
	return retval
}

func (t *TreeMap[K, V]) bind(key K, val V) *TreeMap[K, V] {
	if t == nil {
		return &TreeMap[K, V]{nil, nil, MapEntry[K, V]{key, val}}
	}
	if key.Lt(t.data.key) {
		return &TreeMap[K, V]{t.left.bind(key, val), t.right, t.data}
	} else {
		return &TreeMap[K, V]{t.left, t.right.bnd(t, key, val), t.data}
	}
}

func (t *TreeMap[K, V]) bnd(last *TreeMap[K, V], key K, val V) *TreeMap[K, V] {
	// bug
	if t == nil {
		if last.data.key.Eq(key) && last.data.val == val {
			panic(alreadyThere)
		}
		return &TreeMap[K, V]{nil, nil, MapEntry[K, V]{key, val}}
	}
	if key.Lt(t.data.key) {
		return &TreeMap[K, V]{t.left.bnd(last, key, val), t.right, t.data}
	} else {
		return &TreeMap[K, V]{t.left, t.right.bnd(t, key, val), t.data}
	}
}

// first let's be able to sequence the map
func (t *TreeMap[K,V]) Left() AbsTree[interfaces.MapEntry[K,V]] {
	return t.left
}
func (t *TreeMap[K,V]) Right() AbsTree[interfaces.MapEntry[K,V]] {
	return t.right
}
func (t *TreeMap[K,V]) Data() interfaces.MapEntry[K,V] {
	return t.data
}

type KeySeq[T Ordered[T], V any] struct {
	stack interfaces.Stack[interfaces.MapEntry[T, V]]
}

func (t *TreeMap[K, V])	Sequence() interfaces.Stack[interfaces.MapEntry[K, V]] {
	return Sequence(t)
}
