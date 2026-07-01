package queue

type Queue[T any] interface {
	IsEmpty() bool
	Cons(a T) Queue[T]
	Head() T
	Tail() Queue[T]
}
