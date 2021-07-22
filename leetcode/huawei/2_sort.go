package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
 *  @ClassName:2_sort
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/20 下午2:39
 */

func main() {
	var str string

	for  {
		n,_ := fmt.Scan(&str)

		if n == 0 {
			break
		}
		num,_:= strconv.Atoi(str)
		res := make([]int, num)
		for i := 0; i < num; i++ {
			fmt.Scanln(&res[i])
		}

		sort.Ints(res)
		tag := 0
		for i := 1; i < len(res); i++ {
			if res[i] == res[tag] {
				continue
			}
			tag++
			res[tag] = res[i]
		}
		res = res[:tag+1]
		for i := 0; i < len(res); i++ {
			fmt.Println(res[i])
		}
	}
}
