package _1

/**
 *  @ClassName:12_70_climbStairs
 *  @Description:爬楼梯
 *  @Author:jackey
 *  @Create:2021/8/3 下午6:29
 */

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	first , second := 1,1

	for i := 2; i <= n; i++ {
		tmp := first+second
		second = first
		first = tmp
	}

	return first
}