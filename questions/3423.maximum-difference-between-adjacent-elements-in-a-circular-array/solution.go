/*
 * @lc app=leetcode id=3423 lang=golang
 *
 * [3423] Maximum Difference Between Adjacent Elements in a Circular Array
 */

// @lc code=start
func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	gap := abs(nums[0] - nums[n-1])

	for i := 1; i < n; i++ {
		d := abs(nums[i] - nums[i-1])
		if d > gap {
			gap = d
		}
	}

	return gap
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

