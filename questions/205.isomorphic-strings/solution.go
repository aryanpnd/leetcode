/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	n := len(s)

	for i := 0; i < n; i++ {
		sChar := s[i]
		tChar := t[i]

		if mapped, exist := sMap[sChar]; exist {
			if mapped != tChar {
				return false
			}
		} else {
			sMap[sChar] = tChar
		}

		if mapped, exist := tMap[tChar]; exist {
			if mapped != sChar {
				return false
			}
		} else {
			tMap[tChar] = sChar
		}
	}

	return true
}

// @lc code=end

