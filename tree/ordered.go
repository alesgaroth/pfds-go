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

type OrderedUint64 uint64

func (i OrderedUint64) Eq(j OrderedUint64) bool {
	return i == j
}
func (i OrderedUint64) Lt(j OrderedUint64) bool {
	return i < j
}
func (i OrderedUint64) Leq(j OrderedUint64) bool {
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
