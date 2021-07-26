package __DfsBfs

import "sort"

/**
 *  @ClassName:47_permuteUnique
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/26 下午9:19
 */
// 放置重复的元素继续dfs
// 需要对数组进行排序
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	sort.Ints(nums) // 这⾥是去重的关键逻辑
	generatePermutation47(nums, 0, p, &res, &used)
	return res
}
func generatePermutation47(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] { // 这⾥是去重的关键逻辑 //和前面的相同，且前面没有被访问（当前也不会访问）
				continue
			}
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation47(nums, index+1, p, res, used)
			p = p[:len(p)-1]	// 回溯
			(*used)[i] = false	// 回溯
		}
	}
	return
}
