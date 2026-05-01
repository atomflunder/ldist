# ldist
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
w := ldist.GetWeights()

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

w := ldist.GetWeights()

// Uses the ToLowercase and RemoveWhitespace options to pre-process the strings before calculating the distance.
// This would convert "  SITTING  " into "sitting".
dist := ldist.Distance(s1, s2, weights, ldist.ToLowercase, ldist.RemoveWhitespace)

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

w := ldist.GetWeights()

// A normalized distance of 0 means the strings are identical, while a normalized distance of 1 means they are completely different.
// The normalized distance is calculated as the actual distance divided by the maximum possible distance for the given strings.
normalizedDist := ldist.NormalizedDistance(s1, s2, w, ldist.ToLowercase, ldist.RemoveWhitespace)

// A normalized similarity of 1 means the strings are identical, while a normalized similarity of 0 means they are completely different.
// The normalized similarity is calculated as 1 - normalized distance.
normalizedSim := ldist.NormalizedSimilarity(s1, s2, w, ldist.ToLowercase, ldist.RemoveWhitespace)

fmt.Printf("Normalized Distance: %.2f\n", normalizedDist)
fmt.Printf("Normalized Similarity: %.2f\n", normalizedSim)
// Output:
// Normalized Distance: 0.18
// Normalized Similarity: 0.82
```

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
BenchmarkDistance-32                	11266522	        95.96 ns/op
BenchmarkNormalizedDistance-32      	12367748	        96.25 ns/op
BenchmarkNormalizedSimilarity-32    	12120586	        98.67 ns/op
BenchmarkLongStrings-32             	  988434	      1214 ns/op
```

## Contributing

Contributions are welcome! If you have any ideas for improvements or new features, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
