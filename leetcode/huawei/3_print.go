package main

import (
	"fmt"
	"math"
	"strings"
)

/**
 *  @ClassName:print
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/20 下午12:50
 */

// 考查大数运算

func main() {
	var list = []string{}
	for {
		var x string
		fmt.Scanln(&x)
		if x == "" {
			break
		}
		list = append(list, x)
	}

	for _, item := range list {
		test5(item)

	}
}

func test5(x string) {
	x = strings.Replace(x, "0x", "", 1)
	// rune 通常用于处理字符串
	runlist := []rune(x)
	//
	var total float64 = 0
	for i, r := range runlist {
		var c float64

		if r >= 'A' {
			c = float64(r) - 55	// 因为A代表10
		} else {
			c = float64(r) - 48	// 48 == 0
		}
		total += math.Pow(16,float64(len(runlist)-i-1))*c
	}
	fmt.Println(total)
}
