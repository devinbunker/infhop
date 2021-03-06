package pancake

import (
	"testing"
)

func TestAdd(t *testing.T) {
	in := []Pancake{{false}, {true}, {false}}
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
		if in[i] != val {
			t.Fatalf("incorrect value at position %d: wanted %v, got %v\n", i, in[i], outVal)
		}
	}
}

func TestFlip(t *testing.T) {
	s := Stack{}
	ok := s.Flip(0)
	if ok {
		t.Fatalf("successfully flipped an empty stack\n")
	}
	ok = s.Flip(1)
	if ok {
		t.Fatalf("flipped past the end of a stack\n")
	}

	testData := []struct {
		in     []Pancake
		pos    int
		expect string
	}{
		{[]Pancake{{true}, {true}, {false}}, 2, "+--"},
		{[]Pancake{{true}, {true}, {false}}, 1, "---"},
		{[]Pancake{{true}, {true}, {false}}, 0, "-+-"},
	}

	for _, data := range testData {
		s := Stack{}
		s.Add(data.in...)
		s.Flip(data.pos)
		outStr := s.String()
		if outStr != data.expect {
			t.Errorf("wanted %s, got %s\n", data.expect, outStr)
		}

	}
}

func TestLen(t *testing.T) {
	in := []Pancake{{false}, {true}, {false}}
	s := Stack{}
	s.Add(in...)
	if s.Len() != len(in) {
		t.Fatalf("wanted %d, got %d\n", len(in), s.Len())
	}
}

func TestPeek(t *testing.T) {
	in := []Pancake{{false}, {true}, {false}}
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
	testData := [][]Pancake{
		[]Pancake{{true}, {true}, {true}, {true}},
		[]Pancake{{false}, {false}, {false}, {false}},
		[]Pancake{{false}, {true}, {false}, {true}},
		[]Pancake{{false}, {false}, {false}, {true}},
	}

	for _, in := range testData {
		s := Stack{}
		s.Add(in...)
		_, ok := s.Normalize()
		if !ok {
			t.Fatalf("failed while flipping\n")
		}
		outStr := s.String()
		if len(outStr) != len(in) {
			t.Errorf("expected %d length output, got %d\n", len(in), len(outStr))
		}
		for _, val := range outStr {
			if val != '+' {
				t.Errorf("failed normalization %s", outStr)
				break
			}
		}
	}
}

func TestString(t *testing.T) {
	in := []Pancake{{false}, {true}, {false}}
	expectStr := "-+-"
	s := Stack{}
	s.Add(in...)

	str := s.String()
	if str != expectStr {
		t.Fatalf("wanted %s, got %s\n", expectStr, str)
	}
}
