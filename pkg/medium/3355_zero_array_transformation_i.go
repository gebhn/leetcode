package medium

func isZeroArray_Naive(nums []int, queries [][]int) bool {
	totals := make(map[int]int)
	for _, query := range queries {
		for i := query[0]; i <= query[1]; i++ {
			totals[i]++
		}
	}
	for i := range nums {
		if sub, ok := totals[i]; ok {
			nums[i] = nums[i] - sub
		}
		if nums[i] > 0 {
			return false
		}
	}
	return true
}

func isZeroArray(nums []int, queries [][]int) bool {
	deltas := make([]int, len(nums)+1)
	operations := make([]int, len(deltas))
	current := 0

	for _, query := range queries {
		deltas[query[0]] += 1
		deltas[query[1]+1] -= 1
	}
	for i, d := range deltas {
		current += d
		operations[i] = current
	}
	for i := range nums {
		if operations[i] < nums[i] {
			return false
		}
	}
	return true
}
