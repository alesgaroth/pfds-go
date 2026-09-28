package interfaces

type Key any
type Value any

type Map[K, V any] interface {
	EmptyMap() Map[K, V]
	IsEmpty() bool
	Bind(K, V) Map[K, V]
	Lookup(K) (V, error)
	Sequence() Stack[MapEntry[K, V]]
}

type MapEntry[K, V any] interface {
	Key() K
	Value() V
}
