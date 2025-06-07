/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start

// greedy approach
func maxProfit(prices []int) int {
	profit, n := 0, len(prices)
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit = profit + (prices[i] - prices[i-1])
		}
	}
	return profit
}

// @lc code=end

