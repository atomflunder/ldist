package ldist

// matchingBlock represents a contiguous block of matching characters between s1 and s2,
// defined by the starting positions in both strings and the length of the block.
type matchingBlock struct {
	srcPos  int
	destPos int
	length  int
}

// getMatchingBlocks returns a slice of MatchingBlock representing the matching blocks between r1 and r2.
func getMatchingBlocks(r1, r2 []rune) []matchingBlock {
	editops := getEditops(r1, r2)

	matchingBlocks := []matchingBlock{}

	srcPos, destPos := 0, 0
	for _, op := range editops {
		if srcPos < op.srcPos || destPos < op.destPos {
			l := min(op.srcPos-srcPos, op.destPos-destPos)
			if l > 0 {
				matchingBlocks = append(matchingBlocks, matchingBlock{srcPos: srcPos, destPos: destPos, length: l})
			}
			srcPos = op.srcPos
			destPos = op.destPos
		}

		switch op.tag {
		case tagReplace:
			srcPos++
			destPos++
		case tagDelete:
			srcPos++
		case tagInsert:
			destPos++
		}
	}

	if srcPos < len(r1) || destPos < len(r2) {
		l := min(len(r1)-srcPos, len(r2)-destPos)
		if l > 0 {
			matchingBlocks = append(matchingBlocks, matchingBlock{srcPos: srcPos, destPos: destPos, length: l})
		}
	}

	matchingBlocks = append(matchingBlocks, matchingBlock{srcPos: len(r1), destPos: len(r2), length: 0})

	return matchingBlocks
}
