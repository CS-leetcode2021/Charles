package main

import "math"

// 剑指offer-14-I 剪绳子 DP问题
// 同 leetcode-343 一样
// 积分证明：https://leetcode-cn.com/problems/jian-sheng-zi-lcof/solution/mian-shi-ti-14-i-jian-sheng-zi-tan-xin-si-xiang-by/

// 贪心：核心思路是：尽可能把绳子分成长度为3的小段，这样乘积最大
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	a, b := n/3, n%3

	switch b {
	case 0:
		return int(math.Pow(float64(3), float64(a)))
	case 1:
		return int(math.Pow(float64(3), float64(a-1)) * 4)
	}
	return int(math.Pow(float64(3), float64(a)) * 2)
}
func cuttingRope1(n int) int {
	if n < 4 {
		return n - 1
	}

	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}

	return res * n

}

// 动态规划---完全背包问题（类似）
// 1、想要长度为n的绳子剪掉后的最大乘积，可以从绳子长度为n-1的绳子转化过来
// 2、dp[n]，表示长度为n的绳子剪掉后的最大乘积，dp[i] 表示长度为i的绳子剪掉后的最大乘积
// 3、剪掉第一段j，如果j为1，则对乘积无影响，所以就从j=2开始。
// 4、剪掉j长度后的最大成绩为，剩下的i-j可以选择剪掉还是不减掉，如果不剪的话长度乘积即为j * (i - j)；如果剪的话长度乘积即为j * dp[i - j]。
// 取两者最大值max(j * (i - j), j * dp[i - j])
// 5、第一段长度j可以取的区间为[2,i)，对所有j不同的情况取最大值，因此最终dp[i]的转移方程为 dp[i] = max(dp[i], max(j * (i - j), j * dp[i - j]))

func cuttingRope2(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 2; j < i; j++ {
			dp[i] = MaxRope(dp[i],MaxRope(j*(i-j),j*dp[i-j]))
		}
	}
	return dp[n]
}

func MaxRope(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 数论证明
//
//任何大于1的数都可由2和3相加组成（根据奇偶证明）
//因为2*2=1*4，2*3>1*5, 所以将数字拆成2和3，能得到的积最大
//因为2*2*2<3*3, 所以3越多积越大 时间复杂度O(n/3)，用幂函数可以达到O(log(n/3))。 空间复杂度常数复杂度O(1)

// -------------------------------------------
// 剑指offer-14-II
//  快速幂求余：https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/solution/mian-shi-ti-14-ii-jian-sheng-zi-iitan-xin-er-fen-f/

func cuttingRope3(n int) int {
	if n <= 3 {
		return n-1
	}
	// n = 3a + b
	ret := 1
	// b 等于0，1，2情况最大乘4,所以n > 4
	// 每次都要进行取余操作
	for n > 4 {
		ret = ret * 3 % 1000000007
		n -= 3
	}
	// 最后结果只会剩下2,3,4 所以直接乘以n再取余1000000007
	return ret * n % 1000000007

}
