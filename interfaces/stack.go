package interfaces

type Stack[T any] interface {
	GetEmpty() Stack[T]
	IsEmpty() bool
	Cons(t T) Stack[T]
	Head() T
	Tail() Stack[T]
}



