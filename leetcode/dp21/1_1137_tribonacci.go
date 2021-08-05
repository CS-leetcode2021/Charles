package dp21

/**
 *  @ClassName:1_1137_tribonacci
 *  @Description:第 N 个泰波那契数
 *  @Author:jackey
 *  @Create:2021/8/5 下午6:49
 */

// fib的变形，保存三格临时数值变量
func tribonacci(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	tmp1, tmp2, tmp3 := 1, 1, 0
	for i := 3; i <= n; i++ {
		tmp1, tmp2, tmp3 = tmp1+tmp2+tmp3, tmp1, tmp2
	}
	return tmp1
}
