package main

/**
 *  @ClassName:fib
 *  @Description:offer-10-1 斐波那契数列
 *  @Author:jackey
 *  @Create:2021/7/6 下午8:02
 */

// 动态规划
func fib(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	res := make([]int,n+1)

	res[0] = 0
	res[1] = 1

	for i := 2; i <= n; i++ {
		res[i] = (res[i-1]+res[i-2])%1000000007
	}

	return res[n]

}

// 滑动数组
func fib2(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	first := 0
	second := 1

	for i := 2; i <= n; i++ {
		tmp := (first+second)%1000000007
		first = second
		second = tmp
	}

	return second
}