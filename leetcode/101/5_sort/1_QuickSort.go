package __sort

/**
 *  @ClassName:1_Quicksort
 *  @Description:快速排序
 *  @Author:jackey
 *  @Create:2021/7/23 下午3:33
 */

// 基准
// 左侧寻找比基准大的数字，后侧寻找比基准小的数字，然后交换
// 双向指针接触后，与基准发生互换
// 结合二分
func quickSort(nums []int, l int, r int) {

	pivot := nums[l]
	start := l
	end := r
	for start < end {
		for start < end && pivot < nums[end] {
			end--
		}
		for start < end && pivot > nums[start] {
			start++
		}
		if start < end {
			swap(nums, start, end)
		}

	}
	swap(nums, start, l)
	return
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
