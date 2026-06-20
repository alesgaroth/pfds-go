package tree

import (
	//"github.com/alesgaroth/pfds-go/interfaces"
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

