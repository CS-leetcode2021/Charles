package __sort

import "sort"

/**
 *  @ClassName:215_findKthLargest
 *  @Description:寻找第k大元素
 *  @Author:jackey
 *  @Create:2021/7/23 下午4:34
 */

// 100/100
func findKthLargest(nums []int, k int) int {

	if k > len(nums) {
		return -1
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[k-1]
}

// 快速选择

func findKthLargest2(nums []int, k int) int {
	k = len(nums) - k
	return quickSort215(nums, 0, len(nums)-1, k)
}

func quickSort215(nums []int, l, r, k int) int {
	for l < r {
		index := _quickSort215(nums, l, r)
		if index == k {
			return nums[k]
		}
		if index < k {
			l = index + 1
		} else {
			r = index - 1
		}
	}
	return nums[l]
}

func _quickSort215(nums []int, l, r int) int {
	start := l+1
	end := r
	pivot := nums[l]

	for {
		// 寻找第一个比基准小的。换到前面
		for start < end && pivot <= nums[end] {
			end--
		}
		// 寻找第一个比基准大的。换到后面
		for start < end && pivot >= nums[start] {
			start++
		}

		if start < end {
			swap215(nums, start, end)
		}else {
			break
		}
	}
	swap215(nums, l, start)
	return end
}

func swap215(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
