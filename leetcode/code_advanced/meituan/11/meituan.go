package _11

import "fmt"

/**
 *  @ClassName:meituan
 *  @Description:美团第一题目
 *  @Author:jackey
 *  @Create:2021/8/8 下午12:50
 */

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N , K int
		fmt.Scan(&N,&K)
		nums := make([]int,N+1)

		for j := 0; j < N; j++ {
			tmp := 0
			fmt.Scan(&tmp)
			nums[tmp]++
		}

		// fmt.Println(nums)
		sum := 0
		for j := 1; j < K; j++ {
			 sum += nums[j]
		}

		if sum == K {
			fmt.Println("YES")
			fmt.Println(sum)
		}else {
			fmt.Println("NO")
		}

	}
}