package interfaces

type Stack[T any] interface {
	IsEmpty() bool
	Cons(t T) Stack[T]
	Head() T
	Tail() Stack[T]
}



