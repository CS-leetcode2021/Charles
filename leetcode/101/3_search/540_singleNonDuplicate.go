package __search

/**
 *  @ClassName:540_singleNonDuplicate
 *  @Description:寻找出现一次的数字
 *  @Author:jackey
 *  @Create:2021/7/23 下午3:02
 */

// 100/100
func singleNonDuplicate(nums []int) int {
	if len(nums) == 0 || len(nums) == 2 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 0; i < len(nums); i = i + 2 {
		if i == len(nums)-1 {
			return nums[i]
		}
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// 使用二分查找

func singleNonDuplicate2(nums []int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if mid == low || mid == high {
			return nums[mid]
		}
		if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
			return nums[mid]
		} else if (mid%2 == 0 && nums[mid] == nums[mid+1]) ||
			(mid%2 == 1 && nums[mid] == nums[mid-1]) {
			low = mid + mid%2
		} else {
			high = mid - mid%2
		}
	}
	return 0
}
// 利用异或 但是时间复杂度不合适

func singleNonDuplicate3(nums []int) int {
	res := 0

	for _, v := range nums {
		res ^= v
	}
	return res
}

//