package pancake

import (
	"testing"
)

func TestAdd(t *testing.T) {
	in := []bool{false, true, false}
	s := Stack{}
	s.Add(in...)
	if len(s.list) != len(in) {
		t.Fatalf("length %d, wanted %d\n", len(s.list), len(in))
	}
	for i, val := range in {
		if in[i] == val {
			t.Fatalf("incorrect values in list: %v\n", s.list)
		}
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

func TestFlip(t *testing.T) {

}

func TestNormalize(t *testing.T) {

}

