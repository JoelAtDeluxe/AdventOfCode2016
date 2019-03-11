package tooling

import (
	"testing"
)

func TestFindLetter(t *testing.T) {
	base := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	copy := Copy(base)
	pos := FindLetter(&base, "c")

	if pos != 2 {
		t.Error("Find can't find something that's there")
	}
	if !Eq(base, copy) {
		t.Error("Base seems to be modified during the find")
	}

	pos = FindLetter(&base, "z")
	if pos != -1 {
		t.Error("We are finding things that are not present")
	}
	if !Eq(base, copy) {
		t.Error("Base seems to be modified during the find")
	}
}

func TestReversePorition(t *testing.T) {
	base := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	expected := []rune{'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}
	ReversePortion(&base, 0, len(base)-1)

	if !Eq(base, expected) {
		t.Error("Reverse Pos not working for whole string")
	}
	base = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	expected = []rune{'a', 'b', 'f', 'e', 'd', 'c', 'g', 'h'}
	ReversePortion(&base, 2, 5)
	if !Eq(base, expected) {
		t.Error("Reverse Pos not working for portion of string")
	}

	base = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	expected = []rune{'a', 'b', 'f', 'e', 'd', 'c', 'g', 'h'}
	ReversePortion(&base, 5, 2)
	if !Eq(base, expected) {
		t.Error("Reverse Pos not properlly swapping start, end")
	}

	base = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	expected = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	ReversePortion(&base, 0, 0)
	if !Eq(base, expected) {
		t.Error("Reverse Pos is doing something weird for single letter portions")
	}
}

func TestSwapIdx(t *testing.T) {
	base := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	expected := []rune{'a', 'b', 'f', 'd', 'e', 'c', 'g', 'h'}
	SwapIdx(&base, 2, 5)
	if !Eq(base, expected) {
		t.Error("Swap Idx not working")
	}
}

func TestRotate(t *testing.T) {
	base := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	expected := []rune{'c', 'd', 'e', 'f', 'g', 'h', 'a', 'b'}
	Rotate(&base, 2, 0, len(base)-1, true)
	if !Eq(base, expected) {
		t.Error("Rot left not working on whole string")
	}

	base = []rune{'c', 'd', 'e', 'f', 'g', 'h', 'a', 'b'}
	expected = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	Rotate(&base, 2, 0, len(base)-1, false)
	if !Eq(base, expected) {
		t.Error("Rot right not working on whole string")
	}

	base = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	expected = []rune{'a', 'b', 'd', 'e', 'f', 'c', 'g', 'h'}
	Rotate(&base, 1, 2, 5, true)
	if !Eq(base, expected) {
		t.Error("Rot left not working on partial string")
	}

	base = []rune{'a', 'b', 'd', 'e', 'f', 'c', 'g', 'h'}
	expected = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	Rotate(&base, 1, 2, 5, false)
	if !Eq(base, expected) {
		t.Error("Rot right not working on partial string")
	}

	base = []rune{'a', 'b', 'e', 'f', 'c', 'd', 'g', 'h'}
	expected = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	Rotate(&base, 2, 5, 2, false)
	if !Eq(base, expected) {
		t.Error("Rotate swap endpoints not working")
	}

	base = []rune{'a', 'b', 'e', 'f', 'c', 'd', 'g', 'h'}
	expected = Copy(base)
	Rotate(&base, 0, 2, 5, false)
	if !Eq(base, expected) {
		t.Error("Rot 0 has an effect on base")
	}

	base = []rune{'a', 'b', 'e', 'f', 'c', 'd', 'g', 'h'}
	expected = Copy(base)
	Rotate(&base, 1, 2, 2, false)
	if !Eq(base, expected) {
		t.Error("Rot on single letter range has an effect on base")
	}
}

func TestEq(t *testing.T) {
	base := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	copy := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	if !Eq(base, copy) {
		t.Error("Equality test is failing!")
	}
	base = []rune{}
	copy = []rune{}
	if !Eq(base, copy) {
		t.Error("Equality test is failing!")
	}

	base = []rune{'a'}
	copy = []rune{'a'}
	if !Eq(base, copy) {
		t.Error("Equality test is failing!")
	}

	base = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	copy = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	if Eq(base, copy) {
		t.Error("Equality test is failing!")
	}

	base = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	copy = []rune{'z', 'b', 'c', 'd', 'e', 'f', 'g'}
	if Eq(base, copy) {
		t.Error("Equality test is failing!")
	}
}

func TestCopy(t *testing.T) {
	base := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	copy := Copy(base)
	if !Eq(base, copy) {
		t.Error("Copy not producing an exact copy")
	}

	copy[0] = 'z'
	if Eq(base, copy) {
		t.Error("Copy seems to be pointing to the original, not to new memory")
	}
}

// Helpers
func Copy(original []rune) []rune {
	copy := make([]rune, len(original))

	for i := range original {
		copy[i] = original[i]
	}
	return copy
}

func Eq(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
