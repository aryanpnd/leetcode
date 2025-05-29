/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reverseNum := 0
	originalNum := x
	for x > 0 {
		temp := x % 10
		reverseNum = (reverseNum * 10) + temp
		x = x / 10
	}

	return originalNum == reverseNum

}

// @lc code=end

