package _4

/**
 *  @ClassName:3_509_fin
 *  @Description:斐波那契数
 *  @Author:jackey
 *  @Create:2021/7/26 下午4:41
 */

func fib(n int) int {
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
