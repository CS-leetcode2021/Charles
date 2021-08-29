package dp21

/**
 *  @ClassName:3_53_maxSubArray
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/29 下午9:15
 */

func maxSubArray53(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	max, curSum := nums[0], nums[0]

	for i := 1; i < n; i++ {
		if nums[i] < curSum+nums[i] {
			curSum += nums[i]
		} else {
			curSum = nums[i]
		}
		max = Max(curSum, max)
	}
	return max
}
