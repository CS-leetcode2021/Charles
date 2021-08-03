package _1

/**
 *  @ClassName:12_198_rob
 *  @Description:打家劫舍
 *  @Author:jackey
 *  @Create:2021/8/3 下午6:34
 */

// dp
// i 可以选择偷也可以选择不偷
// 如果偷，金额为当前金额加i-2的总金额
// 如果不偷，那就是i-1的金额
// 100/57
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 只有一家，必须偷
	if len(nums) == 1 {
		return nums[0]
	}

	// 使用一维数组
	dp := make([]int, len(nums)+1)

	dp[1] = nums[0]
	dp[2] = Max198(nums[0], nums[1])

	for i := 2; i <= len(nums); i++ {
		dp[i] = Max198(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}

func Max198(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 100/100
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	tag1, tag2 := nums[0], Max198(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tag1, tag2 = tag2, Max198(tag2, tag1+nums[i])
	}

	return tag2

}
