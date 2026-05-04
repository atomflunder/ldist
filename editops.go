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
// Uses the bit-parallel alignment algorithm by Heikki Hyyrö for efficient computation,
// falling back to standard DP for strings longer than 64 bytes after affix trimming.
func getEditops(s1, s2 string) []editop {
	prefixLen, suffixLen := commonAffixes(s1, s2)

	s1Trimmed := s1[prefixLen : len(s1)-suffixLen]
	s2Trimmed := s2[prefixLen : len(s2)-suffixLen]

	if len(s1Trimmed) > 64 {
		return getEditopsDP(s1Trimmed, s2Trimmed, prefixLen)
	}

	return getEditopsBitParallel(s1Trimmed, s2Trimmed, prefixLen)
}

// getEditopsBitParallel uses the Hyyrö bit-parallel algorithm. Only valid when len(s1) <= 64.
func getEditopsBitParallel(s1Trimmed, s2Trimmed string, prefixLen int) []editop {
	dist, VP, VN := getMatrix(s1Trimmed, s2Trimmed)

	if dist == 0 {
		return []editop{}
	}

	editopList := make([]editop, dist)
	col := len(s1Trimmed)
	row := len(s2Trimmed)
	d := dist

	for row != 0 && col != 0 {
		if row > 0 && col > 0 && (VP[row-1]&(1<<uint(col-1)) != 0) {
			d--
			col--
			editopList[d] = editop{tag: tagDelete, srcPos: col + prefixLen, destPos: row + prefixLen}
		} else {
			row--

			if row > 0 && col > 0 && (VN[row-1]&(1<<uint(col-1)) != 0) {
				d--
				editopList[d] = editop{tag: tagInsert, srcPos: col + prefixLen, destPos: row + prefixLen}
			} else {
				col--

				if col >= 0 && row >= 0 && s1Trimmed[col] != s2Trimmed[row] {
					d--
					editopList[d] = editop{tag: tagReplace, srcPos: col + prefixLen, destPos: row + prefixLen}
				}
			}
		}
	}

	for col != 0 {
		d--
		col--
		editopList[d] = editop{tag: tagDelete, srcPos: col + prefixLen, destPos: row + prefixLen}
	}

	for row != 0 {
		d--
		row--
		editopList[d] = editop{tag: tagInsert, srcPos: col + prefixLen, destPos: row + prefixLen}
	}

	return editopList
}

// getEditopsDP uses a standard Wagner-Fischer DP matrix with backtrace.
// Handles strings of any length, used as a fallback when s1 exceeds 64 bytes.
func getEditopsDP(s1, s2 string, prefixLen int) []editop {
	n := len(s1)
	m := len(s2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				sub := dp[i-1][j-1] + 1
				del := dp[i-1][j] + 1
				ins := dp[i][j-1] + 1
				v := min(ins, min(del, sub))
				dp[i][j] = v
			}
		}
	}

	dist := dp[n][m]
	if dist == 0 {
		return []editop{}
	}

	editopList := make([]editop, dist)
	i, j, d := n, m, dist

	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			i--
			j--
		} else if dp[i-1][j-1] <= dp[i-1][j] && dp[i-1][j-1] <= dp[i][j-1] {
			i--
			j--
			d--
			editopList[d] = editop{tag: tagReplace, srcPos: i + prefixLen, destPos: j + prefixLen}
		} else if dp[i-1][j] <= dp[i][j-1] {
			i--
			d--
			editopList[d] = editop{tag: tagDelete, srcPos: i + prefixLen, destPos: j + prefixLen}
		} else {
			j--
			d--
			editopList[d] = editop{tag: tagInsert, srcPos: i + prefixLen, destPos: j + prefixLen}
		}
	}

	for i > 0 {
		i--
		d--
		editopList[d] = editop{tag: tagDelete, srcPos: i + prefixLen, destPos: j + prefixLen}
	}

	for j > 0 {
		j--
		d--
		editopList[d] = editop{tag: tagInsert, srcPos: i + prefixLen, destPos: j + prefixLen}
	}

	return editopList
}
