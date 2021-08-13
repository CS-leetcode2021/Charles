package _1_21

import "math"

/**
 *  @ClassName:05_209_minSubArrayLen
 *  @Description:长度最小的子数组
 *  @Author:jackey
 *  @Create:2021/8/13 下午7:29
 */

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum, l := 0, 0

	res := math.MaxInt32
	for r := 0; r < len(nums); r++ {
		sum += nums[r]

		for sum >= target {
			sum -= nums[l]
			l++
		}

		res = Min209(res,l-r+2)
	}

	if res == math.MaxInt32{	// 可能自始至终都没有进入循环
		return 0
	}
	return res

}

func Min209(i, j int) int {
	if i > j {
		return j
	}

	return i
}

// 官网

func minSubArrayLen2(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = Min209(ans, end - start + 1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}