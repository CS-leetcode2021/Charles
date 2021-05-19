package array

import "math"

/**
 *  @ClassName:maxProfit
 *  @Description:leetcode 121 买卖股票的最佳时机
 *  @Author:jackey
 *  @Create:2021/5/19 下午8:33
 */

/*
 *  @Description:   双循环，找出历史最大的差价
 *  @Param:         数组
 *  @Return:        最大利润
 *  @problem:		超出时间限制
 */

func maxProfit(prices []int) int {
	n := len(prices)
	maxProfit := 0
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if prices[j] > prices[i] {
				maxProfit = maxDistance(maxProfit,prices[j]-prices[i])
			}
		}
	}

	return maxProfit
}

// 个别案例没通过
func maxProfit2(prices []int) int {
	index := 0
	min := prices[0]
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < min {
			min = prices[i]
			index = i
		}
	}

	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if i > index && prices[i]-min > 0 {
			maxProfit = maxDistance(maxProfit,prices[i]-min)
		}
	}

	return maxProfit
}


func maxProfit3(prices []int) int {
	minValue := math.MaxInt64
	maxValue := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		} else if prices[i]-minValue > maxValue {
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}

// 股票通用模板
/*
三种状态：
    天数:  n 取值范围： 0 <= i<= n-1
    最大交易次数: K     1 <= k <= K
    当前持有的状态（1：持有 0：非持有）

状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0],dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1],dp[i-1][k-1][0] - prices[i])

base case:
//天数的初始状态为-1、交易次数初始状态为0
dp[-1][k][0] = dp[i][0][0] = 0
dp[-1][k][1] = dp[i][0][1] = -infinity


股票问题一:
    k = 1；交易次数

所以上式状态转移方程为：
    dp[i][0] = max(dp[i-1][0],dp[i-1][1] + prices[i])
    dp[i][1] = max(dp[i-1][1],dp[i-1][0][0] - prices[i])
             = max(dp[i-1][1],0-prices[i])
             = max(dp[i-1][1],-prices[i])

base case:
dp[-1][k][0] = dp[i][0][0] =  0
dp[-1][k][1] = dp[i][0][1] = -infinity

*/

func maxProfit4(prices []int) int {

	n := len(prices)
	//init dp
	dp := make([][]int,n)
	for i := 0;i < n;i++{
		dp[i] = make([]int,2)
	}

	//prices[i]
	for i := 0; i < n;i++{
		//base case
		if i - 1 < 0{
			dp[0][0] = 0
			dp[0][1] = - prices[i]
			continue
		}
		//当前持有的状态 1：持有 0:没有持有
		for j := 0;j < 2;j++{
			dp[i][0] = MaxProfit(dp[i-1][0],dp[i-1][1] + prices[i])
			dp[i][1] = MaxProfit(dp[i-1][1],-prices[i])


		}

	}

	return dp[n-1][0]
}

func MaxProfit(a,b int)int{
	if a > b{
		return a
	}
	return b
}

