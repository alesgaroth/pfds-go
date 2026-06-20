package tree

import (
	"github.com/alesgaroth/pfds-go/interfaces"
)

func (mapentry MapEntry[K, V]) Eq(other MapEntry[K, V]) bool {
	return mapentry.Key.Eq(other.Key)
}
func (mapentry MapEntry[K, V]) Leq(other MapEntry[K, V]) bool {
	return mapentry.Key.Leq(other.Key)
}
func (mapentry MapEntry[K, V]) Lt(other MapEntry[K, V]) bool {
	return mapentry.Key.Lt(other.Key)
}

type RbTreeMap[K Ordered[K], V comparable] RbTree[MapEntry[K, V]]

func (t *RbTreeMap[K, V]) IsEmpty() bool {
	return t == nil
}

func (t *RbTreeMap[K, V]) EmptyMap() interfaces.Map[K, V] {
	return nil
}
func (t *RbTreeMap[K, V]) Bind(key K, val V) interfaces.Map[K, V] {
	rb := (*RbTree[MapEntry[K, V]])(t)
	set := rb.Insert(MapEntry[K, V]{key, val})
	rb = set.(*RbTree[MapEntry[K, V]])
	return (*RbTreeMap[K, V])(rb)
}

// this is identical to what's in Bsm.go. (Aside from the casts)
func (t *RbTreeMap[K, V]) Lookup(elem K) (V, error) {
	var v V
	actual, err := (*RbTree[MapEntry[K, V]])(t).Lookup(MapEntry[K, V]{elem, v})
	if err != nil {
		return v, err
	} else {
		return actual.Val, nil
	}
}
