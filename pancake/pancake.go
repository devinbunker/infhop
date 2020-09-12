package pancake

type Stack struct {
	list []bool
}

// Add appends new pancakes to the stack, with
// each boolean argument specifying whether the
// pancake is face up.
func (s Stack) Add(faceUp ...bool) {

}

// Peek examines an individual pancake in the list,
// returning whether it is face up. If the i
// parameter is beyond the current extents of the
// stack, then false will be returned in the second
// return value.
func (s Stack) Peek(i int) (faceUp bool, ok bool) {
	return false, false
}

// Flip flips over the stack from the top to the
// pancake at position i.
// If the i parameter is beyond the extents of the
// stack, then false is returned.
func (s Stack) Flip(i int) (ok bool) {
	return false
}

// Normalize performs the necessary operations to
// make sure every pancake in the stack is face up.
func (s Stack) Normalize() {

}

