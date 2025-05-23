package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("1 Two Sum", func(t *testing.T) {
		type testCase struct {
			Nums   []int
			Target int
		}
		ex1 := &testCase{
			Nums:   []int{2, 7, 11, 15},
			Target: 9,
		}
		ex2 := &testCase{
			Nums:   []int{3, 2, 4},
			Target: 6,
		}
		ex3 := &testCase{
			Nums:   []int{3, 3},
			Target: 6,
		}

		assert.Equal(t, []int{0, 1}, twoSum(ex1.Nums, ex1.Target))
		assert.Equal(t, []int{1, 2}, twoSum(ex2.Nums, ex2.Target))
		assert.Equal(t, []int{0, 1}, twoSum(ex3.Nums, ex3.Target))
	})
}
