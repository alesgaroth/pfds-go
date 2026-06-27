package tree

import (
	"github.com/alesgaroth/pfds-go/interfaces"
	"github.com/alesgaroth/pfds-go/list"
	"testing"
)

type arrayStack []OrderedInt

func buildStack(xs ...OrderedInt) interfaces.Stack[OrderedInt] {
	return arrayStack(xs)
}

func (as arrayStack) GetEmpty() interfaces.Stack[OrderedInt] {
	return arrayStack{}
}
func (as arrayStack) IsEmpty() bool {
	return len(as) == 0
}
func (as arrayStack) Cons(t OrderedInt) interfaces.Stack[OrderedInt] {
	return list.Prepend[OrderedInt](t, as)
}
func (as arrayStack) Head() OrderedInt {
	if as.IsEmpty() {
		var t OrderedInt
		return t
	}
	return as[0]
}
func (as arrayStack) Tail() interfaces.Stack[OrderedInt] {
	if as.IsEmpty() {
		return as
	}
	return as[1:]
}

func TestEmptyOrderedToRbTree(t *testing.T) {
	s := buildStack()
	rbt := OrderedToRbTree(s)
	if rbt != nil {
		t.Fatalf("Empty stack should give empty tree")
	}
	allBlack(t, rbt)
}

func TestSingleOrderedToRbTree(t *testing.T) {
	s := buildStack(1)
	rbt := OrderedToRbTree(s)
	if rbt == nil {
		t.Fatalf("Singleton stack should not give empty tree")
	}
	if !rbt.Member(1) {
		t.Fatalf("Should have found 1 in tree")
	}
	allBlack(t, rbt)
}

func TestTripleOrderedToRbTree(t *testing.T) {
	s := buildStack(1, 2, 3)
	rbt := OrderedToRbTree(s)
	if !rbt.Member(1) {
		t.Fatalf("Should have found 1 in tree")
	}
	if !rbt.Member(2) {
		t.Fatalf("Should have found 2 in tree")
	}
	if !rbt.Member(3) {
		t.Fatalf("Should have found 3 in tree")
	}
	checkInvariants(t, rbt)
	allBlack(t, rbt)
}

func TestSeptupleOrderedToRbTree(t *testing.T) {
	s := buildStack(1, 2, 3, 4, 5, 6, 7)
	rbt := OrderedToRbTree(s)
	if !rbt.Member(1) {
		t.Fatalf("Should have found 1 in tree")
	}
	if !rbt.Member(2) {
		t.Fatalf("Should have found 2 in tree")
	}
	if !rbt.Member(3) {
		t.Fatalf("Should have found 3 in tree")
	}
	checkInvariants(t, rbt)
	allBlack(t, rbt)
}

func TestDoubleOrderedToRbTree(t *testing.T) {
	s := buildStack(1, 2)
	rbt := OrderedToRbTree(s)
	if !rbt.Member(1) {
		t.Fatalf("Should have found 1 in tree")
	}
	if !rbt.Member(2) {
		t.Fatalf("Should have found 2 in tree")
	}
	if rbt.Member(3) {
		t.Fatalf("Should not have found 3 in tree")
	}
	checkInvariants(t, rbt)
}

func TestQuadrupleOrderedToRbTree(t *testing.T) {
	s := buildStack(1, 2, 3, 4)
	rbt := OrderedToRbTree(s)
	if !rbt.Member(1) {
		t.Fatalf("Should have found 1 in tree")
	}
	if !rbt.Member(2) {
		t.Fatalf("Should have found 2 in tree")
	}
	if !rbt.Member(4) {
		t.Fatalf("Should have found 4 in tree")
	}
	checkInvariants(t, rbt)
}

func TestLength(t *testing.T) {
	if 0 != Length(buildStack()) {
		t.Fatalf("empty buildStack should return length 0")
	}
	if 1 != Length(buildStack(1)) {
		t.Fatalf("singleton buildStack should return length 1")
	}
	if 2 != Length(buildStack(1, 2)) {
		t.Fatalf("dual buildStack should return length 2")
	}
}

func allBlack(t *testing.T, tr *RbTree[OrderedInt]) {
	if tr == nil {
		return
	}
	allBlack(t, tr.left)
	allBlack(t, tr.right)
	if tr.colour != Black {
		t.Fatalf("%v is not black", tr)
	}
}

type virtualStack struct {
	current, top uint
}

func fftest(t *testing.T, top uint) {
	vs := virtualStack{0, top}
	rbt := OrderedToRbTree(vs)
	checkInvariants(t, rbt)
}

func (as virtualStack) GetEmpty() interfaces.Stack[OrderedInt] {
	return virtualStack{0, 0}
}
func (as virtualStack) IsEmpty() bool {
	return as.current >= as.top
}
func (as virtualStack) Cons(t OrderedInt) interfaces.Stack[OrderedInt] {
	return list.Prepend[OrderedInt](t, as)
}
func (as virtualStack) Head() OrderedInt {
	if as.IsEmpty() {
		var t OrderedInt
		return t
	}
	return OrderedInt(as.current)
}
func (as virtualStack) Tail() interfaces.Stack[OrderedInt] {
	if as.IsEmpty() {
		return as
	}
	return virtualStack{as.current + 1, as.top}
}

func TestSomeThing(t *testing.T) {
	fftest(t, 20)
	fftest(t, 200)
}


func FuzzOrderedToRbTree(f *testing.F) {
	f.Add(uint(30))
	f.Add(uint(40))
	f.Fuzz(fftest)
}
