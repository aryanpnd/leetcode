/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start

// using loops
// func fib(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}

// 	a, b, sum := 0, 1, 0

// 	for i := 1; i < n; i++ {
// 		sum = a + b
// 		a = b
// 		b = sum
// 	}

// 	return sum
// }

// using recursion
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)

}

// @lc code=end

