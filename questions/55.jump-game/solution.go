/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	checkpoint := 0

	for i := 0; i < len(nums); i++ {
		if i > checkpoint {
			return false
		}

		checkpoint = max(checkpoint, nums[i]+i)
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

