package tree

import (
	"fmt"
	"github.com/alesgaroth/pfds-go/interfaces"
)

type MapEntry[K any, V comparable] struct {
	Key K
	Val V
}

type TreeMap[K Ordered[K], V comparable] struct {
	left, right *TreeMap[K, V]
	data        MapEntry[K, V]
}

func (l *TreeMap[K, V]) GetEmpty() interfaces.Map[K, V] {
	return nil
}

func EmptyTreeMap[K Ordered[K], V comparable]() *TreeMap[K, V] {
	return nil
}

var NotFound = fmt.Errorf("That value is already in the set")

func (t *TreeMap[K, V]) Lookup(elem K) (V, error) {
	if t == nil {
		var v V
		return v, NotFound
	}
	if elem.Lt(t.data.Key) {
		return t.left.Lookup(elem)
	} else {
		return t.right.lookup(t, elem)
	}
}

func (t *TreeMap[K, V]) lookup(last *TreeMap[K, V], elem K) (V, error) {
	if t == nil {
		if last.data.Key.Eq(elem) {
			return last.data.Val, nil
		} else {
			var v V
			return v, NotFound
		}
	}
	if elem.Lt(t.data.Key) {
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
	if key.Lt(t.data.Key) {
		return &TreeMap[K, V]{t.left.bind(key, val), t.right, t.data}
	} else {
		return &TreeMap[K, V]{t.left, t.right.bnd(t, key, val), t.data}
	}
}

func (t *TreeMap[K, V]) bnd(last *TreeMap[K, V], key K, val V) *TreeMap[K, V] {
	if t == nil {
		if last.data.Key.Eq(key) && last.data.Val == val {
			panic(alreadyThere)
		}
		return &TreeMap[K, V]{nil, nil, MapEntry[K, V]{key, val}}
	}
	if key.Lt(t.data.Key) {
		return &TreeMap[K, V]{t.left.bnd(last, key, val), t.right, t.data}
	} else {
		return &TreeMap[K, V]{t.left, t.right.bnd(t, key, val), t.data}
	}
}
