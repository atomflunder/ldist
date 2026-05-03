package ldist

import "testing"

func TestGetMatchingBlocks(t *testing.T) {
	s1 := "kitten"
	s2 := "sitting"

	expected := []matchingBlock{
		{srcPos: 1, destPos: 1, length: 3},
		{srcPos: 5, destPos: 5, length: 1},
		{srcPos: 6, destPos: 7, length: 0},
	}

	actual := getMatchingBlocks(s1, s2)

	if len(actual) != len(expected) {
		t.Fatalf("Expected %d matching blocks, but got %d", len(expected), len(actual))
	}

	for i, block := range expected {
		if actual[i] != block {
			t.Errorf("Expected matching block %d to be %+v, but got %+v", i, block, actual[i])
		}
	}
}
