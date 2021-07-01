package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 *  @ClassName:replaceSpace
 *  @Description:剑指offer-05
 *  @Author:jackey
 *  @Create:2021/7/1 下午7:27
 */

/*
 *  @Description:   byte类型可以直接转化为string类型
 *  @Param:
 *  @Return:
 */
func main() {
	// 标准输入
	input_buffer := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := input_buffer.ReadString('\n')

	if err != nil {
		return
	}
	fmt.Println(input)
}

func replaceSpace(s string) string {
	n := len(s)

	res := []byte{}

	rep := []byte{'%','2','0'}
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			res = append(res,rep...)
			continue
		}
		res = append(res, s[i])
	}


	return string(res)
}

func replaceSpace01(s string) string {
	var res strings.Builder
	for i:=range s{
		if s[i]==' '{
			res.WriteString("%20")
		}else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}