package ldist

// trimPrefix trims the common prefix from r1 and r2, returning the trimmed slices.
func trimPrefix(r1, r2 []rune) ([]rune, []rune) {
	start := 0
	for start < len(r1) && start < len(r2) && r1[start] == r2[start] {
		start++
	}
	return r1[start:], r2[start:]
}

// trimSuffix trims the common suffix from r1 and r2, returning the trimmed slices.
func trimSuffix(r1, r2 []rune) ([]rune, []rune) {
	for len(r1) > 0 && len(r2) > 0 && r1[len(r1)-1] == r2[len(r2)-1] {
		r1 = r1[:len(r1)-1]
		r2 = r2[:len(r2)-1]
	}
	return r1, r2
}

// commonPrefixLen returns the length of the common prefix of r1 and r2.
func commonPrefixLen(r1, r2 []rune) int {
	m := min(len(r1), len(r2))
	l := 0

	for l < m && r1[l] == r2[l] {
		l++
	}

	return l
}

// commonSuffixLen returns the length of the common suffix of r1 and r2.
func commonSuffixLen(r1, r2 []rune) int {
	m := min(len(r1), len(r2))
	l := 0

	s1 := r1[len(r1)-m:]
	s2 := r2[len(r2)-m:]

	for l < m && s1[len(s1)-1-l] == s2[len(s2)-1-l] {
		l++
	}

	return l
}

// commonAffixes returns the lengths of the common prefix and suffix of r1 and r2.
// The suffix is computed on the remaining part after the prefix to prevent overlap.
func commonAffixes(r1, r2 []rune) (int, int) {
	pre := commonPrefixLen(r1, r2)
	suf := commonSuffixLen(r1[pre:], r2[pre:])
	return pre, suf
}
