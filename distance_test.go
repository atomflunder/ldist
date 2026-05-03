package ldist

import (
	"fmt"
	"math"
	"testing"
)

const f64Epsilon = 1e-9

func TestBasicDistance(t *testing.T) {
	weights := GetWeights()

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
	weights := GetWeights()

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
	weights := GetWeights()

	s1 := "kitten"
	s2 := "sitting"

	expected := 3.0 / 13.0
	actual := NormalizedDistance(s1, s2, weights)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected normalized distance between %s and %s to be %f, but got %.8f", s1, s2, expected, actual)
	}
}

func TestNormalizedSimilarity(t *testing.T) {
	weights := GetWeights()

	s1 := "kitten"
	s2 := "sitting"

	expected := 1.0 - (3.0 / 13.0)
	actual := NormalizedSimilarity(s1, s2, weights)

	if math.Abs(expected-actual) > f64Epsilon {
		t.Errorf("Expected normalized similarity between %s and %s to be %f, but got %f", s1, s2, expected, actual)
	}
}

func TestWithOptions(t *testing.T) {
	weights := GetWeights()

	s1 := "kitten"
	s2 := "SITTING  "

	expected := 3
	actual := Distance(s1, s2, weights, ToLowercase, RemoveWhitespace)

	if actual != expected {
		t.Errorf("Expected distance between %s and %s to be %d, but got %d", s1, s2, expected, actual)
	}
}

func ExampleDistance() {
	weights := GetWeights()

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
	weights := GetWeights()

	s1 := "kitten"
	s2 := "sitting"

	fmt.Printf("Normalized Distance: %.8f\n", NormalizedDistance(s1, s2, weights))

	// Output:
	// Normalized Distance: 0.23076923
}

func ExampleNormalizedSimilarity() {
	weights := GetWeights()

	s1 := "kitten"
	s2 := "sitting"

	fmt.Printf("Normalized Similarity: %.8f\n", NormalizedSimilarity(s1, s2, weights))

	// Output:
	// Normalized Similarity: 0.76923077
}

func BenchmarkDistance(b *testing.B) {
	weights := GetWeights()
	s1 := "kitten"
	s2 := "sitting"

	for b.Loop() {
		Distance(s1, s2, weights)
	}
}

func BenchmarkNormalizedDistance(b *testing.B) {
	weights := GetWeights()
	s1 := "kitten"
	s2 := "sitting"

	for b.Loop() {
		NormalizedDistance(s1, s2, weights)
	}
}

func BenchmarkNormalizedSimilarity(b *testing.B) {
	weights := GetWeights()
	s1 := "kitten"
	s2 := "sitting"

	for b.Loop() {
		NormalizedSimilarity(s1, s2, weights)
	}
}

func BenchmarkLongStrings(b *testing.B) {
	weights := GetWeights()
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
	weights := GetWeights()
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
