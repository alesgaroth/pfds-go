package list_test

import (
	"github.com/alesgaroth/pfds-go/interfaces"
	"github.com/alesgaroth/pfds-go/list"
	"testing"
)

func TestAddToListEtc(t *testing.T) {
	var l interfaces.Stack[int]
	l = list.EmptyList[int]()
	l = l.Cons(3)
	l = l.Cons(4)
	if l.Head() != 4 {
		t.Fatalf("expected 4 got %v", l.Head())
	}
	l = l.Tail()
	if l.Head() != 3 {
		t.Fatalf("expected 3 got %v", l.Head())
	}
}



