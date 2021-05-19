package main

import (
	"unicode"
)

/**
 *  @ClassName:	calculate
 *  @Description:	需要用栈实现操作符的存储，是当前数字与上一位数字做运算
 *  @Author:	jackey
 *  @Create:	2021/3/10 下午9:30
 */

func calculate(s string) int {

	/*字符床中只有括号、加减法、空格*/

	/*创建一个Stack*/
	stack := make([]int, len(s))

	// 1为正 -1为负
	sign := 1

	/*res表示当前层的计算结果*/
	res := 0

	/*保存数字*/
	num := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			// 构建多位数（如果有）比如 "12" = 1*10+2 = 12
			j := i // j指针从当前的i开始
			num = 0
			for j < len(s) && unicode.IsDigit(rune(s[j])) { // 如果当前j字符是数字
				num = num*10 + int(s[j]-'0') //计算最新的num
				j++                          // j步进
			}
			res += sign * num
			i = j - 1
		case '(':
			stack = append(stack, res)
			stack = append(stack, sign)
			res = 0
			sign = 1

		case ')':
			sign = stack[len(stack)-1]
			prevRes := stack[len(stack)-2]
			/*出栈*/
			stack = stack[:len(stack)-2]
			res = prevRes + sign*res
		case '+':
			sign = 1
		case '-':
			sign = -1

		}

	}
	return res
}

/*双栈实现： TODO*/
