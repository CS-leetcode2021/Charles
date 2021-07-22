package main

// 剑指offer-15  二进制中1的个数
// &操作
func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			count++
		}
	}
	return count
}

// 优化
// 观察这个运算：)n & (n−1)，其预算结果恰为把 n 的二进制位中的最低位的 1 变为 0 之后的结果。
func hammingWeight2(num uint32) int {
	ones := 0
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return ones
}