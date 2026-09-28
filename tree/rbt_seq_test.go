package tree

import (
	"github.com/alesgaroth/pfds-go/interfaces"
	"testing"
)

//

func TestSeq(t *testing.T) {
	var currDir interfaces.Map[OrderedString, uint64] = EmptyRbTreeMap[OrderedString, uint64]()
	currDir = currDir.Bind(".", 0)
	currDir = currDir.Bind("..", 0)
	items := currDir.Sequence()
	if items.IsEmpty() {
		t.Errorf("Suprised to see the sequence is empty")
	}
	var lines []OrderedString
	for ;!items.IsEmpty(); items = items.Tail() {
		mapentry := items.Head()
		if mapentry == nil {
			t.Errorf("This should not happen %v", items)
			continue
		}
		lines = append(lines, mapentry.Key())
		
	}
	if len(lines) != 2  || lines[0] != "." || lines[1] != ".." {
		t.Fatalf("expected 2 entires, . and .. got %v", lines)
	}
}
