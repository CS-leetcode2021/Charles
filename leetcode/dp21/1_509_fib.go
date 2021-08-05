package dp21

/**
 *  @ClassName:1_509_fib
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/5 下午6:45
 */

func fib(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	// dp问题，创建两个变量保存中间值
	tmp1, tmp2 := 1, 1
	for i := 3; i <= n; i++ {
		tmp1, tmp2 = tmp1+tmp2, tmp1
	}
	return tmp1
}

// 比之前的效果更好，代码简练
func fib2(n int) int {
	if n < 1 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	first := 1
	second := 1
	for i := 3; i <= n; i++ {
		sum := first + second
		second = first
		first = sum
	}
	return first
}
