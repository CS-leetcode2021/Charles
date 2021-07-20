package main

import (
	"fmt"
	"os"
)

/**
 *  @ClassName:1_BuySoda
 *  @Description:买汽水
 *  @Author:jackey
 *  @Create:2021/7/20 下午1:03
 */

func main() {
	tmp := make([]int,0)
	var n int
	for {

		fmt.Scanln(&n)
		if n == 0 {
			break
		}

		tmp = append(tmp,n)
	}

	for i := 0; i < len(tmp); i++ {
		fmt.Println(help(tmp[i]))
	}
	os.Exit(0)

}

func help(n int) int {

	if n == 1 || n == 0 {
		return 0
	}
	res := make([]int,n+1)
	res[0] = 0
	res[1] = 0

	for i := 2; i <= n; i++ {
		res[i] = res[i-2] + 1
	}
	return res[n]
}