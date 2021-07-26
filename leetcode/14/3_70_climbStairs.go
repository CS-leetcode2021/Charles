package _4

/**
 *  @ClassName:3_70_climbStairs
 *  @Description:爬楼梯
 *  @Author:jackey
 *  @Create:2021/7/26 下午4:46
 */

// 100/100 双指针
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	first := 1	// n == 1
	second := 1	// n == 0

	for i := 2; i <= n; i++ {
		sum := first + second
		second = first
		first = sum
	}

	return first
}

// 添加辅助数组进行记录