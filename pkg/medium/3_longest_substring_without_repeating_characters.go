package medium

func lengthOfLongestSubstring(s string) int {
	var low, result int
	seen := map[rune]int{}

	for high, char := range s {
		if exists, ok := seen[char]; ok && exists >= low {
			low = exists + 1
		} else {
			if high-low+1 > result {
				result = high - low + 1
			}
		}
		seen[char] = high
	}

	return result
}
