package __search

/**
 *  @ClassName:34_searchRange
 *  @Description:查找在排序数组中查找元素的第一个和最后一个位置  二分查找
 *  @Author:jackey
 *  @Create:2021/7/23 上午11:07
 */

// 这⼀题是经典的⼆分搜索变种题。⼆分搜索有 4 ⼤基础变种题：
//1. 查找第⼀个值等于给定值的元素
//2. 查找最后⼀个值等于给定值的元素
//3. 查找第⼀个⼤于等于给定值的元素
//4. 查找最后⼀个⼩于等于给定值的元素

// 100/100
func searchRange(nums []int, target int) []int {
	return []int{searchFirst(nums, target), searchSecond(nums, target)}
}

func searchFirst(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) >> 1

		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func searchSecond(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (high + low) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// 寻找第一个大于等于target的元素
func searchEqual(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high + low) >> 1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 寻找最后一个小于等于target的元素

func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] <= target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
