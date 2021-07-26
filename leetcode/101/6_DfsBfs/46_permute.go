package __DfsBfs

/**
 *  @ClassName:46_permute
 *  @Description:全排列
 *  @Author:jackey
 *  @Create:2021/7/25 下午5:52
 */

// 100/92
// 对于每一个当前位置i，我们可以将其于之后的任意位置交换，然后继续处理位置i+1，直到处理到最后一位，为了防止我们每次遍历
// 时都需要新建一个子数字储存位置i之前已经换好的数字，可以利用回溯法，只对原数组进行修改
func permute(nums []int) [][]int {

	res := make([][]int, 0)
	backtracking46(nums,0,&res)
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
