package lazy

import (
	"testing"
)

func TestStream(t *testing.T) {
	var s Stream[int] = MakeStream[int]()
	s = s.Cons(5)
	s = s.Cons(6)
	if s.Head() != 6 {
		t.Fatalf("hello?")
	}
	u := s.Tail()
	if u.Head() != 5 {
		t.Fatalf("so yeah")
	}

	s = s.Append(s)
	if s.IsEmpty() {
		t.Fatalf("why is it empty")
	}

	if s.Head() != 6 {
		t.Fatalf("sss")
	}
	u = s.Tail()
	if u.Head() != 5 {
		t.Fatalf("ssss")
	}
	u = u.Tail()
	if u.Head() != 6 {
		t.Fatalf("sssss")
	}
	u = u.Tail()
	if u.Head() != 5 {
		t.Fatalf("sssssss")
	}
	u = u.Tail()
	if !u.IsEmpty() {
		t.Fatalf("why isn't it empty")
	}

	u = s.Take(3)
	if u.IsEmpty() {
		t.Fatalf("why would take(3) return an empty stream?")
	}
	if u.Head() != 6 {
		t.Fatalf("What?")
	}
	if u.Tail().Head() != 5 {
		t.Fatalf("What?")
	}
	if u.Tail().Tail().Head() != 6 {
		t.Fatalf("What?")
	}
	if !u.Tail().Tail().Tail().IsEmpty() {
		t.Fatalf("I took three from a 4 count stream, then took off the last one, should be empty")
	}

	u = s.Drop(3)
	if u.IsEmpty() {
		t.Fatalf("why would drop(3) return an empty stream?")
	}
	if u.Head() != 5 {
		t.Fatalf("Whaaat? %v", u.Head())
	}
	if !u.Tail().IsEmpty() {
		t.Fatalf("I dropped three from a 4 count stream, then took off the last one, should be empty")
	}

	u = s.Reverse()
	if u.IsEmpty() {
		t.Fatalf("why would reverse return an empty stream?")
	}
	if u.Head() != 5 {
		t.Fatalf("Whoa %v", u.Head())
	}
	if u.Tail().Head() != 6 {
		t.Fatalf("What?")
	}
	if u.Tail().Tail().Head() != 5 {
		t.Fatalf("What?")
	}
	if u.Tail().Tail().Tail().Head() != 6 {
		t.Fatalf("What?")
	}
	if !u.Tail().Tail().Tail().Tail().IsEmpty() {
		t.Fatalf("SSS")
	}

}
