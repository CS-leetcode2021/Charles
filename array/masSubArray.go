package array

/**
 *  @ClassName:masSubArray
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/3/16 下午6:19
 */


func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1]>nums[i]{
			nums[i] += nums[i-1]
		}
		if nums[i]>max{
			max = nums[i]
		}
	}
	return max
}

func max(nums []int) int {
	res := nums[0]

	for i :=1; i < len(nums); i++ {
		if nums[i]>res {
			res = nums[i]
		}

	}
	return res
}