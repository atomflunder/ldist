package ldist

// trimPrefix trims the common prefix from s1 and s2, modifying the strings in place.
func trimPrefix(s1, s2 *string) {
	// Trims common prefix, speeds up calculations for long strings.
	start := 0
	for start < len(*s1) && start < len(*s2) && (*s1)[start] == (*s2)[start] {
		start++
	}
	if start > 0 {
		*s1 = (*s1)[start:]
		*s2 = (*s2)[start:]
	}

}

// trimSuffix trims the common suffix from s1 and s2, modifying the strings in place.
func trimSuffix(s1, s2 *string) {
	for len(*s1) > 0 && len(*s2) > 0 && (*s1)[len(*s1)-1] == (*s2)[len(*s2)-1] {
		*s1 = (*s1)[:len(*s1)-1]
		*s2 = (*s2)[:len(*s2)-1]
	}
}

// commonPrefixLen returns the length of the common prefix of s1 and s2.
func commonPrefixLen(s1, s2 string) int {
	m := min(len(s1), len(s2))
	l := 0

	for l < m && s1[l] == s2[l] {
		l++
	}

	return l
}

// commonSuffixLen returns the length of the common suffix of s1 and s2.
func commonSuffixLen(s1, s2 string) int {
	m := min(len(s1), len(s2))
	l := 0

	s1 = s1[len(s1)-m:]
	s2 = s2[len(s2)-m:]

	for l < m && s1[len(s1)-1-l] == s2[len(s2)-1-l] {
		l++
	}

	return l
}

// commonAffixes returns the lengths of the common prefix and suffix of s1 and s2.
func commonAffixes(s1, s2 string) (int, int) {
	pre := commonPrefixLen(s1, s2)
	suf := commonSuffixLen(s1[pre:], s2[pre:])
	return pre, suf
}
