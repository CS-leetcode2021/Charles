package array

import (
	"math"
)

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
func maxProfit02(prices []int) int {
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

/*
 *  @Description:   成功代码：一次遍历，记录最低的值，同时利用最低的值找到最大的差值，就是最大的利润
 *  @Param:
 *  @Return:        
 */

func maxProfit03(prices []int) int {
	minValue := math.MaxInt64
	maxValue := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			// 寻找最低的值
			minValue = prices[i]
		} else if prices[i]-minValue > maxValue {
			// 按照当前最低的值往后寻找最大的差值
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}

// 股票通用模板：使用动态规划进行解题
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

func maxProfit04(prices []int) int {

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

/*
 *  @Description:   leetcode 122 利润分阶段实现，只要后者的价格大于前一位的价格，就是利润
 *  @Param:
 *  @Return:
 */
func maxProfit2(prices []int) int {
	n := len(prices)
	maxValue := 0

	sum := 0
	// 数据分两种，一种是递增的数据，利润应该是当前的减去之前的
	// 另一种是变低的，应该是重新买入，并且更新利润
	// 其实不用那么麻烦，直接把增量数据叠加起来，也就是只考虑第一种数据

	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			maxValue = prices[i]-prices[i-1]
			sum += maxValue
		}
	}
	return sum
}

/*
 *  @Description:   leetcode 123 买卖股票III  动态规划求解
 *  1、确定dp数组的规划：一种只有五种状态，0：没有任何操作；1：第一次买入；2：第一次卖出；3：第二次买入；4：第二次卖出
 *  2、dp[i][j]：表示第i天，j为她的五个状态，dp[i][j]表示第i天状态j中所剩的最大现金
 *  3、确定递推公式：
 * 				* dp[i][1]：两种操作：
 * 				* 			操作一： 第i天买入股票了，那么dp[i][1] = dp[i-1][0] - prices[i] // 因为之前没有买，所以所剩的资金大于现在第i天的
 * 				* 			操作二： 第i天没有操作，那么沿用i-1天的买入的状态， dp[i][i] = dp[i-1][1]
 * 				* 			所剩资金一定是最大的 dp[i][1] = max(dp[i-1][0]-prices[i],dp[i-1][i])
 * 				*
 * 				* dp[i][2]：两种操作：
 * 							操作一：第i天卖出股票了，那么dp[i][2] = dp[i - 1][1] + prices[i]
 * 							操作二：第i天没有操作，沿用前一天卖出股票的状态，即：dp[i][2] = dp[i - 1][2]
 *							所以dp[i][2] = max(dp[i - 1][1] + prices[i], dp[i - 1][2])
 * 				* dp[i][3]：两种操作：
 * 							操作一：第i天第二次买入， dp[i][3] = dp[i-1][2] -prices[i]
 * 							操作二：第i天没有进行买入，dp[i][3] = dp[i-1][3]
 * 							dp[i][3] = max(dp[i-1][2] - prices[i],dp[i-1][3])
 * 				* dp[i][4]：两种操作：
 * 							操作一：第i天第二次买入， dp[i][4] = dp[i-1][3] + prices[i]
 * 							操作二：第i天没有进行买入，dp[i][4] = dp[i-1][4]
 * 							dp[i][4] = max(dp[i-1][3] + prices[i],dp[i-1][4])
 *
 * 4、初始化dp数组：
 * 			第0天没有擦欧欧，所以			dp[0][0] = 0
 *          第0天有第一次买入的操作，所以	dp[0][1] = -prices[0]
 *          第0天第一次卖出操作，			dp[0][2] = 0
 * 			第0天第二次卖出的操作，			dp[0][3] = -prices[0]
 * 			第0天第二次卖出操作，			dp[0][2] = 0
 *
 *
 *
 */

func maxProfit3(prices []int) int {
	dp := make([][]int,len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int,5)
	}

	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = maxDistance(dp[i-1][1],dp[i-1][0]-prices[i])
		dp[i][2] = maxDistance(dp[i-1][2],dp[i-1][1]+prices[i])
		dp[i][3] = maxDistance(dp[i-1][3],dp[i-1][2]-prices[i])
		dp[i][4] = maxDistance(dp[i-1][4],dp[i-1][3]+prices[i])
	}

	return dp[len(prices)-1][4]
}


/*
 *  @Description:   另一种空间复杂度较低的实现
 *  @Param:
 *  @Return:
 */
func maxProfit4(prices []int) int {
	if(len(prices) < 2){
		return 0
	}
	sell1 := 0
	sell2 := 0
	buy1 := math.MinInt32
	buy2 := math.MinInt32
	for _, p := range prices {
		buy1 = maxPro(buy1, -p)
		sell1 = maxPro(sell1, buy1 + p)
		buy2 = maxPro(buy2, sell1 - p)
		sell2 = maxPro(sell2, buy2 + p)
	}
	return sell2
}

func maxPro(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

