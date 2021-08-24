package dp21

/**
 *  @ClassName:08_309_maxProfit
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/24 下午2:41
 */

func maxProfitIII(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
	}

	dp[0][0] = 0          // 不持有在冷冻期的最大收益
	dp[0][1] = -prices[0] // 持有的最大收益
	dp[0][2] = 0          // 不持有不在冷冻期的收益

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][1]+prices[i]
		dp[i][1] = Max(dp[i-1][1],dp[i-1][2]-prices[i])
		dp[i][2] = Max(dp[i-1][0],dp[i-1][2])
	}
	return Max(dp[n-1][0],dp[n-1][2])
}
