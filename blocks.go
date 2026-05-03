package ldist

// matchingBlock represents a contiguous block of matching characters between s1 and s2,
// defined by the starting positions in both strings and the length of the block.
type matchingBlock struct {
	srcPos  int
	destPos int
	length  int
}

// getMatchingBlocks returns a slice of MatchingBlock representing the matching blocks between s1 and s2.
func getMatchingBlocks(s1, s2 string) []matchingBlock {
	editops := getEditops(s1, s2)

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

	if srcPos < len(s1) || destPos < len(s2) {
		l := min(len(s1)-srcPos, len(s2)-destPos)
		if l > 0 {
			matchingBlocks = append(matchingBlocks, matchingBlock{srcPos: srcPos, destPos: destPos, length: l})
		}
	}

	matchingBlocks = append(matchingBlocks, matchingBlock{srcPos: len(s1), destPos: len(s2), length: 0})

	return matchingBlocks
}
