package medium

func maxArea(height []int) int {
	res := 0

	l := 0
	r := len(height) - 1

	for l < r {
		m := r - l
		h := min(height[l], height[r])
		a := m * h

		if int(a) > res {
			res = int(a)
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
