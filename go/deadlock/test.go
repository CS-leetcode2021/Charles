package main

import "fmt"

/**
 *  @ClassName:test
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/19 下午7:22
 */

func main() {
	n := 1024*8

	count := n%48
	res := n/48
	fmt.Println(res)
	fmt.Println(count)
}
