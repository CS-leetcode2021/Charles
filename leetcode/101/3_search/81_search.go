package __search

import (
	"sort"
)

/**
 *  @ClassName:81_search
 *  @Description:TODO 旋转数组
 *  @Author:jackey
 *  @Create:2021/7/23 上午11:41
 */

// 可以暴力解，只是时间复杂度过大，利用本身的排序性质
//
// 93/54
func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			return true
		}
		// 如果左边和中点相同，我们不能确定那边是相同的区间，可以简单的将左边下标移动一位
		if nums[low] == nums[mid] {
			low++
		} else if nums[mid] <= nums[high] {
			// 右边是递增序列
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			// nums[mid] > nums[high]
			// 右边是排好序的
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return false
}

func search2(nums []int, target int) bool {
	sort.Ints(nums)
	return BinSearch(nums, target)
}

// 利用重新排序后的数组进行判定，空间复杂度过高
func BinSearch(nums []int, target int) bool {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (high + low) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}
