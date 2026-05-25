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

	cycle := numRows + numRows - 2

	batchCount := (len(s) + cycle - 1) / cycle

	zigzag := make([][]string, numRows)
	for i := range zigzag {
		zigzag[i] = make([]string, batchCount*numRows)
	}

	done := false
	for i := 0; i < batchCount; i++ {
		for j := 0; j < numRows; j++ {
			if cycle*i+j == len(s) {
				done = true
				break
			}
			zigzag[j][i*(numRows-1)] = string(s[cycle*i+j])
		}
		if done {
			break
		}
		for j := 0; j < numRows-2; j++ {
			if cycle*i+numRows+j == len(s) {
				break
			}
			zigzag[numRows-j-2][i*(numRows-1)+j+1] = string(s[cycle*i+numRows+j])
		}
	}

	converted := ""
	for i := 0; i < numRows; i++ {
		converted += strings.Join(zigzag[i], "")
	}

	return converted
}
