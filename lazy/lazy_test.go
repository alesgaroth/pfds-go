package lazy

import (
	"testing"
)

func TestMe(t *testing.T) {
	runs := 0
	l := Lzy(func() int {
		runs += 1
		return 3
	})
	if l == nil {
		t.Fatalf("Expected to be able to force a lazy value")
	}
	if runs != 0 {
		t.Fatalf("Didn't expect the lazy value to be evaluated yet")
	}
	q := l.Force()
	if runs != 1 {
		t.Fatalf("Did expect the lazy value to be evaluated")
	}
	if q != 3 {
		t.Fatalf("it's no good to get the wrong return value")
	}
l.Force()
	if runs != 1 {
		t.Fatalf("Did expect the lazy value to be evaluated")
	}
}
