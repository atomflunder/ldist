package ldist

import (
	"fmt"
	"math"
	"testing"
)

const f64Epsilon = 1e-9

func TestBasicDistance(t *testing.T) {
	weights := DefaultWeights()

	s1 := "kitten"
	s2 := "sitting"

	expected := 3
	actual := Distance(s1, s2, weights)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}

	s1 = "eh"
	s2 = "a really, really, really different and long string to compare it to"

	expected = 66
	actual = Distance(s1, s2, weights)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}
}

func TestEdgeCases(t *testing.T) {
	weights := DefaultWeights()

	s1 := ""
	s2 := "abc"

	expected := 3
	actual := Distance(s1, s2, weights)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}

	s1 = "abc"
	s2 = ""

	expected = 3
	actual = Distance(s1, s2, weights)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}

	s1 = "abc"
	s2 = "abc"

	expected = 0
	actual = Distance(s1, s2, weights)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}
}

func TestDistanceWithCustomWeights(t *testing.T) {
	weights := Weights{
		Substitution: 35,
		Insertion:    1,
		Deletion:     35,
	}

	s1 := "kitten"
	s2 := "sitting"

	expected := 71
	actual := Distance(s1, s2, weights)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}
}

func TestNormalizedDistance(t *testing.T) {
	weights := DefaultWeights()

	s1 := "kitten"
	s2 := "sitting"

	expected := 3.0 / 13.0
	actual := NormalizedDistance(s1, s2, weights)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected normalized distance between %s and %s to be %f, but got %.8f", s1, s2, expected, actual)
	}
}

func TestNormalizedSimilarity(t *testing.T) {
	weights := DefaultWeights()

	s1 := "kitten"
	s2 := "sitting"

	expected := 1.0 - (3.0 / 13.0)
	actual := NormalizedSimilarity(s1, s2, weights)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected normalized similarity between %s and %s to be %f, but got %f", s1, s2, expected, actual)
	}
}

func TestWithOptions(t *testing.T) {
	weights := DefaultWeights()

	s1 := "kitten"
	s2 := "SITTING  "

	expected := 3
	actual := Distance(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}
}

func TestCommonSuffix(t *testing.T) {
	weights := DefaultWeights()

	s1 := "kitten"
	s2 := "mitten"

	expected := 1
	actual := Distance(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}
}

func TestSameString(t *testing.T) {
	weights := DefaultWeights()

	s1 := "kitten"
	s2 := "kitten"

	expected := 0
	actual := Distance(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}
}

func TestPartialSimilarity(t *testing.T) {
	weights := IndelWeights()

	s1 := "a test"
	s2 := "this is a test"

	expected := 0.85
	actual := PartialSimilarity(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected distance between %s and %s to be %f, but got %f", s1, s2, expected, actual)
	}

	s1 = "this"
	s2 = "this is a test"

	expected = 0.75
	actual = PartialSimilarity(s1, s2, weights)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected distance between %s and %s to be %.9f, but got %.9f", s1, s2, expected, actual)
	}

	s1 = "test124"
	s2 = "93210"

	expected = 0.158333333
	actual = PartialSimilarity(s1, s2, weights)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected distance between %s and %s to be %.9f, but got %.9f", s1, s2, expected, actual)
	}

	s1 = "ab"
	s2 = "dabuz"

	expected = 0.95
	actual = PartialSimilarity(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected distance between %s and %s to be %f, but got %f", s1, s2, expected, actual)
	}

	s1 = "d"
	s2 = "dabuz"

	expected = 0.85
	actual = PartialSimilarity(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected distance between %s and %s to be %f, but got %f", s1, s2, expected, actual)
	}

	s1 = "dabuz"
	s2 = "d"

	expected = 0.85
	actual = PartialSimilarity(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected distance between %s and %s to be %f, but got %f", s1, s2, expected, actual)
	}
}

func TestPartialEdgeCases(t *testing.T) {
	weights := IndelWeights()

	s1 := ""
	s2 := "dabuz"

	expected := 0.0
	actual := PartialSimilarity(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected distance between %s and %s to be %f, but got %f", s1, s2, expected, actual)
	}

	s1 = "dabuz"
	s2 = "dabuz"

	expected = 1.0
	actual = PartialSimilarity(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected distance between %s and %s to be %f, but got %f", s1, s2, expected, actual)
	}

	s1 = "a "
	s2 = "this is a really really damn long string, wow"

	expected = 0.65
	actual = PartialSimilarity(s1, s2, weights)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected distance between %s and %s to be %f, but got %f", s1, s2, expected, actual)
	}
}

func ExampleDistance() {
	weights := DefaultWeights()

	s1 := "kitten"
	s2 := "  SITTING  "

	fmt.Printf("Normal Distance: %d\n", Distance(s1, s2, weights))
	fmt.Printf("Distance With Options: %d\n", Distance(s1, s2, weights, ToLowercase, RemoveWhitespace))

	weightsCustom := Weights{
		Substitution: 35,
		Insertion:    1,
		Deletion:     35,
	}
	s2Clean := "sitting"
	fmt.Printf("Distance With Custom Weights: %d\n", Distance(s1, s2Clean, weightsCustom))

	// Output:
	// Normal Distance: 11
	// Distance With Options: 3
	// Distance With Custom Weights: 71
}

func ExampleNormalizedDistance() {
	weights := DefaultWeights()

	s1 := "kitten"
	s2 := "sitting"

	fmt.Printf("Normalized Distance: %.8f\n", NormalizedDistance(s1, s2, weights))

	// Output:
	// Normalized Distance: 0.23076923
}

func ExampleNormalizedSimilarity() {
	weights := DefaultWeights()

	s1 := "kitten"
	s2 := "sitting"

	fmt.Printf("Normalized Similarity: %.8f\n", NormalizedSimilarity(s1, s2, weights))

	// Output:
	// Normalized Similarity: 0.76923077
}

func ExamplePartialSimilarity() {
	weights := IndelWeights()

	s1 := "a test"
	s2 := "this is a test"

	fmt.Printf("Partial Similarity: %.8f\n", PartialSimilarity(s1, s2, weights))
	// Output:
	// Partial Similarity: 0.85000000
}

func BenchmarkDistance(b *testing.B) {
	weights := DefaultWeights()
	s1 := "kitten"
	s2 := "sitting"

	for b.Loop() {
		Distance(s1, s2, weights)
	}
}

func BenchmarkNormalizedDistance(b *testing.B) {
	weights := DefaultWeights()
	s1 := "kitten"
	s2 := "sitting"

	for b.Loop() {
		NormalizedDistance(s1, s2, weights)
	}
}

func BenchmarkNormalizedSimilarity(b *testing.B) {
	weights := DefaultWeights()
	s1 := "kitten"
	s2 := "sitting"

	for b.Loop() {
		NormalizedSimilarity(s1, s2, weights)
	}
}

func BenchmarkLongStrings(b *testing.B) {
	weights := DefaultWeights()
	s1 := `a really, really, really different and long string to compare it to
with multiple lines and some special characters!@#$%^&*()_+
and some more text to make it even longer and more different from the other string
and some more text to make it even longer and more different from the other string
==================================================================================
and some more text to make it even longer and more different from the other string`
	s2 := "eh"

	for b.Loop() {
		Distance(s1, s2, weights)
	}
}

func BenchmarkSimilarLongStrings(b *testing.B) {
	weights := DefaultWeights()
	s1 := `a really, really, really different and long string to compare it to
with multiple lines and some special characters!@#$%^&*()_+
and some more text to make it even longer and more different from the other string
and some more text to make it even longer and more different from the other string
==================================================================================
and some more text to make it even longer and more different from the other string`
	s2 := `a really, really, really different and long string to compare it to
with multiple lines and some special characters!@#$%^&*()_+
and some more text to make it even longer and more different from the other string
and this line is totally different and will increase the distance a lot, i hope so
==================================================================================
and some more text to make it even longer and more different from the other string`

	for b.Loop() {
		Distance(s1, s2, weights)
	}
}

func BenchmarkPartialSimilarity(b *testing.B) {
	weights := IndelWeights()
	s1 := "a test"
	s2 := "this is a test"

	for b.Loop() {
		PartialSimilarity(s1, s2, weights)
	}
}

func TestUnicodeDistance(t *testing.T) {
	weights := DefaultWeights()

	tests := []struct {
		s1, s2   string
		expected int
	}{
		{"é", "e", 1},
		{"café", "cafe", 1},
		{"日本語", "日本人", 1},
		{"über", "uber", 1},
		{"naïve", "naive", 1},
		{"日本", "中国", 2},
		{"", "日本語", 3},
		{"日本語", "", 3},
		{"日本語", "日本語", 0},
	}

	for _, tc := range tests {
		actual := Distance(tc.s1, tc.s2, weights)
		if actual != tc.expected {
			t.Errorf("Distance(%q, %q) = %d, want %d", tc.s1, tc.s2, actual, tc.expected)
		}
	}
}

func TestUnicodeNormalizedDistance(t *testing.T) {
	weights := DefaultWeights()

	expected := 1.0 / 8.0
	actual := NormalizedDistance("café", "cafe", weights)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("NormalizedDistance(\"café\", \"cafe\") = %f, want %f", actual, expected)
	}
}

func TestUnicodePartialSimilarity(t *testing.T) {
	weights := IndelWeights()

	actual := PartialSimilarity("日本", "私は日本人です", weights)
	if actual <= 0.0 {
		t.Errorf("PartialSimilarity(\"日本\", \"私は日本人です\") = %f, expected > 0", actual)
	}
}

func TestLongStringDistance(t *testing.T) {
	weights := DefaultWeights()

	s1 := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmn"
	s2 := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmo"

	expected := 1
	actual := Distance(s1, s2, weights)

	if actual != expected {
		t.Errorf("Distance on 66-char strings differing by 1 = %d, want %d", actual, expected)
	}
}

func TestLongStringPartialSimilarity(t *testing.T) {
	weights := IndelWeights()

	s1 := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmn"
	s2 := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmo"

	actual := PartialSimilarity(s1, s2, weights)
	if actual <= 0.0 {
		t.Errorf("PartialSimilarity on 66-char strings = %f, expected > 0", actual)
	}
}

func TestPartialSimilarityBoundsClamping(t *testing.T) {
	weights := IndelWeights()

	// "xxab" (longer) has "ab" matching at srcPos=2, destPos=0 with "abc" (shorter).
	// start=2, end=2+3=5 > len("xxab")=4, so the clamping branch is hit.
	actual := PartialSimilarity("abc", "xxab", weights)
	if actual < 0.0 || actual > 1.0 {
		t.Errorf("PartialSimilarity(\"abc\", \"xxab\") = %f, expected value in [0, 1]", actual)
	}
}

func TestNormalizedDistanceTabVsSpace(t *testing.T) {
	weights := DefaultWeights()

	s1 := "\t"
	s2 := " "

	expected := 0.5
	actual := NormalizedDistance(s1, s2, weights)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("NormalizedDistance(%q, %q) = %f, want %f", s1, s2, actual, expected)
	}

	actual = NormalizedDistance(s1, s2, weights, RemoveWhitespace)

	if actual != 0.0 {
		t.Errorf("NormalizedDistance(%q, %q, RemoveWhitespace) = %f, want 0.0", s1, s2, actual)
	}
}
