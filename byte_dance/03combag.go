package main

import (
	"fmt"
	"math"
)

/**
 *  @ClassName:03combag
 *  @Description:03 完全背包
 *  @Author:jackey
 *  @Create:2021/6/8 下午8:07
 */

/*
 *  @Description:   所有中从前i个物品中选，总体积不超过j的方案所有集合
 *  @Param:         可以一直选则，只能没有空间可以选择为止，最长的比较结果会有N项比较值
 *  @Return:
 */
func main()  {
	N , V := 0,0
	fmt.Scanln(&N,&V)

	v := make([]int,N+1)
	w := make([]int,N+1)
	for i := 1; i <=N; i++ {
		fmt.Scanln(&v[i],&w[i])
	}
	res := bagCom2(N,V,v,w)
	fmt.Println(res)
}

func bagCom(num int , cap int , vol []int , worth []int) int {

	dp := make([]int,cap+1)

	for i := 1; i <= num; i++ {
		for j := vol[i]; j <= cap ; j++ {
			dp[j] = int(math.Max(float64(dp[j]),float64(dp[j-vol[i]]+worth[i])))

		}
	}

	return dp[cap]


}

func bagCom2(num int, cap int, vol []int, worth []int) int {
	dp := make([][]int,num+1)

	for i := 0; i <= num; i++ {
		dp[i] = make([]int,cap+1)
	}


	for i := 1; i <= num ; i++ {
		for j := 1; j <= cap; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= vol[i] {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]),float64(dp[i][j-vol[i]]+worth[i])))
			}
		}
	}

	return dp[num][cap]
}