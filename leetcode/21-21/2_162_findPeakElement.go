package _1_21

/**
 *  @ClassName:1_162_findPeakElement
 *  @Description:选找峰值
 *  @Author:jackey
 *  @Create:2021/8/6 下午1:39
 */

func findPeakElement(nums []int) int {
	// 大于左表同时大于右边
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if i == 1 && nums[i-1] > nums[i] {
			return 0
		} else if nums[i] > nums[i-1] {
			if i == len(nums)-1 {
				return i
			}
			if nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return -1
}

// 二分
// 多个峰值不影响相关的升序降序，可以使用二分
func findPeakElement2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) >> 1
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
