package main

import (
	"sort"
)


/**
 *  @ClassName:lengthOfLongstSubstring
 *  @Description:剑指offer 03 数组中重复的数字
 *  @Author:jackey
 */

func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	l,r := 0,1

	for i := 0; i < len(nums); i++ {
		if nums[l] == nums[r] {
			return  nums[l]
		}else {
			l++
			r++
		}
	}
	return  -1
}

// 优化了时间复杂度
func findRepeatNumber2(nums []int) int {
	var sign [100000]bool
	for _,num := range nums {
		if sign[num] {return num}
		sign[num] = true
	}
	return 0
}

