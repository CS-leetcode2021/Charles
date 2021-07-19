package main

/**
 *  @ClassName:39_majorityElement
 *  @Description:剑指offer-39 数组中出现次数超过一半的数字 同 leetcode-169 
 *  @Author:jackey
 *  @Create:2021/7/19 下午3:03
 */

func majorityElement(nums []int) int {
	m := len(nums)
	tag := 1
	res := nums[0]

	for i := 1; i < m; i++ {
		if nums[i] == res {
			tag++
		}else if tag == 0 {
			res = nums[i]
			tag = 1
		}else {
			tag--
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		if nums[i] == res {
			count++
		}
	}
	if count <= m/2 {
		return -1
	}
	return res
}
