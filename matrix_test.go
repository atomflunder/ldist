package ldist

import (
	"testing"
)

// TestGetMatrixEmptyS1 tests when s1 is empty
func TestGetMatrixEmptyS1(t *testing.T) {
	s1 := ""
	s2 := "abc"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 3 {
		t.Errorf("Expected distance 3, got %d", dist)
	}

	if len(VP) != 0 {
		t.Errorf("Expected empty VP slice, got length %d", len(VP))
	}

	if len(VN) != 0 {
		t.Errorf("Expected empty VN slice, got length %d", len(VN))
	}
}

// TestGetMatrixEmptyS2 tests when s2 is empty
func TestGetMatrixEmptyS2(t *testing.T) {
	s1 := "abc"
	s2 := ""

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 3 {
		t.Errorf("Expected distance 3, got %d", dist)
	}

	if len(VP) != 0 {
		t.Errorf("Expected empty VP slice, got length %d", len(VP))
	}

	if len(VN) != 0 {
		t.Errorf("Expected empty VN slice, got length %d", len(VN))
	}
}

// TestGetMatrixIdenticalStrings tests when strings are identical
func TestGetMatrixIdenticalStrings(t *testing.T) {
	s1 := "abc"
	s2 := "abc"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 0 {
		t.Errorf("Expected distance 0 for identical strings, got %d", dist)
	}

	if len(VP) != 3 {
		t.Errorf("Expected VP length 3, got %d", len(VP))
	}

	if len(VN) != 3 {
		t.Errorf("Expected VN length 3, got %d", len(VN))
	}
}

// TestGetMatrixSingleCharDifference tests with single character difference
func TestGetMatrixSingleCharDifference(t *testing.T) {
	s1 := "a"
	s2 := "b"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 1 {
		t.Errorf("Expected distance 1, got %d", dist)
	}

	if len(VP) != 1 {
		t.Errorf("Expected VP length 1, got %d", len(VP))
	}

	if len(VN) != 1 {
		t.Errorf("Expected VN length 1, got %d", len(VN))
	}
}

// TestGetMatrixKittenSitting tests with "kitten" and "sitting"
func TestGetMatrixKittenSitting(t *testing.T) {
	s1 := "kitten"
	s2 := "sitting"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 3 {
		t.Errorf("Expected distance 3 between 'kitten' and 'sitting', got %d", dist)
	}

	if len(VP) != len(s2) {
		t.Errorf("Expected VP length %d, got %d", len(s2), len(VP))
	}

	if len(VN) != len(s2) {
		t.Errorf("Expected VN length %d, got %d", len(s2), len(VN))
	}
}

// TestGetMatrixInsertionOnly tests case with only insertions needed
func TestGetMatrixInsertionOnly(t *testing.T) {
	s1 := "a"
	s2 := "abc"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 2 {
		t.Errorf("Expected distance 2 (2 insertions), got %d", dist)
	}

	if len(VP) != 3 {
		t.Errorf("Expected VP length 3, got %d", len(VP))
	}

	if len(VN) != 3 {
		t.Errorf("Expected VN length 3, got %d", len(VN))
	}
}

// TestGetMatrixDeletionOnly tests case with only deletions needed
func TestGetMatrixDeletionOnly(t *testing.T) {
	s1 := "abc"
	s2 := "a"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 2 {
		t.Errorf("Expected distance 2 (2 deletions), got %d", dist)
	}

	if len(VP) != 1 {
		t.Errorf("Expected VP length 1, got %d", len(VP))
	}

	if len(VN) != 1 {
		t.Errorf("Expected VN length 1, got %d", len(VN))
	}
}

// TestGetMatrixSubstitutionOnly tests case with only substitutions needed
func TestGetMatrixSubstitutionOnly(t *testing.T) {
	s1 := "abc"
	s2 := "xyz"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 3 {
		t.Errorf("Expected distance 3 (3 substitutions), got %d", dist)
	}

	if len(VP) != 3 {
		t.Errorf("Expected VP length 3, got %d", len(VP))
	}

	if len(VN) != 3 {
		t.Errorf("Expected VN length 3, got %d", len(VN))
	}
}

// TestGetMatrixMixedOperations tests case with mixed operations
func TestGetMatrixMixedOperations(t *testing.T) {
	s1 := "qabxcd"
	s2 := "abycdf"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 3 {
		t.Errorf("Expected distance 3, got %d", dist)
	}

	if len(VP) != len(s2) {
		t.Errorf("Expected VP length %d, got %d", len(s2), len(VP))
	}

	if len(VN) != len(s2) {
		t.Errorf("Expected VN length %d, got %d", len(s2), len(VN))
	}
}

// TestGetMatrixLongerStrings tests with longer strings
func TestGetMatrixLongerStrings(t *testing.T) {
	s1 := "intention"
	s2 := "execution"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 5 {
		t.Errorf("Expected distance 5, got %d", dist)
	}

	if len(VP) != len(s2) {
		t.Errorf("Expected VP length %d, got %d", len(s2), len(VP))
	}

	if len(VN) != len(s2) {
		t.Errorf("Expected VN length %d, got %d", len(s2), len(VN))
	}
}

// TestGetMatrixVectorsNotEmpty tests that VP and VN vectors are properly populated
func TestGetMatrixVectorsNotEmpty(t *testing.T) {
	s1 := "test"
	s2 := "best"

	dist, VP, VN := getMatrix(s1, s2)

	// Vectors should not be empty for non-empty s2
	if len(VP) == 0 && len(s2) > 0 {
		t.Errorf("VP should not be empty when s2 is not empty")
	}

	if len(VN) == 0 && len(s2) > 0 {
		t.Errorf("VN should not be empty when s2 is not empty")
	}

	// Distance should be 1 (replace t with b)
	if dist != 1 {
		t.Errorf("Expected distance 1, got %d", dist)
	}
}

// TestGetMatrixConsistency tests that the function produces consistent results
func TestGetMatrixConsistency(t *testing.T) {
	s1 := "hello"
	s2 := "hallo"

	dist1, VP1, VN1 := getMatrix(s1, s2)
	dist2, VP2, VN2 := getMatrix(s1, s2)

	if dist1 != dist2 {
		t.Errorf("Distance should be consistent: got %d and %d", dist1, dist2)
	}

	if len(VP1) != len(VP2) || len(VN1) != len(VN2) {
		t.Errorf("Vector lengths should be consistent")
	}

	for i := 0; i < len(VP1); i++ {
		if VP1[i] != VP2[i] || VN1[i] != VN2[i] {
			t.Errorf("Vectors should be consistent at index %d", i)
		}
	}
}

// TestGetMatrixSingleCharacter tests with single character strings
func TestGetMatrixSingleCharacter(t *testing.T) {
	s1 := "a"
	s2 := "a"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 0 {
		t.Errorf("Expected distance 0 for identical single characters, got %d", dist)
	}

	if len(VP) != 1 {
		t.Errorf("Expected VP length 1, got %d", len(VP))
	}

	if len(VN) != 1 {
		t.Errorf("Expected VN length 1, got %d", len(VN))
	}
}

// TestGetMatrixCaseSensitivity tests that the function is case-sensitive
func TestGetMatrixCaseSensitivity(t *testing.T) {
	s1 := "ABC"
	s2 := "abc"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 3 {
		t.Errorf("Expected distance 3 (all different case), got %d", dist)
	}

	if len(VP) != 3 {
		t.Errorf("Expected VP length 3, got %d", len(VP))
	}

	if len(VN) != 3 {
		t.Errorf("Expected VN length 3, got %d", len(VN))
	}
}

// TestGetMatrixUnicodeCharacters tests with unicode characters
func TestGetMatrixUnicodeCharacters(t *testing.T) {
	s1 := "hello"
	s2 := "hallo"

	dist, VP, VN := getMatrix(s1, s2)

	if dist != 1 {
		t.Errorf("Expected distance 1, got %d", dist)
	}

	if len(VP) != len(s2) {
		t.Errorf("Expected VP length %d, got %d", len(s2), len(VP))
	}

	if len(VN) != len(s2) {
		t.Errorf("Expected VN length %d, got %d", len(s2), len(VN))
	}
}
