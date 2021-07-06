package main

import "sort"

/**
 *  @ClassName:minArray
 *  @Description:offer-11 旋转数组的最小数字 && leetcode 154
 *  @Author:jackey
 *  @Create:2021/7/6 下午8:16
 */


// 双指针
func minArray(numbers []int) int {
	index := 0
	for i, j := 0, 1; j < len(numbers); i, j = i+1, j+1 {
		if numbers[i] > numbers[j] {
			index = j
			break
		}
	}
	return numbers[index]
}

// 利用go语言内置函数
func minArray1(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}


// 一次扫描
func minArray2(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}

	return numbers[0]
}

// 二分法
/*
 *  @Description:   1 有序数组分成了左右2个小的有序数组，而实际上要找的是右边有序数组的最小值
					2 如果中间值大于右边的最大值，说明中间值还在左边的小数组里，需要left向右移动
					3 如果中间值小于等于当前右边最大值，至少说明了当前右边的right值不是最小值了或者不是唯一的最小值，需要慢慢向左移动一位
 *  @Param:
 *  @Return:        
 */
func minArray3(numbers []int) int {
	left := 0
	right := len(numbers)-1

	for left < right {
		mid := left + (right-left)>>1

		if numbers[mid] > numbers[right] {
			left =  mid +1
		}else if numbers[mid] <= numbers[right] {
			right = right -1
		}
	}

	return numbers[left]
}