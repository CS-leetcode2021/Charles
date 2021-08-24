package dp21

/**
 *  @ClassName:08_714_maxProfit
 *  @Description:leetcode 714
 *  @Author:jackey
 *  @Create:2021/8/24 下午3:02
 */

// 手续只计算一次
func maxProfitIV(prices []int, fee int) int {
	n := len(prices)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0                // 不持有的收益
	dp[0][1] = -prices[0] - fee // 持有的最大收益

	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}

	return dp[n-1][0]
}
