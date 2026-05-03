package ldist

import "testing"

func TestTrimSuffix(t *testing.T) {
	s1 := "kitten"
	s2 := "mitten"

	trimSuffix(&s1, &s2)

	if s1 != "k" || s2 != "m" {
		t.Errorf("Expected s1 and s2 to be trimmed to 'k' and 'm', but got '%s' and '%s'", s1, s2)
	}

	s1 = "kitten"
	s2 = "sitting"

	trimSuffix(&s1, &s2)

	if s1 != "kitten" || s2 != "sitting" {
		t.Errorf("Expected s1 and s2 to be unchanged, but got '%s' and '%s'", s1, s2)
	}
}

func TestTrimPrefix(t *testing.T) {
	s1 := "kitten"
	s2 := "kitchen"

	trimPrefix(&s1, &s2)

	if s1 != "ten" || s2 != "chen" {
		t.Errorf("Expected s1 and s2 to be trimmed to 'itten' and 'chen', but got '%s' and '%s'", s1, s2)
	}

	s1 = "kitten"
	s2 = "sitting"

	trimPrefix(&s1, &s2)

	if s1 != "kitten" || s2 != "sitting" {
		t.Errorf("Expected s1 and s2 to be unchanged, but got '%s' and '%s'", s1, s2)
	}
}

func TestCommonAffixes(t *testing.T) {
	s1 := "kitten"
	s2 := "kitchen"

	pre, suf := commonAffixes(s1, s2)

	if pre != 3 || suf != 2 {
		t.Errorf("Expected common prefix and suffix lengths to be 3 and 2, but got %d and %d", pre, suf)
	}

	s1 = "kitten"
	s2 = "sitting"

	pre, suf = commonAffixes(s1, s2)

	if pre != 0 || suf != 0 {
		t.Errorf("Expected common prefix and suffix lengths to be 0 and 0, but got %d and %d", pre, suf)
	}
}
