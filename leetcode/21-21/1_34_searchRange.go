package _1_21

/**
 *  @ClassName:1_34_searchRange
 *  @Description:在排序数组中查找元素的第一个和最后一个位置
 *  @Author:jackey
 *  @Create:2021/8/5 下午6:56
 */

// 二分查找；条件：升序数组

func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

// 辅助函数
// 查找第一个

func findFirst(nums []int, target int) int {
	// 二分
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (high + low) >> 1

		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != nums[mid] {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func findLast(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != nums[mid] {
				return mid
			}
			low = mid + 1
		}

	}
	return -1
}
