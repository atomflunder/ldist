package ldist

import "slices"

// Match returns true if the normalized similarity between s1 and s2 is greater than or equal to the cutoff.
func Match(s1, s2 string, weights Weights, cutoff float64, similarityFunc SimilarityFunc, opts ...Option) bool {
	sim := similarityFunc(s1, s2, weights, opts...)

	return sim >= cutoff
}

type SimilarityFunc func(s1, s2 string, weights Weights, opts ...Option) float64

// BestMatchResult represents a candidate string and its similarity score to the input string.
type BestMatchResult struct {
	Candidate  string  `json:"candidate"`
	Similarity float64 `json:"similarity"`
}

// GetBestMatch returns the candidate string with the highest similarity to s1 that meets the cutoff threshold.
// If no candidate meets the cutoff, it returns nil.
func GetBestMatch(s1 string, candidates []string, weights Weights, cutoff float64, similarityFunc SimilarityFunc, opts ...Option) *BestMatchResult {
	best := ""
	bestSim := 0.0

	for _, candidate := range candidates {
		sim := similarityFunc(s1, candidate, weights, opts...)
		if sim > bestSim {
			bestSim = sim
			best = candidate
		}
	}

	if bestSim >= cutoff {
		return &BestMatchResult{Candidate: best, Similarity: bestSim}
	}
	return nil
}

// GetBestMatches returns a slice of BestMatchResult for all candidates that have a similarity to s1 greater than or equal to the cutoff.
func GetBestMatches(s1 string, candidates []string, weights Weights, cutoff float64, similarityFunc SimilarityFunc, opts ...Option) []BestMatchResult {
	results := []BestMatchResult{}

	for _, candidate := range candidates {
		sim := similarityFunc(s1, candidate, weights, opts...)
		if sim >= cutoff {
			results = append(results, BestMatchResult{Candidate: candidate, Similarity: sim})
		}
	}

	return results
}

// GetBestMatchesSorted returns a slice of BestMatchResult for all candidates that have a similarity to s1 greater than or equal to the cutoff,
// sorted in descending order of similarity.
func GetBestMatchesSorted(s1 string, candidates []string, weights Weights, cutoff float64, similarityFunc SimilarityFunc, opts ...Option) []BestMatchResult {
	results := GetBestMatches(s1, candidates, weights, cutoff, similarityFunc, opts...)

	slices.SortFunc(results, func(a, b BestMatchResult) int {
		if a.Similarity > b.Similarity {
			return -1
		} else if a.Similarity < b.Similarity {
			return 1
		}
		return 0
	})

	return results
}
