package _1_14

import "sort"

/**
 *  @ClassName:14_136_singleNumber
 *  @Description:只出现一次数字,其它的都是两次
 *  @Author:jackey
 *  @Create:2021/8/5 下午6:23
 */

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i = i + 2 {
		if nums[i] != nums[i-1] {
			return nums[i-1]
		} else if i == len(nums)-2 {
			return nums[i+1]
		}
	}
	return -1
}

func singleNumber2(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); {
		if i == len(nums)-1 {
			return nums[i]
		}
		j := i + 1
		if nums[i] != nums[j] {
			return nums[i]
		}
		i += 2
	}
	return -1
}

func singleNumber3(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
