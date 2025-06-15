/*
 * @lc app=leetcode id=1432 lang=golang
 *
 * [1432] Max Difference You Can Get From Changing an Integer
 */

// @lc code=start
func maxDiff(num int) int {
	numStr := strconv.Itoa(num)

	var maxStr string
	found := false
	for _, ch := range numStr {
		if ch != '9' {
			maxStr = strings.ReplaceAll(numStr, string(ch), "9")
			found = true
			break
		}
	}
	if !found {
		maxStr = numStr
	}

	var minStr string
	if numStr[0] != '1' {
		minStr = strings.ReplaceAll(numStr, string(numStr[0]), "1")
	} else {
		found = false
		for i := 1; i < len(numStr); i++ {
			if numStr[i] != '0' && numStr[i] != '1' {
				minStr = strings.ReplaceAll(numStr, string(numStr[i]), "0")
				found = true
				break
			}
		}
		if !found {
			minStr = numStr
		}
	}

	maxNum, _ := strconv.Atoi(maxStr)
	minNum, _ := strconv.Atoi(minStr)
	return maxNum - minNum
}

// @lc code=end

