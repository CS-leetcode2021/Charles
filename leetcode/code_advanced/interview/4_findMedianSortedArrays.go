package interview

import "fmt"

/**
 *  @ClassName:4_findMedianSortedArrays
 *  @Description:寻找两个升序数组的中位数
 *  @Author:jackey
 *  @Create:2021/8/5 下午9:51
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength/2
		return float64(getNum(nums1, nums2, midIndex + 1))
	} else {
		midIndex1, midIndex2 := totalLength/2 - 1, totalLength/2
		return float64(getNum(nums1, nums2, midIndex1 + 1) + getNum(nums1, nums2, midIndex2 + 1)) / 2.0
	}
	return 0.0
}

func getNum(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0

	for {
		if len(nums1) == index1 {
			// 去掉index2自身
			return nums2[index2+k-1]
		}

		if len(nums2) == index2 {
			return nums1[index1+k-1]
		}

		if k == 1 {
			return min17(nums1[index1], nums2[index2])
		}

		mid := k >> 1

		newIndex1 := min17(index1+mid, len(nums1))-1
		newIndex2 := min17(index2+mid, len(nums2))-1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}

	}
	return 0.0
}

func min17(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func main() {
	m := []int{1,2}
	n := []int{3,4}
	res := findMedianSortedArrays(m,n)
	fmt.Println(res)

}