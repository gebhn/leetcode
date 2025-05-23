package medium

import "math"

func reverse(x int) int {
	var result int
	var negative bool
	if x < 0 {
		negative = true
		x = -x
	}

	for x > 0 {
		digit := x % 10
		// Edge cases for EXACTLY MaxInt32
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if negative && (result > (math.MaxInt32+1)/10 || (result == (math.MaxInt32+1)/10 && digit > 8)) {
			return 0
		}
		result = result*10 + digit
		x /= 10
	}

	if negative {
		result *= -1
	}

	return result
}
