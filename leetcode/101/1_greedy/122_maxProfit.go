package __greedy

/**
 *  @ClassName:122_maxProfit
 *  @Description:maxProfit 股票交易
 *  @Author:jackey
 *  @Create:2021/7/22 下午4:41
 */

// 贪心算法 ： 局部最优解


func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		res += MaxPro(prices[i] - prices[i-1],0)	// 这里只要跟0比较就可以了，只要不是负数就行
	}
	return res
}

func MaxPro(i, j int) int {
	if i > j {
		return i
	}
	return j
}