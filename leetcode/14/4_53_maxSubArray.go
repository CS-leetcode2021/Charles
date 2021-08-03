package _4

/**
 *  @ClassName:3_53_maxSubArray
 *  @Description:最大子序列和
 *  @Author:jackey
 *  @Create:2021/7/26 下午4:49
 */

func maxSubArray(nums []int) int {
	// 记录当前i的值与前置sum相加 并与原值比较，不会添加一个比我自身还小的数值
	if len(nums) <= 0 {
		return 0
	}
	n := len(nums)
	max, currSum := nums[0], nums[0] // 特例：nums[] = {1},max == 0 不会进入循环
	for i := 1; i < n; i++ {
		if nums[i] < nums[i]+currSum {
			currSum += nums[i]
		} else {
			currSum = nums[i]
		}
		max = Max53(max, currSum)
	}
	return max
}

func Max53(i, j int) int {
	if i > j {
		return i
	}
	return j
}
