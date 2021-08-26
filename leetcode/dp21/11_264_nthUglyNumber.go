package dp21

/**
 *  @ClassName:11_264_nthUglyNumber
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/26 下午2:07
 */

// 相当于3个数组，分别是能被2、3、5整除的递增数组，且每个数组的第一个数都为1。
//
// 然后就简单了，维护三个指针，将三个数组合并为一个严格递增的数组。就是传统的双指针法，只是这题是三个指针。
//
// 然后优化一下，不要一下子列出这3个数组，因为并不知道数组预先算出多少合适。
//
// 这样就一边移指针，一边算各个数组的下一个数，一边merge，就变成了题解的动态规划法的代码
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)

	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5

		dp[i] = Min(Min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}
