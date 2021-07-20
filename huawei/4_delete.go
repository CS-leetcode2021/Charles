package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 *  @ClassName:4_delete
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/20 下午3:43
 */

func main() {
	var str string
	for {
		n, _ := fmt.Scan(&str)
		if n == 0 {
			os.Exit(0)
		}

		N, _ := strconv.Atoi(str)

		res := make([]int, N)

		for i := 0; i < len(res); i++ {
			res[i] = i
		}

		step := 0
		last := 0
		for N != 1 {
			for i := 0; i < len(res); i++ {
				if step == 2 && res[i] != -1 {
					step = 0
					N--
					res[i] = -1
				} else if res[i] == -1 {
					continue
				} else {
					step++
				}
			}
		}
		for i := 0; i < len(res); i++ {
			if res[i] != -1 {
				last = res[i]
			}
		}
		fmt.Println(last)
	}
}
