package dp21

/**
 *  @ClassName:2_climbStairs
 *  @Description:爬楼梯
 *  @Author:jackey
 *  @Create:2021/8/6 下午3:31
 */

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	tmp1, tmp2 := 1, 1

	for i := 2; i <= n; i++ {
		tmp1, tmp2 = tmp1+tmp2, tmp1
	}
	return tmp1
}
