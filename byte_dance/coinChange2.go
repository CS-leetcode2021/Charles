package byte_dance

/**
 *  @ClassName:coinChange2.go
 *  @Description:leetcode 518 零钱兑换2
 *  @Author:jackey
 *  @Create:2021/6/7 下午9:14
 */

func coinChange2(amount int, coins []int) int {
	dp := make([][]int,len(coins)+1)

	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int,amount+1)
	}

	for i := 0; i <= len(coins); i++ {
		dp[i][0]= 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[i][j]=dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(coins)][amount]
}

func change(amount int, coins []int) int {
	var n = len(coins)
	var dp = make([]int, amount+1) // dp[i]代表能组合的总金额i的组合数
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}

	//for i := 0; i <= amount; i++ {
	// fmt.Println(dp)
	//}
	// fmt.Println(dp[amount])
	return dp[amount]
}