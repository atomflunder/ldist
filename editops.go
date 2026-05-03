package ldist

type editopTag string

const (
	tagReplace = "replace"
	tagInsert  = "insert"
	tagDelete  = "delete"
)

// editop represents a single edit operation (replace, insert, or delete) needed to transform one string into another,
// along with the positions in the source and destination strings where the operation occurs.
type editop struct {
	tag     editopTag
	srcPos  int
	destPos int
}

// getEditops returns a slice of editop representing the edit operations needed to transform s1 into s2.
// Uses the bit-parallel alignment algorithm by Heikki Hyyrö for efficient computation.
func getEditops(s1, s2 string) []editop {
	prefixLen, suffixLen := commonAffixes(s1, s2)

	s1Trimmed := s1[prefixLen : len(s1)-suffixLen]
	s2Trimmed := s2[prefixLen : len(s2)-suffixLen]

	dist, VP, VN := getMatrix(s1Trimmed, s2Trimmed)

	editops := []editop{}

	if dist == 0 {
		return editops
	}

	// Backtrace to find the edit operations
	editopList := make([]editop, dist)
	col := len(s1Trimmed)
	row := len(s2Trimmed)
	d := dist

	for row != 0 && col != 0 {
		// deletion
		if row > 0 && col > 0 && (VP[row-1]&(1<<uint(col-1)) != 0) {
			d--
			col--
			editopList[d] = editop{tag: tagDelete, srcPos: col + prefixLen, destPos: row + prefixLen}
		} else {
			row--

			// insertion
			if row > 0 && col > 0 && (VN[row-1]&(1<<uint(col-1)) != 0) {
				d--
				editopList[d] = editop{tag: tagInsert, srcPos: col + prefixLen, destPos: row + prefixLen}
			} else {
				col--

				// replace (Matches are not recorded)
				if col >= 0 && row >= 0 && s1Trimmed[col] != s2Trimmed[row] {
					d--
					editopList[d] = editop{tag: tagReplace, srcPos: col + prefixLen, destPos: row + prefixLen}
				}
			}
		}
	}

	// Handle remaining deletions
	for col != 0 {
		d--
		col--
		editopList[d] = editop{tag: tagDelete, srcPos: col + prefixLen, destPos: row + prefixLen}
	}

	// Handle remaining insertions
	for row != 0 {
		d--
		row--
		editopList[d] = editop{tag: tagInsert, srcPos: col + prefixLen, destPos: row + prefixLen}
	}

	return editopList
}
