package pancake

type Pancake struct {
	FaceUp bool
}

type Stack struct {
	list []Pancake
}

// Add appends new pancakes to the stack.
func (s *Stack) Add(p ...Pancake) {
	s.list = append(s.list, p...)
}

// Flip flips over the stack from the top to the
// pancake at position pos.
// If the pos parameter is beyond the extents of the
// stack, then false is returned.
func (s *Stack) Flip(pos int) (ok bool) {
	if pos < 0 || pos >= len(s.list) {
		return false
	}

	// create a temporary list of pancakes to
	// hold the old values, sized to contain
	// just the substack we're interested in
	buf := make([]Pancake, pos+1)

	// copy out the existing values
	for i := 0; i < pos+1; i++ {
		buf[i] = s.list[i]
	}

	// put the values back, but in reverse order
	// and flipped
	for i := 0; i < pos+1; i++ {
		buf[pos-i].FaceUp = !buf[pos-i].FaceUp
		s.list[i] = buf[pos-i]
	}

	return true
}

// Len returns the number of pancakes in the stack.
func (s *Stack) Len() int {
	return len(s.list)
}

// Peek examines an individual pancake in the list,
// returning whether it is face up. If the i
// parameter is beyond the current extents of the
// stack, then false will be returned in the second
// return value.
func (s *Stack) Peek(i int) (p Pancake, ok bool) {
	if i < 0 || i >= len(s.list) {
		return Pancake{}, false
	}
	return s.list[i], true
}

// Normalize performs the necessary operations to
// make sure every pancake in the stack is face up.
// Returns the number of flip operations performed
// in the course of doing this.
func (s *Stack) Normalize() (ops int, ok bool) {
	ops = 0

	// start from the bottom of the stack and
	// work our way up
	for i := len(s.list) - 1; i >= 0; i-- {
		// is it already face up? then skip
		p, ok := s.Peek(i)
		if !ok {
			return ops, false
		}
		if p.FaceUp {
			continue
		}

		// we need to flip, so let's first
		// make sure the top is face down
		// so it will be face up after the
		// flip
		top, ok := s.Peek(0)
		if !ok {
			return ops, false
		}
		if top.FaceUp {
			s.Flip(0)
			ops++
		}

		// do the flip!
		s.Flip(i)
		ops++
	}

	return ops, true
}

// String returns a stringified version of the
// entire pancake stack, with face up pancakes
// represented by + and face down by -. The
// left-most value is the top of the stack.
func (s *Stack) String() string {
	buf := make([]byte, len(s.list))
	for i, p := range s.list {
		if p.FaceUp {
			buf[i] = '+'
		} else {
			buf[i] = '-'
		}
	}
	return string(buf)
}
