/*
 * @lc app=leetcode id=2873 lang=golang
 *
 * [2873] Maximum Value of an Ordered Triplet I
 */

// @lc code=start
func maximumTripletValue(nums []int) int64 {
	numsLen := len(nums)

	if numsLen <= 0 {
		return 0
	}

	maxVal := 0

	for i := 0; i < numsLen-2; i++ {
		for j := i + 1; j < numsLen-1; j++ {
			for k := j + 1; k < numsLen; k++ {
				currentVal := (nums[i] - nums[j]) * nums[k]
				if currentVal > maxVal {
					maxVal = currentVal
				}
			}
		}
	}

	return int64(maxVal)
}

// @lc code=end

