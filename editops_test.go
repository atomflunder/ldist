package ldist

import "testing"

func TestGetEditops(t *testing.T) {
	s1 := "Lorem ipsum."
	s2 := "XYZLorem ABC iPsum"

	expected := []editop{
		{tag: tagInsert, srcPos: 0, destPos: 0},
		{tag: tagInsert, srcPos: 0, destPos: 1},
		{tag: tagInsert, srcPos: 0, destPos: 2},
		{tag: tagInsert, srcPos: 6, destPos: 9},
		{tag: tagInsert, srcPos: 6, destPos: 10},
		{tag: tagInsert, srcPos: 6, destPos: 11},
		{tag: tagInsert, srcPos: 6, destPos: 12},
		{tag: tagReplace, srcPos: 7, destPos: 14},
		{tag: tagDelete, srcPos: 11, destPos: 18},
	}

	actual := getEditops(s1, s2)

	if len(actual) != len(expected) {
		t.Fatalf("Expected %d edit operations, but got %d", len(expected), len(actual))
	}

	for i, op := range expected {
		if actual[i] != op {
			t.Errorf("Expected edit operation %d to be %+v, but got %+v", i, op, actual[i])
		}
	}
}

func TestEdgeCasesGetEditops(t *testing.T) {
	s1 := ""
	s2 := ""

	expected := []editop{}
	actual := getEditops(s1, s2)

	if len(actual) != len(expected) {
		t.Fatalf("Expected %d edit operations, but got %d", len(expected), len(actual))
	}

}
