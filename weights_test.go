package ldist

import "testing"

func TestGetWeights(t *testing.T) {
	weights := GetWeights()

	if weights.Substitution != 1 {
		t.Errorf("Expected Substitution weight to be 1, but got %d", weights.Substitution)
	}
	if weights.Insertion != 1 {
		t.Errorf("Expected Insertion weight to be 1, but got %d", weights.Insertion)
	}
	if weights.Deletion != 1 {
		t.Errorf("Expected Deletion weight to be 1, but got %d", weights.Deletion)
	}
}

func TestGetIndelWeights(t *testing.T) {
	weights := GetIndelWeights()

	if weights.Substitution != 2 {
		t.Errorf("Expected Substitution weight to be 2, but got %d", weights.Substitution)
	}
	if weights.Insertion != 1 {
		t.Errorf("Expected Insertion weight to be 1, but got %d", weights.Insertion)
	}
	if weights.Deletion != 1 {
		t.Errorf("Expected Deletion weight to be 1, but got %d", weights.Deletion)
	}
}
