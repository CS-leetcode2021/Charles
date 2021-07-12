package main

import "strconv"

/**
 *  @ClassName:17_printNumbers
 *  @Description:剑指offer-17 打印从1到最大的n位数 99/85
 *  @Author:jackey
 *  @Create:2021/7/12 下午6:21
 */
// 面试时千万不要这么写，本题实际是考查的大数解法
func printNumbers(n int) []int {
	l := 1

	for i := 0; i < n; i++ {
		l *= 10
	}

	res := make([]int, l)

	for i, j := 1, l-1; i < j; i, j = i+1, j-1 {
		res[i],res[j] = i,j
	}
	res[l/2] = l/2
	return res[1:]
}


// 正常解法

/**
防止大数字溢出，使用字符串模拟数字加法，
低位加到等于10就进1，进到最高位的时候停止进位（999进1变成1000，此时停止进位）
*/
func printNumbers2(n int) []int {
	res := []int{}

	if n <= 0 {
		return res
	}

	number := make([]byte, n)

	for i := 0; i < n; i++ {
		number[i] = '0'
	}

	for !increment(&number, n) {
		if number[0] == '0' {
			tmp := make([]byte, n)
			copy(tmp, number)
			tmp = tmp[1:]
			add, _ := strconv.Atoi(string(tmp))
			res = append(res, add)
		} else {
			add, _ := strconv.Atoi(string(number))
			res = append(res, add)
		}
	}
	return res
}

func increment(number *[]byte, length int) bool {
	// 只有最高位进位的时候才会超过最大值
	isOverflow := false

	// 进位
	var takeOver byte
	takeOver = 0

	for i := length - 1; i >= 0; i-- {
		nSum := ((*number)[i] - '0') + takeOver

		// 从低位开始+1
		if i == length-1 {
			nSum++
		}

		if nSum >= 10 {
			// 需要进位

			if i == 0 {
				// 最高位不能进位， 返回进位结束
				isOverflow = true
			} else {
				nSum -= 10

				takeOver = 1
				(*number)[i] = '0' + nSum
			}
		} else {
			// 执行一次+1操作就退出
			(*number)[i] = '0' + nSum
			break
		}
	}
	return isOverflow
}
