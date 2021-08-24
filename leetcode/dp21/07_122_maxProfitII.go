package dp21

/**
 *  @ClassName:07_122_maxProfitII
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/24 上午11:55
 */
// dp[i][j] 代表第i天j状态下的收益
// dp[i][0] = Max(dp[i-1][0],dp[i-1][1]+price[i])
// dp[i][1] = Max(dp[i-1][1],dp[i-1][0]-price[i])
//
func maxProfitII(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// 优化error
func maxProfitII2(prices []int) int {
	n := len(prices)
	tmp0 := 0
	tmp1 := -prices[0]
	for i := 1; i < n; i++ {
		tmp1, tmp0 = Max(tmp0, tmp1+prices[i]), Max(tmp1, tmp0-prices[i])
	}
	return tmp0
}
