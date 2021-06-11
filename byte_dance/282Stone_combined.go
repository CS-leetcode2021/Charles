package main

import (
	"fmt"
	"math"
)

/**
 *  @ClassName:282Stone_combined
 *  @Description:282 石子合并问题
 *  @Author:jackey
 *  @Create:2021/6/11 下午5:58
 */

func main() {
	num := 0

	fmt.Scanln(&num)
	sum := make([]int, num+1)
	for i := 1; i <= num; i++ {
		fmt.Scan(&sum[i])		// Scan 只是按照空格分割数据
		sum[i] += sum[i-1]
	}

	weight := make([][]int, num+1)
	for i := 0; i <= num; i++ {
		weight[i] = make([]int, num+1)
	}

	res := Combined(num, weight, sum)
	fmt.Println(res)
}

func Combined(num int, weight [][]int, sum []int) int {

	for length := 2; length <= num; length++ { // 枚举区间长度，从2开始
		for l := 1; l +length- 1 <= num; l++ { // 区间左端点
			r := l + length - 1 // 区间右端点
			weight[l][r] = math.MaxInt32
			for k := l; k < r; k++ {
				weight[l][r] = int(math.Min(float64(weight[l][r]), float64(weight[l][k]+weight[k+1][r]+sum[r]-sum[l-1])))
			}

		}
	}

	return weight[1][num]
}
