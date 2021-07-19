package main

import "strconv"

/**
 *  @ClassName:44_findNthDigit
 *  @Description:剑指offer-44 数字序列中某一位的数字 同 leetcode-400
 *  @Author:jackey
 *  @Create:2021/7/19 下午4:55
 */

// 1、将101112...中的每一位称为数位，记为n；
// 2、将10、11、12 称为数字，记为num；
// 3、数字10 是一个两位数，称此数字的位数为2,记为digit；
// 4、每一个digit的起始数字为start，1，10,100
func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9

	for n > count {
		n -= count
		digit++
		start *= 10
		count = digit*start*9
	}
	num := start + (n-1)/digit
	index := (n-1)%digit

	numstr := strconv.Itoa(num)
	return int(numstr[index]-'0')

}
