# ldist
[![Go Reference](https://pkg.go.dev/badge/github.com/atomflunder/ldist.svg)](https://pkg.go.dev/github.com/atomflunder/ldist) [![Test / CI](https://github.com/atomflunder/ldist/actions/workflows/ci.yml/badge.svg)](https://github.com/atomflunder/ldist/actions/workflows/ci.yml) [![codecov](https://codecov.io/gh/atomflunder/ldist/graph/badge.svg?token=ONLF7D1OBG)](https://codecov.io/gh/atomflunder/ldist)

A Go implementation of a Levenshtein distance based string matching algorithm. Designed to be easy, fast and flexible to use with no dependencies.

## Installation

```bash
go get github.com/atomflunder/ldist
```

## Usage

### Basic Usage

You can calculate the Levenshtein distance between two strings using the `Distance` function.  
By default, it uses the standard weights for substitution, insertion, and deletion.

```go
import (
    "fmt"
    "github.com/atomflunder/ldist"
)

s1 := "kitten"
s2 := "  SITTING  "

// Gets you the default weights for the Levenshtein algorithm.
// The default weights are: Substitution: 1, Insertion: 1, Deletion: 1.
w := ldist.DefaultWeights()

dist := ldist.Distance(s1, s2, w)
fmt.Printf("Normal Distance: %d\n", dist)
// Output: Normal Distance: 11
```

### Custom Weights

You can also specify custom weights for the Levenshtein algorithm by creating a `Weights` struct and passing it to the `Distance` function.  
This allows you to assign different costs to substitutions, insertions, and deletions.

```go
import (
    "fmt"
    "github.com/atomflunder/ldist"
)

s1 := "kitten"
s2 := "  SITTING  "

// Create custom weights for the Levenshtein algorithm.
w := ldist.Weights{
    Substitution: 3,
    Insertion: 1,
    Deletion: 2,
}

dist := ldist.Distance(s1, s2, w)
fmt.Printf("Custom Distance: %d\n", dist)
// Output: Custom Distance: 23
```

### With Options

You can also specify some options for pre-processing the strings before calculating the distance.  
For example, you can choose to ignore case and whitespace.

```go
import (
    "fmt"
    "github.com/atomflunder/ldist"
)

s1 := "kitten"
s2 := "  SITTING  "

w := ldist.DefaultWeights()

// Uses the ToLowercase and RemoveWhitespace options to pre-process the strings before calculating the distance.
// This would convert "  SITTING  " into "sitting".
dist := ldist.Distance(s1, s2, w, ldist.ToLowercase, ldist.RemoveWhitespace)
fmt.Printf("Distance with Options: %d\n", dist)
// Output: Distance with Options: 3
```

### Normalized Functions

The package also provides normalized versions of the distance and similarity functions, which return values between 0 and 1 based on the maximum possible distance for the given strings.

```go
import (
    "fmt"
    "github.com/atomflunder/ldist"
)

s1 := "kitten"
s2 := "  SITTING  "

w := ldist.DefaultWeights()

// A normalized distance of 0 means the strings are identical, while a normalized distance of 1 means they are completely different.
// The normalized distance is calculated as the actual distance divided by the maximum possible distance for the given strings.
normalizedDist := ldist.NormalizedDistance(s1, s2, w, ldist.ToLowercase, ldist.RemoveWhitespace)

// A normalized similarity of 1 means the strings are identical, while a normalized similarity of 0 means they are completely different.
// The normalized similarity is calculated as 1 - normalized distance.
normalizedSim := ldist.NormalizedSimilarity(s1, s2, w, ldist.ToLowercase, ldist.RemoveWhitespace)

fmt.Printf("Normalized Distance: %.2f\n", normalizedDist)
fmt.Printf("Normalized Similarity: %.2f\n", normalizedSim)
// Output:
// Normalized Distance: 0.23
// Normalized Similarity: 0.77
```

### Partial Similarity

You can also calculate the distance between two strings using partial matching, which finds the best matching substring in the longer string and calculates the distance based on that, with a penalty based on differing lengths.

```go
import (
    "fmt"
    "github.com/atomflunder/ldist"
)

s1 := "test"
s2 := "this is a test"

w := ldist.DefaultWeights()

// The normal similarity for comparision.
normalizedSim := ldist.NormalizedSimilarity(s1, s2, w)

// The partial similarity finds the best matching substring in the longer string and calculates the similarity based on that.
partialSim := ldist.PartialSimilarity(s1, s2, w)

fmt.Printf("Normalized Similarity: %.2f\n", normalizedSim)
fmt.Printf("Partial Similarity: %.2f\n", partialSim)
// Output:
// Normalized Similarity: 0.44
// Partial Similarity: 0.75
```

### Matches

There are also helper functions to check for matches based on a similarity threshold, and to find the best match from a list of candidates.

```go
import (
    "fmt"
    "github.com/atomflunder/ldist"
)

s1 := "test"
s2 := "this is a test"

w := ldist.DefaultWeights()

// You can check if two strings are a match based on a similarity threshold using the Match function.
// You can specify the similarity function to use (e.g. NormalizedSimilarity or PartialSimilarity) and the threshold for a match (e.g. 0.4).
isMatch := ldist.Match(s1, s2, w, 0.4, ldist.NormalizedSimilarity)

fmt.Printf("Is Match: %t\n", isMatch)
// Output:
// Is Match: true

// Or you can find the best match from a list of candidates using the BestMatch function, 
// which returns the candidate with the highest similarity above the given threshold.
bestMatch := ldist.GetBestMatch(s1, []string{"this", "is", "a", "test"}, w, 0.8, ldist.NormalizedSimilarity)

fmt.Printf("Best Match: %+v\n", bestMatch)
// Output:
// Best Match: &{Candidate:test Similarity:1}

// Or you can get all matches above a certain threshold.
// (There is also a GetBestMatchesSorted function to auto-sort the matches by similarity.)
matches := ldist.GetBestMatches(s1, []string{"this", "is", "a", "test", "and", "test2"}, w, 0.8, ldist.NormalizedSimilarity)
fmt.Printf("Matches: %+v\n", matches)
// Output:
// Matches: [{Candidate:test Similarity:1} {Candidate:test2 Similarity:0.8888888888888888}]
```

### Custom Options & Similarity Functions

Any Option is a function that takes in two string pointers to modify them in place. You can create your own custom options to pre-process the strings in any way you like before calculating the distance.

```go
import (
    "fmt"
    "github.com/atomflunder/ldist"
)

s1 := "kitten"
s2 := "  SITTING  "

w := ldist.DefaultWeights()

// Custom option to modify each string in place.
// In a real use case you would probably want to do something more useful.
myOption := func (s1, s2 *string) {
	*s1, *s2 = "Option 1", "Option 2"
}

dist := ldist.Distance(s1, s2, w, myOption)
fmt.Printf("Distance with Custom Option: %d\n", dist)
// Output: Distance with Custom Option: 1
```

Akin to the Options, you can also create your own custom similarity functions to use with the matching functions.

```go
import (
    "fmt"
    "github.com/atomflunder/ldist"
)

s1 := "test"
sl := []string{"this", "is", "a", "test"}

// Custom similarity function that returns 1 if the strings are exactly equal, and 0 otherwise.
mySimilarity := func(s1, s2 string, w ldist.Weights, options ...ldist.Option) float64 {
    if s1 == s2 {
        return 1
    }
    return 0
}

bestMatch := ldist.GetBestMatch(s1, sl, ldist.DefaultWeights(), 0.5, mySimilarity)
fmt.Printf("Best Match with Custom Similarity: %+v\n", bestMatch)
// Output:
// Best Match with Custom Similarity: &{Candidate:test Similarity:1}
```

## API Reference

Can be found in the [GoDoc](https://pkg.go.dev/github.com/atomflunder/ldist).

## Testing & Coverage

You can run the tests using the following command:

```bash
go test -v
```

Check the code coverage with:

```bash
go test -coverprofile -cover.out .
go tool cover -html cover.out -o cover.html
```

## Performance & Benchmarks

You can run the benchmarks with:

```bash
go test -bench=.
```

Results (on my machine):

```bash
BenchmarkDistance-32                	11110988	        98.05 ns/op
BenchmarkNormalizedDistance-32      	12607795	        96.33 ns/op
BenchmarkNormalizedSimilarity-32    	12237752	        98.92 ns/op
BenchmarkLongStrings-32             	  981541	         1227 ns/op
BenchmarkSimilarLongStrings-32      	  141498	         8510 ns/op
BenchmarkPartialSimilarity-32       	 4566796	        260.0 ns/op
```

## Contributing

Contributions are welcome! If you have any ideas for improvements or new features, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
