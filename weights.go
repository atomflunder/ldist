package ldist

// Weights defines the weights for substitution, insertion, and deletion operations.
type Weights struct {
	// Substitution is the cost of substituting one character for another.
	// By default it should be set to 1.
	Substitution int
	// Insertion is the cost of inserting a character into a string.
	// By default it should be set to 1.
	Insertion int
	// Deletion is the cost of deleting a character from a string.
	// By default it should be set to 1.
	Deletion int
}

// GetWeights returns the default distance weights for substitution, insertion, and deletion.
// Substitution: 1, Insertion: 1, Deletion: 1
func GetWeights() Weights {
	return Weights{
		Substitution: 1,
		Insertion:    1,
		Deletion:     1,
	}
}

// GetIndelWeights returns the distance weights for substitution, insertion, and deletion where substitutions are more expensive than insertions and deletions.
// Substitution: 2, Insertion: 1, Deletion: 1
func GetIndelWeights() Weights {
	return Weights{
		Substitution: 2,
		Insertion:    1,
		Deletion:     1,
	}
}
