/*
 * @lc app=leetcode id=1711 lang=golang
 *
 * [1711] Count Good Meals
 */

// @lc code=start
func countPairs(deliciousness []int) int {
	const MOD = 1000000007
	seen := make(map[int]int)
	count := 0

	// for _, num := range deliciousness {
	// 	seen[num] = seen[num] + 1
	// 	// or
	// 	// seen[num]++
	// }
	for _, num := range deliciousness {
		for i := 0; i <= 21; i++ {
			target := 1 << i
			if seenNumCount, found := seen[target-num]; found {
				count = (count + seenNumCount) % MOD
			}
		}
		seen[num]++
	}

	return count

}

// @lc code=end

