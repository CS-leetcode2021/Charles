package main

/**
 *  @ClassName:climbStairs.go
 *  @Description:leetcode 70 爬楼梯
 *  @Author:jackey
 *  @Create:2021/6/7 下午9:30
 */

/*
 *  @Description:   problem(i) = subproblem(i-1)+subproblem(i-2)
 *  @Param:
 *  @Return:
 */

func climbStairs(n int) int {
	dp := make([]int,n+1)

	for i := 0; i <= n; i++ {
		dp[i] = 0
	}

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] +dp[i-2]
	}
	return dp[n]
}