package __search

import "sort"

/**
 *  @ClassName:153_findMin
 *  @Description:153 旋转数组 二分查找
 *  @Author:jackey
 *  @Create:2021/7/23 下午2:37
 */

func FindMin(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

// 使用二分查找
// 数组中元素各不相同

func FindMin2(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := (low + high) >> 1

		if nums[mid] < nums[high] {
			high = mid
		}else {
			low = mid +1
		}
	}
	return nums[low]
}
