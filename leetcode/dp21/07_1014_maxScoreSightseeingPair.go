package dp21

/**
 *  @ClassName:7_1014_maxScoreSightseeingPair
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/24 上午10:27
 */

func maxScoreSightseeingPair(values []int) int {
	dp := make([]int, len(values))
	// 必须两个观景台
	dp[0] = 0
	max := values[0]
	for i := 1; i < len(values); i++ {
		dp[i] = Max1014(values[i]-i+max, dp[i-1])
		max = Max1014(max, values[i]+i)
	}
	return dp[len(values)-1]
}

func Max1014(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// error
func maxScoreSightseeingPair2(values []int) int {
	max := values[0]
	ans := 0

	for i := 1; i < len(values); i++ {
		ans = Max1014(values[i]-i + max,ans)
		max = Max1014(max, values[i]+i)
	}
	return ans
}
