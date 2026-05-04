package ldist

// getMatrix computes the edit distance and returns the distance along with VP and VN vectors for each column.
// These vectors are used to backtrace and find the actual edit operations.
func getMatrix(r1, r2 []rune) (int, []uint64, []uint64) {
	// This implementation is taken from the RapidFuzz Python library:
	// https://github.com/rapidfuzz/RapidFuzz
	if len(r1) == 0 {
		return len(r2), []uint64{}, []uint64{}
	}

	VP := uint64((1 << uint(len(r1))) - 1)
	VN := uint64(0)
	currDist := len(r1)
	mask := uint64(1 << uint(len(r1)-1))

	block := make(map[rune]uint64)
	for i, ch := range r1 {
		block[ch] |= 1 << uint(i)
	}

	matrixVP := make([]uint64, len(r2))
	matrixVN := make([]uint64, len(r2))

	for i, ch2 := range r2 {
		// Step 1: Computing D0
		PM_j := block[ch2]
		X := PM_j
		D0 := (((X & VP) + VP) ^ VP) | X | VN

		// Step 2: Computing HP and HN
		HP := VN | ^(D0 | VP)
		HN := D0 & VP

		// Step 3: Computing the value D[m,j]
		if (HP & mask) != 0 {
			currDist++
		}
		if (HN & mask) != 0 {
			currDist--
		}

		// Step 4: Computing VP and VN
		HP = (HP << 1) | 1
		HN = HN << 1
		VP = HN | ^(D0 | HP)
		VN = HP & D0

		matrixVP[i] = VP
		matrixVN[i] = VN
	}

	return currDist, matrixVP, matrixVN
}
