package main

import (
	"fmt"
	"math"
)

/**
 *  @ClassName:01bag.go
 *  @Description:01背包问题
 *  @Author:jackey
 *  @Create:2021/6/8 下午5:16
 */
func main01()  {
	N , V := 0,0
	fmt.Scanln(&N,&V)

	v := make([]int,N+1)
	w := make([]int,N+1)
	for i := 1; i <=N; i++ {
		fmt.Scanln(&v[i],&w[i])
	}

	/*创建二维数组*/
	dp := make([][]int,N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int,V+1)
	}
	dp[0][0] = 0

	for i := 1; i <= N; i++ {
		for j := 0; j <= V; j++ {
			dp[i][j]= dp[i-1][j]

			if j >= v[i] {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]),float64(dp[i-1][j-v[i]]+w[i])))
			}
		}
	}

	fmt.Println(dp[N][V])
}

func main()  {
	N , V := 0,0
	fmt.Scanln(&N,&V)

	v := make([]int,N+1)
	w := make([]int,N+1)
	for i := 1; i <=N; i++ {
		fmt.Scanln(&v[i],&w[i])
	}
	res := bag(N,V,v,w)
	fmt.Println(res)
}

func bag(num int , cap int , vol []int , worth []int) int {

	dp := make([]int,cap+1)

	for i := 1; i <= num; i++ {
		for j := cap; j >= vol[i] ; j-- {
			dp[j] = int(math.Max(float64(dp[j]),float64(dp[j-vol[i]]+worth[i])))

		}
	}

	return dp[cap]


}