package dp21

/**
 *  @ClassName:5_53_maxSubArray
 *  @Description:最大子序列和
 *  @Author:jackey
 *  @Create:2021/8/10 下午3:47
 */

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}

	n := len(nums)

	max , currSum := 0,nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1]+currSum {
			currSum += nums[i]
		}else {
			currSum = nums[i]
		}
		max = Max53(max,currSum)
	}
	return max
}

func Max53(i, j int) int {
	if i > j {
		return i
	}
	return j
}