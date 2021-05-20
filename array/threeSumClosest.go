package array

import (
	"math"
	"sort"
)

/**
 *  @ClassName:threeSumClosest
 *  @Description:leetcode 16 最接近的三数之和
 *  @Author:jackey
 *  @Create:2021/5/20 下午2:00
 */

/*
 *  @Description:   暴力解会超时，O（n3）的时间复杂度
 * 	@ 	1、
 *  @Param:         
 *  @Return:        
 */

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	best := math.MaxInt32

	// 需要写一个函数来计算三元组的和与target的差值，返回最小的那组

	update := func(cur int) {

		if abs(cur-target) < abs(best-target) {
			best = cur
		}

	}
	// 进行枚举
	for i := 0; i < n; i++ {

		// 保证每次枚举的值不一样
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i+1
		k := n-1
		// 双指针判定
		for j < k {
			sum := nums[i]+nums[j]+nums[k]
			if sum == target{
				 return sum
			}

			update(sum)
			if sum > target {
				k0 := k-1
				// 移动下一个不想等的元素
				for j < k0 && nums[k] == nums[k0] {
					k0--
				}
				k = k0
			}else{
				j0:=j+1
				// 移动到下一个不想等的元素
				for k > j0 && nums[j] == nums[j0] {
					j0++
				}
				j = j0
			}

		}

	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}else {
		return x
	}
}