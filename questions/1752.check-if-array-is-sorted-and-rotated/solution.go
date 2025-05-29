/*
 * @lc app=leetcode id=1752 lang=golang
 *
 * [1752] Check if Array Is Sorted and Rotated
 */

// @lc code=start
func check(nums []int) bool {
	count, arrLen := 0, len(nums)
	for i := 0; i < arrLen; i++ {
		next := (i + 1) % arrLen
		if nums[i] > nums[next] {
			count++
		}
	}

	return count <= 1
}

// @lc code=end

