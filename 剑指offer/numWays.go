package main

/**
 *  @ClassName:numWays
 *  @Description:offer-10-2 青蛙跳跃台阶问题，同leetcode 70
 *  @Author:jackey
 *  @Create:2021/7/6 下午8:10
 */

// 动态规划
func numWays(n int) int {
	if n <= 1 {
		return 1
	}

	tmp := make([]int, n+1)

	tmp[0] = 1
	tmp[1] = 1
	for i := 2; i <= n ; i++ {
		tmp[i] = tmp[i-1]+tmp[i-2]
	}

	return tmp[n]
}

// 滑动数组
func numWays1(n int) int {
	if n <= 1 {
		return 1
	}

	first := 1
	second := 1

	for i := 2; i <= n; i++ {
		tmp := (first+second)%1000000007
		first = second
		second = tmp
	}

	return second
}
