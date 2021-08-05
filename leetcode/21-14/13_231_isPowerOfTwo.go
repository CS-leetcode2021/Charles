package _1_14

/**
 *  @ClassName:13_231_isPowerOfTwo
 *  @Description:2的幂
 *  @Author:jackey
 *  @Create:2021/8/3 下午7:47
 */

// 如果 n 是正整数并且 n & (n - 1) = 0，那么 n 就是 2 的幂。
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
