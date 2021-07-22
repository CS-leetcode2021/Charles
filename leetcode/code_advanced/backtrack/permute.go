package backtrack

import "fmt"

/**
 *  @ClassName:permute
 *  @Description:46 leetcode 全排列
 *  @Author:jackey
 *  @Create:2021/4/8 上午9:57
 */

func permute01(nums []int) [][]int {
	if len(nums)==1{
		return [][]int{nums}
	}

	res := [][]int{}

	for i,num := range nums{
		tmp := make([]int,len(nums)-1)
		copy(tmp[0:],nums[0:i])
		copy(tmp[i:],nums[i+1:])

		sub := permute01(tmp)
		for _,s := range sub{
			res = append(res,append(s,num))
		}
	}
	return res
}

var result [][]int
func backtrack(nums, pathNums []int, used []bool) {
	if len(nums) == len(pathNums) {
		tmp := make([]int,len(nums))
		copy(tmp,pathNums)
		result = append(result,tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i]= true

			pathNums= append(pathNums,nums[i])
			backtrack(nums,pathNums,used)

			/*撤销*/
			pathNums= pathNums[:len(pathNums)-1]
			used[i]=false
		}
	}
}

func permute(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	// 清空全局数组（leetcode多次执行全局变量不会消失）
	result = [][]int{}
	backtrack(nums, pathNums, used)
	return result
}

func mainTest() {
	a := []int{1,2,3}

	res :=permute(a)
	fmt.Println(res)
}