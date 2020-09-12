package pancake

type Stack struct {
	list []bool
}

// Add appends new pancakes to the stack, with
// each boolean argument specifying whether the
// pancake is face up.
func (s Stack) Add(faceUp ...bool) {

}

// Flip flips over the stack from the top to the
// pancake at position i.
// If the i parameter is beyond the extents of the
// stack, then false is returned.
func (s Stack) Flip(i int) (ok bool) {
	return false
}

// Len returns the number of pancakes in the stack.
func (s Stack) Len() int {
	return len(s.list)
}

// Peek examines an individual pancake in the list,
// returning whether it is face up. If the i
// parameter is beyond the current extents of the
// stack, then false will be returned in the second
// return value.
func (s Stack) Peek(i int) (faceUp bool, ok bool) {
	return false, false
}

// Normalize performs the necessary operations to
// make sure every pancake in the stack is face up.
func (s Stack) Normalize() {

}

// String returns a stringified version of the
// entire pancake stack, with face up pancakes
// represented by + and face down by -. The
// left-most value is the top of the stack.
func (s Stack) String() string {
	buf := make([]byte, len(s.list))
	for i, val := range s.list {
		if val {
			buf[i] = '+'
		} else {
			buf[i] = '-'
		}
	}
	return string(buf)
}
