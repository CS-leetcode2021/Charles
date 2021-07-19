package main

/**
 *  @ClassName:42_maxSubArray
 *  @Description:剑指 Offer 42. 连续子数组的最大和 同 leetcode-53
 *  @Author:jackey
 *  @Create:2021/7/19 下午3:48
 */

// dp

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

