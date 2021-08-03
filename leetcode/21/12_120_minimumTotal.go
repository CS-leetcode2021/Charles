package _1

import "math"

/**
 *  @ClassName:12_120_minimumTotal
 *  @Description:三角形最小路径和
 *  @Author:jackey
 *  @Create:2021/8/3 下午7:00
 */

// dp问题
// f[i][j] 从顶部走到（i,j）的最小路径和
// 当前的位置只能从（i-1,j）(i-1,j-1)转化过来
// 从顶到下，0-》i
func minimumTotal(nums [][]int) int {
	n := len(nums)
	// 访问数组
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}

	f[0][0] = nums[0][0]

	for i := 1; i < n; i++ {
		f[i][0] = f[i-1][0] + nums[i][0]
		for j := 1; j < i; j++ {
			f[i][j] = Min120(f[i-1][j], f[i-1][j-1]) + nums[i][j]
		}
		f[i][i] = f[i-1][i-1] + nums[i][i]
	}

	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = Min120(ans, f[n-1][i])
	}
	return ans
}

func Min120(i, j int) int {
	if i < j {
		return i
	}
	return j
}


// 从下到顶，i-》0
func minimumTotal2(triangle [][]int) int {
	h := len(triangle)
	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	for i := h - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == h-1 {
				dp[i][j] = triangle[i][j]
			} else {
				dp[i][j] = Min120(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}
