package array

import "sort"

/**
 *  @ClassName:merge
 *  @Description:leetcode 88 合并两个有序数组
 *  @Author:jackey
 *  @Create:2021/5/19 下午7:45
 */

/*
 *  @Description:   直接使用append进行拼接
 *  @Param:         数组
 *  @Return:        nil
 *  @problem:		给定的数组是固定长度的，不能使用append直接进行添加。使用copy函数，将nums2直接拼接到nums1的尾部，然后使用sort进行排序
 */

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:],nums2)
	// go语言内部实现的排序是快排，递归找基准是logn，遍历所有的数据是n，所以时间复杂度是 nlogn
	// {@link} https://blog.csdn.net/weshjiness/article/details/8660583
	sort.Ints(nums1)
}

/*
 *  @Description:   双指针遍历，再将结果覆盖给nums1
 *  @Param:         数组
 *  @Return:        nil
 */

func merge2(nums1 []int, m int, nums2 []int, n int) {
	res := []int{}

	p1 := 0
	p2 := 0

	for p1 < m && p2 < n {
		if nums1[p1]<nums2[p2] {
			res = append(res,nums1[p1])
			p1++
		}else{
			res = append(res,nums2[p2])
			p2++
		}
	}

	if p1 == m {
		for i := p2; i < n; i++ {
			res = append(res,nums2[i])
		}
	}else{
		for i := p1; i <m ; i++ {
			res = append(res,nums1[i])
		}
	}

	copy(nums1,res)
}