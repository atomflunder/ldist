package ldist

import "testing"

func TestGetEditops(t *testing.T) {
	r1 := []rune("Lorem ipsum.")
	r2 := []rune("XYZLorem ABC iPsum")

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

	actual := getEditops(r1, r2)

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
	r1 := []rune("")
	r2 := []rune("")

	expected := []editop{}
	actual := getEditops(r1, r2)

	if len(actual) != len(expected) {
		t.Fatalf("Expected %d edit operations, but got %d", len(expected), len(actual))
	}

}

func TestLongStringGetEditops(t *testing.T) {
	// 66 runes, triggers DP fallback (>64)
	r1 := []rune("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmn")
	r2 := []rune("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmo")

	ops := getEditops(r1, r2)

	if len(ops) != 1 {
		t.Fatalf("Expected 1 edit operation, got %d", len(ops))
	}

	if ops[0].tag != tagReplace {
		t.Errorf("Expected replace operation, got %s", ops[0].tag)
	}

	if ops[0].srcPos != 65 || ops[0].destPos != 65 {
		t.Errorf("Expected operation at position 65, got srcPos=%d destPos=%d", ops[0].srcPos, ops[0].destPos)
	}
}

func TestLongStringGetEditopsDPFallback(t *testing.T) {
	// Completely different 65-rune strings: no common prefix/suffix to trim,
	// so r1Trimmed stays >64 and the DP fallback path is exercised.
	r1 := make([]rune, 65)
	r2 := make([]rune, 65)
	for i := range r1 {
		r1[i] = 'a'
		r2[i] = 'b'
	}

	ops := getEditops(r1, r2)

	if len(ops) != 65 {
		t.Fatalf("Expected 65 edit operations, got %d", len(ops))
	}

	for i, op := range ops {
		if op.tag != tagReplace {
			t.Errorf("Expected replace at index %d, got %s", i, op.tag)
		}
	}
}

func TestGetEditopsDPDistZero(t *testing.T) {
	// Identical inputs → dist == 0 → return empty
	r := []rune("abc")
	r2 := make([]rune, len(r))
	copy(r2, r)

	ops := getEditopsDP(r, r2, 0)
	if len(ops) != 0 {
		t.Errorf("Expected 0 ops for identical strings, got %d", len(ops))
	}
}

func TestGetEditopsDPAllReplace(t *testing.T) {
	// All characters differ → all replaces in the main loop
	ops := getEditopsDP([]rune("abc"), []rune("xyz"), 0)
	if len(ops) != 3 {
		t.Fatalf("Expected 3 ops, got %d", len(ops))
	}
	for i, op := range ops {
		if op.tag != tagReplace {
			t.Errorf("Op %d: expected replace, got %s", i, op.tag)
		}
	}
}

func TestGetEditopsDPRemainingDeletions(t *testing.T) {
	// r1="abx", r2="x": match 'x', then remaining deletes for 'a','b'
	ops := getEditopsDP([]rune("abx"), []rune("x"), 0)
	if len(ops) != 2 {
		t.Fatalf("Expected 2 ops, got %d", len(ops))
	}
	for _, op := range ops {
		if op.tag != tagDelete {
			t.Errorf("Expected delete, got %+v", op)
		}
	}
}

func TestGetEditopsDPRemainingInsertions(t *testing.T) {
	// r1="x", r2="abx": match 'x', then remaining inserts for 'a','b'
	ops := getEditopsDP([]rune("x"), []rune("abx"), 0)
	if len(ops) != 2 {
		t.Fatalf("Expected 2 ops, got %d", len(ops))
	}
	for _, op := range ops {
		if op.tag != tagInsert {
			t.Errorf("Expected insert, got %+v", op)
		}
	}
}

func TestGetEditopsDPMainLoopDelete(t *testing.T) {
	// r1="aab", r2="ba": backtrace hits delete branch in main loop
	// DP matrix yields: replace at (0,0), match 'a' at (1,1), delete at (2,2)
	ops := getEditopsDP([]rune("aab"), []rune("ba"), 0)
	if len(ops) != 2 {
		t.Fatalf("Expected 2 ops, got %d", len(ops))
	}
	if ops[0].tag != tagReplace {
		t.Errorf("Op 0: expected replace, got %s", ops[0].tag)
	}
	if ops[1].tag != tagDelete {
		t.Errorf("Op 1: expected delete, got %s", ops[1].tag)
	}
}

func TestGetEditopsDPMainLoopInsert(t *testing.T) {
	// r1="ba", r2="aab": backtrace hits insert branch in main loop
	// DP matrix yields: replace at (0,0), match 'a' at (1,1), insert at (2,2)
	ops := getEditopsDP([]rune("ba"), []rune("aab"), 0)
	if len(ops) != 2 {
		t.Fatalf("Expected 2 ops, got %d", len(ops))
	}
	if ops[0].tag != tagReplace {
		t.Errorf("Op 0: expected replace, got %s", ops[0].tag)
	}
	if ops[1].tag != tagInsert {
		t.Errorf("Op 1: expected insert, got %s", ops[1].tag)
	}
}

func TestGetEditopsDPPrefixOffset(t *testing.T) {
	// Verify prefixLen offsets positions correctly
	ops := getEditopsDP([]rune("a"), []rune("b"), 5)
	if len(ops) != 1 {
		t.Fatalf("Expected 1 op, got %d", len(ops))
	}
	if ops[0].srcPos != 5 || ops[0].destPos != 5 {
		t.Errorf("Expected positions offset by 5, got srcPos=%d destPos=%d", ops[0].srcPos, ops[0].destPos)
	}
}

func TestUnicodeGetEditops(t *testing.T) {
	r1 := []rune("café")
	r2 := []rune("cafe")

	ops := getEditops(r1, r2)

	if len(ops) != 1 {
		t.Fatalf("Expected 1 edit operation for café→cafe, got %d", len(ops))
	}

	if ops[0].tag != tagReplace {
		t.Errorf("Expected replace, got %s", ops[0].tag)
	}

	// Position 3: é → e
	if ops[0].srcPos != 3 || ops[0].destPos != 3 {
		t.Errorf("Expected operation at position 3, got srcPos=%d destPos=%d", ops[0].srcPos, ops[0].destPos)
	}
}
