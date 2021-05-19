package array

import "sort"

/**
 *  @ClassName:singleNumber
 *  @Description:126 leetcode
 *  @Author:jackey
 *  @Create:2021/4/8 上午11:17
 */

func singleNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); {
		if i == len(nums)-1{
			return nums[i]
		}
		j := i+1
		if nums[i] != nums[j] {
			return nums[i]
		}
		i +=2
	}
	return -1
}

/*异或解法*/
/*所有数字：自己和自己做异或结果为0,任何数字和0做异或结果都是它本身，并且异或具有交换律*/
func singleNumber01(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}