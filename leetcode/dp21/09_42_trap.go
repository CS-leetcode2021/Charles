package dp21

/**
 *  @ClassName:09_42_trap
 *  @Description:leetcode-42 接雨水
 *  @Author:jackey
 *  @Create:2021/8/24 下午2:26
 */

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	// 维护下标i左右两边的最大高度
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]

	for i := 1; i < n; i++ {
		leftMax[i] = Max(height[i], leftMax[i-1])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = Max(height[i], rightMax[i+1])
	}

	res := 0
	for i, h := range height {
		res += Min(leftMax[i], rightMax[i]) - h
	}
	return res
}
