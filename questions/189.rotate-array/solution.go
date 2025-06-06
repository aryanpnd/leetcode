/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start

// func rotate(nums []int, k int) {
// 	numsLen := len(nums)
// 	if numsLen <= 1 || k == 0 {
// 		return
// 	}

// 	k = k % numsLen // handle k > n
// 	if k == 0 {
// 		return
// 	}

// 	arr := []int{}

// 	for i := numsLen - k; i < numsLen; i++ {
// 		arr = append(arr, nums[i])
// 	}

// 	for i := 0; i < numsLen-k; i++ {
// 		arr = append(arr, nums[i])
// 	}

// 	for i := 0; i < numsLen; i++ {
// 		nums[i] = arr[i]
// 	}

// }

func rotate(nums []int, k int) {
	numsLen := len(nums)
	if numsLen <= 1 || k == 0 {
		return
	}

	k = k % numsLen // handle k > n
	if k == 0 {
		return
	}

	reverse(nums, 0, numsLen-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, numsLen-1)
}

func reverse(arr []int, left int, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// @lc code=end

