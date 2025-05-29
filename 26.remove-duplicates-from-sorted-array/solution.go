/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[count] != nums[i] {
			count++
			nums[count] = nums[i]
		}
	}
	return count + 1
}

// @lc code=end

