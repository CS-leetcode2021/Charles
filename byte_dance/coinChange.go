package main

import "math"

/**
 *  @ClassName:coinChange.go
 *  @Description:322 leetcode 零钱兑换
 *  @Author:jackey
 *  @Create:2021/6/7 下午7:29
 */

// 类似于动态规划（青蛙跳楼），dp[表示能实现的最少数量]
//
func coinChange(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}

	// 最少的数量
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32

		for _, c := range coins {
			if i >= c && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] +1
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

/*
 *  @Description:   不压缩空间实现，也就是使用二维数组实现动态规划
 *  @Param:
 *  @Return:
 */

func coinsChange(coins []int, amount int) int {
	dp := make([][]int,len(coins)+1)

	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int,amount+1)
	}
	for i := 0; i <= amount; i++ {
		dp[0][i] = amount+1
	}

	dp[0][0] = 0

	for i := 1; i <= len(coins) ; i++ {
		for j := 0; j <=amount; j++ {
			// j 理解为所余空间，如果比所选币种要大，那你就可以 选或者不选
			if j >= coins[i-1] {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]),float64(dp[i][j-coins[i-1]]+1)))
			}else {
				// 所余空间没有当前选择的币种大，那一定是不选择的
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	if dp[len(coins)][amount] > amount {
		return -1
	}else {
		return dp[len(coins)][amount]
	}
}

/*
 *  @Description:   压缩空间，采用一维数组
 *  @Param:			我们发现，dp[i][j]只和dp[i-1][j]与dp[i-1][j-coins[i]]有关
 *  @Return:
 */
func CoinChange(coins []int, amount int) int {
	dp:=make([]int,amount+1)
	//初始值
	for i:=0;i<=amount;i++{
		dp[i]=amount+1
	}
	dp[0]=0

	for i:=1;i<=len(coins);i++{
		for j:=0;j<=amount;j++{
			if j >= coins[i-1]{
				dp[j]=int(math.Min(float64(dp[j]),float64(dp[j-coins[i-1]]+1)))
			}
		}
	}
	if dp[amount] > amount{
		return -1
	}else{
		return dp[amount]
	}

	return 0
}




















