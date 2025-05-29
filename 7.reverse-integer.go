/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	reverseNum, negativeFlag := 0, 1

	if x < 0 {
		negativeFlag = -1
		x = negativeFlag * x
	}

	for x > 0 {
		temp := x % 10
		reverseNum = (reverseNum * 10) + temp
		x = x / 10
	}

	if reverseNum < -2147483648 || reverseNum > 2147483647 {
		return 0
	}

	return reverseNum * negativeFlag
}

// @lc code=end

