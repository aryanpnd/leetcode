/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	minP, maxPft := int(^uint(0)>>1), 0

	for _, price := range prices {
		if price < minP {
			minP = price
		} else {
			maxPft = max(maxPft, price-minP)
		}
	}

	return maxPft
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// @lc code=end

