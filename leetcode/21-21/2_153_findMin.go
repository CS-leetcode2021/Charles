package _1_21

import "sort"

/**
 *  @ClassName:1_153_findMin
 *  @Description:和53题不一样，53是寻找旋转数组的指定值，该题是寻找最小值
 *  @Author:jackey
 *  @Create:2021/8/6 下午1:30
 */

func findMin(nums []int) int {

	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) >> 1
		// 分左右半段

		if nums[high] > nums[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return nums[low]
}

func findMin2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sort.Ints(nums)
	return nums[0]
}
