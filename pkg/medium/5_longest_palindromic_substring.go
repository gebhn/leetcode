package medium

func longestPalindrome(s string) string {
	var longest string

	expand := func(s string, l int, r int) string {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return s[l+1 : r]
	}
	for i := range s {
		// The longest palindrome within a string can be any length, so you must
		// calculate both odd and even if you're expanding from the center of
		// the substring.
		odd := expand(s, i, i)
		even := expand(s, i, i+1)

		if len(odd) > len(longest) {
			longest = odd
		}
		if len(even) > len(longest) {
			longest = even
		}
	}
	return longest
}
