package pancake

import (
	"testing"
)

func TestAdd(t *testing.T) {
	in := []bool{false, true, false}
	s := Stack{}
	s.Add(in...)
	if s.Len() != len(in) {
		t.Fatalf("length %d, wanted %d\n", s.Len(), len(in))
	}
	for i, val := range in {
		outVal, ok := s.Peek(i)
		if !ok {
			t.Errorf("failed to read value at position %d\n", i)
		}
		if in[i] == val {
			t.Fatalf("incorrect values in list: %v\n", outVal)
		}
	}
}

func TestFlip(t *testing.T) {

}

func TestLen(t *testing.T) {
	in := []bool{false, true, false}
	s := Stack{}
	s.Add(in...)
	if s.Len() != len(in) {
		t.Fatalf("wanted %d, got %d\n", len(in), s.Len())
	}
}

func TestPeek(t *testing.T) {
	in := []bool{false, true, false}
	s := Stack{}
	s.Add(in...)

	// try peeking outside the stack extents
	var pos int
	pos = -1
	_, ok := s.Peek(pos)
	if ok {
		t.Fatalf("peek succeeded when it should have failed at position %d\n", pos)
	}
	pos = len(in)
	_, ok = s.Peek(pos)
	if ok {
		t.Fatalf("peek succeeded when it should have failed at position %d\n", pos)
	}

	for i, val := range in {
		peekedVal, ok := s.Peek(i)
		if !ok {
			t.Errorf("peek at position %d failed\n", i)
		}
		if val != peekedVal {
			t.Errorf("bad value at position %d: wanted %v, got %v\n", i, val, peekedVal)
		}
	}
}

func TestNormalize(t *testing.T) {

}

func TestString(t *testing.T) {
	in := []bool{false, true, false}
	expectStr := "-+-"
	s := Stack{}
	s.Add(in...)

	str := s.String()
	if str != expectStr {
		t.Fatalf("wanted %s, got %s\n", expectStr, str)
	}
}

