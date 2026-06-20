package tree

import (
	"testing"
)

func TestTT(t *testing.T) {

	rb := EmptyRbTree[MapEntry[OrderedInt, int]]()
	rbm := (*RbTreeMap[OrderedInt, int])(rb)

	tr := rbm.Bind(2, 1)
	tr = tr.Bind(1, 3)
	tr = tr.Bind(3, 2)

	if val, err := tr.Lookup(2); err == nil && val != 1 {
		t.Fatalf("should have found 2 => 1")
	}
	if val, err := tr.Lookup(1); err == nil && val != 3 {
		t.Fatalf("should have found 1 => 3")
	}
	if val, err := tr.Lookup(3); err == nil && val != 2 {
		t.Fatalf("should have found 3 => 2")
	}
	if val, err := tr.Lookup(4); err == nil {
		t.Fatalf("should not have found 4 got %v", val)
	}
}
