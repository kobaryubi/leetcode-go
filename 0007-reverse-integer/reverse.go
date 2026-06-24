package reverse

import (
	"fmt"
	"math"
	"slices"
)

func countDigits(x int) (count int) {
	for x > 0 {
		x = x / 10
		count++
	}

	return
}

// split integer into slice of single digits
func splitInt(n int) []int {
	slc := []int{}
	for n != 0 {
		r := n % 10
		if r < 0 {
			r = -r
		}
		slc = append(slc, r)
		n = n / 10
	}

	for i, j := 0, len(slc)-1; i<j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}

	return slc
}

func removeTrailingZeros(x int) int {
	if x == 0 {
		return 0
	}
	for x%10 == 0 {
		x /= 10
	}
	return x
}

func sliceToInt(s []int) int {
	res := 0	
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

// reverse returns x with its digits reversed. If the reversed value falls
// outside the signed 32-bit integer range, it returns 0.
//
// https://leetcode.com/problems/reverse-integer/
func reverse(x int) int {
	fmt.Println("=== reverse ===")
	fmt.Println("x:", x)

	// remove 0 from x
	px := removeTrailingZeros(x)

	// MaxInt32 digit count
	maxIntDigitCount := countDigits(math.MaxInt32)

	// MinInt32 digit count
	minIntDigitCount := countDigits(-math.MinInt32)

	// x digit count
	var xDigitCount int
	if (px >= 0) {
		xDigitCount = countDigits(px)
	} else {
		xDigitCount = countDigits(-px)
	}

	fmt.Println("px:", px)

	// check x is larger than 0 or not
	if px >= 0 {
		reversedXDigits := splitInt(px)
		slices.Reverse(reversedXDigits)
		fmt.Println("reversedXDigits:", reversedXDigits)

		// check length is the same with max int
		if xDigitCount == maxIntDigitCount {

			maxIntDigits := splitInt(math.MaxInt32)

			for i := 0; i < len(reversedXDigits); i++ {
				xd := reversedXDigits[i]
				md := maxIntDigits[i]
				if xd == md {
					continue
				}
				if xd < md {
					break
				}
				if xd > md {
					return 0
				}
			}
		}

		return sliceToInt(reversedXDigits)
	} else {
		reversedXDigits := splitInt(px)
		slices.Reverse(reversedXDigits)
		fmt.Println("reversedXDigits:", reversedXDigits)

		// check length is the same with min int
		if (xDigitCount == minIntDigitCount) {
			minIntDigits := splitInt(math.MinInt32)
			fmt.Println("minIntDigits:", minIntDigits)

			for i := 0; i < len(reversedXDigits); i++ {
				xd := reversedXDigits[i]
				md := minIntDigits[i]
				if xd == md {
					continue
				}
				if xd < md {
					break
				}
				if xd > md {
					return 0
				}
			}
		}

		return -1 * sliceToInt(reversedXDigits)
	}
}
