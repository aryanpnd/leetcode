/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start

// using recursion and two pointers
// func isPalindrome(s string) bool {
// 	trimmedString := trimString(s)
// 	stringLen := len(trimmedString)
// 	return isPalindromeCheck(0, stringLen-1, trimmedString)
// }

// func trimString(str string) string {
// 	var filteredString []rune
// 	for _, char := range str {
// 		if unicode.IsLetter(char) || unicode.IsDigit(char) {
// 			filteredString = append(filteredString, unicode.ToLower(char))
// 		}
// 	}
// 	return string(filteredString)
// }

// func isPalindromeCheck(left int, right int, inputString string) bool {
// 	if left >= right {
// 		return true
// 	}

// 	if inputString[left] != inputString[right] {
// 		return false
// 	}

// 	return isPalindromeCheck(left+1, right-1, inputString)
// }

// using loop and two-pointers
func isPalindrome(s string) bool {
	var filtered []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered = append(filtered, unicode.ToLower(ch))
		}
	}

	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end

