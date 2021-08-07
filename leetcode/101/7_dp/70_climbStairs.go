package __dp

/**
 *  @ClassName:70_climbStairs
 *  @Description:爬梯子
 *  @Author:jackey
 *  @Create:2021/8/7 下午8:45
 */

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	pre, cur := 1, 1
	for i := 2; i <= n; i++ {
		pre, cur = cur, cur+pre
	}

	return cur
}
