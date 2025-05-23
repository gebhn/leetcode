package easy

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, value := range nums {
		compliment := target - value
		if index, ok := seen[compliment]; ok {
			return []int{index, i}
		}
		seen[value] = i
	}
	return nil
}
