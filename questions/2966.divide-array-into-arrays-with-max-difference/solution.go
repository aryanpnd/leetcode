/*
 * @lc app=leetcode id=2966 lang=golang
 *
 * [2966] Divide Array Into Arrays With Max Difference
 */

// @lc code=start
func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	n := len(nums)
	for i := 0; i < n; i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}

		ans = append(ans, []int{nums[i], nums[i+1], nums[i+2]})
	}
	return ans
}

// @lc code=end

