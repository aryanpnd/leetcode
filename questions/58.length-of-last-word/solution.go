/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		if count == 0 && !(unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i]))) {
			continue
		}

		if unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) {
			count++
		} else {
			break
		}

	}

	return count
}

// @lc code=end
