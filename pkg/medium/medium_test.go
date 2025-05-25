package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("2 Add Two Numbers", func(t *testing.T) {
		type testCase struct {
			l1 *ListNode
			l2 *ListNode
		}
		ex1 := &testCase{
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		}
		ex2 := &testCase{
			l1: &ListNode{Val: 0, Next: nil},
			l2: &ListNode{Val: 0, Next: nil},
		}
		assert.Equal(t, &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}, addTwoNumbers(ex1.l1, ex1.l2))
		assert.Equal(t, &ListNode{Val: 0}, addTwoNumbers(ex2.l1, ex2.l2))
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

	t.Run("7 Reverse Integer", func(t *testing.T) {
		ex1 := 123
		ex2 := -123
		ex3 := 120

		assert.Equal(t, 321, reverse(ex1))
		assert.Equal(t, -321, reverse(ex2))
		assert.Equal(t, 21, reverse(ex3))
	})

	t.Run("11 Container With Most Water", func(t *testing.T) {
		ex1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		ex2 := []int{1, 1}

		assert.Equal(t, 49, maxArea(ex1))
		assert.Equal(t, 1, maxArea(ex2))
	})

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
}
