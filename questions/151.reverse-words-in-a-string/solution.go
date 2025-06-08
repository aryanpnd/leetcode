/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
// func reverseWords(s string) string {
// 	cleanedStr := removeTrailingSpaces(s)
// 	splitStrArr := splitStringInWords(cleanedStr)
// 	reverseWordsArr(splitStrArr)
// 	final := joinWords(splitStrArr, " ")

// 	return final
// }

// func removeTrailingSpaces(s string) string {
// 	runes := []rune(s)

// 	start := 0
// 	for start < len(runes) && !unicode.IsLetter(runes[start]) {
// 		start++
// 	}

// 	end := len(runes) - 1
// 	for end >= 0 && !unicode.IsLetter(runes[end]) {
// 		end--
// 	}

// 	if start > end {
// 		return ""
// 	}

// 	return string(runes[start : end+1])
// }

// func splitStringInWords(s string) []string {
// 	words := []string{}
// 	word := ""

// 	for _, char := range s {
// 		if char == ' ' {
// 			if word != "" {
// 				words = append(words, word)
// 				word = ""
// 			}
// 		} else {
// 			word += string(char)
// 		}
// 	}

// 	if word != "" {
// 		words = append(words, word)
// 	}

// 	return words
// }

// func reverseWordsArr(wordArr []string) {
// 	for i, j := 0, len(wordArr)-1; i < j; i, j = i+1, j-1 {
// 		wordArr[i], wordArr[j] = wordArr[j], wordArr[i]
// 	}
// }

// func joinWords(a []string, sep string) string {
// 	if len(a) == 0 {
// 		return ""
// 	}

// 	result := a[0]
// 	for i := 1; i < len(a); i++ {
// 		result += sep + a[i]
// 	}
// 	return result
// }

func reverseWords(s string) string {
	words := strings.Fields(s)
	reverse(words)
	return strings.Join(words, " ")
}

func reverse(words []string) {
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
}

// @lc code=end

