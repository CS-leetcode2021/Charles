package main

import (
	"fmt"
	"os"
)

/**
 *  @ClassName:5_string
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/20 下午4:23
 */

func main() {

	var strs string

	for  {
		n, _ := fmt.Scan(&strs)
		if n == 0 {
			os.Exit(0)
		}
		m := len(strs)

		tmp := map[byte]int{}
		res := []byte{}
		for i := 0; i < m; i++ {
			if _, ok := tmp[strs[i]]; !ok {
				tmp[strs[i]] = 1
				res = append(res,strs[i])
			}else {
				continue
			}
		}
	// byte 数组可以直接转化为切片
		fmt.Println(string(res))
	}
}
