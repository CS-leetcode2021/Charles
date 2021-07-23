package __search

/**
 *  @ClassName:154_findMin
 *  @Description:旋转数组n次，寻找最小值
 *  @Author:jackey
 *  @Create:2021/7/23 下午12:04
 */

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}

	return nums[0]
}