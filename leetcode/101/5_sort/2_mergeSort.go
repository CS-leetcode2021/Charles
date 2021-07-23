package __sort

/**
 *  @ClassName:2_merge
 *  @Description:归并排序
 *  @Author:jackey
 *  @Create:2021/7/23 下午4:00
 */

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) >> 1

	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	l, r := len(left), len(right)

	i, j := 0, 0
	res := make([]int, 0)
	for i < l && j < r {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
