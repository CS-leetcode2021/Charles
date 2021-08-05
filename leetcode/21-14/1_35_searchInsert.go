package _1_14

/**
 *  @ClassName:1_35_searchInsert
 *  @Description:寻找插入的位置
 *  @Author:jackey
 *  @Create:2021/7/26 上午10:43
 */

// 100/100
// 二分
func searchInsert(nums []int, target int) int {
	return binSearch35(nums, 0, len(nums)-1, target)
}

func binSearch35(nums []int, l, r, target int) int {
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
	return l
}
