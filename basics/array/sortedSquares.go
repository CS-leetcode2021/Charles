package array

import "sort"

/**
 *  @ClassName:sortedSquares
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/3/16 下午5:05
 */

/*时间复杂度：O(nlogn)*/
func sortedSquares(nums []int) []int {
	var res = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		res[i] = nums[i] * nums[i]
	}

	sort.Ints(res)

	return res
}

func sortedS(nums []int) []int {
	n := len(nums)
	neg := -1

	for i := 0; i < n && nums[i] < 0; i++ {
		neg = i
	}

	ans := make([]int, 0, n)
	for i, j := neg, neg+1; i >= 0 || j < n; {
		if i < 0 {
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == n {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] <= nums[j]*nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}
	return ans
}

func sortedSQU(nums []int) []int {
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
