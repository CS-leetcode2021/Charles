package _1_14

/**
 *  @ClassName:1_search
 *  @Description:704 二分查找
 *  @Author:jackey
 *  @Create:2021/7/26 上午10:24
 */

func search(nums []int, target int) int {
	return binSearch704(nums, 0, len(nums)-1, target)
}

func binSearch704(nums []int, l, r, target int) int {
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
