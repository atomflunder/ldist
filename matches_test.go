package ldist

import (
	"fmt"
	"math"
	"testing"
)

func TestMatchBasic(t *testing.T) {
	weights := IndelWeights()

	if !Match("a test", "this is a test", weights, 0.5, NormalizedSimilarity) {
		t.Error("Expected Match to return true for 'a test' and 'this is a test' with cutoff 0.5")
	}

	if Match("a test", "this is a test", weights, 0.9, NormalizedSimilarity) {
		t.Error("Expected Match to return false for 'a test' and 'this is a test' with cutoff 0.9")
	}
}

func TestMatchWithOptions(t *testing.T) {
	weights := IndelWeights()

	if !Match("A TEST", "this is a test", weights, 0.5, NormalizedSimilarity, ToLowercase) {
		t.Error("Expected Match to return true with ToLowercase option")
	}
}

func TestMatchEdgeCases(t *testing.T) {
	weights := IndelWeights()

	if Match("", "abc", weights, 0.5, NormalizedSimilarity) {
		t.Error("Expected Match to return false for empty s1")
	}

	if !Match("abc", "abc", weights, 1.0, NormalizedSimilarity) {
		t.Error("Expected Match to return true for identical strings with cutoff 1.0")
	}

	if !Match("abc", "abc", weights, 0.0, NormalizedSimilarity) {
		t.Error("Expected Match to return true for identical strings with cutoff 0.0")
	}
}

func TestGetBestMatchBasic(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"kitten", "sitting", "bitten", "written"}
	result := GetBestMatch("kittens", candidates, weights, 0.5, NormalizedSimilarity)

	if result == nil {
		t.Fatal("Expected GetBestMatch to return a result, got nil")
	}

	if result.Candidate != "kitten" {
		t.Errorf("Expected best match to be 'kitten', got '%s'", result.Candidate)
	}
}

func TestGetBestMatchNoCandidateMeetsCutoff(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"xyz", "abc", "123"}
	result := GetBestMatch("completely different string", candidates, weights, 0.99, NormalizedSimilarity)

	if result != nil {
		t.Errorf("Expected GetBestMatch to return nil, got %+v", result)
	}
}

func TestGetBestMatchEmptyCandidates(t *testing.T) {
	weights := IndelWeights()

	result := GetBestMatch("test", []string{}, weights, 0.5, NormalizedSimilarity)

	if result != nil {
		t.Errorf("Expected GetBestMatch to return nil for empty candidates, got %+v", result)
	}
}

func TestGetBestMatchWithOptions(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"KITTEN", "SITTING", "BITTEN"}
	result := GetBestMatch("kittens", candidates, weights, 0.5, NormalizedSimilarity, ToLowercase)

	if result == nil {
		t.Fatal("Expected GetBestMatch to return a result with ToLowercase, got nil")
	}

	if result.Candidate != "KITTEN" {
		t.Errorf("Expected best match to be 'KITTEN', got '%s'", result.Candidate)
	}
}

func TestGetBestMatchesBasic(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"kitten", "sitting", "bitten", "written", "xyz"}
	results := GetBestMatches("kittens", candidates, weights, 0.5, NormalizedSimilarity)

	if len(results) == 0 {
		t.Fatal("Expected GetBestMatches to return at least one result")
	}

	for _, r := range results {
		if r.Similarity < 0.5 {
			t.Errorf("Result '%s' has similarity %f, which is below cutoff 0.5", r.Candidate, r.Similarity)
		}
	}
}

func TestGetBestMatchesNoneMatchCutoff(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"xyz", "123", "!!!"}
	results := GetBestMatches("completely different", candidates, weights, 0.99, NormalizedSimilarity)

	if len(results) != 0 {
		t.Errorf("Expected 0 results, got %d", len(results))
	}
}

func TestGetBestMatchesEmptyCandidates(t *testing.T) {
	weights := IndelWeights()

	results := GetBestMatches("test", []string{}, weights, 0.5, NormalizedSimilarity)

	if len(results) != 0 {
		t.Errorf("Expected 0 results for empty candidates, got %d", len(results))
	}
}

func TestGetBestMatchesWithOptions(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"KITTEN", "SITTING", "BITTEN"}
	results := GetBestMatches("kittens", candidates, weights, 0.5, NormalizedSimilarity, ToLowercase)

	if len(results) == 0 {
		t.Fatal("Expected GetBestMatches to return results with ToLowercase option")
	}

	found := false
	for _, r := range results {
		if r.Candidate == "KITTEN" {
			found = true
		}
	}
	if !found {
		t.Error("Expected 'KITTEN' to be among the results")
	}
}

func TestGetBestMatchesSortedBasic(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"kitten", "sitting", "bitten", "written"}
	results := GetBestMatchesSorted("kittens", candidates, weights, 0.5, NormalizedSimilarity)

	if len(results) == 0 {
		t.Fatal("Expected GetBestMatchesSorted to return at least one result")
	}

	for i := 1; i < len(results); i++ {
		if results[i].Similarity > results[i-1].Similarity {
			t.Errorf("Results not sorted in descending order: index %d (%f) > index %d (%f)",
				i, results[i].Similarity, i-1, results[i-1].Similarity)
		}
	}
}

func TestGetBestMatchesSortedNoneMatchCutoff(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"xyz", "123"}
	results := GetBestMatchesSorted("completely different", candidates, weights, 0.99, NormalizedSimilarity)

	if len(results) != 0 {
		t.Errorf("Expected 0 results, got %d", len(results))
	}
}

func TestGetBestMatchesSortedWithOptions(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"KITTEN", "SITTING", "BITTEN", "WRITTEN"}
	results := GetBestMatchesSorted("kittens", candidates, weights, 0.5, NormalizedSimilarity, ToLowercase)

	if len(results) == 0 {
		t.Fatal("Expected results with ToLowercase option")
	}

	for i := 1; i < len(results); i++ {
		if results[i].Similarity > results[i-1].Similarity {
			t.Errorf("Results not sorted in descending order: index %d (%f) > index %d (%f)",
				i, results[i].Similarity, i-1, results[i-1].Similarity)
		}
	}
}

func TestGetBestMatchesSortedFirstIsBest(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"kitten", "sitting", "bitten", "written"}
	results := GetBestMatchesSorted("kittens", candidates, weights, 0.5, NormalizedSimilarity)

	if len(results) == 0 {
		t.Fatal("Expected at least one result")
	}

	bestSingle := GetBestMatch("kittens", candidates, weights, 0.5, NormalizedSimilarity)
	if bestSingle == nil {
		t.Fatal("Expected GetBestMatch to return a result")
	}

	if math.Abs(results[0].Similarity-bestSingle.Similarity) > f64Epsilon {
		t.Errorf("First sorted result (%f) should match GetBestMatch result (%f)",
			results[0].Similarity, bestSingle.Similarity)
	}
}

func TestMatchUnicode(t *testing.T) {
	weights := IndelWeights()

	if !Match("café", "café latte", weights, 0.5, NormalizedSimilarity) {
		t.Error("Expected Match to return true for 'café' and 'café latte'")
	}

	if !Match("日本", "日本語", weights, 0.5, NormalizedSimilarity) {
		t.Error("Expected Match to return true for '日本' and '日本語'")
	}
}

func TestGetBestMatchUnicode(t *testing.T) {
	weights := IndelWeights()

	candidates := []string{"café", "caffe", "coffee"}
	result := GetBestMatch("cafe", candidates, weights, 0.5, NormalizedSimilarity)

	if result == nil {
		t.Fatal("Expected GetBestMatch to return a result for unicode candidates")
	}
}

func ExampleMatch() {
	weights := IndelWeights()

	fmt.Println(Match("a test", "this is a test", weights, 0.5, NormalizedSimilarity))
	fmt.Println(Match("a test", "this is a test", weights, 0.9, NormalizedSimilarity))

	// Output:
	// true
	// false
}

func ExampleGetBestMatch() {
	weights := IndelWeights()

	candidates := []string{"kitten", "sitting", "bitten", "written"}
	result := GetBestMatch("kittens", candidates, weights, 0.5, NormalizedSimilarity)

	if result != nil {
		fmt.Printf("Best Match: %s (Similarity: %.8f)\n", result.Candidate, result.Similarity)
	}

	noMatch := GetBestMatch("xyz", candidates, weights, 0.99, NormalizedSimilarity)
	fmt.Printf("No Match: %v\n", noMatch)

	// Output:
	// Best Match: kitten (Similarity: 0.92307692)
	// No Match: <nil>
}

func ExampleGetBestMatches() {
	weights := IndelWeights()

	candidates := []string{"kitten", "sitting", "bitten", "written"}
	results := GetBestMatches("kittens", candidates, weights, 0.5, NormalizedSimilarity)

	fmt.Printf("Matches: %d\n", len(results))
	for _, r := range results {
		fmt.Printf("  %s: %.8f\n", r.Candidate, r.Similarity)
	}

	// Output:
	// Matches: 4
	//   kitten: 0.92307692
	//   sitting: 0.57142857
	//   bitten: 0.76923077
	//   written: 0.71428571
}

func ExampleGetBestMatchesSorted() {
	weights := IndelWeights()

	candidates := []string{"kitten", "sitting", "sitting", "bitten", "written"}
	results := GetBestMatchesSorted("kittens", candidates, weights, 0.5, NormalizedSimilarity)

	fmt.Printf("Sorted Matches: %d\n", len(results))
	for _, r := range results {
		fmt.Printf("  %s: %.8f\n", r.Candidate, r.Similarity)
	}

	// Output:
	// Sorted Matches: 5
	//   kitten: 0.92307692
	//   bitten: 0.76923077
	//   written: 0.71428571
	//   sitting: 0.57142857
	//   sitting: 0.57142857
}

func BenchmarkMatch(b *testing.B) {
	weights := IndelWeights()
	s1 := "a test"
	s2 := "this is a test"

	for b.Loop() {
		Match(s1, s2, weights, 0.5, NormalizedSimilarity)
	}
}

func BenchmarkGetBestMatch(b *testing.B) {
	weights := IndelWeights()
	candidates := []string{"kitten", "sitting", "bitten", "written"}

	for b.Loop() {
		GetBestMatch("kittens", candidates, weights, 0.5, NormalizedSimilarity)
	}
}

func BenchmarkGetBestMatches(b *testing.B) {
	weights := IndelWeights()
	candidates := []string{"kitten", "sitting", "bitten", "written"}

	for b.Loop() {
		GetBestMatches("kittens", candidates, weights, 0.5, NormalizedSimilarity)
	}
}

func BenchmarkGetBestMatchesSorted(b *testing.B) {
	weights := IndelWeights()
	candidates := []string{"kitten", "sitting", "bitten", "written"}

	for b.Loop() {
		GetBestMatchesSorted("kittens", candidates, weights, 0.5, NormalizedSimilarity)
	}
}
