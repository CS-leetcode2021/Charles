package _1_14

/**
 *  @ClassName:11_46_permute
 *  @Description:全排列
 *  @Author:jackey
 *  @Create:2021/8/3 下午4:38
 */

func permute(nums []int) [][]int {
	res := make([][]int,0)

	return res
}

func backtracking46(nums []int, level int, res *[][]int) {
	if level == len(nums)-1 {
		tmp := make([]int,len(nums))
		copy(tmp,nums)
		*res = append(*res,tmp)
		return
	}

	for i := level; i < len(nums); i++ {
		swap46(nums,i,level)
		backtracking46(nums,level+1,res)
		swap46(nums,i,level)
	}
	return
}

func swap46(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}