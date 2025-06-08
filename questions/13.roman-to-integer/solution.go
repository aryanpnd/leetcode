/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	rtiMap := map[byte]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		curr := rtiMap[s[i]]

		if i+1 < len(s) && curr < rtiMap[s[i+1]] {
			sum -= curr
		} else {
			sum += curr
		}
	}
	return sum
}

// @lc code=end

