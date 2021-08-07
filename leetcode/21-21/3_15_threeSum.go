package _1_21

import (
	"sort"
)

/**
 *  @ClassName:3_15_threeSum
 *  @Description:三数之和
 *  @Author:jackey
 *  @Create:2021/8/7 下午5:07
 */

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	sort.Ints(nums)

	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -1 * nums[i]
		for j, k := i+1, n-1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for j < k && nums[k]+nums[j] > target {
				k--
			}

			if j == k {
				break
			}

			if nums[j]+nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}

		}

	}
	return res
}
