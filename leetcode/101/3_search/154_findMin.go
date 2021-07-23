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

// 二分法查找

func findMin2(nums []int) int {
	low , high := 0, len(nums)-1
	for low <= high {
		mid := (low+high) >>1

		if nums[mid] < nums[high] {
			high = mid
		}else if nums[mid] > nums[high] {
			low = mid +1
		}else {
			high--
		}
	}
	return nums[low]
}
