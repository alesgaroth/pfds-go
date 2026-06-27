package interfaces

type Elem any

type Set[T any] interface {
	GetEmpty() Set[T]
	IsEmpty() bool
	Insert(T) Set[T]
	Member(T) bool
	Sequence() Stack[T]
	Delete(v T) Set[T]
	Merge(Set[T]) Set[T]
}
