/*
 * @lc app=leetcode id=2566 lang=golang
 *
 * [2566] Maximum Difference by Remapping a Digit
 */

// @lc code=start
func minMaxDifference(num int) int {
	numStr := strconv.Itoa(num)

	var maxValTemp string
	for _, ch := range numStr {
		if ch != '9' {
			maxValTemp = strings.ReplaceAll(numStr, string(ch), "9")
			break
		}
	}

	if maxValTemp == "" {
		maxValTemp = numStr
	}

	maxVal, _ := strconv.Atoi(maxValTemp)

	toReplaceDigit := string(numStr[0])
	minValTemp := strings.ReplaceAll(numStr, toReplaceDigit, "0")
	minVal, _ := strconv.Atoi(minValTemp)

	return maxVal - minVal

}

// @lc code=end

