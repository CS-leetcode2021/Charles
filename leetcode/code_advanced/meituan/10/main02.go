package main

import (
	"fmt"
	"sort"
)

/**
 *  @ClassName:main02
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/8 下午1:37
 */

//输入第一行仅包含一个正整数n，表示任意序列的长度。(1<=n<=20000)
//
//输入第二行包含n个整数，表示给出的序列，每个数的绝对值都小于10000。

func main02() {
	var N int

	fmt.Scanln(&N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	// fmt.Println(nums) ok
	sort.Ints(nums)

	count := 0
	for i := 0; i < N; i++ {
		if nums[i] != i+1 {
			tmp := nums[i] - (i + 1)
			if tmp < 0 {
				tmp =- tmp
			}
			count += tmp
		}
	}
	fmt.Println(count)

}
