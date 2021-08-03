package _4

/**
 *  @ClassName:4_416_canPartition
 *  @Description:分割等和子集
 *  @Author:jackey
 *  @Create:2021/8/3 下午8:23
 */

// 长度小于2返回false
// 数组总和是奇数返回false
func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	sum, max := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if max < nums[i] {
			max = nums[i]
		}
	}
	if sum&1 == 1 {	// 奇数
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}

	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}

	// dp[i][0] = true
	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}

	dp[0][nums[0]] = true

	for i := 1; i < len(nums); i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}
