/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 {
		if seen[n] {
			return false
		}
		seen[n] = true
		n = sumOfSquares(n)
	}
	return true
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		t := n % 10
		sum += t * t
		n = n / 10
	}
	return sum
}

// @lc code=end

