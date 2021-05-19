package main

import "fmt"

/**
 *  @ClassName: fib
 *  @Description: 斐波那契数
 *  @Author: jackey
 *  @Create:2021/3/16 上午10:14
 */

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fib(n-1)+fib(n-2)
}

func main() {
	fmt.Println(fib3(3))
	fmt.Println(fib3(4))
	fmt.Println(fib3(5))
	fmt.Println(fib3(6))
}

func fib2(n int)int  {
	if n < 1 {
		return 0
	}
	// 创建保存中间值的数组，初始化为0
	var memo = make([]int,n+1)
	//fmt.Println(memo)

	// 进行带备忘录的递归查询
	return hepler(memo,n)
}

func hepler(memo []int,n int)  int {

	/*基础计算*/
	if n == 1 || n == 2 {
		return 1
	}

	/*直接从数组取值*/
	if memo[n] != 0	{
		return memo[n]
	}

	/*递归运算*/
	memo[n]=hepler(memo,n-1)+hepler(memo,n-2)
	return memo[n]
}

/*自底向上*/
func fib3(n int) int {
	if n < 1 {
		return 0
	}

	var table = make([]int,n+1)
	table[1]=1
	table[2]=1
	for i := 3; i < len(table); i++ {
		table[i]=table[i-1]+table[i-2]
	}
	return table[n]
}

/*状态压缩*/

func fib4(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	pre1 :=1
	pre2 :=1
	sum :=0
	for i := 3; i <= n; i++ {
		sum = pre1+pre2
		pre1 = pre2
		pre2= sum
	}
	return sum
}