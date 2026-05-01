package ldist

// Distance calculates the distance between two strings s1 and s2 using the provided weights for substitution, insertion, and deletion.
// Can use options to modify the input strings before calculating the distance, such as converting to lowercase, removing whitespace, or removing punctuation.
func Distance(s1, s2 string, weights Weights, opts ...Option) int {
	for _, opt := range opts {
		opt(&s1, &s2)
	}

	if len(s1) == 0 {
		return len(s2) * weights.Insertion
	} else if len(s2) == 0 {
		return len(s1) * weights.Insertion
	}

	if s1 == s2 {
		return 0
	}
	n := len(s1)
	m := len(s2)

	prev := make([]int, m+1)
	cur := make([]int, m+1)
	for j := 0; j <= m; j++ {
		prev[j] = j * weights.Insertion
	}

	for i := 1; i <= n; i++ {
		cur[0] = i * weights.Deletion
		for j := 1; j <= m; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = weights.Substitution
			}
			del := prev[j] + weights.Deletion
			ins := cur[j-1] + weights.Insertion
			sub := prev[j-1] + cost
			min := del
			if ins < min {
				min = ins
			}
			if sub < min {
				min = sub
			}
			cur[j] = min
		}
		prev, cur = cur, prev
	}

	return prev[m]
}

// NormalizedDistance calculates the normalized distance between two strings s1 and s2 using the provided weights for substitution, insertion, and deletion.
// The normalized distance is the distance divided by the sum of the lengths of the two strings, resulting in a value between 0 and 1.
// This means that the normalized distance is 0 when the strings are identical and approaches 1 as the strings become more different.
func NormalizedDistance(s1, s2 string, weights Weights, opts ...Option) float64 {
	return float64(Distance(s1, s2, weights, opts...)) / float64(len(s1)+len(s2))
}

// NormalizedSimilarity calculates the normalized similarity between two strings s1 and s2 using the provided weights for substitution, insertion, and deletion.
// The normalized similarity is 1 minus the normalized distance, resulting in a value between 0 and 1.
// This means that the normalized similarity is 1 when the strings are identical and approaches 0 as the strings become more different.
func NormalizedSimilarity(s1, s2 string, weights Weights, opts ...Option) float64 {
	return 1.0 - NormalizedDistance(s1, s2, weights, opts...)
}
