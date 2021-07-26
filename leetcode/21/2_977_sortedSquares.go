package _1

/**
 *  @ClassName:2_977
 *  @Description:有序数组的平方
 *  @Author:jackey
 *  @Create:2021/7/26 上午10:53
 */


func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i := 0
	j := n - 1

	for pos := n - 1; pos >= 0; pos-- {
		if nums[i]*nums[i]<nums[j]*nums[j]{
			ans[pos] = nums[j]*nums[j]
			j--
		}else{
			ans[pos]= nums[i]*nums[i]
			i++
		}
	}
	return ans
}