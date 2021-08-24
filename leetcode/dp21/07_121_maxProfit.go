package dp21

/**
 *  @ClassName:07_121_maxProfit
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/24 上午11:21
 */
// dp[i][j] 代表第i天j状态下的收益
// dp[i][0] = Max(dp[i-1][0],dp[i-1][1]+price[i])
// dp[i][1] = Max(dp[i-1][1],-price[i])
// 优化只要两个变量
// tmp0 表示不持有状态
// tmp1 表示持有状态

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

// 优化一下

func maxProfit2(prices []int) int {
	n := len(prices)
	tmp0 := 0
	tmp1 := -prices[0]
	for i := 1; i < n; i++ {
		tmp0 = Max(tmp0, tmp1+prices[i])
		tmp1 = Max(tmp1, -prices[i])
	}
	return tmp0
}
