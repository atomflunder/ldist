package ldist

// Distance calculates the distance between two strings s1 and s2 using the provided weights for substitution, insertion, and deletion.
// Can use options to modify the input strings before calculating the distance, such as converting to lowercase, removing whitespace, or removing punctuation.
func Distance(s1, s2 string, weights Weights, opts ...Option) int {
	for _, opt := range opts {
		opt(&s1, &s2)
	}

	trimPrefix(&s1, &s2)
	trimSuffix(&s1, &s2)

	if len(s1) == 0 {
		return len(s2) * weights.Insertion
	} else if len(s2) == 0 {
		return len(s1) * weights.Deletion
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
			if ins < del {
				del = ins
			}
			if sub < del {
				del = sub
			}
			cur[j] = del
		}
		prev, cur = cur, prev
	}

	return prev[m]
}

// NormalizedDistance calculates the normalized distance between two strings s1 and s2 using the provided weights for substitution, insertion, and deletion.
// The normalized distance is the distance divided by the sum of the lengths of the two strings, resulting in a value between 0 and 1.
// This means that the normalized distance is 0 when the strings are identical and approaches 1 as the strings become more different.
func NormalizedDistance(s1, s2 string, weights Weights, opts ...Option) float64 {
	for _, opt := range opts {
		opt(&s1, &s2)
	}

	// Options could influence the length of the strings, so we calculate the length after applying options.
	totalLen := len(s1) + len(s2)
	if totalLen == 0 {
		return 0.0
	}

	return float64(Distance(s1, s2, weights)) / float64(totalLen)
}

// NormalizedSimilarity calculates the normalized similarity between two strings s1 and s2 using the provided weights for substitution, insertion, and deletion.
// The normalized similarity is 1 minus the normalized distance, resulting in a value between 0 and 1.
// This means that the normalized similarity is 1 when the strings are identical and approaches 0 as the strings become more different.
func NormalizedSimilarity(s1, s2 string, weights Weights, opts ...Option) float64 {
	return 1.0 - NormalizedDistance(s1, s2, weights, opts...)
}

// PartialSimilarity calculates the partial similarity between two strings s1 and s2 using the provided weights for substitution, insertion, and deletion.
// The partial similarity is a measure of how similar the shorter string is to any substring of the longer string, with a penalty based on differing lengths.
// This may yield more desirable results when comparing strings of vastly differing lengths, depending on the use-case.
func PartialSimilarity(s1, s2 string, weights Weights, opts ...Option) float64 {
	for _, opt := range opts {
		opt(&s1, &s2)
	}

	if len(s1) == 0 || len(s2) == 0 {
		return 0.0
	}

	if s1 == s2 {
		return 1.0
	}

	var longer, shorter string
	if len(s1) >= len(s2) {
		longer, shorter = s1, s2
	} else {
		longer, shorter = s2, s1
	}

	blocks := getMatchingBlocks(longer, shorter)

	filteredBlocks := make([]matchingBlock, 0, len(blocks))
	for _, block := range blocks {
		if block.length > 1 || (block.length == 1 && block.srcPos == 0 && block.destPos == 0) {
			filteredBlocks = append(filteredBlocks, block)
		}
	}

	diff := len(longer) - len(shorter)

	mult := 1.0
	switch {
	case diff >= 20:
		mult = 0.65
	case diff >= 10:
		mult = 0.75
	case diff >= 4:
		mult = 0.85
	case diff >= 1:
		mult = 0.95
	}

	scores := make([]float64, 0, len(filteredBlocks)+1)

	for _, block := range filteredBlocks {
		start := max(0, block.srcPos-block.destPos)
		end := start + len(shorter)
		if end > len(longer) {
			end = len(longer)
			start = end - len(shorter)
		}
		sub := longer[start:end]

		scores = append(scores, NormalizedSimilarity(sub, shorter, weights))
	}

	scores = append(scores, NormalizedSimilarity(longer, shorter, weights))

	var maxScore float64
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore * mult
}
