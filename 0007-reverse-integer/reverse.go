package reverse

import "math"

// reverse returns x with its digits reversed. If the reversed value falls
// outside the signed 32-bit integer range, it returns 0.
//
// https://leetcode.com/problems/reverse-integer/
func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}

	return rev
}
