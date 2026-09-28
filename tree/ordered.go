package tree

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

type OrderedString string

func (i OrderedString) Eq(j OrderedString) bool {
	return i == j
}
func (i OrderedString) Lt(j OrderedString) bool {
	return i < j
}
func (i OrderedString) Leq(j OrderedString) bool {
	return i <= j
}
