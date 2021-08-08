package main

import "fmt"

/**
 *  @ClassName:main03
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/8 下午1:43
 */

func main() {
	var T int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var N, M int
		fmt.Scanln(&N)
		var sits, gender string
		sit := make([]int, N)
		fmt.Scanln(&sits)
		for j := 0; j < N; j++ {
			sit[j] = int(sits[j] - '0')
		}
		fmt.Scanln(&M)
		fmt.Scanln(&gender)

		//fmt.Println(sits)
		//fmt.Println(gender)

		res := prePlan(sit, gender)
		for j := 0; j < len(res); j++ {
			fmt.Println(res[j])
		}
	}
}

func prePlan(sits []int, gen string) []int {
	res := make([]int, 0)
	for i := 0; i < len(gen); i++ {
		if gen[i] == 'M' {
			index := judeNum(sits, 1)
			if index != -1 {
				sits[index] = 2
				res = append(res, index+1)
			} else {
				index = judeNum(sits, 0)
				sits[index] = 1
				res = append(res, index+1)
			}
		} else {
			index := judeNum(sits, 0)
			if index != -1 {
				sits[index] = 1
				res = append(res, index+1)
			} else {
				index = judeNum(sits, 1)
				sits[index] = 2
				res = append(res, index+1)
			}
		}
	}
	return res
}

func judeNum(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
