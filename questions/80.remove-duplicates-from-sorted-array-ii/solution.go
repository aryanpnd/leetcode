/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	k, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if k < 2 {
			k++
			continue
		}
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// @lc code=end

