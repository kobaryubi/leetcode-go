package zigzag

import (
	"strings"
)

// convert returns s read off in a zigzag pattern over numRows rows.
//
// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	batchCount := len(s) / (numRows + numRows - 2)
	batchRest := len(s) % (numRows + numRows - 2)
	if (batchRest > 0) {
		batchCount++
	}

	zigzag := make([][]string, numRows)
	for i := range zigzag {
		zigzag[i] = make([]string, batchCount * numRows)
	}

	done := false
	for i := 0; i < batchCount; i++ {
		for j := 0; j < numRows; j++ {
			if (numRows + numRows - 2) * i + j == len(s) {
				done = true
				break
			}
			zigzag[j][i * (numRows - 1)] = string(s[(numRows + numRows - 2) * i + j])
		}
		if (done) {
			break
		}
		for j := 0; j < numRows - 2; j++ {
			if (numRows + numRows - 2) * i + numRows + j == len(s) {
				break
			}
			zigzag[numRows - j - 2][i * (numRows - 1) + j + 1] = string(s[(numRows + numRows - 2) * i + numRows + j])
		}
	}

	converted := ""
	for i := 0; i < numRows; i++ {
		converted += strings.Join(zigzag[i], "")
	}

	return converted
}
