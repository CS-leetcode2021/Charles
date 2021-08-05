package _1_21

/**
 *  @ClassName:1_33_search
 *  @Description：搜索旋转排序数组
 *  @Author:jackey
 *  @Create:2021/8/5 下午7:08
 */

// 1、先判断mid在前半段还是在后半段
// 2、如果在前半段，利用前半段的升序关系递归求解
// 3、后半段同理
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] {
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

	}
	return -1
}
