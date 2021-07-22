package main

import "math"

/**
 *  @ClassName:coin_change
 *  @Description: 零钱兑换
 *  @Author:jackey
 *  @Create:2021/3/16 下午1:52
 */
/** comments
1、01背包：如果每个物品最多拿一次，则称之为01背包；反之，可以拿取无数次的称为完全不背包
2、完全背包的状态可以定义为如下：dp[i][j]:当考虑放入第i个物品时，占用空间j所能得到的最大的价值；状态转移方程为：dp[i][j]=Max(dp[i-1][j],dp[i][j-vi]+wi)
3、每次放入一个物体是只有两种选择：要么放，要么不放。
	3.1 不放，dp[i-1][j]
	3.2 放，dp[i][j-vi]+wi;那么就需要把背包从当前容量为j的状态给它掏出去vi的量，这样才能将物品i放进去
4、总金额就是背包容量，硬币就是物品占用的体积coins[i]，数量就是价值wi=1
5、改进公式 dp[i][j]=min(dp[i-1][j],dp[i][j-coins[i]]+1)

*/
func coinChange(coins []int, amount int) int {
	/*初始化二维数组状态*/
	dp := make([][]int, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}

	/*初始化条件*/
	for j := 0; j <= amount; j++ {
		dp[0][j] = amount + 1
	}

	dp[0][0] = 0

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-coins[i-1]]+1)))
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	if dp[len(coins)][amount] > amount {
		return -1
	} else {
		return dp[len(coins)][amount]
	}
}

/*输入一个目标金额 n，返回凑出目标金额 n 的最少硬币数量。*/

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-coins[i-1]]+1)))
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

func coinChange3(coins []int, amount int) int {
	dp := make([]int,amount+1)
	for i := 0; i <= amount ; i++ {
		dp[i]=amount+1
	}

	dp[0]=0

	for  {

	}


}
