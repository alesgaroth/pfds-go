package tree

import (
	"fmt"
	"github.com/alesgaroth/pfds-go/interfaces"
	"github.com/alesgaroth/pfds-go/list"
	"math/bits"
)

func OrderedToRbTree[T Ordered[T]](s interfaces.Stack[T]) *RbTree[T] {
	ln := list.Length(s)
	if ln < 1 {
		return nil
	}
	k := bits.Len64(uint64(ln+1)) - 1
	fmt.Printf("ln=%d k=%d\n", ln, k)
	tree, leftovers := orderedToRbTree(s, ln, k-1)
	if leftovers != nil && !leftovers.IsEmpty() {
		panic(fmt.Sprintf("oops, the stack should be empty after building the tree  from %v items, left %v", list.Length(s), list.Length(leftovers)))
	}
	fmt.Printf("ln=%d k=%d ===> ", ln, k)
	fmt.Printf("%v\n", tree)
	return tree
}

func orderedToRbTree[T Ordered[T]](s interfaces.Stack[T], ln, k int) (*RbTree[T], interfaces.Stack[T]) {
	if ln < 1 {
		fmt.Printf("    ln=%d k=%d\n", ln, k)
		return nil, s
	}
	if k < 0 {
		fmt.Printf("Red ln=%d k=%d data=%v\n", ln, k, s.Head())
		return &RbTree[T]{
			colour: Red,
			left:   nil,
			right:  nil,
			data:   s.Head(),
		}, s.Tail()
	}
	left, s := orderedToRbTree(s, ln/2, k-1)
	val := s.Head()
	right, s := orderedToRbTree(s.Tail(), (ln-1)/2, k-1)
	fmt.Printf("Blk ln=%d k=%d data=%v\n", ln, k, val)
	return &RbTree[T]{
		colour: Black,
		left:   left,
		right:  right,
		data:   val,
	}, s
}
