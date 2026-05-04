package ldist

import "testing"

func TestTrimSuffix(t *testing.T) {
	r1 := []rune("kitten")
	r2 := []rune("mitten")

	r1, r2 = trimSuffix(r1, r2)

	if string(r1) != "k" || string(r2) != "m" {
		t.Errorf("Expected r1 and r2 to be trimmed to 'k' and 'm', but got '%s' and '%s'", string(r1), string(r2))
	}

	r1 = []rune("kitten")
	r2 = []rune("sitting")

	r1, r2 = trimSuffix(r1, r2)

	if string(r1) != "kitten" || string(r2) != "sitting" {
		t.Errorf("Expected r1 and r2 to be unchanged, but got '%s' and '%s'", string(r1), string(r2))
	}
}

func TestTrimPrefix(t *testing.T) {
	r1 := []rune("kitten")
	r2 := []rune("kitchen")

	r1, r2 = trimPrefix(r1, r2)

	if string(r1) != "ten" || string(r2) != "chen" {
		t.Errorf("Expected r1 and r2 to be trimmed to 'ten' and 'chen', but got '%s' and '%s'", string(r1), string(r2))
	}

	r1 = []rune("kitten")
	r2 = []rune("sitting")

	r1, r2 = trimPrefix(r1, r2)

	if string(r1) != "kitten" || string(r2) != "sitting" {
		t.Errorf("Expected r1 and r2 to be unchanged, but got '%s' and '%s'", string(r1), string(r2))
	}
}

func TestCommonAffixes(t *testing.T) {
	r1 := []rune("kitten")
	r2 := []rune("kitchen")

	pre, suf := commonAffixes(r1, r2)

	if pre != 3 || suf != 2 {
		t.Errorf("Expected common prefix and suffix lengths to be 3 and 2, but got %d and %d", pre, suf)
	}

	r1 = []rune("kitten")
	r2 = []rune("sitting")

	pre, suf = commonAffixes(r1, r2)

	if pre != 0 || suf != 0 {
		t.Errorf("Expected common prefix and suffix lengths to be 0 and 0, but got %d and %d", pre, suf)
	}
}
