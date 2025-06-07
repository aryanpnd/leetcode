/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
// using map approach
// func majorityElement(nums []int) int {
// 	m := make(map[int]int)
// 	c := len(nums) / 2

// 	for _, num := range nums {
// 		m[num]++
// 		if m[num] > c {
// 			return num
// 		}
// 	}
// 	return -1
// }

// Boyer-Moore Voting Algorithm
func majorityElement(nums []int) int {
	count, current := 0, 0

	for _, num := range nums {
		if count == 0 {
			current = num
		}

		if current == num {
			count++
		} else {
			count--
		}
	}

	return current
}

// @lc code=end

