package tree

import (
	"fmt"
	"slices"
	"github.com/alesgaroth/pfds-go/interfaces"
	"testing"
)

func TestTen(t *testing.T) {
	rb := EmptyRbTree[OrderedInt]()
	tr := rb.Insert(0)
	for x := OrderedInt(1); x < 10; x += 1 {
		tr = tr.Insert(x)
	}

	for x := OrderedInt(0); x < 10; x += 1 {
		if !tr.Member(x) {
			t.Fatalf("oops missing %v", x)
		}
	}
	for x := OrderedInt(10); x < 12; x += 1 {
		if tr.Member(x) {
			t.Fatalf("oops has %v", x)
		}
	}
}

func (t *RbTree[T]) String() string {
	if t == nil {
		return "."
	}
	var colour string = "?"
	if t.colour == Black {
		colour = "B"
	} else if t.colour == Red {
		colour = "R"
	}
	return fmt.Sprint("[", colour, t.left, t.data, t.right, "]")
}

func TestTenBack(t *testing.T) {
	rb := EmptyRbTree[OrderedInt]()
	tr := rb.Insert(0)
	for x := OrderedInt(9); x >= 0; x -= 1 {
		tr = tr.Insert(x)
	}

	for x := OrderedInt(0); x < 10; x += 1 {
		if !tr.Member(x) {
			t.Fatalf("oops missing %v", x)
		}
	}
	for x := OrderedInt(10); x < 12; x += 1 {
		if tr.Member(x) {
			t.Fatalf("oops has %v", x)
		}
	}
}

func TestTenUpFromFive(t *testing.T) {
	rb := EmptyRbTree[OrderedInt]()
	tr := rb.Insert(0)
	for x := OrderedInt(6); x < 10; x += 1 {
		tr = tr.Insert(x)
	}
	for x := OrderedInt(5); x >= 0; x -= 1 {
		tr = tr.Insert(x)
	}

	for x := OrderedInt(0); x < 10; x += 1 {
		if !tr.Member(x) {
			t.Fatalf("oops missing %v", x)
		}
	}
	for x := OrderedInt(10); x < 12; x += 1 {
		if tr.Member(x) {
			t.Fatalf("oops has %v", x)
		}
	}
}

func TestTenDownFromFive(t *testing.T) {
	rb := EmptyRbTree[OrderedInt]()
	tr := rb.Insert(0)
	for x := OrderedInt(5); x >= 0; x -= 1 {
		tr = tr.Insert(x)
	}

	for x := OrderedInt(6); x < 10; x += 1 {
		tr = tr.Insert(x)
	}

	for x := OrderedInt(0); x < 10; x += 1 {
		if !tr.Member(x) {
			t.Fatalf("oops missing %v", x)
		}
	}
	for x := OrderedInt(10); x < 12; x += 1 {
		if tr.Member(x) {
			t.Fatalf("oops has %v", x)
		}
	}
}

func ftest(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11 int) {
	slice := []OrderedInt{OrderedInt(x1), OrderedInt(x2), OrderedInt(x3), OrderedInt(x4), OrderedInt(x5), OrderedInt(x6), OrderedInt(x7), OrderedInt(x8), OrderedInt(x9)}
	notslice := []OrderedInt{OrderedInt(x10), OrderedInt(x11)}
	var tr interfaces.Set[OrderedInt] = EmptyRbTree[OrderedInt]()
	if slices.Contains(slice, notslice[0]) ||slices.Contains(slice, notslice[1]) {
		return // bad test case
	}

	for _, k := range slice {
		tr = tr.Insert(k)
	}

	for _, x := range slice {
		if !tr.Member(x) {
			t.Fatalf("oops missing %v", x)
		}
	}

	for _, x := range notslice {
		if tr.Member(x) {
			t.Fatalf("oops has %v", x)
		}
	}
}

func FuzzRbTree(f *testing.F) {
	f.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	f.Fuzz(ftest)
}
