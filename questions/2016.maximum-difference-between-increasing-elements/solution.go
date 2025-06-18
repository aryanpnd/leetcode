/*
 * @lc app=leetcode id=2016 lang=golang
 *
 * [2016] Maximum Difference Between Increasing Elements
 */

// @lc code=start

// brute force
// func maximumDifference(nums []int) int {
// 	m := -1
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j <= n-1; j++ {
// 			if nums[i] < nums[j] {
// 				m = max(m, nums[j]-nums[i])
// 			}
// 		}
// 	}
// 	return m
// }

// func max(n1, n2 int) int {
// 	if n1 > n2 {
// 		return n1
// 	}
// 	return n2
// }

// @lc code=end

