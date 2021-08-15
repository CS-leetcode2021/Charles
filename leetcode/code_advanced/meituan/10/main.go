package main

import (
	"fmt"
	"sort"
)

/**
 *  @ClassName:server.go
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/8 下午1:17
 */

// 输入第一行仅包含三个正整数n,x,y，分别表示参赛的人数和晋级淘汰人数区间。(1<=n<=50000,1<=x,y<=n)
//
//输入第二行包含n个整数，中间用空格隔开，表示从1号选手到n号选手的成绩。(1<=|a_i|<=1000)

func main01() {
	var N, X, Y int
	fmt.Scan(&N, &X, &Y)

	Fraction := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&Fraction[i])
	}

	// fmt.Println(Fraction)
	sort.Ints(Fraction)

	res := -1
	for i := X; i < N; i++ {
		if N-i < Y {
			res = Fraction[i-1]
		}
		break
	}

	fmt.Println(res)

}
