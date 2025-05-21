package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("3355 Zero Array Transformation I", func(t *testing.T) {
		type testCase struct {
			nums    []int
			queries [][]int
		}
		ex1 := &testCase{
			nums:    []int{1, 0, 1},
			queries: [][]int{{0, 2}},
		}
		ex2 := &testCase{
			nums:    []int{4, 3, 2, 1},
			queries: [][]int{{1, 3}, {0, 2}},
		}

		assert.True(t, isZeroArray_Naive(ex1.nums, ex1.queries))
		assert.False(t, isZeroArray_Naive(ex2.nums, ex2.queries))

		assert.True(t, isZeroArray(ex1.nums, ex1.queries))
		assert.False(t, isZeroArray(ex2.nums, ex2.queries))
	})

	t.Run("3 Longest Substring Without Repeating Characters", func(t *testing.T) {
		ex1 := "abcabcbb"
		ex2 := "bbbbb"
		ex3 := "pwwkew"

		assert.Equal(t, 3, lengthOfLongestSubstring(ex1))
		assert.Equal(t, 1, lengthOfLongestSubstring(ex2))
		assert.Equal(t, 3, lengthOfLongestSubstring(ex3))
	})

	t.Run("5 Longest Palindromic Substring", func(t *testing.T) {
		ex1 := "babad"
		ex2 := "cbbd"

		assert.Equal(t, "bab", longestPalindrome(ex1))
		assert.Equal(t, "bb", longestPalindrome(ex2))
	})

	t.Run("6 ZigZag Conversion", func(t *testing.T) {
		type testCase struct {
			s string
			r int
		}
		ex1 := &testCase{
			s: "PAYPALISHIRING",
			r: 3,
		}
		ex2 := &testCase{
			s: "PAYPALISHIRING",
			r: 4,
		}
		assert.Equal(t, "PAHNAPLSIIGYIR", convertNaive(ex1.s, ex1.r))
		assert.Equal(t, "PINALSIGYAHRPI", convertNaive(ex2.s, ex2.r))
	})
}
