package dp21

/**
 *  @ClassName:10_413_numberOfArithmeticSlices
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/26 下午1:35
 */


// 思路：等差数组的长度至少大于等于3,小于3的长度无结果
// dp：dp[i]代表以i位置为结束时nums[0:i]共有多少个等差数列个数
// dp[i] = dp[i-1] + 1
// 等差位置是可以在任意一个位置终结的，所以需要对其求和
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)

	if n < 3 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 0
	res := 0
	for i := 2; i < n; i++ {

		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			res += dp[i]
		}
	}
	return res
}
