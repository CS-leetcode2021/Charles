package __pointer

import "sort"

/**
 *  @ClassName:88_merge
 *  @Description:easy	合并升序数组
 *  @Author:jackey
 *  @Create:2021/7/22 上午10:44
 */

// 100/71
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 使用双指针
// 100/100
func merge2(nums1 []int, m int, nums2 []int, n int) {

	pos := m + n - 1
	m--
	n--
	for m > 0 && n > 0 {
		if nums1[m] > nums2[n] {
			nums1[pos] = nums1[m]
			pos--
			m--
		}else {
			nums1[pos] = nums2[n]
			pos--
			n--
		}
	}

	for n > 0 {
		nums1[pos] = nums2[n]
		n--
		pos--
	}

}
