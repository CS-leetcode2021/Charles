package dp21

/**
 *  @ClassName:11_96_numTrees
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/26 下午2:21
 */

// dp问题
// G(n): 长度为 nn 的序列能构成的不同二叉搜索树的个数。
// F(i, n)F(i,n): 以 i 为根、序列长度为 n 的不同二叉搜索树个数 (1≤i≤n)
//	G[n] = 求和 G[i-1]*G[n-i]

func numTrees(n int) int {

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
