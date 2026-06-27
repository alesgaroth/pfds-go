package tree

import (
	"testing"
	"github.com/alesgaroth/pfds-go/interfaces"
)

type OrderedInt int

func (i OrderedInt) Eq(j OrderedInt) bool {
	return i == j
}
func (i OrderedInt) Lt(j OrderedInt) bool {
	return i < j
}
func (i OrderedInt) Leq(j OrderedInt) bool {
	return i <= j
}

func TestOne(t *testing.T) {
	var tr interfaces.Set[OrderedInt] = EmptyTree[OrderedInt]()
	tr = tr.Insert(2)
	tr = tr.Insert(1)
	tr = tr.Insert(3)
	if !tr.Member(2) {
		t.Fatalf("should have found 2")
	}
	if !tr.Member(1) {
		t.Fatalf("should have found 1")
	}
	if !tr.Member(3) {
		t.Fatalf("should have found 3")
	}
	if tr.Member(4) {
		t.Fatalf("should not have found 4")
	}
}
