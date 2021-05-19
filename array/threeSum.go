package array

import (
	"sort"
)

/**
 *  @ClassName:threeSum
 *  @Description:leetcode 15 mid
 *  @Author:jackey
 *  @Create:2021/5/19 上午11:58
 */
/**
给定一个无序数组，找出所有和为0的三元组
思路：	1、先进行数组排序
	  	2、先固定一个target，后续双指针遍历
		3、注意判定相等的条件
*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	// 双循环,指针滑动
	for first := 0; first < n; first++ {
		// 相同的跳过
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return res
}
