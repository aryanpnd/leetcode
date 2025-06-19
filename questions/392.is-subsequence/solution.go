/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	current := 0

	if len(s) == 1 {
		return false
	}

	for i := 0; i < len(t); i++ {
		if t[i] == s[current] {
			current++
			if current == len(s) {
				return true
			}
		}
	}

	return false
}

// @lc code=end

