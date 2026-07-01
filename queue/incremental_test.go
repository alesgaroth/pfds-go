package queue

import (
	"fmt"
	"github.com/alesgaroth/pfds-go/interfaces"
	"github.com/alesgaroth/pfds-go/list"
	"testing"
)

func TestCanCreateIncrementalQueue(t *testing.T) {
	var q Queue[int] = EmptyIncrementalQueue[int]()
	if !q.IsEmpty() {
		t.Fatalf("how did a new empty IncrementalQueue have something in it?")
	}
	q2 := q.Cons(1).Cons(2)
	if q2.IsEmpty() {
		t.Fatalf("how did a IncrementalQueue after 2 additions have nothing in it?")
	}
	q3 := q2.Tail()
	if q3.IsEmpty() {
		t.Fatalf("how did a IncrementalQueue after 2 additions and a removal have nothing in it?")
	}
	q4 := q3.Tail()
	if !q4.IsEmpty() {
		t.Fatalf("how did a emptied IncrementalQueue have something in it?")
	}

}

func TestOneItem(t *testing.T) {
	var q Queue[int] = EmptyIncrementalQueue[int]()
	q = q.Cons(1)
	data := q.Head()
	if data != 1 {
		t.Fatalf("you should be able to add something the queue and get it back")
	}
}

func TestTwoItems(t *testing.T) {
	var q = EmptyIncrementalQueue[int]().(*IncrementalQueue[int])
	q = q.Cons(2).(*IncrementalQueue[int])
	checkInvariants(t, q)
	if q.dequeue.IsEmpty() {
		t.Fatalf("It is surprising that dequeue is empty")
	}
	q = q.Cons(1).(*IncrementalQueue[int])
	checkInvariants(t, q)
	if q.dequeue.IsEmpty() {
		t.Fatalf("It is surprising that dequeue is empty")
	}
	if q.enqueue.IsEmpty() {
		t.Fatalf("It is surprising that enqueue is empty")
	}
	data := q.Head()
	if data != 2 {
		t.Fatalf("you should be able to add something the queue and get it back")
	}
	if q.num != 2 {
		t.Errorf("two items in the queue, but %v counted", q.num)
	}
	q = q.Tail().(*IncrementalQueue[int])
	checkInvariants(t, q)
	if q.IsEmpty() {
		t.Fatalf("one remaining item in the queue, empty")
	}
	if q.num != 1 {
		t.Errorf("one remaining item in the queue, but %v counted", q.num)
	}
	data = q.Head()
	if data != 1 {
		t.Errorf("you should be able to add two things the queue and get both back (expected %v got %v)", 1, data)
	}
	if q.dequeue.IsEmpty() {
		t.Errorf("It is surprising that dequeue is empty")
	}
	if !q.enqueue.IsEmpty() {
		t.Errorf("It is surprising that enqueue is not empty")
	}
}

func TestThreeItems(t *testing.T) {
	var q = EmptyIncrementalQueue[int]().(*IncrementalQueue[int])
	q = q.Cons(3).(*IncrementalQueue[int])
	q = q.Cons(2).(*IncrementalQueue[int])
	if q.dequeue.IsEmpty() {
		t.Fatalf("It is surprising that dequeue is empty")
	}
	q = q.Cons(1).(*IncrementalQueue[int])
	if q.dequeue.IsEmpty() {
		t.Fatalf("It is surprising that dequeue is empty")
	}
	data := q.Head()
	if data != 3 {
		t.Fatalf("you should be able to add something the queue and get it back")
	}
	if q.num != 3 {
		t.Errorf("three items in the queue, but %v counted", q.num)
	}
	q = q.Tail().(*IncrementalQueue[int])
	if q.num != 2 {
		t.Errorf("two remaining items in the queue, but %v counted", q.num)
	}
	data = q.Head()
	if data != 2 {
		t.Errorf("you should be able to add three things the queue and get the second back (expected %v got %v)", 2, data)
	}

	if q.dequeue.IsEmpty() {
		t.Errorf("It is surprising that dequeue is empty")
	}
	if !q.enqueue.IsEmpty() {
		t.Errorf("It is surprising that enqueue is not empty")
	}
	q = q.Tail().(*IncrementalQueue[int])
	if q.num != 1 {
		t.Errorf("one remaining item in the queue, but %v counted", q.num)
	}
	data = q.Head()
	if data != 1 {
		t.Errorf("you should be able to add three things the queue and get the third back (expected %v got %v)", 1, data)
	}
	if q.dequeue.IsEmpty() {
		t.Errorf("It is surprising that dequeue is empty")
	}
	if !q.enqueue.IsEmpty() {
		t.Errorf("It is surprising that enqueue is not empty")
	}
}

func checkInvariants(t *testing.T, q *IncrementalQueue[int]) {
	// q.dequeue is not empty unless the whole thing is empty
	if q.IsEmpty() {
		if q != nil {
			if !q.dequeue.IsEmpty() {
				t.Errorf("weird that q is empty but its dequeue is not")
			}
			if !q.enqueue.IsEmpty() {
				t.Errorf("weird that q is empty but its enqueue is not")
			}
			if q.progress != nil {
				t.Errorf("weird that q is empty but its progress is not")
			}
		}
		return
	}
	// length of q.enqueue is <= q.num/2
	if Length[int](q.enqueue) > q.num/2 {
		t.Errorf("enqueue should never be bigger than half the queued items, or you won't be able to reverse it!")
	}

	// length of q.enqueue + q.enSpace == (q.num/2) + 1
	if Length[int](q.enqueue)+q.enSpace != ((q.num+q.enSpace)/2)+1 {
		t.Errorf("len(enqueue)  + enSpace should be equal to one more than half the full number, or you won't be able to reverse it!")
	}
	if q.progress != nil {
		// length of q.progress.input < q.enSpace
		lenProgInput := Length[int](q.progress.input)

		if lenProgInput >= q.enSpace {
			t.Errorf("len(progress.input) should be less than the space left in enqueue")
		}
		// length of q.progress.input < length of q.dequeue + length of all items in internal
		if !q.progress.internal.IsEmpty()  {
		  internalnum := 0
			for internal := q.progress.internal; !internal.IsEmpty(); internal = internal.Tail() {
				internalnum += Length(internal.Head().(interfaces.Stack[int]))
			}
		  
			if lenProgInput > Length(q.dequeue) + internalnum {
				t.Errorf("len(progress.input) %v should be less than what's in already in the correct order %v + %v", lenProgInput, Length(q.dequeue), internalnum)
			}
		} else {
			if lenProgInput > Length(q.dequeue) {
				t.Errorf("len(progress.input) %v should be less than what's in already in the correct order %v", lenProgInput, Length(q.dequeue))
			}
		}
		// which is the same as
		// q.num - length of q.enqueue - length of q.progress.{input, reversed}
		if lenProgInput > q.num-Length[int](q.enqueue)-lenProgInput-Length[int](q.progress.reversed) {
			t.Errorf("len(progress.input) %v should be less than what's already in the correct order %v-%v-%v-%v", lenProgInput,
				q.num,Length[int](q.enqueue),lenProgInput,Length[int](q.progress.reversed))
		}
	}

}

func Length[T any](s interfaces.Stack[T]) int {
	return list.Length(s)
}

func ftest(t *testing.T, cnt int) {
	q := EmptyIncrementalQueue[int]()
	for k := 0; k < cnt; k += 1 {
		q = q.Cons(k)
	}
	checkInvariants(t, q.(*IncrementalQueue[int]))
	for k := 0; k < cnt; k += 1 {
		data := q.Head()
		if data != k {
			t.Errorf("expected %v got %v when working with %d items", k, data, cnt)
			t.Errorf("%v", q)
		}
		q = q.Tail()
	}
	if !q.IsEmpty() {
		t.Errorf("expected the queue to be empty after removing as many (%d) as I added", cnt)
	}
}

func TestLarger(t *testing.T) {
	for j:= 6; j < 9; j += 1 {
		ftest(t, j)
	}
}

func (p *proGress[T]) String() string {
	return fmt.Sprintf("%v %v %v", Length(p.input), Length(p.reversed), p.internal)
}
func (q *IncrementalQueue[T]) String() string {
	return fmt.Sprintf("%v %v %v %v <%v>", Length(q.enqueue), Length(q.dequeue), q.num, q.enSpace, q.progress)
}
