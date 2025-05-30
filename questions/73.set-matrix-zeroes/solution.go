/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start

// brute force, setting the zeros to -97 and then resetting them back to zero
// func setZeroes(matrix [][]int) {
// 	rows := len(matrix)
// 	if rows == 0 {
// 		return
// 	}
// 	cols := len(matrix[0])

// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			if matrix[i][j] == 0 {
// 				setRow(matrix, i)
// 				setColumn(matrix, j)
// 			}
// 		}
// 	}

// 	finalize(matrix)
// }

// func setRow(matrix [][]int, i int) {
// 	for j := 0; j < len(matrix[i]); j++ {
// 		if matrix[i][j] != 0 {
// 			matrix[i][j] = -97
// 		}
// 	}
// }

// func setColumn(matrix [][]int, j int) {
// 	for i := 0; i < len(matrix); i++ {
// 		if j < len(matrix[i]) && matrix[i][j] != 0 {
// 			matrix[i][j] = -97
// 		}
// 	}
// }

// func finalize(matrix [][]int) {
// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[0]); j++ {
// 			if matrix[i][j] == -97 {
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}
// }

// @lc code=end

