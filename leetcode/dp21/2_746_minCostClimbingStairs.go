package dp21

/**
 *  @ClassName:2_746_minCostClimbingStairs
 *  @Description:使用最小花费爬楼梯
 *  @Author:jackey
 *  @Create:2021/8/6 下午3:34
 */

// dp问题
// dp[i] 表示当前i次行走所花费的花销总和
// dp[i] = min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
// 100/30
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 || n == 1 {
		return 0
	}
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = min746(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]

}

func min746(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// 优化空间
// 100/100
func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	if n == 0 || n == 1 {
		return 0
	}

	tmp1, tmp2 := 0, 0

	for i := 2; i <= n; i++ {
		tmp1, tmp2 = min746(tmp1+cost[i-1], tmp2+cost[i-2]), tmp1
	}
	return tmp1
}
